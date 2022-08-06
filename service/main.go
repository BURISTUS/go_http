package main

import (
	"log"
	"net/http"
	"test/db"
	"test/routes"
)

func main() {
	database, err := db.NewDatabase("localhost:6379")
	if err != nil {
		log.Fatalf("Failed to connect to redis: %s", err.Error())
	}

	//http.HandleFunc("/", handler)
	//log.Fatal(http.ListenAndServe(":8080", nil))
	mux := http.NewServeMux()
	mux.HandleFunc("/test1", routes.Test1(database))
	//mux.HandleFunc("/test2", test2(database))
	//mux.HandleFunc("/test3", test1(database))

	err2 := http.ListenAndServe(":4000", mux)
	log.Fatal(err2)
}
