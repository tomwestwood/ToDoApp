package jsonRead

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"todo/todo"
)

func TestConsoleItems(t *testing.T) {
	cases := []struct {
		items []todo.TodoItem
		fileContents string 
	}{
		{ items: []todo.TodoItem{ { Instruction: "single item" } }, fileContents: `[{"Instruction":"single item"}]` },
		{ items: []todo.TodoItem{ { Instruction: "item one" }, { Instruction:  "item two"}, { Instruction: "item three" } }, fileContents: `[{"Instruction":"item one"},{"Instruction":"item two"},{"Instruction":"item three"}]`},
	}

	for _, test := range cases {		
		t.Run(fmt.Sprintf("%q reads the following to json 'file': %q", test.items, test.fileContents), func(t *testing.T) {
			var b bytes.Buffer			
			b.WriteString(test.fileContents)
			read := JsonRead {}
			read.Reader = &b
			
			i, e := read.Read()
			if(e != nil) {
				fmt.Errorf("there's been a problem %w", e)
			}

			j, e := json.Marshal(i)

			if(e != nil) {
				fmt.Errorf("could not encode the returned json content")
			}

			if(string(j) != test.fileContents) {
				fmt.Errorf("these don't match")
			}
		})
	}
}