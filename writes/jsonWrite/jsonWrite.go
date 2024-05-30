package jsonWrite

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"todo/todo"
)

type JsonWrite struct {
	//WriteFileFunction func(data []byte) error
	Writer io.Writer
}


func (o *JsonWrite) Write(items ...todo.TodoItem) error {	
	itemsJson, err := json.Marshal(items)	
	if err != nil {
 		return fmt.Errorf("error encoding JSON: %w", err)
 	}

	//file, err = newFunction()

	// err1 := o.WriteFileFunction([]byte(itemsJson))

	// if(err1 != nil) {
	// 	return fmt.Errorf("error writing file: %w", err1)
	// }

	//encoder := json.NewEncoder(o.Writer)
	//err = encoder.Encode(itemsJson)
	//if (err != nil) {
	//	return fmt.Errorf("failed to encode To Do List as JSON %v", err)
	//}

	o.Writer.Write(itemsJson)

	return nil
}

func CreateFile() (file *os.File, err error) {
	f, e := os.Create("items.json")
	if e != nil {
		fmt.Errorf("Failed to create file: %v", err)
	}
	return f, e
}