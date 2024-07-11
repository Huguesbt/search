package search

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

var DbObj DbEntity

type DbEntity struct {
	DriverName     string
	DataSourceName string
	db             *sql.DB
}

func (d *DbEntity) connectDB() error {
	if db, err := sql.Open(d.DriverName, d.DataSourceName); err != nil {
		return err
	} else {
		d.db = db
		return nil
	}
}

func (d *DbEntity) GetDb() *sql.DB {
	return d.db
}

func InitDB() error {
	DbObj.DriverName = "sqlite"
	DbObj.DataSourceName = "search.db"

	err := DbObj.connectDB()
	if err != nil {
		return err
	}

	query := `
	CREATE TABLE IF NOT EXISTS documents (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		text TEXT NOT NULL,
		tags TEXT
	);
	`
	_, err = DbObj.db.Exec(query)
	return err
}
