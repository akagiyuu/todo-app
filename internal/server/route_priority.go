package server

import (
	"fmt"
	"net/http"
	"todo-app/internal/database"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetPriorities(c *gin.Context) {
	queries := database.New(s.db)

	priorities, err := queries.GetPriorities(c)
	if err !=  nil {
		s.logger.Error(err)
		c.Error(fmt.Errorf("Failed to get all priorities"))
		return
	}

	c.JSON(http.StatusOK, priorities)
}


