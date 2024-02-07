package categorycontroller

import (
	"database/sql"
	"github.com/MuhammadIbraAlfathar/go-web-native/models/categorymodel"
	"html/template"
	"net/http"
)

func Index(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		category := categorymodel.GetAll(db)
		data := map[string]any{
			"categories": category,
		}

		files, err := template.ParseFiles("views/categorypage/index.html")
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = files.Execute(w, data)
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
