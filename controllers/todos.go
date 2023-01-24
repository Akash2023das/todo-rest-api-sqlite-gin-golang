package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Akash2023das/todo-rest-api/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// var todos []Todo //use slice to store todos
type TodoStore struct {
	Todos []Todo
	db    *sql.DB
}

func NewTodoStore(dbIns *sql.DB) *TodoStore {
	return &TodoStore{
		Todos: []Todo{},
		db:    dbIns,
	}
}

// take user input and store todos
// func init() {
// 	todos = append(todos, Todo{ID: 1, Title: "Todo 1", Status: "active"})
// 	todos = append(todos, Todo{ID: 2, Title: "Todo 2", Status: "active"})
// 	todos = append(todos, Todo{ID: 3, Title: "Todo 3", Status: "active"})

// }

// this is the function that gets all the todos
// func GetAllTodos(c *gin.Context) {
// 	c.JSON(200, todos)
// }

// create a GetAllTodos method that gets all the todos from the database

func (t TodoStore) GetAllTodos(c *gin.Context) {
	rows, err := t.db.Query("SELECT * FROM todo")
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error getting Todos")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Status)
		if err != nil {
			utils.Response(c, http.StatusInternalServerError, nil, "Error getting todos")
			return
		}
		t.Todos = append(t.Todos, todo)
	}
	utils.Response(c, http.StatusOK, t.Todos, "Todos found")
}

// create a GetATodo method that gets a todo from the database
func (t TodoStore) GetATodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error getting todo")
		return
	}

	row := t.db.QueryRow("SELECT * FROM todo WHERE id = ?", id)
	var todo Todo
	err = row.Scan(&todo.ID, &todo.Title, &todo.Status)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error getting todo")
		return
	}

	utils.Response(c, http.StatusNotFound, todo, "Todo found")
}

// this is the function that creates a new todo
func (t TodoStore) PostTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error creating Todo")
		return
	}

	// Save the blog to the database
	stmt, err := t.db.Prepare("INSERT INTO todo(title, status) VALUES(?, ?)")
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error creating todo")
		return
	}
	res, err := stmt.Exec(todo.Title, todo.Status)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error creating todo")
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error creating todo")
		return
	}
	// Marshal the blog into json
	utils.Response(c, http.StatusCreated, id, "Todo created successfully")
}

// this is the function that updates the todo
func (t TodoStore) UpdateTodo(c *gin.Context) {
	// Get blog id from url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error getting todo")
		return
	}

	// Get the blog from the request body
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error updating todo")
		return
	}

	// Update the todo in the database
	stmt, err := t.db.Prepare("UPDATE todo SET title = ?, status = ? WHERE id = ?")
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error updating todo")
		return
	}
	_, err = stmt.Exec(todo.Title, todo.Status, id)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error updating todo")
		return
	}

	// If the todo is not found, return 404
	utils.Response(c, http.StatusNoContent, nil, "Todo Updated")
}

// this is the function that deletes the todo
func (t TodoStore) DeleteTodo(c *gin.Context) {
	// Get todo id from url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error getting todo")
		return
	}
	// Delete the todo from the database
	stmt, err := t.db.Prepare("DELETE FROM todo WHERE id = ?")

	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error deleting todo")
		return
	}

	_, err = stmt.Exec(id)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error deleting todo")
		return
	}

	// If the todo is not found, return 404
	utils.Response(c, http.StatusNoContent, nil, "Todo Deleted")
}

// create a function that only shows the completed todos
func (t TodoStore) CompletedTodos(c *gin.Context) {

	// Get todo id from url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error getting todo")
		return
	}
	status, err := strconv.ParseBool(c.Param("status"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error getting todo")
		return
	}
	// Delete the todo from the database
	stmt, err := t.db.Prepare("UPDATE todo SET status=? where id = ?")

	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error deleting todo")
		return
	}

	_, err = stmt.Exec(status, id)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error deleting todo")
		return
	}

	// If the todo is not found, return 404
	utils.Response(c, http.StatusNoContent, nil, "Todo Deleted")
}

func (t TodoStore) GetAllDoneTodos(c *gin.Context) {
	rows, err := t.db.Query("SELECT * FROM todo where status=true")
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, nil, "Error getting Todos")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Status)
		if err != nil {
			utils.Response(c, http.StatusInternalServerError, nil, "Error getting todos")
			return
		}
		t.Todos = append(t.Todos, todo)
	}
	utils.Response(c, http.StatusOK, t.Todos, "Todos found")
}
