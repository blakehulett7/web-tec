package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Jesus is Lord!")

	db, err := gorm.Open(sqlite.Open("app.db"))
	if err != nil {
		panic("failed to connect to database")
	}

	b := Bridge{
		Db: db,
	}

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "index.html") })
	router.HandleFunc("/new-instance", b.new_instance_handler)
	router.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) { fmt.Println("Made it") })

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
