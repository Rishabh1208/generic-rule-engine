package rules

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Rule holds an expression and metadata (for example, “criticality”).
type Rule struct {
	ID          string `json:"id,omitempty"`
	Condition   string `json:"condition"`
	Criticality string `json:"criticality,omitempty"`
	// Add more fields if needed (e.g. Priority, etc.)
}

// RuleRepository is the interface for loading rules from somewhere.
type RuleRepository interface {
	FetchRules() ([]Rule, error)
}

// FileRuleRepository loads rules from a JSON file.
type FileRuleRepository struct {
	FilePath string
}

func (fr *FileRuleRepository) FetchRules() ([]Rule, error) {
	file, err := os.Open(fr.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var data struct {
		Rules []Rule `json:"rules"`
	}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	return data.Rules, nil
}
