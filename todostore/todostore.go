package todostore

import (
	"fmt"
	"todo/todo"
)

type TodoStore struct {
	todoList []todo.TodoItem
}

func NewInMemoryTodoStore() *TodoStore {
	return &TodoStore{
		[]todo.TodoItem {},
	}
}

func (t *TodoStore) GetTodoList() []todo.TodoItem {
	return t.todoList
}

func (t *TodoStore) AddTodoItem(item todo.TodoItem) {
 	t.todoList = append(t.todoList, item)
}

func (t *TodoStore) RemoveTodoItem(item todo.TodoItem) error {
 	for i, todoItem := range t.todoList {
        if todoItem.Instruction == item.Instruction {
			t.todoList = append(t.todoList[:i], t.todoList[i+1:]...)
			return nil
        }
    }
	
    return fmt.Errorf("couldn't find that todo item.")
}