package todos

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/tengrommel/GO_Cloud_Computing/Cloud_Computing_Overview/All_Go/First/libs/database"
	"github.com/tengrommel/GO_Cloud_Computing/Cloud_Computing_Overview/All_Go/First/models"
	"github.com/jinzhu/gorm"
	"fmt"
)

var todo_items []string

func init()  {
	todo_items = make([]string, 5)
	todo_items[0] = "123"
	todo_items[1] = "456"
}


func All(ctx *gin.Context) {
	db := database.Open()
	defer db.Close()

	var todoItems []models.Todo
	//db.Model(&models.Todo{}).Where("created_at = ?", time.Now()).Find(&todoItems)
	if err := db.Find(&todoItems).Error; err != nil{
		if err == gorm.ErrRecordNotFound{
			ctx.Status(http.StatusNotFound)
		}else {
			ctx.Status(http.StatusInternalServerError)
		}
		return
	}
	ctx.JSON(http.StatusOK, todoItems)
}

func One(ctx *gin.Context)  {
	id := ctx.Param("id")

	db := database.Open()
	defer db.Close()

	var todoItem models.Todo
	if err := db.Model(&todoItem).Where("id = ?", id).First(&todoItem).Error; err !=nil{
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, todoItem)
}

type postCreate struct {
	Title  		string		`json:"title" binding:"required"`
}

func Create(ctx *gin.Context)  {
	var postData postCreate
	if err := ctx.BindJSON(&postData); err!=nil{
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	db := database.Open()
	defer db.Close()

	todoItem := models.Todo{
		Title: "Hello",
	}

	if err := db.Create(&todoItem).Error; err != nil{
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": todoItem.ID})
}

func Update(ctx *gin.Context)  {
	id := ctx.Param("id")

	db := database.Open()
	defer db.Close()

	var todoItem models.Todo
	if err := db.Model(&todoItem).Where("id = ?", id).First(&todoItem).Error; err != nil{
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	//todoItem.Title = "World"
	//if err := db.Save(&todoItem).Error; err != nil{
	//	fmt.Println(err)
	//	ctx.Status(http.StatusInternalServerError)
	//	return
	//}
	if err :=db.Model(&todoItem).Update("title", "direct update yo!").Error; err != nil{
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}