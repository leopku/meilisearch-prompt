package meilisearch

import (
	"fmt"

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
	for _, s := range settings.Synonyms {
		fmt.Println("\t", s)
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
		for _, s := range *resp {
			fmt.Println("\t", s)
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
