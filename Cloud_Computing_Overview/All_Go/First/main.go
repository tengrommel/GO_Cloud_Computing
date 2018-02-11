package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tengrommel/GO_Cloud_Computing/Cloud_Computing_Overview/All_Go/First/api/todos"
)

func main() {
	app := gin.Default()

	app.GET("/api/todos", todos.All)
	app.POST("/api/todos/:id", todos.One)
	app.POST("/api/todos", todos.Create)
	app.PUT("/api/todos/:id", todos.Update)

	app.Run(":8081")
}
