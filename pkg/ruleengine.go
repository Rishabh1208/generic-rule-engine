package pkg

import (
	"github.com/Rishabh1208/generic-rule-engine/pkg/engine"
	"github.com/Rishabh1208/generic-rule-engine/pkg/ingestion"
	"github.com/Rishabh1208/generic-rule-engine/pkg/rules"
)

// Process is an optional helper that performs the entire flow in one go.
func Process(dataFilePath, rulesFilePath string) ([]map[string]interface{}, error) {
	// Ingestion
	fi := ingestion.FileIngestion{FilePath: dataFilePath}
	dataRecords, err := fi.FetchData()
	if err != nil {
		return nil, err
	}

	// Rules
	rr := rules.FileRuleRepository{FilePath: rulesFilePath}
	loadedRules, err := rr.FetchRules()
	if err != nil {
		return nil, err
	}

	// Engine
	re := engine.GovaluateRuleEngine{}

	// Evaluate each record
	var results []map[string]interface{}
	for _, record := range dataRecords {
		updated, _ := re.Evaluate(record, loadedRules)
		results = append(results, updated)
	}

	return results, nil
}
