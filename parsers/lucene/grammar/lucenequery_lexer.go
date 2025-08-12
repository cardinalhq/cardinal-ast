// Code generated from LuceneQuery.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type LuceneQueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var LuceneQueryLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lucenequerylexerLexerInit() {
	staticData := &LuceneQueryLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'", "':'", "'.'", "'AND'", "'OR'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "AND", "OR", "IDENTIFIER", "PHRASE", "WORD", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "AND", "OR", "IDENTIFIER", "PHRASE",
		"WORD", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 10, 66, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 5, 6, 39, 8, 6, 10, 6, 12, 6, 42, 9, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 5, 7, 48, 8, 7, 10, 7, 12, 7, 51, 9, 7, 1, 7, 1, 7, 1,
		8, 4, 8, 56, 8, 8, 11, 8, 12, 8, 57, 1, 9, 4, 9, 61, 8, 9, 11, 9, 12, 9,
		62, 1, 9, 1, 9, 0, 0, 10, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 1, 0, 4, 3, 0, 65, 90, 95, 95, 97, 122, 5, 0, 45, 45,
		48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 34, 34, 92, 92, 3, 0, 9, 10, 13,
		13, 32, 32, 70, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0,
		7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0,
		0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 1, 21, 1, 0, 0,
		0, 3, 23, 1, 0, 0, 0, 5, 25, 1, 0, 0, 0, 7, 27, 1, 0, 0, 0, 9, 29, 1, 0,
		0, 0, 11, 33, 1, 0, 0, 0, 13, 36, 1, 0, 0, 0, 15, 43, 1, 0, 0, 0, 17, 55,
		1, 0, 0, 0, 19, 60, 1, 0, 0, 0, 21, 22, 5, 40, 0, 0, 22, 2, 1, 0, 0, 0,
		23, 24, 5, 41, 0, 0, 24, 4, 1, 0, 0, 0, 25, 26, 5, 58, 0, 0, 26, 6, 1,
		0, 0, 0, 27, 28, 5, 46, 0, 0, 28, 8, 1, 0, 0, 0, 29, 30, 5, 65, 0, 0, 30,
		31, 5, 78, 0, 0, 31, 32, 5, 68, 0, 0, 32, 10, 1, 0, 0, 0, 33, 34, 5, 79,
		0, 0, 34, 35, 5, 82, 0, 0, 35, 12, 1, 0, 0, 0, 36, 40, 7, 0, 0, 0, 37,
		39, 7, 1, 0, 0, 38, 37, 1, 0, 0, 0, 39, 42, 1, 0, 0, 0, 40, 38, 1, 0, 0,
		0, 40, 41, 1, 0, 0, 0, 41, 14, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 43, 49,
		5, 34, 0, 0, 44, 48, 8, 2, 0, 0, 45, 46, 5, 92, 0, 0, 46, 48, 9, 0, 0,
		0, 47, 44, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 48, 51, 1, 0, 0, 0, 49, 47,
		1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 52, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0,
		52, 53, 5, 34, 0, 0, 53, 16, 1, 0, 0, 0, 54, 56, 7, 1, 0, 0, 55, 54, 1,
		0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58,
		18, 1, 0, 0, 0, 59, 61, 7, 3, 0, 0, 60, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0,
		0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 65,
		6, 9, 0, 0, 65, 20, 1, 0, 0, 0, 6, 0, 40, 47, 49, 57, 62, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// LuceneQueryLexerInit initializes any static state used to implement LuceneQueryLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLuceneQueryLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LuceneQueryLexerInit() {
	staticData := &LuceneQueryLexerLexerStaticData
	staticData.once.Do(lucenequerylexerLexerInit)
}

// NewLuceneQueryLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLuceneQueryLexer(input antlr.CharStream) *LuceneQueryLexer {
	LuceneQueryLexerInit()
	l := new(LuceneQueryLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &LuceneQueryLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "LuceneQuery.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LuceneQueryLexer tokens.
const (
	LuceneQueryLexerT__0       = 1
	LuceneQueryLexerT__1       = 2
	LuceneQueryLexerT__2       = 3
	LuceneQueryLexerT__3       = 4
	LuceneQueryLexerAND        = 5
	LuceneQueryLexerOR         = 6
	LuceneQueryLexerIDENTIFIER = 7
	LuceneQueryLexerPHRASE     = 8
	LuceneQueryLexerWORD       = 9
	LuceneQueryLexerWS         = 10
)
