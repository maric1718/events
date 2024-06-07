package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ThrowStatusOk(i interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, i)
}

func ThrowStatusInternalServerError(msg string, c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": msg,
		"error":   msg,
	})
}
