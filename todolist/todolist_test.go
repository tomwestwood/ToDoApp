package main

import (
	"bytes"
	"testing"
)

func TestTodoOutput(t *testing.T) {
	buffer := bytes.Buffer{}
	readToDos(&buffer)

	got := buffer.String()
	want := `Instruction: item one
---> Status: Done
Instruction: item two
---> Status: Not Done
Instruction: item three
---> Status: Doing
Instruction: item four
---> Status: Done
`

	if(want != got) {
		t.Errorf("got %q want %q", got, want)
	}
}