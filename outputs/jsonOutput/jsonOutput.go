package jsonOutput

import (
	"encoding/json"
	"fmt"
	"io"
	"todo/todo"
)

type JsonOutput struct {
	Writer io.Writer
}

func (o *JsonOutput) Output(items ...todo.TodoItem) error {
	itemsJson, err := json.Marshal(items)
	
	if err != nil {
 		return fmt.Errorf("error encoding JSON: %w", err)
 	}

	fmt.Fprint(o.Writer, string(itemsJson))
	return nil
}