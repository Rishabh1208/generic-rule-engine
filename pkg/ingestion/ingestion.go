package ingestion

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	// etc.
)

// DataIngestion is the interface that all ingestion methods must implement.
type DataIngestion interface {
	FetchData() ([]map[string]interface{}, error)
}

// FileIngestion reads data from a local JSON file (example).
type FileIngestion struct {
	FilePath string
}

// FetchData implements the DataIngestion interface.
func (f *FileIngestion) FetchData() ([]map[string]interface{}, error) {
	fmt.Println(f.FilePath)
	file, err := os.Open(f.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var data []map[string]interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	return data, nil
}
