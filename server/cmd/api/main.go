package main

import (
	"log"
	"os"

	// Correct imports using your GitHub module path
	"github.com/BimaPDev/MixMatch/db"
	"github.com/BimaPDev/MixMatch/internal/api"
	"github.com/BimaPDev/MixMatch/internal/database" // <--- This fixes "undefined: database"
)

func main() {
	// 1. Connect to DB
	conn := database.Connect()
	defer conn.Close(nil)

	// 2. Create the Store (sqlc)
	store := db.New(conn)

	// 3. Create the Server (Inject the store)
	server := api.NewServer(store)

	// 4. Start the Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	err := server.Start(":" + port)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
