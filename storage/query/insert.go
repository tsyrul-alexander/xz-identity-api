package query

import (
	"database/sql"
	"identity-web-api/storage/column"
	"strings"
)

type Insert struct {
	TableName    string
	ColumnValues map[string]column.ColumnValue
	Db           *sql.DB
}
//insert ...
func (i *Insert) Execute(db *sql.DB) (sql.Result, error) {
	var sqlText = i.getInsertSql()
	var result, err = db.Exec(sqlText)
	if err != nil {
		return nil, err
	}
	return result, nil
}
//getInsertSql ...
func (i *Insert) getInsertSql() string {
	var columns []string
	var columnValues []string
	for key, value := range i.ColumnValues {
		columns = append(columns, "\"" + key + "\"")
		columnValues = append(columnValues, value.GetValue())
	}
	var columnsSql = strings.Join(columns, ", ")
	var columnValuesSql = strings.Join(columnValues, ", ")
	return "INSERT INTO \"" + i.TableName + "\"(" + columnsSql + ") VALUES (" + columnValuesSql + ")"
}