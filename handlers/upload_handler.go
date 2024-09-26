package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-tika/tika"
)

// var index bleve.Index

func UploadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error retrieving the file", http.StatusBadRequest)
			return
		}
		defer file.Close()
		fmt.Println("filenameee", handler.Filename)
		bucketName := r.FormValue("bucketName")
		// objName := r.FormValue("objName")
		fmt.Println("bucketName", bucketName)
		client := tika.NewClient(nil, "http://localhost:9998")
		body, err := client.ParseRecursive(context.Background(), file)
		fmt.Printf("text extracted %+v \n", body)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Files indexed successfully"))
	}
}
