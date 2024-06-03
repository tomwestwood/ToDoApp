package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/todo"
	"todo/todostore"
)

func TestAddingTodoListItems(t *testing.T) {
	store := todostore.NewInMemoryTodoStore()
	server := NewTodoServer(store)

	item1 := todo.TodoItem { Instruction: "todo item 1", Status: "complete" }
	item2 := todo.TodoItem { Instruction: "todo item 2", Status: "incomplete"}

	server.ServeHTTP(httptest.NewRecorder(), newPostTodoItemRequest(item1))
	server.ServeHTTP(httptest.NewRecorder(), newPostTodoItemRequest(item2))

	t.Run("get todo list", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newTodolistRequest())
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "[{\"Instruction\":\"todo item 1\",\"Status\":\"complete\"},{\"Instruction\":\"todo item 2\",\"Status\":\"incomplete\"}]\n")
	})

	t.Run("test remove item 1", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newRemoveTodoItemRequest(item1))
		assertStatus(t, response.Code, http.StatusAccepted)

		server.ServeHTTP(response, newTodolistRequest())
		assertStatus(t, response.Code, http.StatusAccepted)

		assertResponseBody(t, response.Body.String(), "[{\"Instruction\":\"todo item 2\",\"Status\":\"incomplete\"}]\n")
	})
}