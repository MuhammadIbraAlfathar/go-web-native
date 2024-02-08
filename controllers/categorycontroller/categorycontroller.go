package categorycontroller

import (
	"database/sql"
	"github.com/MuhammadIbraAlfathar/go-web-native/entities"
	"github.com/MuhammadIbraAlfathar/go-web-native/models/categorymodel"
	"html/template"
	"net/http"
	"strconv"
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

func Create(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {

			var categories entities.Category

			err := r.ParseForm()

			if err != nil {
				panic(err)
			}

			categories.Name = r.Form["name"][0]

			categorymodel.CreateCategory(db, categories)

			if err != nil {
				_, _ = w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/categories", http.StatusMovedPermanently)

		} else if r.Method == "GET" {
			files, err := template.ParseFiles("views/categorypage/create.html")
			if err != nil {
				panic(err)
			}

			err = files.Execute(w, nil)
			if err != nil {
				panic(err)
			}
		}
	}
}

func Edit(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			files, err := template.ParseFiles("views/categorypage/edit.html")

			idString := r.URL.Query().Get("id")
			id, _ := strconv.Atoi(idString)

			category := categorymodel.DetailCategory(id, db)

			data := map[string]any{
				"category": category,
			}

			if err != nil {
				panic(err)
			}

			err = files.Execute(w, data)
			if err != nil {
				panic(err)
			}
		} else if r.Method == "POST" {
			idString := r.URL.Query().Get("id")
			id, _ := strconv.Atoi(idString)

			_ = r.ParseForm()

			name := r.Form["name"][0]

			categorymodel.EditCategory(id, name, db)

			http.Redirect(w, r, "/categories", http.StatusMovedPermanently)
		}
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {

}
