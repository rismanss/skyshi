package routes

import (
	"skyshi/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Route(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "wellcome",
		})
	})

	// activities route
	r.GET("/activity-groups", controllers.GetAllActivity)
	r.POST("/activity-groups", controllers.CreateActivity)
	r.GET("/activity-groups/:id", controllers.GetByIdActivity)
	r.PATCH("/activity-groups/:id", controllers.UpdateActivity)
	r.DELETE("/activity-groups/:id", controllers.DeleteActivity)

	// todos route
	r.GET("/todo-items", controllers.GetAllTodo)
	r.GET("/todo-items/:id", controllers.GetByIdTodo)
	r.POST("/todo-items", controllers.CreateTodo)
	r.PATCH("/todo-items/:id", controllers.UpdateTodo)
	r.DELETE("/todo-items/:id", controllers.DeleteTodo)

	return r
}
