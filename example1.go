//
package main

import (
	"github.com/eriksonJAguiar/DP-Tool-blockchain/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"fmt"
	"bufio"
	"os"
)

type translatorListener struct {
	*parser.BaseTranslatorListener

	Json string
}

func (t *translatorListener) ExitSelect1 (c *parser.Select1Context) {
	t.Json = "{\n" + t.Json + "\n}"
}

func (t *translatorListener) ExitSet_list(c *parser.Set_listContext) {
	t.Json = "\"fields\":[\"" + c.Attribute().GetText() + "\"]"
}

func (t *translatorListener) ExitRelation(c *parser.RelationContext) {
	t.Json =  "\"doctype\":\"" + c.GetText() + "\"},\n" + t.Json 
}

func (t *translatorListener) ExitCondition(c *parser.ConditionContext) {
	t.Json = "\"" + c.Attribute().GetText() + "\":" + c.STRING_LITERAL().GetText()+ "," + t.Json
	t.Json = "\"selector\":{" + t.Json
}



func main() {
	// Setup the input
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	
	// SELECT AVG(id) FROM dicom WHERE nome="joao"
	is := antlr.NewInputStream(text)

	// Create the Lexer
	lexer := parser.NewTranslatorLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewTranslatorParser(stream)

	var listener translatorListener
	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())
	fmt.Println(listener.Json)
}