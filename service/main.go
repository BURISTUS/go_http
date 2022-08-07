package main

import (
	"flag"
	"log"
	"net/http"
	"test/db"
	"test/routes"
)

func main() {
	host := flag.String("h", "localhost", "There should be host")
	port := flag.String("p", "6379", "There should be port")
	database, databaseErr := db.NewDatabase(*host + ":" + *port)

	if databaseErr != nil {
		log.Fatalf("Failed to connect to redis: %s", databaseErr.Error())
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/test1", routes.JsonSum(database))
	mux.HandleFunc("/test2", routes.GetHashFromJson)
	mux.HandleFunc("/test3", routes.TcpClient)

	err2 := http.ListenAndServe(":4000", mux)
	log.Fatal(err2)
}
