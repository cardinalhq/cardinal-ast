package lucene

import (
	"encoding/json"
	"fmt"
)

func ConvertLuceneToJSONDSL(luceneQuery string) (string, error) {
	baseExpr, err := MapLuceneQueryToCardinal(luceneQuery)
	if err != nil {
		return "", fmt.Errorf("failed to parse Lucene query: %w", err)
	}

	if baseExpr == nil {
		return "", fmt.Errorf("failed to parse query: no result generated")
	}

	jsonObj := baseExpr.ToJsonObj()

	jsonBytes, err := json.Marshal(jsonObj)
	if err != nil {
		return "", fmt.Errorf("failed to marshal to JSON: %w", err)
	}

	return string(jsonBytes), nil
}

// Lucene query to JSON DSL map
func ConvertLuceneToJSONDSLMap(luceneQuery string) (map[string]interface{}, error) {
	baseExpr, err := MapLuceneQueryToCardinal(luceneQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Lucene query: %w", err)
	}

	if baseExpr == nil {
		return nil, fmt.Errorf("failed to parse query: no result generated")
	}

	jsonObj := baseExpr.ToJsonObj()
	return jsonObj, nil
}
