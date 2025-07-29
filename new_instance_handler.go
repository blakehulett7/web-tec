package main

import (
	"fmt"
	"net/http"
)

func (b Bridge) NewInstanceHandler(w http.ResponseWriter, r *http.Request) {
	state := NewState()
	b.Db.Create(&state)

	redirect_url := fmt.Sprintf("/instance/%v", state.Id)
	http.Redirect(w, r, redirect_url, http.StatusSeeOther)
}
