package main

import (
	"github.com/MuhammadIbraAlfathar/go-web-native/config"
	"github.com/MuhammadIbraAlfathar/go-web-native/routes"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err2 := godotenv.Load()
	if err2 != nil {
		panic(err2)
	}

	db := config.ConnectionDB()

	server := http.NewServeMux()

	routes.MapRoutes(server, db)

	port := os.Getenv("PORT")
	log.Println("Server running in port" + port)
	err := http.ListenAndServe(port, server)
	if err != nil {
		panic(err)
	}

}
