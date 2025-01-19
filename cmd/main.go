package main

import (
	"fmt"
	"log"

	"github.com/Rishabh1208/generic-rule-engine/pkg"
)

func main() {
	results, err := pkg.Process("firms.json", "rules.json")
	if err != nil {
		log.Fatalf("Error processing: %v", err)
	}

	for _, r := range results {
		fmt.Printf("Result: %+v\n", r)
	}
}
