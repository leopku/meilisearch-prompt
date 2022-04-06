package completer

import (
	"fmt"
	"os"

	"github.com/leopku/meilisearch-prompt/pkg/meilisearch"
	"github.com/looplab/fsm"

	"github.com/c-bata/go-prompt"
	"github.com/dlclark/regexp2"
	"github.com/gookit/filter"
	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/strutil"
	"github.com/phuslu/log"
)

type Completer struct {
	MS           *meilisearch.Meilisearch
	Host         string
	CurrentIndex string
	FSM          *fsm.FSM
}

func NewCompleter(ms *meilisearch.Meilisearch, host string) *Completer {
	c := &Completer{
		MS:           ms,
		Host:         host,
		CurrentIndex: "",
	}
	c.FSM = fsm.NewFSM(
		"root",
		fsm.Events{
			{Name: "in", Src: []string{"root", "index"}, Dst: "index"},
			{Name: "out", Src: []string{"index"}, Dst: "root"},
			{Name: "ls", Src: []string{"root"}, Dst: "root"},
			{Name: "ls", Src: []string{"index"}, Dst: "index"},
			{Name: "create", Src: []string{"root"}, Dst: "root"},
			{Name: "create", Src: []string{"index"}, Dst: "index"},
			{Name: "update", Src: []string{"index"}, Dst: "index"},
			{Name: "delete", Src: []string{"index"}, Dst: "root"},
			{Name: "settings", Src: []string{"index"}, Dst: "index"},
		},
		fsm.Callbacks{
			"in":       c.inFunc,
			"out":      c.outFunc,
			"ls":       c.lsFunc,
			"create":   c.createFunc,
			"update":   c.updateFunc,
			"delete":   c.deleteFunc,
			"settings": c.settingsFunc,
		},
	)
	return c
}

func (c *Completer) inFunc(e *fsm.Event) {
	c.CurrentIndex = e.Args[0].(string)
}

func (c *Completer) outFunc(e *fsm.Event) {
	c.CurrentIndex = ""
}

func (c *Completer) lsFunc(e *fsm.Event) {
	indexes, err := c.MS.GetAllIndexes()
	if err != nil {
		fmt.Println("Get indexes error: ", err)
		return
	}
	// fmt.Println(strings.Join(indexes, "\n"))
	for _, s := range indexes {
		fmt.Println("\t", s)
	}
}

func (c *Completer) createFunc(e *fsm.Event) {
	// if len(e.Args) == 0 {
	// 	return
	// }
	log.Debug().Interface("args", e.Args[0].(string)).Msg("")
	// args, ok := e.FSM.Metadata("args")
	// log.Debug().Bool("ok", ok).Msg("")
	// if ok {
	// 	fmt.Println(args.([]string))
	// }
	currentIndex := e.Args[0].(string)
	if arrutil.Contains([]string{"/", ".."}, currentIndex) {
		currentIndex = ""
	}
	c.CurrentIndex = currentIndex
}

func (c *Completer) updateFunc(e *fsm.Event) {

}

func (c *Completer) deleteFunc(e *fsm.Event) {

}

func (c *Completer) settingsFunc(e *fsm.Event) {
	args := e.Args[0].([]string)
	switch len(args) {
	case 1:
		c.getSettings()
	case 2:
		c.getSettingsItem(args[1])
	}
}

func (c *Completer) indexSuggestions(in string) (suggestions []prompt.Suggest) {
	indexes, err := c.MS.Cli.GetAllIndexes()
	if err != nil {
		return []prompt.Suggest{}
	}
	suggestions = append(suggestions, prompt.Suggest{Text: "..", Description: "Clear current selected index and change to root"})
	for _, idx := range indexes {
		suggestions = append(suggestions, prompt.Suggest{Text: idx.UID, Description: fmt.Sprintf("primary key: %s", idx.PrimaryKey)})
	}
	return
}

func (c *Completer) rootCommandSuggestion(in string) (suggestions []prompt.Suggest) {
	return []prompt.Suggest{
		{Text: "ls", Description: "list all indexes"},
		{Text: "cd", Description: "select an index as operation target"},
		{Text: "create", Description: "create a new index"},
	}
}

func (c *Completer) indexCommandSuggestions(in string) (suggestions []prompt.Suggest) {
	return []prompt.Suggest{
		{Text: "update", Description: "update index with new primary index"},
		{Text: "delete", Description: "delete current index"},
		{Text: "settings", Description: "show settings of current index"},
	}
}

func (c *Completer) settingsCommandSuggestions(in string) (suggestions []prompt.Suggest) {
	// args := filter.StrToSlice(in, " ")
	return []prompt.Suggest{
		{Text: "displayed-attributes"},
		{Text: "searchable-attributes"},
		{Text: "filterable-attributes"},
		{Text: "sortable-attributes"},
		{Text: "ranking-rules"},
		{Text: "stop-words"},
		{Text: "synonyms"},
		{Text: "distinct-attribute"},
	}
}

func (c *Completer) Executor(in string) {
	// in = strings.TrimSpace(in)
	// args := strings.Split(in, " ")
	args := filter.StrToSlice(in, " ")
	if len(args) == 0 {
		return
	}
	// log.Debug().Str("args 0", args[0]).Msg("")

	switch args[0] {
	case "ls":
		c.FSM.Event("ls")
	case "cd":
		indexes, err := c.MS.GetAllIndexes()
		if err != nil {
			fmt.Println("Error: %w", err)
			return
		}
		if len(args) != 2 || !arrutil.Contains(indexes, args[1]) {
			fmt.Println("Wrong parameter number.")
			return
		}
		currentIndex := args[1]
		if arrutil.Contains([]string{"/", ".."}, currentIndex) {
			c.FSM.Event("out")
		} else {
			c.FSM.Event("in", currentIndex)
		}
	case "settings":
		// switch len(args) {
		// case 1:
		// 	c.getSettings()
		// case 2:
		// 	c.getSettingsItem(args[1])
		// }
		c.FSM.Event("settings", args)
	case "quite", "exit", "q":
		fmt.Println("Bye!")
		os.Exit(0)
	}
}

func (c *Completer) Completer(in prompt.Document) (suggestions []prompt.Suggest) {
	tbc := in.TextBeforeCursor()
	suggestions = c.rootCommandSuggestion(tbc)
	args := filter.StrToSlice(tbc, " ")
	cmd, filtered := "", ""
	if len(args) > 0 {
		cmd = args[0]
		filtered = args[len(args)-1]
	}
	if strutil.IsNotBlank(c.CurrentIndex) {
		suggestions = append(suggestions, c.indexCommandSuggestions(filtered)...)
	}
	if strutil.HasOnePrefix(cmd, []string{"cd"}) {
		suggestions = c.indexSuggestions(filtered)
	}
	if strutil.HasOnePrefix(cmd, []string{"settings"}) {
		suggestions = c.settingsCommandSuggestions(filtered)
	}
	return filterSuggestion(suggestions, filtered)
}

func (c *Completer) PromptPrefix() (string, bool) {
	// log.Debug().Str("current index", c.CurrentIndex).Msg("")
	if strutil.IsBlank(c.CurrentIndex) {
		return fmt.Sprintf("%s>>> ", c.Host), true
	}
	return fmt.Sprintf("%s/indexes/%s>>> ", c.Host, c.CurrentIndex), true
}

func filterSuggestion(in []prompt.Suggest, arg string) (out []prompt.Suggest) {
	out = in
	filtered := prompt.FilterHasPrefix(in, arg, true)
	if len(filtered) > 0 {
		out = filtered
	}
	return
}

func parseCommand(arg string) ([]string, error) {
	re := regexp2.MustCompile(`(\S+)`, 0)
	m, err := re.FindStringMatch(arg)
	if err != nil {
		return nil, err
	}
	var result []string
	for m != nil {
		result = append(result, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return result, nil
}

func (c *Completer) getSettings() {
	if strutil.IsBlank(c.CurrentIndex) {
		return
	}
	i := c.MS.Cli.Index(c.CurrentIndex)
	c.MS.GetSettings(i)
}

func (c *Completer) getSettingsItem(item string) {
	if strutil.IsBlank(c.CurrentIndex) || strutil.IsBlank(item) {
		return
	}
	i := c.MS.Cli.Index(c.CurrentIndex)
	// switch strutil.Lowercase(item) {
	// case "displayed-attributes":

	// }
	c.MS.GetSettingsItem(i, item)
}