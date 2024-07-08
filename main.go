package main

import (
	"log"
	"net/http"

	"bleve.search/handlers"
	"bleve.search/utility"
	"github.com/blevesearch/bleve/v2"
)

var index bleve.Index

func main() {
	// indexMapping := bleve.NewIndexMapping()
	// index, err := bleve.New("files.bleve", indexMapping)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // List of files to index
	// files := []FileInfo{
	// 	{Filename: "file1.txt", Path: "/path/to/file1.txt"},
	// 	{Filename: "file2.txt", Path: "/another/path/file2.txt"},
	// 	// Add more files here
	// }

	// // Index each file
	// for _, file := range files {
	// 	err = index.Index(file.Filename, file)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// // Close the index when done
	// index.Close()

	//open a new index
	// index, err := bleve.Open("files.bleve")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Query for filenames
	// searchFilenames := []string{"file1.txt", "file2.txt", "file", "file2"} // List of filenames to search

	// for _, filename := range searchFilenames {
	// 	// Create a match query for the filename
	// 	//query := bleve.NewMatchQuery(filename)
	// 	query := bleve.NewPrefixQuery(filename)
	// 	//query.SetField("Filename")
	// 	searchRequest := bleve.NewSearchRequest(query)

	// 	// Execute the search
	// 	searchResult, err := index.Search(searchRequest)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	// Display the results
	// 	fmt.Printf("Search Results for %s:\n", filename)
	// 	fmt.Println("Searchresultes", searchResult)
	// 	// for _, hit := range searchResult.Hits {
	// 	// 	fmt.Printf("ID: %s, Score: %f, Fields: %+v\n", hit.ID, hit.Score, hit.Fields)
	// 	// }
	// }

	// // Close the index when done
	// index.Close()

	index, err := utility.OpenOrCreateIndex("files.bleve")
	if err != nil {
		log.Fatal(err)
	}
	defer index.Close()

	http.HandleFunc("/index", handlers.IndexHandler(index))
	http.HandleFunc("/search", handlers.SearchHandler(index))

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
