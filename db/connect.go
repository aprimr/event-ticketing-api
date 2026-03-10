package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool
var ctx = context.Background()
var err error

func ConnectDB() {
	// Init connection
	Pool, err = pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}

	// Verify connection
	err = Pool.Ping(ctx)
	if err != nil {
		log.Fatal("Unable to ping database: ", err)
	}

	log.Println("Connected to database")
}
