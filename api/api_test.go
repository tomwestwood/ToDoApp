package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"todo/todo"
)

type StubTodoStore struct {
	item func()
	itemsAdded []todo.TodoItem
	todoList []todo.TodoItem
}

func (s *StubTodoStore) AddTodoItem(item todo.TodoItem) {
	s.todoList = append(s.todoList, item)
}

func (s *StubTodoStore) RemoveTodoItem(item todo.TodoItem) error {
	for i, todoItem := range s.todoList {
        if todoItem.Instruction == item.Instruction {
			s.todoList = append(s.todoList[:i], s.todoList[i+1:]...)
			return nil
        }
    }
	
    return fmt.Errorf("couldn't find that todo item.")
}

func (s *StubTodoStore) GetTodoList() []todo.TodoItem {
	return s.todoList
}

func TestGetTodoList(t *testing.T) {
	t.Run("it returns the todo list as JSON", func(t *testing.T) {
		wantedTodoList := []todo.TodoItem{
			{"abc", "Done"},
			{"def", "Not started"},
			{"hij", "I didn't even know about this"},
		}

		store := StubTodoStore{nil, nil, wantedTodoList}
		server := NewTodoServer(&store)

		request := newTodolistRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getTodoListFromResponse(t, response.Body)

		assertStatus(t, response.Code, http.StatusOK)
		assertTodoList(t, got, wantedTodoList)
		assertContentType(t, response, jsonContentType)
	})
}

func TestStoreAdditions(t *testing.T) {
	store := StubTodoStore{
		nil,
		nil,
		[]todo.TodoItem{},
	}
	server := NewTodoServer(&store)

	t.Run("it records new todo item on POST", func(t *testing.T) {
		item := todo.TodoItem { Instruction: "New item", Status: "To do"}
		lengthBeforeAdding := len(store.todoList)

		request := newPostTodoItemRequest(item)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		diff := (len(store.todoList) - lengthBeforeAdding)
		if (diff != 1)  {
			t.Fatalf("got %d calls to add todo item, want %d", diff, 1)
		}

		lastAdded := store.todoList[len(store.todoList)-1]
		if lastAdded.Instruction != item.Instruction {
			t.Errorf("did not store correct todo item, got %q want %q", lastAdded.Instruction, item.Instruction)
		}
	})
}

func TestStoreRemovals(t *testing.T) {
	store := StubTodoStore{
		nil,
		nil,
		[]todo.TodoItem{},
	}
	server := NewTodoServer(&store)

	t.Run("it removes a todo item on REMOVAL", func(t *testing.T) {
		item := todo.TodoItem { Instruction: "Test one", Status: "Done"}
		store.todoList = append(store.todoList, item)
		store.todoList = append(store.todoList, todo.TodoItem{Instruction: "Test two", Status: "Done"})
		lengthBeforeAdding := len(store.todoList)

		request := newRemoveTodoItemRequest(item)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		diff := (lengthBeforeAdding  - len(store.todoList))
		if (diff != 1)  {
			t.Fatalf("Wanted to remove %d items, but looks like it was %d", 1, diff)
		}

		canFind := false
		for _, tdi := range store.todoList {
			if(tdi.Instruction == item.Instruction) {
				canFind = true
				break
			}
		}

		if canFind {
			t.Errorf("did not remove the correct todo item, should have removed %q", item.Instruction)
		}
	})
}

func newTodolistRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/todolist", nil)
	return req
}

func newPostTodoItemRequest(item todo.TodoItem) *http.Request {
	body, _ := json.Marshal(item)
	req, _ := http.NewRequest(http.MethodPost, "/todo", bytes.NewReader(body))
	return req
}

func newRemoveTodoItemRequest(item todo.TodoItem) *http.Request {
	body, _ := json.Marshal(item)
	req, _ := http.NewRequest(http.MethodDelete, "/todo", bytes.NewReader(body))
	return req
}

func getTodoListFromResponse(t testing.TB, body io.Reader) (todoList []todo.TodoItem) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&todoList)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertTodoList(t testing.TB, got, want []todo.TodoItem) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}