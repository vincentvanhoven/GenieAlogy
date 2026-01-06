package repositories

import (
	"GenieAlogy/database/migrations"
	"database/sql"
	"errors"
	"log"

	_ "database/sql"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

type DatabaseRepository struct {
	DB *sql.DB
}

var DatabaseRepo = &DatabaseRepository{}

func (repo *DatabaseRepository) Create(filePath string) error {
	err := repo.Fetch(filePath)
	if err != nil {
		return err
	}

	err = repo.Update()

	if err != nil {
		return err
	}

	return nil
}

func (repo *DatabaseRepository) Fetch(filePath string) error {
	if repo.DB != nil {
		err := repo.DB.Close()

		if err != nil {
			log.Fatal(err)
		}
	}

	db, err := sql.Open("sqlite", filePath)

	if err != nil {
		return err
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	
	repo.DB = db
	return nil
}

func (repo *DatabaseRepository) Update() error {
	if repo.DB == nil {
		return errors.New("Database not initialized")
	}

	goose.SetBaseFS(migrations.Migrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		return err
	}

	if err := goose.Up(repo.DB, "."); err != nil {
		return err
	}

	return nil
}
