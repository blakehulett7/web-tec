package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func (b Bridge) LoadInstanceHandler(w http.ResponseWriter, r *http.Request) {
	var app State
	result := b.Db.First(&app, "id = ?", r.PathValue("id"))
	if result.Error != nil {
		fmt.Println("save not found, handle this...")
		http.NotFound(w, r)
		return
	}

	t := template.Must(template.ParseFiles("templates/instance.html"))
	t.Execute(w, app)
}
