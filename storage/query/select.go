package query

import (
	"database/sql"
	"identity-web-api/storage/column"
	"identity-web-api/storage/condition"
	"identity-web-api/storage/join"
)

type Select struct {
	TableName string
	Columns *[]*column.TableColumn
	Joins *[]*join.TableJoin
	Conditions *condition.QueryCondition
}

//GetEntityRows ...
func (s *Select) GetEntityRows(db *sql.DB) (*[]Row, error) {
	var sqlText = s.GenerateSql()
	var sqlRows, err = db.Query(sqlText)
	if err != nil {
		return nil, err
	}
	var sqlColumns, _ = sqlRows.Columns()
	var rows []Row
	for sqlRows.Next() {
		var row, err = parseSqlRows(sqlRows, sqlColumns)
		if err != nil {
			return nil, err
		}
		rows = append(rows, *row)
	}
	return &rows, err
}

//GenerateSql ...
func (s *Select) GenerateSql() string {
	return "SELECT " + getColumnsSql(s.Columns) + "\nFROM " + "\"" + s.TableName + "\"\n" + getJoinSql(s.Joins) + getWhereSql(s.Conditions)
}

func parseSqlRows(sqlRows *sql.Rows, sqlColumns []string) (*Row, error) {
	var row = Row{}
	var sqlRowValues = make([]interface{}, len(sqlColumns))
	var sqlRowPointValues = make([]interface{}, len(sqlColumns))
	for i, _ := range sqlRowValues {
		sqlRowPointValues[i] = &sqlRowValues[i]
	}
	var err = sqlRows.Scan(sqlRowPointValues...)
	if err != nil {
		return nil, err
	}
	for index, sqlColumn := range sqlColumns {
		row[sqlColumn] = sqlRowValues[index]
	}
	return &row, nil
}

func getJoinSql(joins *[]*join.TableJoin) string {
	if len(*joins) == 0 {
		return ""
	}
	var sqlText = ""
	for _, j := range *joins {
		sqlText +=  getJoinTypeSql(j.Type) + " JOIN  \"" + j.JoinTableName + "\" ON " + getConditionsSql(&j.Conditions) + "\n"
	}
	return sqlText
}

func getWhereSql(condition *condition.QueryCondition) string {
	if condition == nil {
		return ""
	}
	return "WHERE " + getConditionsSql(condition) + "\n"
}

func getConditionsSql(c *condition.QueryCondition) string {
	return condition.GetConditionSql(c)
}



func getColumnsSql(columns *[]*column.TableColumn) string {
	var sqlText = ""
	var columnCount = len(*columns)
	for i, c := range *columns {
		sqlText += getTableColumnSql(c)
		if i != (columnCount - 1) {
			sqlText += ", "
		}
	}
	return sqlText
}

func getTableColumnSql(tableColumn *column.TableColumn) string {
	var columnSql = "\"" + tableColumn.ColumnName + "\""
	var alias = tableColumn.Alias
	if alias == "" {
		alias = tableColumn.TableName + tableColumn.ColumnName
	}
	if tableColumn.TableName != "" {
		return "\"" + tableColumn.TableName + "\"." + columnSql + " AS " + alias
	}
	return columnSql
}

func getJoinTypeSql(t join.Type) string {
	switch t {
	case join.LeftJoin:
		return "LEFT"
	case join.InnerJoin:
		return "INNER"
	default:
		panic("not implemented")
	}
}

