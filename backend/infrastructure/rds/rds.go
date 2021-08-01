package rds

import (
	"database/sql"
	"log"
	"time"

	"github.com/kkhan01/caputo/backend/config"

	_ "github.com/lib/pq"
)

func GetConnection(dbConfig *config.DBConfig) *sql.DB {
	// TODO: disable sslmode conditionally
	conn, err := sql.Open("postgres", dbConfig.DSN()+" sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to db:", err.Error())
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal("Failed to ping to db:", err.Error())
	}

	conn.SetConnMaxLifetime(time.Minute)
	return conn
}
