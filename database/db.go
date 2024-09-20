package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	path string
	db   *sql.DB
}

func (d *Database) exists() bool {
	_, err := os.Stat(d.path)
	return !os.IsNotExist(err)
}

func (d *Database) SetupDatabase(path string) error {
	d.path = path
	err := d.open()
	if err != nil {
		return err
	} else {
		if !d.exists() {
			fmt.Println("Warning: ", fmt.Sprintf("Database '%s' does not exist.", d.path), "Creating new database.")
			return d.CreateDefaultDB()

		} else {
			fmt.Println("Database exists")
			return nil
		}
	}
}

func (d *Database) open() error {
	var err error
	d.db, err = sql.Open("sqlite3", d.path)

	return err
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) CreateDefaultDB() error {
	createCommand := `
		CREATE TABLE metadata (
        id INTEGER PRIMARY KEY AUTOINCREMENT,

        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

        version TEXT NOT NULL
    	);
	

		CREATE TABLE categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,

		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,

		name TEXT NOT NULL,
		columns TEXT NOT NULL
		);


		CREATE TABLE datatypes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,

		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)		
		`

	_, err := d.db.Exec(createCommand)

	return err
}
