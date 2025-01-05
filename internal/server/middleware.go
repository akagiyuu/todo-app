package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	if c.Errors.Last() == nil {
		return
	}

	c.JSON(http.StatusBadRequest, c.Errors)
}
