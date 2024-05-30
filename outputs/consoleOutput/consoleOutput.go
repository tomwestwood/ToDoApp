package consoleOutput

import (
	"fmt"
	"io"
	"todo/todo"
)

type ConsoleOutput struct{
	Writer io.Writer
}

func (o *ConsoleOutput) Output(items ...todo.TodoItem ) error {
	for _, item := range items {
		fmt.Fprintln(o.Writer, item.Instruction)
	}
	return nil
}