package homecontroller

import (
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("views/homepage/index.html")
	if err != nil {
		panic(err)
	}

	err = files.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
