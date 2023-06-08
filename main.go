package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// referenced video: https://www.youtube.com/watch?v=d_L64KT3SFM&list=LL&index=2&t=632s

type todo struct {
	ID string `json:"id"`
	Item string `json:"title"`
	Completed bool `json:"completed"`
}

var todos = []todo {
	{ID: "1", Item: "Read Book", Completed: false},
	{ID: "2", Item: "Cook Mornin", Completed: false},
	{ID: "3", Item: "Clean Room", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string)(*todo, error) {
	for i, t := range todos {
		if(t.ID == id) {
			// 目的のTodoが取得出来た場合
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func toggleCompleteState(context *gin.Context) {
	id := context.Param("id")
	todo, err  := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}


func main() {
	fmt.Println("Start Program")
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", addTodo)
	router.PATCH("/todos/:id", toggleCompleteState)
	router.Run("localhost:9090")
}