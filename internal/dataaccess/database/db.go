package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/nguyenhoang711/downloader/internal/configs"
)

type Database interface {
	Delete(table interface{}) *goqu.DeleteDataset
	Dialect() string
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	From(from ...interface{}) *goqu.SelectDataset
	Insert(table interface{}) *goqu.InsertDataset
	Logger(logger goqu.Logger)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ScanStruct(i interface{}, query string, args ...interface{}) (bool, error)
	ScanStructContext(ctx context.Context, i interface{}, query string, args ...interface{}) (bool, error)
	ScanStructs(i interface{}, query string, args ...interface{}) error
	ScanStructsContext(ctx context.Context, i interface{}, query string, args ...interface{}) error
	ScanVal(i interface{}, query string, args ...interface{}) (bool, error)
	ScanValContext(ctx context.Context, i interface{}, query string, args ...interface{}) (bool, error)
	ScanVals(i interface{}, query string, args ...interface{}) error
	ScanValsContext(ctx context.Context, i interface{}, query string, args ...interface{}) error
	Select(cols ...interface{}) *goqu.SelectDataset
	Trace(op string, sqlString string, args ...interface{})
	Update(table interface{}) *goqu.UpdateDataset
}

func InitializeDB(dbConfig configs.Database) (*sql.DB, func(), error) {
	// Format: username:password@tcp(host:port)/dbname?parseTime=true
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Printf("error connecting to the database: %v\n", err)
		return nil, nil, err
	}
	cleanup := func() {
		db.Close()
	}
	return db, cleanup, nil
}

func InitializeGoquDB(db *sql.DB) *goqu.Database {
	return goqu.New("mysql", db)
}
