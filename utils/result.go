package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ERROR   = -1
	SUCCESS = 1
)

func Result(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"resultCode":    code,
		"resultMessage": msg,
	})
}

func ResultData(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"resultCode":    code,
		"resultMessage": msg,
		"data":          data,
	})
}

func SuccessDataResult(c *gin.Context, msg string, data interface{}) {
	ResultData(c, SUCCESS, msg, data)
}

func SuccessResult(c *gin.Context, msg string) {
	Result(c, SUCCESS, msg)
}

func FailedResult(c *gin.Context, msg string) {
	Result(c, ERROR, msg)
}
