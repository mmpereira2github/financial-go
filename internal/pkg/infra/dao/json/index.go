package json

import (
	"encoding/json"
	"financial/internal/pkg/config"
	"financial/internal/pkg/index"
	"financial/internal/pkg/infra/dao"
	"log"
	"os"
)

var indexMap map[string]*index.Index

// JSONFileIndexDao IndexDao implementation whose repo is a file with a JSON  with the whole index and values
type JSONFileIndexDao struct{}

// NewJSONFileIndexDao creates a new JSONFileIndexDao instance
func NewJSONFileIndexDao() dao.IndexDao { return &JSONFileIndexDao{} }

func loadIndex(id string) (*index.Index, error) {
	indexFile := config.AppHome + "/configs/" + id + ".json"
	f, err := os.Open(indexFile)
	if err != nil {
		log.Printf("No file for loading index at %s\n", indexFile)
		return nil, err
	}
	defer f.Close()
	var indexInstance = index.Index{}

	log.Printf("Loading index %s", indexFile)
	dec := json.NewDecoder(f)
	if err := dec.Decode(&indexInstance); err != nil {
		log.Printf("Failed to load the  index at %s\n, error=(%s)", indexFile, err)
		return nil, err
	}
	return &indexInstance, nil
}

// FindByID finds a JSON file with name
func (*JSONFileIndexDao) FindByID(id string) (*index.Index, error) {
	var err error
	var indexInstance *index.Index
	if indexMap == nil {
		indexInstance, err = loadIndex(id)
		if err != nil {
			return nil, err
		}
		indexMap = make(map[string]*index.Index, 1)
		indexMap[id] = indexInstance
	} else {
		indexInstance = indexMap[id]
		if indexInstance == nil {
			indexInstance, err = loadIndex(id)
			if err != nil {
				return nil, err
			}
			indexMap[id] = indexInstance
		}
	}
	return indexInstance, nil
}
