package routes

import (
	"database/sql"
	"github.com/MuhammadIbraAlfathar/go-web-native/controllers/categorycontroller"
	"github.com/MuhammadIbraAlfathar/go-web-native/controllers/homecontroller"
	"net/http"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", homecontroller.HomePage)
	//category
	server.HandleFunc("/categories", categorycontroller.Index(db))
}
