package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func (b Bridge) StartInstanceHandler(w http.ResponseWriter, r *http.Request) {
	var app State
	result := b.Db.Preload("IData").First(&app, "id = ?", r.PathValue("id"))
	if result.Error != nil {
		fmt.Println("save not found, handle this...")
		http.NotFound(w, r)
		return
	}

	t := template.Must(template.ParseFiles("templates/instance.html"))
	w.Header().Add("HX-Push-Url", r.URL.String())
	t.Execute(w, app)
}
