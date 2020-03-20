package pq

import (
	"database/sql"
	_ "github.com/lib/pq"
	"identity-web-api/model"
	"identity-web-api/storage"
	"t-storage/core/condition"
	"t-storage/core/parameter"
	"t-storage/pq/builder"
	"t-storage/query"
)
//Storage ...
type Storage struct {
	Config *storage.Config
}

const (
	UserTableName = "User"
	DefaultIdentityTableName = "DefaultIdentity"
)
//CreatePQStore ...
func CreateStore(config *storage.Config) *Storage {
	return &Storage{Config: config}
}

//CreateUser ...
func (store *Storage) CreateUser(user *model.User) error {
	var userInsert = getCreateUserInsert(user)
	var credentialsUserInsert = getCreateUserCredentialsInsert(user)
	var err = store.OpenDb(func(db *sql.DB) error {
		var tx, err = db.Begin()
		if err != nil {
			return err
		}
		if _, err = credentialsUserInsert.Execute(db); err != nil {
			return tx.Rollback()
		}
		if _, err = userInsert.Execute(db); err != nil {
			return tx.Rollback()
		}
		return tx.Commit()
	})
	return err
}

//GetUser ...
func (store *Storage) GetUser(login string) (*model.User, error) {
	var s = builder.CreateSelect(UserTableName)
	setUserColumns(s)
	setUserJoins(s)
	setUserCondition(s, login)
	var rows, err = store.ExecuteSelect(s)
	if err != nil {
		return nil, err
	}
	if len(*rows) == 0 {
		return nil, nil
	}
	var row = (*rows)[0]
	var user = &model.User{
		ID:              row.GetUuidValue("Id"),
		Name:            "",
		IdentityType:    0,
		DefaultIdentity: model.DefaultIdentity{},
	}
	return user, nil
}

func getCreateUserInsert(user *model.User) *query.Insert {
	var columnValues = query.ColumnValueList{}
	columnValues["Id"] = parameter.CreateGuidParameter(user.ID)
	columnValues["Name"] = parameter.CreateStringParameter(user.Name)
	columnValues["IdentityType"] = parameter.CreateIntParameter(int(user.IdentityType))
	columnValues["DefaultIdentityId"] = parameter.CreateGuidParameter(user.DefaultIdentity.ID)
	return builder.CreateInsert(UserTableName, &columnValues)
}

func getCreateUserCredentialsInsert(user *model.User) *query.Insert {
	if user.IdentityType == model.IdentityTypeDefault {
		return getCreateUserDefaultCredentialsInsert(user)
	}
	panic("not implemented")
}

func getCreateUserDefaultCredentialsInsert(user *model.User) *query.Insert {
	var columnValues = query.ColumnValueList{}
	columnValues["Id"] = parameter.CreateGuidParameter(user.DefaultIdentity.ID)
	columnValues["Login"] = parameter.CreateStringParameter(user.DefaultIdentity.Login)
	columnValues["Password"] = parameter.CreateStringParameter(user.DefaultIdentity.Password.String())
	return builder.CreateInsert(DefaultIdentityTableName, &columnValues)
}

func setUserJoins(s *query.Select) {
	s.AddLeftJoin(DefaultIdentityTableName, "Id", UserTableName, "DefaultIdentityId")
}

func setUserColumns(s *query.Select) {
	s.AddTableColumn(UserTableName, "Id")
	s.AddTableColumn(UserTableName, "Name")
	s.AddTableColumn(UserTableName, "IdentityType")
	s.AddTableColumn(DefaultIdentityTableName, "Id")
	s.AddTableColumn(DefaultIdentityTableName, "Login")
	s.AddTableColumn(DefaultIdentityTableName, "Password")
}

func setUserCondition(s *query.Select, login string) {
	s.AddColumnValueCondition(condition.ComparisonTypeEqual, DefaultIdentityTableName, "Login", login)
}