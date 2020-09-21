//
package main

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type translatorListener struct {
	*BaseTranslatorListener

	// select
	queryString string // json format to query in couchDB
	attributes  []string
	functions   []string
	epsilon		string


	// insert
	insertString string // json format to insert in couchDB
	columns      []string
	values       []string
	docType      string
}

// select parser
// SELECT AVG(id) FROM doctype2 WHERE coluna1="valor1"
// SELECT AVG(coluna1) FROM doctype3 WHERE coluna2=\"1111\"
func (t *translatorListener) ExitSelect1(c *Select1Context) {
	t.queryString = "{\n" + t.queryString + "\n}"
}

func (t *translatorListener) ExitSelect2(c *Select2Context) {
	numFloat := c.NUM_FLOAT();
	if numFloat != nil {
		t.epsilon = numFloat.GetText();
	}	
}

func (t *translatorListener) ExitSet_list(c *Set_listContext) {
	var s string
	firstLine := true
	for _, attribute := range t.attributes {
		if firstLine == true {
			firstLine = false
		} else {
			s = s + ","
		}
		s = s + "\"" + attribute + "\""
	}
	t.queryString = "\"fields\":[" + s + "]"
}

func (t *translatorListener) ExitPair(c *PairContext) {
	t.attributes = append(t.attributes, c.Attribute().GetText())
	t.functions = append(t.functions, c.Function().GetText())
}

func (t *translatorListener) ExitRelation(c *RelationContext) {
	t.docType = c.GetText()
	t.queryString = "\"docType\":\"" + c.GetText() + "\"},\n" + t.queryString
}

func (t *translatorListener) ExitCondition(c *ConditionContext) {
	t.queryString = "\"" + c.Attribute().GetText() + "\":" + c.STRING_LITERAL().GetText() + "," + t.queryString
	t.queryString = "\"selector\":{" + t.queryString
}


//insert parser executeCommand
// INSERT INTO doctype1 (coluna1, coluna2, coluna3) VALUES ("valor1","valor2","valor3")
//INSERT INTO doctype3 (coluna1, coluna2, coluna3) VALUES (\"1111\",\"1111\",\"1111\")
func (t *translatorListener) ExitColumn(c *ColumnContext) {
	t.columns = append(t.columns, c.GetText())
}

func (t *translatorListener) ExitValue(c *ValueContext) {
	t.values = append(t.values, strings.Trim(c.GetText(), "\""))
}

func (t *translatorListener) ExitTable_name(c *Table_nameContext) {
	t.docType = c.GetText()
}

func (t *translatorListener) ExitInsert1(c *Insert1Context) {
	firstMemberWritten := false
	for i, s := range t.columns {
		if firstMemberWritten == true {
			t.insertString = t.insertString + ","
		}
		t.insertString = t.insertString + "\"" + s + "\":\"" + t.values[i] + "\""
		firstMemberWritten = true
	}
	t.insertString = t.insertString + "," + "\"docType\":\"" + t.docType + "\""
	t.insertString = "{" + t.insertString + "}"

	antlr.NewInputStream("af")
}

/*
func main() {
	// Setup the input
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	is := antlr.NewInputStream(text)

	// Create the Lexer
	lexer := parser.NewTranslatorLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewTranslatorParser(stream)

	var listener translatorListener
	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())
	fmt.Println(listener.insertString)
}
*/