package main

import (
	"github.com/dnvriend/postgres-go-test/internal/dao"
	"github.com/dnvriend/postgres-go-test/internal/rest"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	dao.Connect()
	r := mux.NewRouter()
	r.HandleFunc("/api/actors", rest.GetActors).Methods("GET")
	log.Print("Running Server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
