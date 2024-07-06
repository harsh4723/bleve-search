package main

import (
	"fmt"
	"os"

	"github.com/blevesearch/bleve/v2"
)

type FileData struct {
	FileName string `json:"name"`
	FilePath string `json:"path"`
}

func main() {
	// open a new index
	folderPath := "example.bleve"

	err1 := os.RemoveAll(folderPath)
	if err1 != nil {
		fmt.Println("Error deleting folder:", err1)
		return
	}
	mapping := bleve.NewIndexMapping()
	// err := mapping.AddCustomTokenFilter("color_stop_filter", map[string]interface{}{
	// 	"type": "stop_token_map",
	// 	"tokens": []interface{}{
	// 		"red",
	// 		"green",
	// 		"blue",
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	index, err := bleve.New("example.bleve", mapping)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := []FileData{}
	data = append(data, FileData{"copy.go", "/y/j"})
	data = append(data, FileData{"mod.go", "/y/x"})

	// index some data
	index.Index("id", data)

	// search for some text
	query := bleve.NewMatchQuery("/y/*")
	search := bleve.NewSearchRequest(query)
	search.Highlight = bleve.NewHighlight()
	search.Highlight.AddField("name")
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults)
}
