# Todo Rest Api

A Todo Rest Api it uses the following technologies:
1. golang
2. sqlite3
3. gin-gonic

## Installation

1. Clone the repository
git clone https://github.com/Akash2023das/todo-rest-api-sqlite-gin-golang.git

```Install the dependencies
go mod download
```

## Usage

The api has seven commands:

1. r.GET("/todos", newTodo.GetAllTodos)      // this is the function that gets all the todos
2. r.GET("/todo/:id", newTodo.GetATodo)      // this is the function that gets a single todo
3.	r.POST("/todo", newTodo.PostTodo)         // this is the function that creates a new todo
4.	r.PUT("/todo/:id", newTodo.UpdateTodo)    // this is the function that updates the todo
5.	r.DELETE("/todo/:id", newTodo.DeleteTodo) // this is the function that deletes the todo
6.	r.PUT("/comtodo/:status/:id", newTodo.CompletedTodos)
7.	r.GET("/donetodo", newTodo.GetAllDoneTodos)
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)