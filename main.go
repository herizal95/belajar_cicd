package main

import (
	"github.com/gin-gonic/gin"
	"github.com/herizal95/golang_gin_starter/config"
	"github.com/herizal95/golang_gin_starter/routes"
	"github.com/herizal95/golang_gin_starter/utils"
)

func main() {

	config.ConnectDatabase()

	router := gin.Default()

	api := router.Group(utils.BaseUrl)
	{
		routes.BookRoutes(api)
	}

	router.Run(":8080")
}
