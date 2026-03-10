package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func ConnectDB() {
	var err error

	// Init connection
	Pool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}

	// Verify connection
	err = Pool.Ping(context.Background())
	if err != nil {
		log.Fatal("Unable to ping database: ", err)
	}

	log.Println("Connected to database")
}
