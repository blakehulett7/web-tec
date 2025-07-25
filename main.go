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

	db.AutoMigrate(&State{})

	b := Bridge{
		Db: db,
	}

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "index.html") })
	router.HandleFunc("/new-instance", b.NewInstanceHandler)
	router.HandleFunc("/{id}", b.LoadInstanceHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
