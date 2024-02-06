package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	err2 := godotenv.Load()
	if err2 != nil {
		panic(err2)
	}
	port := os.Getenv("PORT")
	fmt.Println("server running in port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}

}
