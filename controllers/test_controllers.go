package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/herizal95/golang_gin_starter/config"
	"github.com/herizal95/golang_gin_starter/models"
	"github.com/herizal95/golang_gin_starter/utils"
)

// get all books
func GetAllBook(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			utils.ResponseError(c, err)
			return
		}
	}()
	var book []models.Book
	config.DB.Find(&book)
	utils.ResponseSuccess(c, book)
}

// get by id
func GetBook(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			utils.ResponseError(c, err)
			return
		}
	}()
	var book []models.Book
	if err := config.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	utils.ResponseSuccess(c, book)
}

// create new book
func CreateBook(c *gin.Context) {

	// validasi input
	var input models.Book

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, err.Error())
		return
	}

	data := models.Book{
		Title:       input.Title,
		Description: input.Description,
	}

	config.DB.Create(&data)

	utils.ResponseSuccess(c, data)
}
