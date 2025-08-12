package lucene

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/cardinalhq/cardinal-ast/core"
	luceneparser "github.com/cardinalhq/cardinal-ast/parsers/lucene/grammar"
)

// Lucene AST to Cardinal AST
type LuceneToCardinalMapper struct {
	*luceneparser.BaseLuceneQueryVisitor
}

func NewLuceneToCardinalMapper() *LuceneToCardinalMapper {
	return &LuceneToCardinalMapper{}
}

func (m *LuceneToCardinalMapper) MapQueryToAPIFormat(queryCtx luceneparser.IQueryContext, limit *int, order *string) map[string]any {
	baseExpr := m.MapQuery(queryCtx)
	if baseExpr == nil {
		return nil
	}

	if limit != nil {
		baseExpr.Limit = limit
	} else if baseExpr.Limit == nil {
		defaultLimit := 1000
		baseExpr.Limit = &defaultLimit
	}

	if order != nil {
		baseExpr.Order = order
	} else if baseExpr.Order == nil {
		defaultOrder := "DESC"
		baseExpr.Order = &defaultOrder
	}

	result := map[string]any{
		"baseExpressions": map[string]any{
			"a": baseExpr.ToJsonObj(),
		},
	}

	return result
}

// Lucene query to Cardinal BaseExpr
func (m *LuceneToCardinalMapper) MapQuery(queryCtx luceneparser.IQueryContext) *core.BaseExpr {
	if queryCtx == nil {
		return nil
	}

	// Get the expression from the query
	exprCtx := queryCtx.Expression()
	if exprCtx == nil {
		return nil
	}

	// Map the expression to a filter
	filter := m.MapExpression(exprCtx)
	if filter == nil {
		return nil
	}

	// Create a BaseExpr with default values
	baseExpr := &core.BaseExpr{
		ID:            "lucene_query",
		Dataset:       core.DATASET_LOGS,
		Filter:        filter,
		ReturnResults: true,
	}

	return baseExpr
}

// Lucene expression to Cardinal QueryClause
func (m *LuceneToCardinalMapper) MapExpression(exprCtx luceneparser.IExpressionContext) core.QueryClause {
	if exprCtx == nil {
		return nil
	}

	// Check if it is a binary expression (AND/OR)
	if len(exprCtx.AllExpression()) == 2 {
		leftExpr := m.MapExpression(exprCtx.Expression(0))
		rightExpr := m.MapExpression(exprCtx.Expression(1))

		if leftExpr != nil && rightExpr != nil {
			op := "and"
			if exprCtx.OR() != nil {
				op = "or"
			}

			return &core.BinaryClause{
				Q1: leftExpr,
				Q2: rightExpr,
				Op: op,
			}
		}
	}

	// k:v pair
	if exprCtx.FieldExpression() != nil {
		return m.MapFieldExpression(exprCtx.FieldExpression())
	}

	// check for term
	if exprCtx.MessageTerm() != nil {
		return m.MapMessageTerm(exprCtx.MessageTerm())
	}

	return nil
}

// Lucene expression to Cardinal Filter
func (m *LuceneToCardinalMapper) MapFieldExpression(fieldExprCtx luceneparser.IFieldExpressionContext) core.QueryClause {
	if fieldExprCtx == nil {
		return nil
	}

	// Get field and value
	fieldCtx := fieldExprCtx.Field()
	valueCtx := fieldExprCtx.Value()

	if fieldCtx == nil || valueCtx == nil {
		return nil
	}

	// Build field name from identifiers
	fieldName := m.BuildFieldName(fieldCtx)
	if fieldName == "" {
		return nil
	}

	value := m.GetValue(valueCtx)
	if value == "" {
		return nil
	}

	filter := &core.Filter{
		K:         fieldName,
		V:         []string{value},
		Op:        core.OP_EQ,
		Extracted: false,
		Computed:  false,
		DataType:  core.DATA_TYPE_STRING,
	}

	return filter
}

// MapMessageTerm maps a Lucene message term to a Cardinal Filter
func (m *LuceneToCardinalMapper) MapMessageTerm(msgTermCtx luceneparser.IMessageTermContext) core.QueryClause {
	if msgTermCtx == nil {
		return nil
	}

	var value string

	// Check if it's a phrase or word
	if msgTermCtx.PHRASE() != nil {
		value = msgTermCtx.PHRASE().GetText()
		// Remove quotes
		if len(value) >= 2 {
			value = value[1 : len(value)-1]
		}
	} else if msgTermCtx.WORD() != nil {
		value = msgTermCtx.WORD().GetText()
	}

	if value == "" {
		return nil
	}

	filter := &core.Filter{
		K:         "_cardinalhq.message",
		V:         []string{value},
		Op:        core.OP_CONTAINS,
		Extracted: false,
		Computed:  false,
		DataType:  core.DATA_TYPE_STRING,
	}

	return filter
}

// BuildFieldName builds a field name from the field context
func (m *LuceneToCardinalMapper) BuildFieldName(fieldCtx luceneparser.IFieldContext) string {
	if fieldCtx == nil {
		return ""
	}

	identifiers := fieldCtx.AllIDENTIFIER()
	if len(identifiers) == 0 {
		return ""
	}

	// Build field name with dots
	fieldName := identifiers[0].GetText()
	for i := 1; i < len(identifiers); i++ {
		fieldName += "." + identifiers[i].GetText()
	}

	return m.MapFieldName(fieldName)
}

func (m *LuceneToCardinalMapper) MapFieldName(fieldName string) string {
	fieldMappings := map[string]string{
		"service":   "resource.service.name",
		"level":     "_cardinalhq.level",
		"message":   "_cardinalhq.message",
		"app":       "resource.service.name",
		"pod":       "resource.k8s.pod.name",
		"namespace": "resource.k8s.namespace.name",
		"container": "resource.k8s.container.name",
		"node":      "resource.k8s.node.name",
		"cluster":   "resource.k8s.cluster.name",
	}

	if mapped, exists := fieldMappings[fieldName]; exists {
		return mapped
	}

	return fieldName
}

// GetValue extracts the value from the value context
func (m *LuceneToCardinalMapper) GetValue(valueCtx luceneparser.IValueContext) string {
	if valueCtx == nil {
		return ""
	}

	var value string

	// Check if it's a phrase or word
	if valueCtx.PHRASE() != nil {
		value = valueCtx.PHRASE().GetText()
		// Remove quotes
		if len(value) >= 2 {
			value = value[1 : len(value)-1]
		}
	} else if valueCtx.WORD() != nil {
		value = valueCtx.WORD().GetText()
	}

	return value
}

// Lucene query string to Cardinal AST
func MapLuceneQueryToCardinal(query string) (*core.BaseExpr, error) {
	input := antlr.NewInputStream(query)

	lexer := luceneparser.NewLuceneQueryLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	parser := luceneparser.NewLuceneQueryParser(stream)

	tree := parser.Query()

	mapper := NewLuceneToCardinalMapper()
	result := mapper.MapQuery(tree)

	return result, nil
}
