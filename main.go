package main

import (
	"fmt"
	"os"
	"todo/outputs/consoleOutput"
	"todo/outputs/jsonOutput"
	"todo/reads/jsonRead"
	"todo/todo"
	"todo/writes/jsonWrite"
)


func main() {
	var todoItems []todo.TodoItem
	todoItems = append(todoItems, todo.TodoItem{"item one", "Done"}, todo.TodoItem{"item two", "Not Done"}, todo.TodoItem{"item three", "Doing"}, todo.TodoItem{"item four", "Done"})

	//outputItemsToConsole(todoItems)
	//outputItemsToConsoleAsJson(todoItems)
	//writeItemsToJsonFile(todoItems)

	readItemsFromJsonFile()
}

func outputItemsToConsole(todoItems []todo.TodoItem) {
	output := consoleOutput.ConsoleOutput{ Writer: os.Stdout }
	outputItems(&output, todoItems)
}

func outputItemsToConsoleAsJson(todoItems []todo.TodoItem) {	
	output := jsonOutput.JsonOutput{ Writer: os.Stdout }
	outputItems(&output, todoItems)
}

func writeItemsToJsonFile(todoItems []todo.TodoItem) {
	// writeFile := func(data []byte) error {
	// 	return ioutil.WriteFile("realdata.json", data, 0644)
	// }

	var f, e = jsonWrite.CreateFile()
	if(e != nil) {
		fmt.Errorf("error")
	}

	write := jsonWrite.JsonWrite {}
	//write.WriteFileFunction = writeFile
	write.Writer = f
	writeItems(&write, todoItems)
}

func readItemsFromJsonFile() {
	// readFile := func(filename string) []byte {
	// 	itemsJson, err := os.ReadFile(filename)
	// 	if(err != nil) {
	// 		panic("problemo")
	// 	}

	// 	return itemsJson
	// }

	f, e := os.Open("items.json")
	if e != nil {
        fmt.Errorf("Failed to open file: %v", e)
    }

	read := jsonRead.JsonRead {}
	read.Reader = f
	items, e := read.Read()

	if(e != nil) {
		fmt.Errorf("Problem reading the file")
	}

	outputItemsToConsole(items)
}

func outputItems(output todo.Output, todoItems []todo.TodoItem) {
	e := output.Output(todoItems...)
	check(e)

	if(e == nil) {
		fmt.Println("All items outputted successfully.")
	}
}

func writeItems(write todo.Write, todoItems []todo.TodoItem) {
	e := write.Write(todoItems...)
	check(e)

	if(e == nil) {
		fmt.Println("All items written successfully.")
	}
}
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}