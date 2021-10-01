package main

import "github.com/teng231/workshop/db"

var (
	ws = &Workshop{}
)

func main() {
	syncDb()
	d := &db.DB{}
	d.ConnectDb("root:my-secret-pw@tcp(localhost:3306)", "test")
	ws.db = d
	startServer()
}

func syncDb() {
	d := &db.DB{}
	d.ConnectDb("root:my-secret-pw@tcp(localhost:3306)", "test")
	d.CreateDb() // create or sync db
}
