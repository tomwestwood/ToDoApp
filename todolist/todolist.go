package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"todo/todo"
)

func main() {
	readToDos(os.Stdout)
}


func readToDos(writer io.Writer) {
	var wg sync.WaitGroup
	instructionStream := make(chan todo.TodoItem)
	statusStream := make(chan bool)
	
	todoItems := []todo.TodoItem{ 
		{Instruction: "item one", Status: "Done"}, 
		{Instruction: "item two", Status: "Not Done"}, 
		{Instruction: "item three", Status: "Doing"}, 
		{Instruction: "item four", Status: "Done"},
	}

    wg.Add(2)

	go printTodoInstructions(&wg, todoItems, instructionStream, statusStream, writer)
	go printTodoStatuses(&wg, instructionStream, statusStream, writer)

	wg.Wait()
}

func printTodoInstructions(wg *sync.WaitGroup, todoItems []todo.TodoItem, instructionStream chan todo.TodoItem, statusStream chan bool, writer io.Writer) {
	defer wg.Done()
	for _, todo := range todoItems {
		fmt.Fprintln(writer, "Instruction:", todo.Instruction)
		instructionStream <- todo
		<-statusStream
	}
	close(instructionStream)
}

func printTodoStatuses(wg *sync.WaitGroup, instructionStream chan todo.TodoItem, statusStream chan bool, writer io.Writer) {
	defer wg.Done()
	for todo := range instructionStream {
		fmt.Fprintln(writer, "---> Status:", todo.Status)
		statusStream <- true
	}
	close(statusStream)
}
