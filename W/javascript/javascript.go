package javascript

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const parserURL = "localhost:3000/parse"

type Node map[string]interface{}

func analyze(astString string) (err error) {

	jsonMap := Node{}
	err = json.Unmarshal([]byte(astString), &jsonMap)
	if err != nil {
		return
	}

	d := Depth(jsonMap)
	fmt.Println("code has a depth of ", d)
	return
}

func Depth(root interface{}) (n int) {
	m, ok := root.(map[string]interface{})
	if !ok {
		return 1
	}

	currentMax := 1
	for _, v := range m {

		switch v.(type) {
		case []interface{}:

			l := v.([]interface{})
			for _, e := range l {

				d := Depth(e)
				if d > currentMax {
					currentMax = d
				}

			}
		case map[string]interface{}:
			d := Depth(v)
			if d > currentMax {
				currentMax = d
			}
		default:
			fmt.Println("hit dead end  skipping...", v)
		}
	}

	n = currentMax + 1

	return

}

func Parse(program string) (err error) {

	resp, err := http.Post(parserURL, "text/plain", strings.NewReader(program))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	_ = analyze(string(body))
	return
}
