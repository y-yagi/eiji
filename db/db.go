package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

func Connect() *pgx.Conn {
	user := getEnv("QOVERY_DATABASE_MY_DB_USERNAME", "postgres")
	password := getEnv("QOVERY_DATABASE_MY_DB_PASSWORD", "postgres")
	addr := getEnv("QOVERY_DATABASE_MY_DB_HOST", "localhost")
	port := getEnv("QOVERY_DATABASE_MY_DB_PORT", "5432")
	database := getEnv("QOVERY_DATABASE_MY_DB_DATABASE", "postgres")
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, addr, port, database)
	conn, err := pgx.Connect(context.Background(), url)

	if err != nil {
		log.Printf("Failed to connect to PotsgreSQL")
		os.Exit(100)
	}

	log.Printf("Connected to PostgreSQL")

	return conn
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
