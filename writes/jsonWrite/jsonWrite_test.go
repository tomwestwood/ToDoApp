package jsonWrite

import (
	"bytes"
	"fmt"
	"testing"
	"todo/todo"
)

func TestConsoleItems(t *testing.T) {
	cases := []struct {
		items []todo.TodoItem
		expected string 
	}{
		{ items: []todo.TodoItem{ { Instruction: "single item" } }, expected: `[{"Instruction":"single item"}]` },
		{ items: []todo.TodoItem{ { Instruction: "item one" }, { Instruction:  "item two"}, { Instruction: "item three" } }, expected: `[{"Instruction":"item one"},{"Instruction":"item two"},{"Instruction":"item three"}]`},
	}

	for _, test := range cases {		
		t.Run(fmt.Sprintf("%q writes the following to jsone 'file': %q", test.items, test.expected), func(t *testing.T) {
			var b bytes.Buffer
			write := JsonWrite {}
			write.Writer = &b
			
			e := write.Write(test.items...)
			if(e != nil) {
				fmt.Errorf("there's been a problem %w", e)
			}

			if b.Len() == 0 {
				t.Fatalf("Expected data to be written, but got none")
			}

			if(b.String() != test.expected) {
				fmt.Errorf("these don't match")
			}
		})
	}
}