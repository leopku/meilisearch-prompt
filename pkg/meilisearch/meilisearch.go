package meilisearch

import (
	"fmt"

	"github.com/duke-git/lancet/validator"
	"github.com/gookit/goutil/strutil"
	"github.com/meilisearch/meilisearch-go"
)

type Meilisearch struct {
	Cli *meilisearch.Client
}

func NewMeilisearch(host string) *Meilisearch {
	cli := meilisearch.NewClient(meilisearch.ClientConfig{Host: host})
	return &Meilisearch{Cli: cli}
}

func (m *Meilisearch) GetAllIndexes() ([]string, error) {
	indexes, err := m.Cli.GetAllIndexes()
	if err != nil {
		return nil, err
	}
	var result []string
	for _, idx := range indexes {
		result = append(result, idx.UID)
	}
	return result, nil
}

func (m *Meilisearch) FetchInfo(idxUid string) {
	i := m.Cli.Index(idxUid)
	resp, err := i.FetchInfo()
	if err != nil {
		fmt.Println("Fetch info failed: ", err)
		return
	}
	fmt.Println(resp)
}

func (m *Meilisearch) GetSettings(i *meilisearch.Index) {
	settings, err := i.GetSettings()
	if err != nil {
		fmt.Println("Get settings failed: ", err)
		return
	}
	fmt.Println("Displayed Attributes:")
	for _, s := range settings.DisplayedAttributes {
		fmt.Println("\t", s)
	}
	fmt.Println("Searchable Attributes:")
	for _, s := range settings.SearchableAttributes {
		fmt.Println("\t", s)
	}
	fmt.Println("Filterable Attributes:")
	for _, s := range settings.FilterableAttributes {
		fmt.Println("\t", s)
	}
	fmt.Println("Sortable Attributes:")
	for _, s := range settings.SortableAttributes {
		fmt.Println("\t", s)
	}
	fmt.Println("Ranking Rules:")
	for _, s := range settings.RankingRules {
		fmt.Println("\t", s)
	}
	fmt.Println("Stop Words:")
	for _, s := range settings.StopWords {
		fmt.Println("\t", s)
	}
	fmt.Println("Synonyms:")
	for k, v := range settings.Synonyms {
		fmt.Println("\t", k, "\t", v)
	}
	fmt.Println("Distinct Attribute:")
	fmt.Println("\t", settings.DistinctAttribute)

}

func (m *Meilisearch) GetSettingsItem(i *meilisearch.Index, item string) {
	switch strutil.Lowercase(item) {
	case "displayed-attributes":
		resp, err := i.GetDisplayedAttributes()
		if err != nil {
			fmt.Println("Get displayed-attributes settings failed:\t", err)
			return
		}
		fmt.Println("Displayed Attributes:")
		for _, s := range *resp {
			fmt.Println("\t", s)
		}
	case "searchable-attributes":
		resp, err := i.GetSearchableAttributes()
		if err != nil {
			fmt.Println("Get searchable-attributes settings failed:\t", err)
			return
		}
		fmt.Println("Searchable Attributes:")
		for _, s := range *resp {
			fmt.Println("\t", s)
		}
	case "filterable-attributes":
		resp, err := i.GetFilterableAttributes()
		if err != nil {
			fmt.Println("Get filterable-attributes settings failed:\t", err)
			return
		}
		fmt.Println("Filterable Attributes:")
		for _, s := range *resp {
			fmt.Println("\t", s)
		}
	case "sortable-attributes":
		resp, err := i.GetSortableAttributes()
		if err != nil {
			fmt.Println("Get sortable-attributes settings failed:\t", err)
			return
		}
		fmt.Println("Sortable Attributes:")
		for _, s := range *resp {
			fmt.Println("\t", s)
		}
	case "ranking-rules":
		resp, err := i.GetRankingRules()
		if err != nil {
			fmt.Println("Get ranking-rules settings failed:\t", err)
			return
		}
		fmt.Println("Ranking Rules:")
		for _, s := range *resp {
			fmt.Println("\t", s)
		}
	case "stop-words":
		resp, err := i.GetStopWords()
		if err != nil {
			fmt.Println("Get stop-words settings failed:\t", err)
			return
		}
		fmt.Println("Stop Words:")
		for _, s := range *resp {
			fmt.Println("\t", s)
		}
	case "synonyms":
		resp, err := i.GetSynonyms()
		if err != nil {
			fmt.Println("Get synonyms settings failed:\t", err)
			return
		}
		fmt.Println("Synonyms:")
		for k, v := range *resp {
			fmt.Println("\t", k, "\t", v)
		}
	case "distinct-attribute":
		resp, err := i.GetDisplayedAttributes()
		if err != nil {
			fmt.Println("Get distinct-attribute settings failed:\t", err)
			return
		}
		fmt.Println("Distinct Attribute:")
		for _, s := range *resp {
			fmt.Println("\t", s)
		}
	}
}

func (m *Meilisearch) SetSettingsItem(i *meilisearch.Index, item string, params *[]string) {
	switch strutil.Lowercase(item) {
	case "displayed-attributes":
		resp, err := i.UpdateDisplayedAttributes(params)
		if err != nil {
			fmt.Println("Set displayed-attributes failed:\t", err)
			return
		}
		fmt.Println("Task UID:\t", resp.UID)
	case "searchable-attributes":
		resp, err := i.UpdateSearchableAttributes(params)
		if err != nil {
			fmt.Println("Set searchable-attributes failed:\t", err)
			return
		}
		fmt.Println("Task UID:\t", resp.UID)
	case "filterable-attributes":
		resp, err := i.UpdateFilterableAttributes(params)
		if err != nil {
			fmt.Println("Set filterable-attributes failed:\t", err)
			return
		}
		fmt.Println("Task UID:\t", resp.UID)
	case "sortable-attributes":
		resp, err := i.UpdateSortableAttributes(params)
		if err != nil {
			fmt.Println("Set sortable-attributes failed:\t", err)
			return
		}
		fmt.Println("Task UID:\t", resp.UID)
	case "ranking-rules":
		resp, err := i.UpdateRankingRules(params)
		if err != nil {
			fmt.Println("Set ranking-rules failed:\t", err)
			return
		}
		fmt.Println("Task UID:\t", resp.UID)
	case "stop-words":
		resp, err := i.UpdateStopWords(params)
		if err != nil {
			fmt.Println("Set stop-words failed:\t", err)
			return
		}
		fmt.Println("Task UID:\t", resp.UID)

	case "synonyms":
		// resp, err := i.UpdateSynonyms(params.(*map[string][]string))
		// if err != nil {
		// 	fmt.Println("Set sortable-attributes failed:\t", err)
		// 	return
		// }
		// fmt.Println("Task UID:\t", resp.UID)
		fmt.Println("NOT Implenment yet.")
	case "distinct-attribute":
		resp, err := i.UpdateDistinctAttribute((*params)[0])
		if err != nil {
			fmt.Println("Set sortable-attributes failed:\t", err)
			return
		}
		fmt.Println("Task UID:\t", resp.UID)
	}
}

func (m *Meilisearch) CreateIndex(uid, primaryKey string) (*meilisearch.Task, error) {
	cfg := meilisearch.IndexConfig{Uid: uid}
	if !validator.IsEmptyString(primaryKey) {
		cfg.PrimaryKey = primaryKey
	}
	return m.Cli.CreateIndex(&cfg)
}

func (m *Meilisearch) UpdateIndex(uid, primaryKey string) (*meilisearch.Key, error) {
	key := &meilisearch.Key{Key: primaryKey}
	return m.Cli.UpdateKey(uid, key)
}

func (m *Meilisearch) GetIndexFields(idxUid string) []string {
	i := m.Cli.Index(idxUid)
	resp, err := i.GetStats()
	if err != nil {
		fmt.Println("Get fields failed: ", err)
		return nil
	}
	var result []string
	for k := range resp.FieldDistribution {
		result = append(result, k)
	}
	return result
}

func (m *Meilisearch) FetchTask(id int64) {
	resp, err := m.Cli.GetTask(id)
	if err != nil {
		fmt.Println("Get task failed: ", err)
		return
	}
	fmt.Println("Task #", id, " :")
	fmt.Println("\tstatus: ", resp.Status)
}
