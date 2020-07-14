// Code generated from Translator.g4 by ANTLR 4.8. DO NOT EDIT.

package main // Translator

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

// EnterInsert1 is called when production insert1 is entered.
func (s *BaseTranslatorListener) EnterInsert1(ctx *Insert1Context) {}

// ExitInsert1 is called when production insert1 is exited.
func (s *BaseTranslatorListener) ExitInsert1(ctx *Insert1Context) {}

// EnterColumns_list is called when production columns_list is entered.
func (s *BaseTranslatorListener) EnterColumns_list(ctx *Columns_listContext) {}

// ExitColumns_list is called when production columns_list is exited.
func (s *BaseTranslatorListener) ExitColumns_list(ctx *Columns_listContext) {}

// EnterValues_list is called when production values_list is entered.
func (s *BaseTranslatorListener) EnterValues_list(ctx *Values_listContext) {}

// ExitValues_list is called when production values_list is exited.
func (s *BaseTranslatorListener) ExitValues_list(ctx *Values_listContext) {}

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

// EnterColumn is called when production column is entered.
func (s *BaseTranslatorListener) EnterColumn(ctx *ColumnContext) {}

// ExitColumn is called when production column is exited.
func (s *BaseTranslatorListener) ExitColumn(ctx *ColumnContext) {}

// EnterValue is called when production value is entered.
func (s *BaseTranslatorListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseTranslatorListener) ExitValue(ctx *ValueContext) {}

// EnterTable_name is called when production table_name is entered.
func (s *BaseTranslatorListener) EnterTable_name(ctx *Table_nameContext) {}

// ExitTable_name is called when production table_name is exited.
func (s *BaseTranslatorListener) ExitTable_name(ctx *Table_nameContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseTranslatorListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseTranslatorListener) ExitAttribute(ctx *AttributeContext) {}

// EnterRelation is called when production relation is entered.
func (s *BaseTranslatorListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BaseTranslatorListener) ExitRelation(ctx *RelationContext) {}
