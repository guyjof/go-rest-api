package models

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var Todos = []Todo{
	{ID: "1", Item: "Learn Go", Completed: false},
	{ID: "2", Item: "Build a RESTful API in Go", Completed: false},
	{ID: "3", Item: "Build a React app", Completed: false},
}
