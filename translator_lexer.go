// Code generated from Translator.g4 by ANTLR 4.8. DO NOT EDIT.

package main

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 16, 117,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 6, 13, 90, 10, 13, 13, 13,
	14, 13, 91, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 6, 14, 100, 10, 14,
	13, 14, 14, 14, 101, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 109, 10,
	15, 12, 15, 14, 15, 112, 11, 15, 3, 15, 3, 15, 3, 16, 3, 16, 2, 2, 17,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23,
	13, 25, 14, 27, 15, 29, 16, 31, 2, 3, 2, 4, 5, 2, 11, 12, 15, 15, 34, 34,
	5, 2, 50, 59, 67, 92, 99, 124, 2, 122, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29,
	3, 2, 2, 2, 3, 33, 3, 2, 2, 2, 5, 45, 3, 2, 2, 2, 7, 47, 3, 2, 2, 2, 9,
	49, 3, 2, 2, 2, 11, 56, 3, 2, 2, 2, 13, 58, 3, 2, 2, 2, 15, 65, 3, 2, 2,
	2, 17, 70, 3, 2, 2, 2, 19, 76, 3, 2, 2, 2, 21, 80, 3, 2, 2, 2, 23, 86,
	3, 2, 2, 2, 25, 89, 3, 2, 2, 2, 27, 99, 3, 2, 2, 2, 29, 103, 3, 2, 2, 2,
	31, 115, 3, 2, 2, 2, 33, 34, 7, 75, 2, 2, 34, 35, 7, 80, 2, 2, 35, 36,
	7, 85, 2, 2, 36, 37, 7, 71, 2, 2, 37, 38, 7, 84, 2, 2, 38, 39, 7, 86, 2,
	2, 39, 40, 7, 34, 2, 2, 40, 41, 7, 75, 2, 2, 41, 42, 7, 80, 2, 2, 42, 43,
	7, 86, 2, 2, 43, 44, 7, 81, 2, 2, 44, 4, 3, 2, 2, 2, 45, 46, 7, 42, 2,
	2, 46, 6, 3, 2, 2, 2, 47, 48, 7, 43, 2, 2, 48, 8, 3, 2, 2, 2, 49, 50, 7,
	88, 2, 2, 50, 51, 7, 67, 2, 2, 51, 52, 7, 78, 2, 2, 52, 53, 7, 87, 2, 2,
	53, 54, 7, 71, 2, 2, 54, 55, 7, 85, 2, 2, 55, 10, 3, 2, 2, 2, 56, 57, 7,
	46, 2, 2, 57, 12, 3, 2, 2, 2, 58, 59, 7, 85, 2, 2, 59, 60, 7, 71, 2, 2,
	60, 61, 7, 78, 2, 2, 61, 62, 7, 71, 2, 2, 62, 63, 7, 69, 2, 2, 63, 64,
	7, 86, 2, 2, 64, 14, 3, 2, 2, 2, 65, 66, 7, 72, 2, 2, 66, 67, 7, 84, 2,
	2, 67, 68, 7, 81, 2, 2, 68, 69, 7, 79, 2, 2, 69, 16, 3, 2, 2, 2, 70, 71,
	7, 89, 2, 2, 71, 72, 7, 74, 2, 2, 72, 73, 7, 71, 2, 2, 73, 74, 7, 84, 2,
	2, 74, 75, 7, 71, 2, 2, 75, 18, 3, 2, 2, 2, 76, 77, 7, 67, 2, 2, 77, 78,
	7, 88, 2, 2, 78, 79, 7, 73, 2, 2, 79, 20, 3, 2, 2, 2, 80, 81, 7, 69, 2,
	2, 81, 82, 7, 81, 2, 2, 82, 83, 7, 87, 2, 2, 83, 84, 7, 80, 2, 2, 84, 85,
	7, 86, 2, 2, 85, 22, 3, 2, 2, 2, 86, 87, 7, 63, 2, 2, 87, 24, 3, 2, 2,
	2, 88, 90, 9, 2, 2, 2, 89, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 89,
	3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 94, 8, 13, 2, 2,
	94, 26, 3, 2, 2, 2, 95, 100, 5, 31, 16, 2, 96, 100, 7, 97, 2, 2, 97, 98,
	7, 94, 2, 2, 98, 100, 7, 36, 2, 2, 99, 95, 3, 2, 2, 2, 99, 96, 3, 2, 2,
	2, 99, 97, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101,
	102, 3, 2, 2, 2, 102, 28, 3, 2, 2, 2, 103, 110, 7, 36, 2, 2, 104, 109,
	5, 31, 16, 2, 105, 109, 7, 97, 2, 2, 106, 107, 7, 94, 2, 2, 107, 109, 7,
	36, 2, 2, 108, 104, 3, 2, 2, 2, 108, 105, 3, 2, 2, 2, 108, 106, 3, 2, 2,
	2, 109, 112, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111,
	113, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 113, 114, 7, 36, 2, 2, 114, 30,
	3, 2, 2, 2, 115, 116, 9, 3, 2, 2, 116, 32, 3, 2, 2, 2, 8, 2, 91, 99, 101,
	108, 110, 3, 8, 2, 2,
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
	"", "'INSERT INTO'", "'('", "')'", "'VALUES'", "','", "'SELECT'", "'FROM'",
	"'WHERE'", "'AVG'", "'COUNT'", "'='",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "WHITESPACE", "IDENTIFIER",
	"STRING_LITERAL",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "WHITESPACE", "IDENTIFIER", "STRING_LITERAL", "NUM_LETTER",
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
	TranslatorLexerT__8           = 9
	TranslatorLexerT__9           = 10
	TranslatorLexerT__10          = 11
	TranslatorLexerWHITESPACE     = 12
	TranslatorLexerIDENTIFIER     = 13
	TranslatorLexerSTRING_LITERAL = 14
)
