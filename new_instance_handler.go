package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (b Bridge) NewInstanceHandler(w http.ResponseWriter, r *http.Request) {
	id := uuid.New()
	state := State{Id: id}

	b.Db.Create(&state)

	redirect_url := fmt.Sprintf("/%v", id)
	http.Redirect(w, r, redirect_url, http.StatusSeeOther)
}
