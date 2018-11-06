package main

import (
	"log"

	"github.com/pisensor/server/internal/dbfactory"
	"github.com/pisensor/server/internal/endpoints"

	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	db, err := dbfactory.NewDB(dbfactory.SQLite)
	if err != nil {
		log.Fatalf("Error getting database connection: %s", err.Error())
	}

	m.Group("/api", func(r martini.Router) {
		endpoints.RegisterReadingsEndpoints(m, db)
	})

	log.Printf("Running REST API ...")

	m.Run()
}
