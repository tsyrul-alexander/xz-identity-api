package pq

import (
	"database/sql"
	"github.com/tsyrul-alexander/go-query-builder/query"
)

const providerName = "postgres"

func (store *DataStorage)ExecuteSelect(s *query.Select) (*[]query.Row, error) {
	var result *[]query.Row
	var err = store.OpenDb(func(db *sql.DB) error {
		var rows, err = s.Execute(db)
		result = rows
		return err
	})
	return result, err
}

func (store *DataStorage)ExecuteInsert(insert *query.Insert) (*sql.Result, error) {
	var result sql.Result
	var err = store.OpenDb(func(db *sql.DB) error {
		var r, err = insert.Execute(db)
		result = r
		return err
	})
	return &result, err
}

func (store *DataStorage) OpenDb(f func(db *sql.DB) error) error  {
	var db, err = sql.Open(providerName, store.Config.ConnectionString)
	if err != nil {
		return err
	}
	if err = f(db); err != nil {
		return err
	}
	if err =  db.Close(); err != nil {
		return err
	}
	return nil
}

