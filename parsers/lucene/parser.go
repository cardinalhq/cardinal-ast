package lucene

import (
	"encoding/json"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	luceneparser "github.com/cardinalhq/cardinal-ast/parsers/lucene/grammar"
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

func ConvertLuceneToAPIFormat(luceneQuery string, limit *int, order *string) (map[string]any, error) {
	input := antlr.NewInputStream(luceneQuery)

	lexer := luceneparser.NewLuceneQueryLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	parser := luceneparser.NewLuceneQueryParser(stream)

	tree := parser.Query()

	mapper := NewLuceneToCardinalMapper()
	result := mapper.MapQueryToAPIFormat(tree, limit, order)

	if result == nil {
		return nil, fmt.Errorf("failed to parse query: no result generated")
	}

	return result, nil
}
