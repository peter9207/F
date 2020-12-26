package javascript

import "github.com/antlr/antlr4/runtime/Go/antlr"
import "github.com/peter9207/F/W/javascript/lexer"
import "github.com/peter9207/F/W/javascript/parser"

func Parse(program string) {

	is := antlr.NewInputStream(program)

	lexer := lexer.NewJavaScriptLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewJavaScriptParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(&parser.BaseJavaScriptParserListener{}, p.Start())
}
