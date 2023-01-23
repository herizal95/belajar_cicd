package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/herizal95/golang_gin_starter/controllers"
)

func BookRoutes(route *gin.RouterGroup) {
	route.GET("/book", controllers.GetAllBook)
	route.GET("/book/:id", controllers.GetBook)
	route.POST("/book", controllers.CreateBook)
	// route.PUT("/book", controllers.UpdateBook)
	// route.DELETE("/book", controllers.DeleteBook)
}
