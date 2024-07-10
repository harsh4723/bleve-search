package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/blevesearch/bleve/v2"
)

func SearchHandler(index bleve.Index) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		query := r.URL.Query().Get("query")
		if query == "" {
			http.Error(w, "Query parameter is required", http.StatusBadRequest)
			return
		}
		matchQuery := bleve.NewWildcardQuery(query)
		searchRequest := bleve.NewSearchRequest(matchQuery)
		searchResult, err := index.Search(searchRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		totalSearchTime := float64(time.Since(startTime).Microseconds())
		fmt.Printf("Time taken for search %s is %f \n", query, totalSearchTime)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(searchResult)
	}
}
