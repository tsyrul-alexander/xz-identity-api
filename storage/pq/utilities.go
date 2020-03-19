package pq

import (
	"database/sql"
	"identity-web-api/storage/column"
	"identity-web-api/storage/condition"
	"identity-web-api/storage/join"
	"identity-web-api/storage/query"
)

const providerName = "postgres"

func (store *Storage)ExecuteSelect(s *query.Select) (*[]query.Row, error) {
	var result *[]query.Row
	var err = store.OpenDb(func(db *sql.DB) error {
		var rows, err = s.GetEntityRows(db)
		result = rows
		return err
	})
	return result, err
}

func (store *Storage)ExecuteInsert(insert *query.Insert) (*sql.Result, error) {
	var result sql.Result
	var err = store.OpenDb(func(db *sql.DB) error {
		var r, err = insert.Execute(db)
		result = r
		return err
	})
	return &result, err
}

func (store *Storage) OpenDb(f func(db *sql.DB) error) error  {
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

//CreateInsert ...
func CreateInsert(tableName string, columnValues map[string]column.ColumnValue) *query.Insert {
	return &query.Insert{TableName: tableName, ColumnValues: columnValues}
}

//CreateSelect ...
func CreateSelect(tableName string, columns *[]*column.TableColumn, joins *[]*join.TableJoin,
		conditions *condition.QueryCondition) *query.Select {
	return &query.Select{TableName: tableName, Columns:columns, Joins:joins, Conditions:conditions}
}

func CreateTableColumn(columnName string, tableName string, alias string) *column.TableColumn {
	return &column.TableColumn{TableName:tableName, ColumnName:columnName, Alias:alias}
}
