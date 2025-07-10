package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func new_instance_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Nomine Patris...")

	id := uuid.New()
	state := init_state(id)
}
