package main

import (
	"bytes"
	"testing"
)

func TestRaceOutput(t *testing.T) {
	buffer := bytes.Buffer{}
	raceCorrect_InOrder(&buffer)

	got := buffer.String()
	want := `1
2
3
4
5
6
7
8
9
10
`

	if(want != got) {
		t.Errorf("got %q want %q", got, want)
	}
}