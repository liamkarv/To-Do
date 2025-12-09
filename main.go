package main

type todo struct {
	ID string
	Item string
	Completed bool
}

var todos = []todo {
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Break my balls", Completed: false},
	{ID: "3", Item: "Breathe", Completed: false},
}