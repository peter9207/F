package parser

import "github.com/robertkrimen/otto/parser"
import "github.com/robertkrimen/otto/ast"

import "io/ioutil"

func Parse(filename string) (program *ast.Program, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	program, err = parser.ParseFile(nil, filename, data, parser.IgnoreRegExpErrors)
	return
}
