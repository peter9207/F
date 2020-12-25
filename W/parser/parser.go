package parser

import "github.com/tdewolff/parse/v2/js"
import "github.com/tdewolff/parse/v2"
import "os"
import "fmt"
import "encoding/json"

func Parse(filename string) (ast *js.AST, err error) {
	reader, err := os.Open(filename)
	if err != nil {
		return
	}

	ast, err = js.Parse(parse.NewInput(reader))
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

	ast, err := Parse(filename)
	if err != nil {
		return
	}

	err = traverseAST(ast)
	return
}
