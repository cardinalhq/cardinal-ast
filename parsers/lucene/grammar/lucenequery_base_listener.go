// Code generated from LuceneQuery.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // LuceneQuery
import "github.com/antlr4-go/antlr/v4"

// BaseLuceneQueryListener is a complete listener for a parse tree produced by LuceneQueryParser.
type BaseLuceneQueryListener struct{}

var _ LuceneQueryListener = &BaseLuceneQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLuceneQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLuceneQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLuceneQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLuceneQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseLuceneQueryListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseLuceneQueryListener) ExitQuery(ctx *QueryContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLuceneQueryListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLuceneQueryListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFieldExpression is called when production fieldExpression is entered.
func (s *BaseLuceneQueryListener) EnterFieldExpression(ctx *FieldExpressionContext) {}

// ExitFieldExpression is called when production fieldExpression is exited.
func (s *BaseLuceneQueryListener) ExitFieldExpression(ctx *FieldExpressionContext) {}

// EnterMessageTerm is called when production messageTerm is entered.
func (s *BaseLuceneQueryListener) EnterMessageTerm(ctx *MessageTermContext) {}

// ExitMessageTerm is called when production messageTerm is exited.
func (s *BaseLuceneQueryListener) ExitMessageTerm(ctx *MessageTermContext) {}

// EnterField is called when production field is entered.
func (s *BaseLuceneQueryListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseLuceneQueryListener) ExitField(ctx *FieldContext) {}

// EnterValue is called when production value is entered.
func (s *BaseLuceneQueryListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseLuceneQueryListener) ExitValue(ctx *ValueContext) {}
