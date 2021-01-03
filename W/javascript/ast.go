package javascript

import (
	"fmt"
)

// type Node map[string]interface{}

type Node struct {
	Value    string
	Type     string `json:"type"`
	meta     map[string]string
	Children []Node
}

type Program struct {
	Body []Statement `json:"body"`
}

type BlockStatement struct {
	Body []Statement `json:"body"`
}

type ExpressionStatement struct {
}

type Identifier struct {
	name string `json:"name"`
}

type Literal struct {
	raw string `json:"raw"`
}
type VariableDeclarator struct {
	id Identifier `json:"id"`
}
type VariableDeclaration struct {
	Kind        string             `json:"kind"`
	Declaration VariableDeclarator `json:"declarator"`
}

type BinaryExpression struct {
}

type IfStatement struct {
	test       string
	subsequent string
	alternate  string
}

type FunctionDeclaration struct {
}

func toNode(root map[string]interface{}) (node *Node) {

	node = &Node{}

	for k, v := range root {

		switch k {
		case "type":
			node.Type = v
		}

	}

	// switch root.(type) {
	// case []interface{}:
	// 	for _, v := range root {

	// 	}
	// case map[string]interface{}:
	// 	for k, v := range root {

	// 		node[k] = toNode(v)
	// 	}
	// 	return
	// default:
	// }
}

// create data nodes

func listToNodes(l []interface{}) (nodes []Node) {

}

type NodeList []Node
