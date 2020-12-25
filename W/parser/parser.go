package parser

import "github.com/tdewolff/parse/v2/js"

import "github.com/tdewolff/parse/v2"
import "os"
import "fmt"

func Parse(filename string) (err error) {
	reader, err := os.Open(filename)
	if err != nil {
		return
	}

	p, err := js.Parse(parse.NewInput(reader))
	// ast, err := js.Parse(parse.NewInputString("if (state == 5) { console.log('In state five'); }"))
	if err != nil {
		return
	}

	fmt.Println(p)

	// for {
	// 	tt, text := lex.Next()
	// 	switch tt {
	// 	case js.ErrorToken:
	// 		if lex.Err() != io.EOF {
	// 			fmt.Println("Error on line", text, ":", lex.Err())
	// 			return lex.Err()
	// 		}
	// 		fmt.Println("finished")
	// 		return
	// 	case js.IdentifierToken:
	// 		fmt.Println("Identifier", string(text))
	// 	case js.NumericToken:
	// 		fmt.Println("Numeric", string(text))
	// 	default:
	// 		fmt.Printf("default : %s :  %T  \n", text, tt)
	// 	}

	// }
	return
}
