// Code generated from Translator.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Translator

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TranslatorListener is a complete listener for a parse tree produced by TranslatorParser.
type TranslatorListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterSelect1 is called when entering the select1 production.
	EnterSelect1(c *Select1Context)

	// EnterSet_list is called when entering the set_list production.
	EnterSet_list(c *Set_listContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterFrom_list is called when entering the from_list production.
	EnterFrom_list(c *From_listContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterRelation is called when entering the relation production.
	EnterRelation(c *RelationContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitSelect1 is called when exiting the select1 production.
	ExitSelect1(c *Select1Context)

	// ExitSet_list is called when exiting the set_list production.
	ExitSet_list(c *Set_listContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitFrom_list is called when exiting the from_list production.
	ExitFrom_list(c *From_listContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitRelation is called when exiting the relation production.
	ExitRelation(c *RelationContext)
}
