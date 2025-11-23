package handlers

import (
	"github.com/BimaPDev/MixMatch/db"
)

// Handler struct holds dependencies (like the database)
// so that all functions can access them.
type Handler struct {
	store *db.Queries
}

// NewHandler creates a new Handler struct with the database connection
func NewHandler(store *db.Queries) *Handler {
	return &Handler{store: store}
}
