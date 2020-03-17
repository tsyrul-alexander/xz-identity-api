package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"identity-web-api/model"
	"strings"
)

const providerName = "postgres"
//PQStorage ...
type PQStorage struct {
	Config *Config
}

func CreatePQStore(config *Config) *PQStorage {
	return &PQStorage{Config:config}
}
//createUser ...
func (store *PQStorage) CreateUser(user *model.User) error {
	var dictionary = map[string]string{}
	dictionary["Id"] = user.ID.String()
	dictionary["Name"] = user.Name
	var _, err = store.insert("User", dictionary)
	return err
}
//getInsertSql ...
func getInsertSql(tableName string, values map[string]string) string {
	var columns []string
	var columnValues []string
	for key, value := range values {
		columns = append(columns, "\"" + key + "\"")
		columnValues = append(columnValues, "'" + value + "'")
	}
	var columnsSql = strings.Join(columns, ", ")
	var columnValuesSql = strings.Join(columnValues, ", ")
	return "INSERT INTO \"" + tableName + "\"(" + columnsSql + ") VALUES (" + columnValuesSql + ")"
}
//insert ...
func (store *PQStorage) insert(tableName string, values map[string]string) (sql.Result, error) {
	var sqlText = getInsertSql(tableName, values)
	var resultFunc sql.Result
	var errFunc = store.openDb(func(db *sql.DB) error {
		var result, err = db.Exec(sqlText)
		if err != nil {
			return err
		}
		resultFunc = result
		return nil
	})
	if errFunc != nil {
		return nil, errFunc
	}
	return resultFunc, nil
}
func (store *PQStorage) openDb(f func(db *sql.DB) error) error  {
	var db, err = sql.Open(providerName, store.Config.ConnectionString)
	if err != nil {
		return err
	}
	err = f(db)
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}
	return nil
}