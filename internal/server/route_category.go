package server

import (
	"fmt"
	"net/http"
	"todo-app/internal/database"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetCategories(c *gin.Context) {
	queries := database.New(s.db)

	categories, err := queries.GetCategories(c)
	if err !=  nil {
		s.logger.Error(err)
		c.Error(fmt.Errorf("Failed to get all categories"))
		return
	}

	c.JSON(http.StatusOK, categories)
}
