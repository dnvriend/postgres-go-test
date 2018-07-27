package dao

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Actor struct {
	ActorID   int
	Firstname string
	Lastname  string
}

var _db *sql.DB

func Close() {
	err := _db.Close()
	checkErr(err)
}

func Connect() {
	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	_db = db
	checkErr(err)
}

func GetActors() []Actor {
	var actors []Actor
	rows, err := _db.Query("SELECT actor_id, first_name, last_name FROM actor")
	checkErr(err)
	for rows.Next() {
		var actor Actor
		err = rows.Scan(&actor.ActorID, &actor.Firstname, &actor.Lastname)
		checkErr(err)
		actors = append(actors, actor)
	}
	return actors
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
