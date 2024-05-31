package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"ur_shortner/db"
	"ur_shortner/repositories"
	"ur_shortner/server"
)

func main(){
	config := db.LoadConfig()
    dbconn := db.InitDB(config)

	repositories.ConnectionPSQL(dbconn)
	
    r := mux.NewRouter()
    r.HandleFunc("/{shortURL}", server.RedirectHandler).Methods("GET")
    r.HandleFunc("/create", server.CreateHandler).Methods("POST")

    http.Handle("/", r)
    log.Println("Server started at localhost:8080")
    log.Fatal(http.ListenAndServe("localhost:8080", nil))

}