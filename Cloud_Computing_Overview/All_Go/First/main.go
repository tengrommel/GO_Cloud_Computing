package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func main() {
	app := gin.Default()
	app.GET("/hello/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		page := ctx.Query("page")
		fmt.Println(page)
		ctx.String(http.StatusOK, "hello %s", name)
	})
	app.Run(":8081")
}
