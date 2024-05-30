package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"todo/todo"
)

var todoList []todo.TodoItem

func main() {
	runApp(os.Stdout)
}

func runApp(writer io.Writer) {
	fmt.Println("Starting...")
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
  		scanner.Scan()
		cliArguments := strings.Split(scanner.Text(), " ")
		
		handleCommand(cliArguments, writer)
	}
}

func handleCommand(cliArguments []string, writer io.Writer) {
	switch strings.ToLower(cliArguments[0]) {
	case "add":
		handleAdd(writer)
	case "read":
		handleRead(writer)
	case "remove":
		handleDelete(writer)
	case "update":
		handleUpdate(writer)
	default:
		handleBadCommand(writer, cliArguments[0])
	}
}

func handleAdd(writer io.Writer) {
	fmt.Fprintln(writer, "*ADD TODO ITEM* -> Please enter the description of your new Todo item")

	scanner := bufio.NewScanner(os.Stdin)
	
	for {
  		scanner.Scan()
		todoList = append(todoList, todo.TodoItem{scanner.Text(), "Not started"})
		fmt.Fprintln(writer, "Added new todo item: '%w'", scanner.Text())
		break
	}
}

func handleRead(writer io.Writer) {
	for _, todoItem := range todoList {
		fmt.Fprintln(writer, todoItem.Instruction)
		fmt.Fprintln(writer, "--->", todoItem.Status)
	}
}

func handleDelete(writer io.Writer) {
	fmt.Fprintln(writer, "*REMOVE TODO ITEM* -> Please enter the description of the todo item you'd like to remove")
	scanner := bufio.NewScanner(os.Stdin)
	var err error = nil
	
	for {
  		scanner.Scan()
		todoList, err = remove(scanner.Text())

		if(err != nil) {
			fmt.Fprintln(writer, err)
			break
		}

		fmt.Fprintln(writer, "Removed todo item: '%w'", scanner.Text())
	}
}

func handleBadCommand(writer io.Writer, command string) {
	fmt.Fprintf(writer, "Invalid operation %s", command)
}

func handleUpdate(writer io.Writer) {
	fmt.Fprintln(writer, "*UPDATE TODO ITEM* -> Please enter the description of the todo item you'd like to update")
	scanner := bufio.NewScanner(os.Stdin)
	var err error = nil
	
	for {
  		scanner.Scan()
		err = update(writer, scanner.Text())

		if(err != nil) {
			fmt.Fprintln(writer, err)
			break
		}

		fmt.Fprintln(writer, "Updated todo item: '%w'", scanner.Text())
		break
	}
}

func remove(instruction string) (items []todo.TodoItem, err error) {
    for i, todoItem := range todoList {
        if todoItem.Instruction == instruction {
            return append(todoList[:i], todoList[i+1:]...), nil
        }
    }


    return todoList, fmt.Errorf("Couldn't find that todo item, pal.")
}

func update(writer io.Writer, instruction string) error {
	for _, todoItem := range todoList {
		if todoItem.Instruction == instruction {
			updateStatusTo(writer, &todoItem)
			return nil
		}
	}

    return fmt.Errorf("Couldn't find that todo item, sorry.")
}

func updateStatusTo(writer io.Writer, item *todo.TodoItem) {	
	fmt.Fprintln(writer, "Please enter the new status for item '%w'", item.Instruction)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		item.Status = scanner.Text()
		fmt.Fprintln(writer, "Status updated to %w", item.Status)
		break
	}
}