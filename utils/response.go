package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(c *gin.Context, body interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"msg":   "success",
		"error": 0,
		"data":  body,
	})
}

func ResponseError(c *gin.Context, msg interface{}) {
	emptyArray := make([]string, 0)
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":   msg,
		"error": 1,
		"data":  emptyArray,
	})
}
