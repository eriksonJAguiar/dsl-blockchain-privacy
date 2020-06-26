// Code generated from Translator.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Translator

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 46, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 2, 2, 10, 2, 4, 6, 8, 10,
	12, 14, 16, 2, 3, 3, 2, 8, 9, 2, 37, 2, 18, 3, 2, 2, 2, 4, 21, 3, 2, 2,
	2, 6, 28, 3, 2, 2, 2, 8, 33, 3, 2, 2, 2, 10, 35, 3, 2, 2, 2, 12, 37, 3,
	2, 2, 2, 14, 41, 3, 2, 2, 2, 16, 43, 3, 2, 2, 2, 18, 19, 5, 4, 3, 2, 19,
	20, 7, 2, 2, 3, 20, 3, 3, 2, 2, 2, 21, 22, 7, 3, 2, 2, 22, 23, 5, 6, 4,
	2, 23, 24, 7, 4, 2, 2, 24, 25, 5, 10, 6, 2, 25, 26, 7, 5, 2, 2, 26, 27,
	5, 12, 7, 2, 27, 5, 3, 2, 2, 2, 28, 29, 5, 8, 5, 2, 29, 30, 7, 6, 2, 2,
	30, 31, 5, 14, 8, 2, 31, 32, 7, 7, 2, 2, 32, 7, 3, 2, 2, 2, 33, 34, 9,
	2, 2, 2, 34, 9, 3, 2, 2, 2, 35, 36, 5, 16, 9, 2, 36, 11, 3, 2, 2, 2, 37,
	38, 5, 14, 8, 2, 38, 39, 7, 10, 2, 2, 39, 40, 7, 13, 2, 2, 40, 13, 3, 2,
	2, 2, 41, 42, 7, 12, 2, 2, 42, 15, 3, 2, 2, 2, 43, 44, 7, 12, 2, 2, 44,
	17, 3, 2, 2, 2, 2,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'SELECT'", "'FROM'", "'WHERE'", "'('", "')'", "'AVG'", "'COUNT'",
	"'='",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "WHITESPACE", "IDENTIFIER", "STRING_LITERAL",
}

var ruleNames = []string{
	"start", "select1", "set_list", "function", "from_list", "condition", "attribute",
	"relation",
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
	TranslatorParserWHITESPACE     = 9
	TranslatorParserIDENTIFIER     = 10
	TranslatorParserSTRING_LITERAL = 11
)

// TranslatorParser rules.
const (
	TranslatorParserRULE_start     = 0
	TranslatorParserRULE_select1   = 1
	TranslatorParserRULE_set_list  = 2
	TranslatorParserRULE_function  = 3
	TranslatorParserRULE_from_list = 4
	TranslatorParserRULE_condition = 5
	TranslatorParserRULE_attribute = 6
	TranslatorParserRULE_relation  = 7
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

func (s *StartContext) Select1() ISelect1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelect1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelect1Context)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(TranslatorParserEOF, 0)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Select1()
	}
	{
		p.SetState(17)
		p.Match(TranslatorParserEOF)
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
	p.EnterRule(localctx, 2, TranslatorParserRULE_select1)

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
		p.SetState(19)
		p.Match(TranslatorParserT__0)
	}
	{
		p.SetState(20)
		p.Set_list()
	}
	{
		p.SetState(21)
		p.Match(TranslatorParserT__1)
	}
	{
		p.SetState(22)
		p.From_list()
	}
	{
		p.SetState(23)
		p.Match(TranslatorParserT__2)
	}
	{
		p.SetState(24)
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

func (s *Set_listContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *Set_listContext) Attribute() IAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
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
	p.EnterRule(localctx, 4, TranslatorParserRULE_set_list)

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
		p.SetState(26)
		p.Function()
	}
	{
		p.SetState(27)
		p.Match(TranslatorParserT__3)
	}
	{
		p.SetState(28)
		p.Attribute()
	}
	{
		p.SetState(29)
		p.Match(TranslatorParserT__4)
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
	p.EnterRule(localctx, 6, TranslatorParserRULE_function)
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
		p.SetState(31)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TranslatorParserT__5 || _la == TranslatorParserT__6) {
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
	p.EnterRule(localctx, 8, TranslatorParserRULE_from_list)

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
		p.SetState(33)
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
	p.EnterRule(localctx, 10, TranslatorParserRULE_condition)

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
		p.SetState(35)
		p.Attribute()
	}
	{
		p.SetState(36)
		p.Match(TranslatorParserT__7)
	}
	{
		p.SetState(37)
		p.Match(TranslatorParserSTRING_LITERAL)
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
	p.EnterRule(localctx, 12, TranslatorParserRULE_attribute)

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
		p.SetState(39)
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
	p.EnterRule(localctx, 14, TranslatorParserRULE_relation)

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
		p.SetState(41)
		p.Match(TranslatorParserIDENTIFIER)
	}

	return localctx
}
