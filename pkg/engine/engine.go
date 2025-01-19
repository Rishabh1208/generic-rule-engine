package engine

import (
	"log"

	"github.com/Knetic/govaluate"
	"github.com/Rishabh1208/generic-rule-engine/pkg/rules"
)

// RuleEngine is the interface for evaluating data against a set of rules.
type RuleEngine interface {
	Evaluate(data map[string]interface{}, allRules []rules.Rule) (map[string]interface{}, error)
}

// GovaluateRuleEngine is a concrete implementation using govaluate.
type GovaluateRuleEngine struct{}

func (gre *GovaluateRuleEngine) Evaluate(
	data map[string]interface{},
	allRules []rules.Rule,
) (map[string]interface{}, error) {
	// default
	data["criticality"] = "UNKNOWN"

	for _, r := range allRules {
		expr, err := govaluate.NewEvaluableExpression(r.Condition)
		if err != nil {
			log.Printf("Error parsing expression '%s': %v", r.Condition, err)
			continue
		}
		result, err := expr.Evaluate(data)
		if err != nil {
			log.Printf("Error evaluating expression '%s': %v", r.Condition, err)
			continue
		}

		boolRes, ok := result.(bool)
		if ok && boolRes {
			data["criticality"] = r.Criticality
			break // stop at first match, or remove if you want multiple matches
		}
	}

	return data, nil
}
