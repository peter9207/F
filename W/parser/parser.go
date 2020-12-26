package parser

import "github.com/tdewolff/parse/v2/js"
import "github.com/tdewolff/parse/v2"
import "os"
import "fmt"
import "encoding/json"
import "io"

func Parse(filename string) (err error) {
	reader, err := os.Open(filename)
	if err != nil {
		return
	}

	l := js.NewLexer(parse.NewInput(reader))
	// fmt.Println(ast)
	for {
		tt, text := l.Next()
		switch tt {
		case js.ErrorToken:
			if l.Err() != io.EOF {
				fmt.Println("Error on line", text, ":", l.Err())
			}
			return
		case js.IdentifierToken:
			fmt.Println("Identifier", string(text))
		case js.NumericToken:
			fmt.Println("Numeric", string(text))
		case js.OpenBraceToken:
			fmt.Println("OpenBrace", string(text))
		case js.CloseBraceToken:
			fmt.Println("CloseBrace", string(text))
		case js.LineTerminatorToken:
		default:
			fmt.Println("default case", tt, string(text))

		}
	}
	return
}

func traverseAST(ast *js.AST) (err error) {

	fmt.Println("starting ast traverse")

	data, err := json.Marshal(ast)
	if err != nil {
		return
	}

	fmt.Println(string(data))
	fmt.Println("finished")
	return
}

func CreateTree(filename string) (err error) {

	err = Parse(filename)
	if err != nil {
		return
	}

	// err = traverseAST(ast)
	return
}
