// internal/handler/handler.go
package handler

import (
	"log"

	"github.com/zkurdi45/mystack-go/internal/data"
)

// Handler holds all application-wide dependencies for our handlers.
type Handler struct {
	Logger *log.Logger
	Models data.Models
}

// New creates a new Handler instance.
func New(
	logger *log.Logger,
	models data.Models,
) *Handler {
	return &Handler{
		Logger: logger,
		Models: models,
	}
}
