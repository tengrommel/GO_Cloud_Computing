package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/tengrommel/GO_Cloud_Computing/Cloud_Computing_Overview/All_Go/First/api/todos"
	"github.com/tengrommel/GO_Cloud_Computing/Cloud_Computing_Overview/All_Go/First/middlewares"
)

func main() {
	app := gin.Default()

	app.GET("/hello/:name", hello)

	todosGrp := app.Group("/api/todos")
	todosGrp.GET("/", middlewares.Log, todos.All)
	todosGrp.GET("/:id", todos.One)
	todosGrp.PUT("/:id")
	todosGrp.POST("/")
	todosGrp.DELETE("/:id")

	app.Run(":8081")
}

func hello(ctx *gin.Context) {
		name := ctx.Param("name")
		page := ctx.Query("page")
		fmt.Println(page)
		ctx.String(http.StatusOK, "hello %s", name)
}
