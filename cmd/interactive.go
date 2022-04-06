/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/leopku/meilisearch-prompt/pkg/completer"
	"github.com/leopku/meilisearch-prompt/pkg/meilisearch"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
)

var subCommands []string
var cmdSuggestions []prompt.Suggest
var c *completer.Completer

// interactiveCmd represents the interactive command
var interactiveCmd = &cobra.Command{
	Use:     "interactive <http://localhost:7700>",
	Aliases: []string{"i"},
	Short:   "enter interactive mode",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, subCmd := range rootCmd.Commands() {
			subCommands = append(subCommands, subCmd.Use)
			cmdSuggestions = append(cmdSuggestions, prompt.Suggest{Text: subCmd.Use, Description: subCmd.Short})
		}

		ms := meilisearch.NewMeilisearch(args[0])
		c = completer.NewCompleter(ms, args[0])
		p := prompt.New(
			c.Executor,
			c.Completer,
			prompt.OptionLivePrefix(c.PromptPrefix),
			prompt.OptionInputTextColor(prompt.Yellow),
		)
		p.Run()
	},
}

func init() {
	rootCmd.AddCommand(interactiveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// interactiveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// interactiveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// func executor(in string) {
// 	// in = strings.TrimSpace(in)
// 	// args := strings.Split(in, " ")
// 	args := filter.StrToSlice(in, " ")
// 	if len(args) == 0 {
// 		return
// 	}
// 	// log.Debug().Str("args 0", args[0]).Msg("")

// 	switch args[0] {
// 	case "ls":
// 		complete.FSM.Event("ls")
// 	case "cd":
// 		if len(args) != 2 {
// 			fmt.Println("Wrong parameter number.")
// 			return
// 		}
// 		currentIndex := args[1]
// 		if arrutil.Contains([]string{"/", ".."}, currentIndex) {
// 			complete.FSM.Event("out")
// 		} else {
// 			complete.FSM.Event("in", currentIndex)
// 		}
// 	case "quite", "exit", "q":
// 		fmt.Println("Bye!")
// 		os.Exit(0)
// 	}
// }
