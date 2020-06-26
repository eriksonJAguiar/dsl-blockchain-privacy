// Code generated from Translator.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Translator

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTranslatorListener is a complete listener for a parse tree produced by TranslatorParser.
type BaseTranslatorListener struct{}

var _ TranslatorListener = &BaseTranslatorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTranslatorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTranslatorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTranslatorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTranslatorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseTranslatorListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseTranslatorListener) ExitStart(ctx *StartContext) {}

// EnterSelect1 is called when production select1 is entered.
func (s *BaseTranslatorListener) EnterSelect1(ctx *Select1Context) {}

// ExitSelect1 is called when production select1 is exited.
func (s *BaseTranslatorListener) ExitSelect1(ctx *Select1Context) {}

// EnterSet_list is called when production set_list is entered.
func (s *BaseTranslatorListener) EnterSet_list(ctx *Set_listContext) {}

// ExitSet_list is called when production set_list is exited.
func (s *BaseTranslatorListener) ExitSet_list(ctx *Set_listContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseTranslatorListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseTranslatorListener) ExitFunction(ctx *FunctionContext) {}

// EnterFrom_list is called when production from_list is entered.
func (s *BaseTranslatorListener) EnterFrom_list(ctx *From_listContext) {}

// ExitFrom_list is called when production from_list is exited.
func (s *BaseTranslatorListener) ExitFrom_list(ctx *From_listContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseTranslatorListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseTranslatorListener) ExitCondition(ctx *ConditionContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseTranslatorListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseTranslatorListener) ExitAttribute(ctx *AttributeContext) {}

// EnterRelation is called when production relation is entered.
func (s *BaseTranslatorListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BaseTranslatorListener) ExitRelation(ctx *RelationContext) {}
