package server

import (
	"database/sql"
	"fmt"
	"github.com/charmbracelet/log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	port   int
	db     *sql.DB
	logger *log.Logger
}

func NewServer() *http.Server {
	db, err := sql.Open("sqlite3", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Error("Failed to open database", err)
		return nil
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})

	NewServer := &Server{
		port:   port,
		db:     db,
		logger: logger,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
