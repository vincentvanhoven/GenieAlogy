package main

import (
	"database/sql"
	"log"

	"GenieAlogy/database/seeders"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "test-database.geniealogy")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(db, "database/migrations"); err != nil {
		log.Fatal(err)
	}

	seeders.RunPeopleSeeder(db)
	seeders.RunFamilySeeder(db)
}
