package database // <--- Check this! It must be "package database", not "package db"

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

// Connect returns a raw database connection
func Connect() *pgx.Conn {
	// Update this connection string if needed (e.g. user/password)
	connString := "postgres://postgres:secret@localhost:5432/stylesphere"

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	log.Println("Connected to Postgres successfully")
	return conn
}
