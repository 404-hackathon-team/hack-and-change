package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Jeno7u/studybud/cmd/api"
	"github.com/Jeno7u/studybud/config"
	"github.com/Jeno7u/studybud/db"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {
	conn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBHost,
		config.Envs.DBName,
	)

	database, err := db.NewPostgresStorage(conn)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	migrateDatabase(database)
	initStorage(database)

	port := fmt.Sprintf(":%s", config.Envs.Port)
	server := api.NewAPIServer(port, database)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// ping db
func initStorage(database *sql.DB) {
	err := database.Ping()
	if err != nil {
		log.Fatal("tried to ping db, got", err)
	}

	log.Println("DB: Successfully connected")
}

// migrate db
func migrateDatabase(database *sql.DB) {
	goose.SetDialect("postgres")

	if err := goose.Up(database, "migrations"); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
}