package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"bleve.search/handlers"
	"bleve.search/model"
	"bleve.search/utility"
	"github.com/blevesearch/bleve/v2"
	"github.com/bxcodec/faker/v3"
)

func generateAndIndex(index bleve.Index, totalFiles int) {
	var fileInfos []model.FileInfo
	fileNames := []string{}
	for i := 0; i < totalFiles; i++ {
		fileName := faker.Word() + ".txt"
		fileNames = append(fileNames, fileName)
		fileInfo := model.FileInfo{
			Filename: fileName,
			Path:     "/" + faker.Word() + "/" + faker.Word() + "/" + fileName,
		}
		fileInfos = append(fileInfos, fileInfo)
	}

	jsonData, err := json.MarshalIndent(fileNames, "", "  ")
	if err != nil {
		fmt.Println("error marshaling")
	}
	err = os.WriteFile("scripts/file_names.json", jsonData, 0644)
	if err != nil {
		fmt.Println("error writing in json file")
	}
	err = utility.IndexFiles(index, fileInfos)
	if err != nil {
		fmt.Println("error indexing")
	}
}

func sizeOfIndex(path string) (int64, error) {
	var size int64

	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return size, nil
}

func main() {
	index, err := utility.OpenOrCreateIndex("files_index.bleve")
	if err != nil {
		log.Fatal(err)
	}
	defer index.Close()
	totalFiles := 100000
	generateAndIndex(index, totalFiles)
	size, err := sizeOfIndex("files_index.bleve")
	if err != nil {
		fmt.Println("Error in calculating size", err)
	}
	fmt.Printf("Size of index for %d files is %d MB \n", totalFiles, size/(1024*1024))

	// http.HandleFunc("/index", handlers.IndexHandler(index))
	http.HandleFunc("/search", handlers.SearchHandler(index))

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
