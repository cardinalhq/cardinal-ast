// Code generated from LuceneQuery.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // LuceneQuery
import "github.com/antlr4-go/antlr/v4"

type BaseLuceneQueryVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLuceneQueryVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuceneQueryVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuceneQueryVisitor) VisitFieldExpression(ctx *FieldExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuceneQueryVisitor) VisitMessageTerm(ctx *MessageTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuceneQueryVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuceneQueryVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}
