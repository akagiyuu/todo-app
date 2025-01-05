package server

import (
	"fmt"
	"net/http"
	"strconv"
	"todo-app/internal/database"

	"github.com/gin-gonic/gin"
)

type CreateTodoParams struct {
	Title       string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
	Priority    string `form:"priority" binding:"required"`
	Category    string `form:"category" binding:"required"`
}

func (s *Server) CreateTodo(c *gin.Context) {
	var params CreateTodoParams
	if err := c.ShouldBind(&params); err != nil {
		s.logger.Error(err)
		c.Error(fmt.Errorf("Invalid data provided"))
		return
	}

	queries := database.New(s.db)

	category_id, err := queries.GetCategory(c, params.Category)
	if err != nil {
		category_id, err = queries.CreateCategory(c, params.Category)
		if err != nil {
			s.logger.Error(err)
			c.Error(fmt.Errorf("Invalid category: %s", params.Category))
		}
	}
	priority_id, err := queries.GetPriority(c, params.Priority)
	if err != nil {
		s.logger.Error(err)
		c.Error(fmt.Errorf("Invalid priority: %s", params.Priority))
	}
	if c.Errors.Last() != nil {
		return
	}

	err = queries.CreateTodo(c, database.CreateTodoParams{
		Title:       params.Title,
		Description: params.Description,
		CategoryID:  category_id,
		PriorityID:  priority_id,
	})
	if err != nil {
		s.logger.Error(err)
		c.Error(fmt.Errorf("Failed to create todo with data: %+v", params))
		return
	}

	c.JSON(http.StatusOK, params)
}

type FilterTodoParams struct {
	Priority string `form:"priority" json:"priority"`
	Category string `form:"category" json:"category"`
}

func (s *Server) FilterTodo(c *gin.Context) {
	var params FilterTodoParams
	if err := c.ShouldBind(&params); err != nil {
		s.logger.Error(err)
		c.Error(fmt.Errorf("Invalid data provided"))
		return
	}

	queries := database.New(s.db)

	todo, err := queries.FilterTodo(c, database.FilterTodoParams{
		Priority: params.Priority,
		Category: params.Category,
	})
	if err != nil {
		s.logger.Error(err)
		c.Error(fmt.Errorf("Failed to delete todo with parameter: %+v", params))
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (s *Server) GetTodo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		s.logger.Error(err)
		c.Error(fmt.Errorf("Invalid todo id: %s", c.Param("id")))
		return
	}

	queries := database.New(s.db)

	todo, err := queries.GetTodo(c, id)
	if err != nil {
		s.logger.Error(err)
		c.Error(fmt.Errorf("Failed to get todo with id %d", id))
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (s *Server) DeleteTodo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		s.logger.Error(err)
		c.Error(fmt.Errorf("Invalid todo id: %s", c.Param("id")))
		return
	}

	queries := database.New(s.db)

	if err := queries.DeleteTodo(c, id); err != nil {
		s.logger.Error(err)
		c.Error(fmt.Errorf("Failed to delete todo with id %d", id))
	}
}

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(ErrorHandler)
	r.POST("/todo", s.CreateTodo)
	r.GET("/todo", s.FilterTodo)
	r.GET("/todo/:id", s.GetTodo)
	r.DELETE("/todo/:id", s.DeleteTodo)

	return r
}
