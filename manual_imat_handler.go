package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func (b Bridge) ManualIMatHander(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	r.ParseForm()
	var imat IMat
	for key, value_array := range r.Form {
		count, err := strconv.Atoi(value_array[0])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		count++

		imat.UserId = id
		imat.Name = key
		imat.Count = count
	}

	t := template.Must(template.ParseFiles("templates/imat.html"))
	t.Execute(w, imat)
}
