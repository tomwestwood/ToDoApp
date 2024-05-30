package jsonRead

import (
	"encoding/json"
	"io"
	"todo/todo"
)

type JsonRead struct {
	//ReadFileFunction func(fileLocation string) []byte
	Reader io.Reader
}


func (o *JsonRead) Read() ([]todo.TodoItem, error) {	
	//fileContents := o.ReadFileFunction(filePath)
	var items []todo.TodoItem
    decoder := json.NewDecoder(o.Reader)
    e := decoder.Decode(&items)

    if e != nil {
        return nil, e
    }
    return items, nil
}