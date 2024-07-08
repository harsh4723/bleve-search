package handlers

import (
	"encoding/json"
	"net/http"

	"bleve.search/model"
	"bleve.search/utility"
	"github.com/blevesearch/bleve/v2"
)

// var index bleve.Index

func IndexHandler(index bleve.Index) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var files []model.FileInfo
		err := json.NewDecoder(r.Body).Decode(&files)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = utility.IndexFiles(index, files)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Files indexed successfully"))
	}
}
