package main

import (
	"crypto/subtle"
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
	r.Use(loggingMiddleware)
	r.Use(basicAuthMiddleware)
	log.Print("Running Server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func basicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if BasicAuth(w, r, "user", "pass", "Provide user name and password") {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(401)
		}
	})
}

func BasicAuth(w http.ResponseWriter, r *http.Request, username, password, realm string) bool {

	user, pass, ok := r.BasicAuth()

	if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
		w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
		w.WriteHeader(401)
		w.Write([]byte("Unauthorised.\n"))
		return false
	}

	return true
}
