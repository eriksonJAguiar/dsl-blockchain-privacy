// Code generated from Translator.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 90, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 10, 6, 10, 63, 10, 10, 13, 10, 14, 10, 64, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 6, 11, 73, 10, 11, 13, 11, 14, 11, 74, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 7, 12, 82, 10, 12, 12, 12, 14, 12, 85, 11, 12, 3, 12, 3,
	12, 3, 13, 3, 13, 2, 2, 14, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9,
	17, 10, 19, 11, 21, 12, 23, 13, 25, 2, 3, 2, 4, 5, 2, 11, 12, 15, 15, 34,
	34, 5, 2, 50, 59, 67, 92, 99, 124, 2, 95, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21,
	3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 3, 27, 3, 2, 2, 2, 5, 34, 3, 2, 2, 2, 7,
	39, 3, 2, 2, 2, 9, 45, 3, 2, 2, 2, 11, 47, 3, 2, 2, 2, 13, 49, 3, 2, 2,
	2, 15, 53, 3, 2, 2, 2, 17, 59, 3, 2, 2, 2, 19, 62, 3, 2, 2, 2, 21, 72,
	3, 2, 2, 2, 23, 76, 3, 2, 2, 2, 25, 88, 3, 2, 2, 2, 27, 28, 7, 85, 2, 2,
	28, 29, 7, 71, 2, 2, 29, 30, 7, 78, 2, 2, 30, 31, 7, 71, 2, 2, 31, 32,
	7, 69, 2, 2, 32, 33, 7, 86, 2, 2, 33, 4, 3, 2, 2, 2, 34, 35, 7, 72, 2,
	2, 35, 36, 7, 84, 2, 2, 36, 37, 7, 81, 2, 2, 37, 38, 7, 79, 2, 2, 38, 6,
	3, 2, 2, 2, 39, 40, 7, 89, 2, 2, 40, 41, 7, 74, 2, 2, 41, 42, 7, 71, 2,
	2, 42, 43, 7, 84, 2, 2, 43, 44, 7, 71, 2, 2, 44, 8, 3, 2, 2, 2, 45, 46,
	7, 42, 2, 2, 46, 10, 3, 2, 2, 2, 47, 48, 7, 43, 2, 2, 48, 12, 3, 2, 2,
	2, 49, 50, 7, 67, 2, 2, 50, 51, 7, 88, 2, 2, 51, 52, 7, 73, 2, 2, 52, 14,
	3, 2, 2, 2, 53, 54, 7, 69, 2, 2, 54, 55, 7, 81, 2, 2, 55, 56, 7, 87, 2,
	2, 56, 57, 7, 80, 2, 2, 57, 58, 7, 86, 2, 2, 58, 16, 3, 2, 2, 2, 59, 60,
	7, 63, 2, 2, 60, 18, 3, 2, 2, 2, 61, 63, 9, 2, 2, 2, 62, 61, 3, 2, 2, 2,
	63, 64, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 66, 3,
	2, 2, 2, 66, 67, 8, 10, 2, 2, 67, 20, 3, 2, 2, 2, 68, 73, 5, 25, 13, 2,
	69, 73, 7, 97, 2, 2, 70, 71, 7, 94, 2, 2, 71, 73, 7, 36, 2, 2, 72, 68,
	3, 2, 2, 2, 72, 69, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2,
	74, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 22, 3, 2, 2, 2, 76, 83, 7,
	36, 2, 2, 77, 82, 5, 25, 13, 2, 78, 82, 7, 97, 2, 2, 79, 80, 7, 94, 2,
	2, 80, 82, 7, 36, 2, 2, 81, 77, 3, 2, 2, 2, 81, 78, 3, 2, 2, 2, 81, 79,
	3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2,
	84, 86, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 87, 7, 36, 2, 2, 87, 24, 3,
	2, 2, 2, 88, 89, 9, 3, 2, 2, 89, 26, 3, 2, 2, 2, 8, 2, 64, 72, 74, 81,
	83, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'SELECT'", "'FROM'", "'WHERE'", "'('", "')'", "'AVG'", "'COUNT'",
	"'='",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "WHITESPACE", "IDENTIFIER", "STRING_LITERAL",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "WHITESPACE",
	"IDENTIFIER", "STRING_LITERAL", "NUM_LETTER",
}

type TranslatorLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewTranslatorLexer(input antlr.CharStream) *TranslatorLexer {

	l := new(TranslatorLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Translator.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TranslatorLexer tokens.
const (
	TranslatorLexerT__0           = 1
	TranslatorLexerT__1           = 2
	TranslatorLexerT__2           = 3
	TranslatorLexerT__3           = 4
	TranslatorLexerT__4           = 5
	TranslatorLexerT__5           = 6
	TranslatorLexerT__6           = 7
	TranslatorLexerT__7           = 8
	TranslatorLexerWHITESPACE     = 9
	TranslatorLexerIDENTIFIER     = 10
	TranslatorLexerSTRING_LITERAL = 11
)
