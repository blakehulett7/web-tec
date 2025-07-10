package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Jesus is Lord!")

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "index.html") })
	router.HandleFunc("/new-instance", new_instance_handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
