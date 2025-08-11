package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (b Bridge) NewInstanceHandler(w http.ResponseWriter, r *http.Request) {
	state := State{Id: uuid.New()}
	b.Db.Create(&state)

	PrettyPrint("STATE", state)

	redirect_url := fmt.Sprintf("/instance/%v", state.Id)
	http.Redirect(w, r, redirect_url, http.StatusSeeOther)
}
