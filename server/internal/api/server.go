package api

import (
	"github.com/BimaPDev/MixMatch/db"
	"github.com/BimaPDev/MixMatch/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer(store *db.Queries) *Server {
	h := handlers.NewHandler(store)
	router := gin.Default()

	// --- User Routes ---
	router.POST("/users", h.CreateUser)
	router.GET("/users/:id", h.GetUser) // <--- NEW

	// --- Item Routes ---
	router.POST("/items", h.CreateItem)
	router.GET("/items", h.ListItems) // <--- NEW (Usage: /items?owner_id=1)

	return &Server{router: router}
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
