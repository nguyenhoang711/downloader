package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/nguyenhoang711/downloader/internal/configs"
)

func InitializeDB(dbConfig configs.Database) (db *sql.DB, cleanup func(), err error) {
	// Format: username:password@tcp(host:port)/dbname?parseTime=true
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Printf("error connecting to the database: %v\n", err)
		return nil, nil, err
	}
	cleanup = func() {
		db.Close()
	}
	return db, cleanup, nil
}

func InitializeGoquDB(db *sql.DB) *goqu.Database {
	return goqu.New("mysql", db)
}
