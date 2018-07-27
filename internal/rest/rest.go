package rest

import (
	"encoding/json"
	"github.com/dnvriend/postgres-go-test/internal/dao"
	"github.com/gorilla/mux"
	"net/http"
)

func respond(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func getParam(r *http.Request, param string) string {
	params := mux.Vars(r)
	return params[param]
}

func GetActors(w http.ResponseWriter, r *http.Request) {
	respond(w, dao.GetActors())
}
