// Code generated from Translator.g4 by ANTLR 4.8. DO NOT EDIT.

package main // Translator
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 18, 113,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 41, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 56, 10, 4, 12,
	4, 14, 4, 59, 11, 4, 3, 5, 3, 5, 3, 5, 7, 5, 64, 10, 5, 12, 5, 14, 5, 67,
	11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 73, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 7, 8, 85, 10, 8, 12, 8, 14, 8, 88,
	11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3,
	16, 3, 17, 3, 17, 3, 17, 2, 2, 18, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 2, 3, 3, 2, 13, 14, 2, 101, 2, 40, 3, 2, 2, 2,
	4, 42, 3, 2, 2, 2, 6, 52, 3, 2, 2, 2, 8, 60, 3, 2, 2, 2, 10, 68, 3, 2,
	2, 2, 12, 74, 3, 2, 2, 2, 14, 81, 3, 2, 2, 2, 16, 89, 3, 2, 2, 2, 18, 94,
	3, 2, 2, 2, 20, 96, 3, 2, 2, 2, 22, 98, 3, 2, 2, 2, 24, 102, 3, 2, 2, 2,
	26, 104, 3, 2, 2, 2, 28, 106, 3, 2, 2, 2, 30, 108, 3, 2, 2, 2, 32, 110,
	3, 2, 2, 2, 34, 35, 5, 10, 6, 2, 35, 36, 7, 2, 2, 3, 36, 41, 3, 2, 2, 2,
	37, 38, 5, 4, 3, 2, 38, 39, 7, 2, 2, 3, 39, 41, 3, 2, 2, 2, 40, 34, 3,
	2, 2, 2, 40, 37, 3, 2, 2, 2, 41, 3, 3, 2, 2, 2, 42, 43, 7, 3, 2, 2, 43,
	44, 5, 28, 15, 2, 44, 45, 7, 4, 2, 2, 45, 46, 5, 6, 4, 2, 46, 47, 7, 5,
	2, 2, 47, 48, 7, 6, 2, 2, 48, 49, 7, 4, 2, 2, 49, 50, 5, 8, 5, 2, 50, 51,
	7, 5, 2, 2, 51, 5, 3, 2, 2, 2, 52, 57, 5, 24, 13, 2, 53, 54, 7, 7, 2, 2,
	54, 56, 5, 24, 13, 2, 55, 53, 3, 2, 2, 2, 56, 59, 3, 2, 2, 2, 57, 55, 3,
	2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 7, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 60,
	65, 5, 26, 14, 2, 61, 62, 7, 7, 2, 2, 62, 64, 5, 26, 14, 2, 63, 61, 3,
	2, 2, 2, 64, 67, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66,
	9, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 68, 72, 5, 12, 7, 2, 69, 70, 7, 8, 2,
	2, 70, 71, 7, 9, 2, 2, 71, 73, 7, 18, 2, 2, 72, 69, 3, 2, 2, 2, 72, 73,
	3, 2, 2, 2, 73, 11, 3, 2, 2, 2, 74, 75, 7, 10, 2, 2, 75, 76, 5, 14, 8,
	2, 76, 77, 7, 11, 2, 2, 77, 78, 5, 20, 11, 2, 78, 79, 7, 12, 2, 2, 79,
	80, 5, 22, 12, 2, 80, 13, 3, 2, 2, 2, 81, 86, 5, 16, 9, 2, 82, 83, 7, 7,
	2, 2, 83, 85, 5, 16, 9, 2, 84, 82, 3, 2, 2, 2, 85, 88, 3, 2, 2, 2, 86,
	84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 15, 3, 2, 2, 2, 88, 86, 3, 2, 2,
	2, 89, 90, 5, 18, 10, 2, 90, 91, 7, 4, 2, 2, 91, 92, 5, 30, 16, 2, 92,
	93, 7, 5, 2, 2, 93, 17, 3, 2, 2, 2, 94, 95, 9, 2, 2, 2, 95, 19, 3, 2, 2,
	2, 96, 97, 5, 32, 17, 2, 97, 21, 3, 2, 2, 2, 98, 99, 5, 30, 16, 2, 99,
	100, 7, 9, 2, 2, 100, 101, 7, 17, 2, 2, 101, 23, 3, 2, 2, 2, 102, 103,
	7, 16, 2, 2, 103, 25, 3, 2, 2, 2, 104, 105, 7, 17, 2, 2, 105, 27, 3, 2,
	2, 2, 106, 107, 7, 16, 2, 2, 107, 29, 3, 2, 2, 2, 108, 109, 7, 16, 2, 2,
	109, 31, 3, 2, 2, 2, 110, 111, 7, 16, 2, 2, 111, 33, 3, 2, 2, 2, 7, 40,
	57, 65, 72, 86,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'INSERT INTO'", "'('", "')'", "'VALUES'", "','", "'EPSILON'", "'='",
	"'SELECT'", "'FROM'", "'WHERE'", "'AVG'", "'HISTOGRAM'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "WHITESPACE", "IDENTIFIER",
	"STRING_LITERAL", "NUM_FLOAT",
}

var ruleNames = []string{
	"start", "insert1", "columns_list", "values_list", "select2", "select1",
	"set_list", "pair", "function", "from_list", "condition", "column", "value",
	"table_name", "attribute", "relation",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type TranslatorParser struct {
	*antlr.BaseParser
}

func NewTranslatorParser(input antlr.TokenStream) *TranslatorParser {
	this := new(TranslatorParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Translator.g4"

	return this
}

// TranslatorParser tokens.
const (
	TranslatorParserEOF            = antlr.TokenEOF
	TranslatorParserT__0           = 1
	TranslatorParserT__1           = 2
	TranslatorParserT__2           = 3
	TranslatorParserT__3           = 4
	TranslatorParserT__4           = 5
	TranslatorParserT__5           = 6
	TranslatorParserT__6           = 7
	TranslatorParserT__7           = 8
	TranslatorParserT__8           = 9
	TranslatorParserT__9           = 10
	TranslatorParserT__10          = 11
	TranslatorParserT__11          = 12
	TranslatorParserWHITESPACE     = 13
	TranslatorParserIDENTIFIER     = 14
	TranslatorParserSTRING_LITERAL = 15
	TranslatorParserNUM_FLOAT      = 16
)

// TranslatorParser rules.
const (
	TranslatorParserRULE_start        = 0
	TranslatorParserRULE_insert1      = 1
	TranslatorParserRULE_columns_list = 2
	TranslatorParserRULE_values_list  = 3
	TranslatorParserRULE_select2      = 4
	TranslatorParserRULE_select1      = 5
	TranslatorParserRULE_set_list     = 6
	TranslatorParserRULE_pair         = 7
	TranslatorParserRULE_function     = 8
	TranslatorParserRULE_from_list    = 9
	TranslatorParserRULE_condition    = 10
	TranslatorParserRULE_column       = 11
	TranslatorParserRULE_value        = 12
	TranslatorParserRULE_table_name   = 13
	TranslatorParserRULE_attribute    = 14
	TranslatorParserRULE_relation     = 15
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TranslatorParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TranslatorParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Select2() ISelect2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelect2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelect2Context)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(TranslatorParserEOF, 0)
}

func (s *StartContext) Insert1() IInsert1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInsert1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInsert1Context)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *TranslatorParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TranslatorParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(38)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TranslatorParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)
			p.Select2()
		}
		{
			p.SetState(33)
			p.Match(TranslatorParserEOF)
		}

	case TranslatorParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Insert1()
		}
		{
			p.SetState(36)
			p.Match(TranslatorParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInsert1Context is an interface to support dynamic dispatch.
type IInsert1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInsert1Context differentiates from other interfaces.
	IsInsert1Context()
}

type Insert1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsert1Context() *Insert1Context {
	var p = new(Insert1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TranslatorParserRULE_insert1
	return p
}

func (*Insert1Context) IsInsert1Context() {}

func NewInsert1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Insert1Context {
	var p = new(Insert1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TranslatorParserRULE_insert1

	return p
}

func (s *Insert1Context) GetParser() antlr.Parser { return s.parser }

func (s *Insert1Context) Table_name() ITable_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Insert1Context) Columns_list() IColumns_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumns_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumns_listContext)
}

func (s *Insert1Context) Values_list() IValues_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValues_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValues_listContext)
}

func (s *Insert1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Insert1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Insert1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.EnterInsert1(s)
	}
}

func (s *Insert1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.ExitInsert1(s)
	}
}

func (p *TranslatorParser) Insert1() (localctx IInsert1Context) {
	localctx = NewInsert1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TranslatorParserRULE_insert1)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(TranslatorParserT__0)
	}
	{
		p.SetState(41)
		p.Table_name()
	}
	{
		p.SetState(42)
		p.Match(TranslatorParserT__1)
	}
	{
		p.SetState(43)
		p.Columns_list()
	}
	{
		p.SetState(44)
		p.Match(TranslatorParserT__2)
	}
	{
		p.SetState(45)
		p.Match(TranslatorParserT__3)
	}
	{
		p.SetState(46)
		p.Match(TranslatorParserT__1)
	}
	{
		p.SetState(47)
		p.Values_list()
	}
	{
		p.SetState(48)
		p.Match(TranslatorParserT__2)
	}

	return localctx
}

// IColumns_listContext is an interface to support dynamic dispatch.
type IColumns_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumns_listContext differentiates from other interfaces.
	IsColumns_listContext()
}

type Columns_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumns_listContext() *Columns_listContext {
	var p = new(Columns_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TranslatorParserRULE_columns_list
	return p
}

func (*Columns_listContext) IsColumns_listContext() {}

func NewColumns_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Columns_listContext {
	var p = new(Columns_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TranslatorParserRULE_columns_list

	return p
}

func (s *Columns_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Columns_listContext) AllColumn() []IColumnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumnContext)(nil)).Elem())
	var tst = make([]IColumnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumnContext)
		}
	}

	return tst
}

func (s *Columns_listContext) Column(i int) IColumnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumnContext)
}

func (s *Columns_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Columns_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Columns_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.EnterColumns_list(s)
	}
}

func (s *Columns_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.ExitColumns_list(s)
	}
}

func (p *TranslatorParser) Columns_list() (localctx IColumns_listContext) {
	localctx = NewColumns_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TranslatorParserRULE_columns_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Column()
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TranslatorParserT__4 {
		{
			p.SetState(51)
			p.Match(TranslatorParserT__4)
		}
		{
			p.SetState(52)
			p.Column()
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IValues_listContext is an interface to support dynamic dispatch.
type IValues_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValues_listContext differentiates from other interfaces.
	IsValues_listContext()
}

type Values_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValues_listContext() *Values_listContext {
	var p = new(Values_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TranslatorParserRULE_values_list
	return p
}

func (*Values_listContext) IsValues_listContext() {}

func NewValues_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Values_listContext {
	var p = new(Values_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TranslatorParserRULE_values_list

	return p
}

func (s *Values_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Values_listContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *Values_listContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Values_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Values_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Values_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.EnterValues_list(s)
	}
}

func (s *Values_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.ExitValues_list(s)
	}
}

func (p *TranslatorParser) Values_list() (localctx IValues_listContext) {
	localctx = NewValues_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TranslatorParserRULE_values_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Value()
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TranslatorParserT__4 {
		{
			p.SetState(59)
			p.Match(TranslatorParserT__4)
		}
		{
			p.SetState(60)
			p.Value()
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISelect2Context is an interface to support dynamic dispatch.
type ISelect2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelect2Context differentiates from other interfaces.
	IsSelect2Context()
}

type Select2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect2Context() *Select2Context {
	var p = new(Select2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TranslatorParserRULE_select2
	return p
}

func (*Select2Context) IsSelect2Context() {}

func NewSelect2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select2Context {
	var p = new(Select2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TranslatorParserRULE_select2

	return p
}

func (s *Select2Context) GetParser() antlr.Parser { return s.parser }

func (s *Select2Context) Select1() ISelect1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelect1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelect1Context)
}

func (s *Select2Context) NUM_FLOAT() antlr.TerminalNode {
	return s.GetToken(TranslatorParserNUM_FLOAT, 0)
}

func (s *Select2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Select2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.EnterSelect2(s)
	}
}

func (s *Select2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.ExitSelect2(s)
	}
}

func (p *TranslatorParser) Select2() (localctx ISelect2Context) {
	localctx = NewSelect2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TranslatorParserRULE_select2)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Select1()
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TranslatorParserT__5 {
		{
			p.SetState(67)
			p.Match(TranslatorParserT__5)
		}
		{
			p.SetState(68)
			p.Match(TranslatorParserT__6)
		}
		{
			p.SetState(69)
			p.Match(TranslatorParserNUM_FLOAT)
		}

	}

	return localctx
}

// ISelect1Context is an interface to support dynamic dispatch.
type ISelect1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelect1Context differentiates from other interfaces.
	IsSelect1Context()
}

type Select1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect1Context() *Select1Context {
	var p = new(Select1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TranslatorParserRULE_select1
	return p
}

func (*Select1Context) IsSelect1Context() {}

func NewSelect1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select1Context {
	var p = new(Select1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TranslatorParserRULE_select1

	return p
}

func (s *Select1Context) GetParser() antlr.Parser { return s.parser }

func (s *Select1Context) Set_list() ISet_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISet_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISet_listContext)
}

func (s *Select1Context) From_list() IFrom_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFrom_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFrom_listContext)
}

func (s *Select1Context) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *Select1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Select1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.EnterSelect1(s)
	}
}

func (s *Select1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.ExitSelect1(s)
	}
}

func (p *TranslatorParser) Select1() (localctx ISelect1Context) {
	localctx = NewSelect1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TranslatorParserRULE_select1)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(TranslatorParserT__7)
	}
	{
		p.SetState(73)
		p.Set_list()
	}
	{
		p.SetState(74)
		p.Match(TranslatorParserT__8)
	}
	{
		p.SetState(75)
		p.From_list()
	}
	{
		p.SetState(76)
		p.Match(TranslatorParserT__9)
	}
	{
		p.SetState(77)
		p.Condition()
	}

	return localctx
}

// ISet_listContext is an interface to support dynamic dispatch.
type ISet_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSet_listContext differentiates from other interfaces.
	IsSet_listContext()
}

type Set_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_listContext() *Set_listContext {
	var p = new(Set_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TranslatorParserRULE_set_list
	return p
}

func (*Set_listContext) IsSet_listContext() {}

func NewSet_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_listContext {
	var p = new(Set_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TranslatorParserRULE_set_list

	return p
}

func (s *Set_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_listContext) AllPair() []IPairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPairContext)(nil)).Elem())
	var tst = make([]IPairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPairContext)
		}
	}

	return tst
}

func (s *Set_listContext) Pair(i int) IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *Set_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.EnterSet_list(s)
	}
}

func (s *Set_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.ExitSet_list(s)
	}
}

func (p *TranslatorParser) Set_list() (localctx ISet_listContext) {
	localctx = NewSet_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TranslatorParserRULE_set_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Pair()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TranslatorParserT__4 {
		{
			p.SetState(80)
			p.Match(TranslatorParserT__4)
		}
		{
			p.SetState(81)
			p.Pair()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TranslatorParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TranslatorParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *PairContext) Attribute() IAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.ExitPair(s)
	}
}

func (p *TranslatorParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TranslatorParserRULE_pair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.Function()
	}
	{
		p.SetState(88)
		p.Match(TranslatorParserT__1)
	}
	{
		p.SetState(89)
		p.Attribute()
	}
	{
		p.SetState(90)
		p.Match(TranslatorParserT__2)
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TranslatorParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TranslatorParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }
func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *TranslatorParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TranslatorParserRULE_function)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TranslatorParserT__10 || _la == TranslatorParserT__11) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFrom_listContext is an interface to support dynamic dispatch.
type IFrom_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFrom_listContext differentiates from other interfaces.
	IsFrom_listContext()
}

type From_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrom_listContext() *From_listContext {
	var p = new(From_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TranslatorParserRULE_from_list
	return p
}

func (*From_listContext) IsFrom_listContext() {}

func NewFrom_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *From_listContext {
	var p = new(From_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TranslatorParserRULE_from_list

	return p
}

func (s *From_listContext) GetParser() antlr.Parser { return s.parser }

func (s *From_listContext) Relation() IRelationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationContext)
}

func (s *From_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *From_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *From_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.EnterFrom_list(s)
	}
}

func (s *From_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.ExitFrom_list(s)
	}
}

func (p *TranslatorParser) From_list() (localctx IFrom_listContext) {
	localctx = NewFrom_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TranslatorParserRULE_from_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Relation()
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TranslatorParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TranslatorParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) Attribute() IAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *ConditionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(TranslatorParserSTRING_LITERAL, 0)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *TranslatorParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TranslatorParserRULE_condition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Attribute()
	}
	{
		p.SetState(97)
		p.Match(TranslatorParserT__6)
	}
	{
		p.SetState(98)
		p.Match(TranslatorParserSTRING_LITERAL)
	}

	return localctx
}

// IColumnContext is an interface to support dynamic dispatch.
type IColumnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnContext differentiates from other interfaces.
	IsColumnContext()
}

type ColumnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnContext() *ColumnContext {
	var p = new(ColumnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TranslatorParserRULE_column
	return p
}

func (*ColumnContext) IsColumnContext() {}

func NewColumnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnContext {
	var p = new(ColumnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TranslatorParserRULE_column

	return p
}

func (s *ColumnContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TranslatorParserIDENTIFIER, 0)
}

func (s *ColumnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.EnterColumn(s)
	}
}

func (s *ColumnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.ExitColumn(s)
	}
}

func (p *TranslatorParser) Column() (localctx IColumnContext) {
	localctx = NewColumnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TranslatorParserRULE_column)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(TranslatorParserIDENTIFIER)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TranslatorParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TranslatorParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(TranslatorParserSTRING_LITERAL, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *TranslatorParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TranslatorParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Match(TranslatorParserSTRING_LITERAL)
	}

	return localctx
}

// ITable_nameContext is an interface to support dynamic dispatch.
type ITable_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTable_nameContext differentiates from other interfaces.
	IsTable_nameContext()
}

type Table_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_nameContext() *Table_nameContext {
	var p = new(Table_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TranslatorParserRULE_table_name
	return p
}

func (*Table_nameContext) IsTable_nameContext() {}

func NewTable_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_nameContext {
	var p = new(Table_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TranslatorParserRULE_table_name

	return p
}

func (s *Table_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TranslatorParserIDENTIFIER, 0)
}

func (s *Table_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.EnterTable_name(s)
	}
}

func (s *Table_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.ExitTable_name(s)
	}
}

func (p *TranslatorParser) Table_name() (localctx ITable_nameContext) {
	localctx = NewTable_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TranslatorParserRULE_table_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(TranslatorParserIDENTIFIER)
	}

	return localctx
}

// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}

type AttributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeContext() *AttributeContext {
	var p = new(AttributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TranslatorParserRULE_attribute
	return p
}

func (*AttributeContext) IsAttributeContext() {}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext {
	var p = new(AttributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TranslatorParserRULE_attribute

	return p
}

func (s *AttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TranslatorParserIDENTIFIER, 0)
}

func (s *AttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.EnterAttribute(s)
	}
}

func (s *AttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.ExitAttribute(s)
	}
}

func (p *TranslatorParser) Attribute() (localctx IAttributeContext) {
	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TranslatorParserRULE_attribute)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(TranslatorParserIDENTIFIER)
	}

	return localctx
}

// IRelationContext is an interface to support dynamic dispatch.
type IRelationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationContext differentiates from other interfaces.
	IsRelationContext()
}

type RelationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationContext() *RelationContext {
	var p = new(RelationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TranslatorParserRULE_relation
	return p
}

func (*RelationContext) IsRelationContext() {}

func NewRelationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationContext {
	var p = new(RelationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TranslatorParserRULE_relation

	return p
}

func (s *RelationContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TranslatorParserIDENTIFIER, 0)
}

func (s *RelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.EnterRelation(s)
	}
}

func (s *RelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TranslatorListener); ok {
		listenerT.ExitRelation(s)
	}
}

func (p *TranslatorParser) Relation() (localctx IRelationContext) {
	localctx = NewRelationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TranslatorParserRULE_relation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(TranslatorParserIDENTIFIER)
	}

	return localctx
}
