package consoleOutput

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
		{ items: []todo.TodoItem{ { Instruction: "single item" } }, expected: `single item
` },
		{ items: []todo.TodoItem{ { Instruction: "item one" }, { Instruction:  "item two"}, { Instruction: "item three" } }, expected: `item one
item two
item three
`},
	}

	for _, test := range cases {		
		t.Run(fmt.Sprintf("%q gets converted to %q", test.items, test.expected), func(t *testing.T) {
			buffer := bytes.Buffer{}
			output := &ConsoleOutput{ Writer: &buffer }
			output.Output(test.items...)
			got := buffer.String()
			if got != test.expected {
				t.Errorf("got %q, want %q", got, test.expected)
			}
		})
	}
}