package todos

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var todo_items []string

func init()  {
	todo_items = make([]string, 5)
	todo_items[0] = "123"
	todo_items[1] = "456"
}


func All(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"todos": todo_items,
	})
}

func One(ctx *gin.Context)  {
	id := ctx.Param("id")
	i, _ := strconv.Atoi(id)
	ctx.JSON(http.StatusOK, gin.H{
		"todo": todo_items[i],
	})
}