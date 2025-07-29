package main

import (
	"fmt"
	"net/http"
)

func (b Bridge) LoadInstanceHandler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("/instance/%v", r.FormValue("hash"))
	http.Redirect(w, r, url, http.StatusSeeOther)
}
