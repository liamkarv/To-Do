package main

import (
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

func getTodos(context *gin.Context){

}

func main(){
	router := gin.Default()
	router.GET("/todos")
	router.Run("localhost:8080")
}
