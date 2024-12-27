package search

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

var DbObj DbEntity

type DbEntity struct {
	driverName     string
	dataSourceName string
	db             *sql.DB
}

func (d *DbEntity) connectDB() error {
	if db, err := sql.Open(d.driverName, d.dataSourceName); err != nil {
		return err
	} else {
		d.db = db
		return nil
	}
}

func (d *DbEntity) GetDb() *sql.DB {
	return d.db
}

func (d *DbEntity) SetDriverName(driverName string) {
	d.driverName = driverName
}

func (d *DbEntity) SetDataSourceName(dataSourceName string) {
	d.dataSourceName = dataSourceName
}

func InitDB(driverName string, dataSourceName string) error {
	DbObj.SetDriverName(driverName)
	DbObj.SetDataSourceName(dataSourceName)
	err := DbObj.connectDB()
	if err != nil {
		return err
	}

	query := `
	CREATE TABLE IF NOT EXISTS documents (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		text TEXT NOT NULL,
		description TEXT NOT NULL,
		notes TEXT NOT NULL,
		tags TEXT
	);
	`
	_, err = DbObj.db.Exec(query)
	return err
}
