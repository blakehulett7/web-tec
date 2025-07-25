package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (b Bridge) new_instance_handler(w http.ResponseWriter, r *http.Request) {
	id := uuid.New()
	state := NewState(id)

	redirect_url := fmt.Sprintf("/%v", id)
	http.Redirect(w, r, redirect_url, http.StatusSeeOther)
}
