package api

import (
	"encoding/json"
	"net/http"
	"todo/todo"
)

const jsonContentType = "application/json"

type TodoStore interface {
	AddTodoItem(item todo.TodoItem)
	RemoveTodoItem(item todo.TodoItem) error
	GetTodoList() []todo.TodoItem
}

type TodoServer struct {
	Store TodoStore
	http.Handler
}

func NewTodoServer(store TodoStore) *TodoServer {
	p := new(TodoServer)
	p.Store = store

	router := http.NewServeMux()
	router.Handle("/todo", http.HandlerFunc(p.todoHandler))
	router.Handle("/todolist", http.HandlerFunc(p.listHandler))

	p.Handler = router
	return p
}

func (p *TodoServer) todoHandler(w http.ResponseWriter, r *http.Request) {
	var todoItem todo.TodoItem
	err := json.NewDecoder(r.Body).Decode(&todoItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPost:
		p.addTodoItem(w, todoItem)
	case http.MethodDelete:
		p.removeTodoItem(w, todoItem)
	}
}


func (p *TodoServer) listHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.Store.GetTodoList())
}

func (p *TodoServer) addTodoItem(w http.ResponseWriter, item todo.TodoItem) {
	p.Store.AddTodoItem(item)
	w.WriteHeader(http.StatusAccepted)
}

func (p *TodoServer) removeTodoItem(w http.ResponseWriter, item todo.TodoItem) {
	p.Store.RemoveTodoItem(item)
	w.WriteHeader(http.StatusAccepted)
}

func (p *TodoServer) getTodoList(w http.ResponseWriter) []todo.TodoItem {
	return p.Store.GetTodoList()
}
