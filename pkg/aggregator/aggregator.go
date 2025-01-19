package aggregator

import "fmt"

// Example aggregator function
func MergeDataByKey(
	primary []map[string]interface{},
	secondary []map[string]interface{},
	key string,
) ([]map[string]interface{}, error) {
	// Implementation that merges two slices of maps on `key`
	// ...
	fmt.Println("Merging data on key:", key)
	return nil, nil
}
