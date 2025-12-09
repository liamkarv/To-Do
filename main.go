package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Client and server communicate through JSON
type todo struct{
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Break my balls", Completed: false},
	{ID: "3", Item: "Breathe", Completed: false},
}

func addToDo(context *gin.Context){
	var newToDo todo

	if err := context.BindJSON(&newToDo); err != nil {
		return
	}

	todos = append(todos, newToDo)
	context.IndentedJSON(http.StatusCreated, newToDo)
}

func getToDos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}

func main(){
	router := gin.Default()
	router.GET("/todos", getToDos)
	router.POST("/todos", addToDo)
	router.Run("localhost:8080")
}
