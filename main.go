package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Jesus is Lord!")

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "index.html") })
	router.HandleFunc("/sign-up", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "sign_up.html") })
	router.HandleFunc("/new-instance", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "new_instance.html") })

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
