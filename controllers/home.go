package controllers

import (
	"net/http"
	"html/template"
)

func HomeIndex(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)

}