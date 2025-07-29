package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Dominus Iesus Christus")

	db, err := gorm.Open(sqlite.Open("app.db"))
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&IMat{})
	db.AutoMigrate(&State{})

	b := Bridge{
		Db: db,
	}

	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("."))

	router.Handle("/", fs)
	router.HandleFunc("/new-instance", b.NewInstanceHandler)
	router.HandleFunc("/load-instance", b.LoadInstanceHandler)
	router.HandleFunc("/instance/{id}", b.StartInstanceHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
