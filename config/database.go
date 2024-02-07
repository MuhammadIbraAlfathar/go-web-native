package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var DB *sql.DB

func ConnectionDB() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	userDB := os.Getenv("DB_USERNAME")
	dbName := os.Getenv("DB_NAME")
	db, err := sql.Open("mysql", userDB+"@/"+dbName)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	log.Print("Database connected")
	DB = db
}
