package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"

	"github.com/Jeno7u/studybud/cmd/api"
	"github.com/Jeno7u/studybud/config"
	"github.com/Jeno7u/studybud/db"
)


func main() {
	db, err := db.NewMySQLStorate(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("tried to ping db, got", err)
	}

	log.Println("DB: Successfully connected")
}