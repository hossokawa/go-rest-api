package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Code a REST API", Completed: false},
	{ID: "2", Item: "Walk the dog", Completed: false},
	{ID: "3", Item: "Wash the dishes", Completed: false},
}

func getTodos(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, todos)
}

func addTodo(ctx *gin.Context) {
	var newTodo todo
	if err := ctx.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	ctx.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	ctx.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("could not find todo with matching ID")
}

func main() {
	r := gin.Default()

	r.GET("/todos", getTodos)
	r.GET("/todos/:id", getTodo)
	r.PATCH("/todos/:id", toggleTodoStatus)
	r.POST("/todos", addTodo)
	r.Run("localhost:8080")
}
