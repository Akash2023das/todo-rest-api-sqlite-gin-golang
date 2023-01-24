package main

import (
	"database/sql"
	"log"

	"github.com/Akash2023das/todo-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the table if it doesn't exist
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS todo (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		status BOOLEAN DEFAULT 0
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s", err, sqlStmt)
		return
	}

	r := gin.Default()

	newTodo := controllers.NewTodoStore(db)

	r.GET("/todos", newTodo.GetAllTodos)      // this is the function that gets all the todos
	r.GET("/todo/:id", newTodo.GetATodo)      // this is the function that gets a single todo
	r.POST("/todo", newTodo.PostTodo)         // this is the function that creates a new todo
	r.PUT("/todo/:id", newTodo.UpdateTodo)    // this is the function that updates the todo
	r.DELETE("/todo/:id", newTodo.DeleteTodo) // this is the function that deletes the todo
	r.PUT("/comtodo/:status/:id", newTodo.CompletedTodos)
	r.GET("/donetodo", newTodo.GetAllDoneTodos)
	
	r.Run()
}

//To test this api
//1. get all todos
//select -> GET http://localhost:8080/todos

//2. get a single todo
//select -> GET http://localhost:8080/todo/1

//3. create a new todo
//select -> POST -H "Content-Type: application/json" -d '{"title":"todo 1","status":"active"}' http://localhost:8080/todo

//4. update a todo
//select ->  PUT -H "Content-Type: application/json" -d '{"title":"todo 1","status":"inactive"}' http://localhost:8080/todo/1

//5. delete a todo
//select -> DELETE http://localhost:8080/todo/1
