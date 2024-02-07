package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func ConnectionDB() *sql.DB {
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

	return db
}
