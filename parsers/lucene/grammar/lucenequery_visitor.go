// Code generated from LuceneQuery.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // LuceneQuery
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by LuceneQueryParser.
type LuceneQueryVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LuceneQueryParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by LuceneQueryParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by LuceneQueryParser#fieldExpression.
	VisitFieldExpression(ctx *FieldExpressionContext) interface{}

	// Visit a parse tree produced by LuceneQueryParser#messageTerm.
	VisitMessageTerm(ctx *MessageTermContext) interface{}

	// Visit a parse tree produced by LuceneQueryParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by LuceneQueryParser#value.
	VisitValue(ctx *ValueContext) interface{}
}
