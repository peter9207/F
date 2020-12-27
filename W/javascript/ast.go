package javascript

import (
	"fmt"
)

type Node map[string]interface{}

func (e *Node) Type() string {

	t, ok := e["type"].(string)
	if !ok {
		fmt.Println("[Warning] expression has no field type")
	}
	return t
}

func parse() {
	t := map[string]interface{}{
		"testing": 123,
	}

	n := Node{}
	for k, v := range t {
		n[k] = v
	}

}

type NodeList []Node
