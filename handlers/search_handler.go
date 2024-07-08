package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blevesearch/bleve/v2"
)

func SearchHandler(index bleve.Index) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		if query == "" {
			http.Error(w, "Query parameter is required", http.StatusBadRequest)
			return
		}
		fmt.Println("Search for query", query)
		matchQuery := bleve.NewWildcardQuery(query)
		searchRequest := bleve.NewSearchRequest(matchQuery)
		searchResult, err := index.Search(searchRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(searchResult)
	}
}
