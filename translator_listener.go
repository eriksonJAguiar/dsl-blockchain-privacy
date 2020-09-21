// Code generated from Translator.g4 by ANTLR 4.8. DO NOT EDIT.

package main // Translator
import "github.com/antlr/antlr4/runtime/Go/antlr"

// TranslatorListener is a complete listener for a parse tree produced by TranslatorParser.
type TranslatorListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterInsert1 is called when entering the insert1 production.
	EnterInsert1(c *Insert1Context)

	// EnterColumns_list is called when entering the columns_list production.
	EnterColumns_list(c *Columns_listContext)

	// EnterValues_list is called when entering the values_list production.
	EnterValues_list(c *Values_listContext)

	// EnterSelect2 is called when entering the select2 production.
	EnterSelect2(c *Select2Context)

	// EnterSelect1 is called when entering the select1 production.
	EnterSelect1(c *Select1Context)

	// EnterSet_list is called when entering the set_list production.
	EnterSet_list(c *Set_listContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterFrom_list is called when entering the from_list production.
	EnterFrom_list(c *From_listContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterColumn is called when entering the column production.
	EnterColumn(c *ColumnContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterTable_name is called when entering the table_name production.
	EnterTable_name(c *Table_nameContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterRelation is called when entering the relation production.
	EnterRelation(c *RelationContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitInsert1 is called when exiting the insert1 production.
	ExitInsert1(c *Insert1Context)

	// ExitColumns_list is called when exiting the columns_list production.
	ExitColumns_list(c *Columns_listContext)

	// ExitValues_list is called when exiting the values_list production.
	ExitValues_list(c *Values_listContext)

	// ExitSelect2 is called when exiting the select2 production.
	ExitSelect2(c *Select2Context)

	// ExitSelect1 is called when exiting the select1 production.
	ExitSelect1(c *Select1Context)

	// ExitSet_list is called when exiting the set_list production.
	ExitSet_list(c *Set_listContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitFrom_list is called when exiting the from_list production.
	ExitFrom_list(c *From_listContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitColumn is called when exiting the column production.
	ExitColumn(c *ColumnContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitTable_name is called when exiting the table_name production.
	ExitTable_name(c *Table_nameContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitRelation is called when exiting the relation production.
	ExitRelation(c *RelationContext)
}
