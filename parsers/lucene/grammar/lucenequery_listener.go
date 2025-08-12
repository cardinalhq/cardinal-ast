// Code generated from LuceneQuery.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // LuceneQuery
import "github.com/antlr4-go/antlr/v4"

// LuceneQueryListener is a complete listener for a parse tree produced by LuceneQueryParser.
type LuceneQueryListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFieldExpression is called when entering the fieldExpression production.
	EnterFieldExpression(c *FieldExpressionContext)

	// EnterMessageTerm is called when entering the messageTerm production.
	EnterMessageTerm(c *MessageTermContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFieldExpression is called when exiting the fieldExpression production.
	ExitFieldExpression(c *FieldExpressionContext)

	// ExitMessageTerm is called when exiting the messageTerm production.
	ExitMessageTerm(c *MessageTermContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
