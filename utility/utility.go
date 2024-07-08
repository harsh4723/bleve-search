package utility

import (
	"bleve.search/model"
	"github.com/blevesearch/bleve/v2"
)

func OpenOrCreateIndex(indexPath string) (bleve.Index, error) {
	index, err := bleve.Open(indexPath)
	if err == bleve.ErrorIndexPathDoesNotExist {
		indexMapping := bleve.NewIndexMapping()
		index, err = bleve.New(indexPath, indexMapping)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	return index, nil
}

func IndexFiles(index bleve.Index, files []model.FileInfo) error {
	for _, file := range files {
		err := index.Index(file.Path, file) // Use filename as ID
		if err != nil {
			return err
		}
	}
	return nil
}
