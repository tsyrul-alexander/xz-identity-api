package pq

import (
	"database/sql"
	_ "github.com/lib/pq"
	"identity-web-api/model"
	"identity-web-api/storage"
	"identity-web-api/storage/column"
	"identity-web-api/storage/condition"
	"identity-web-api/storage/join"
	"identity-web-api/storage/query"
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
	var conditions = getUserCondition(login)
	var s = CreateSelect(UserTableName, getUserColumns(), getUserJoins(), &conditions)
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
	var dictionary = map[string]column.ColumnValue{}
	dictionary["Id"] = column.CreateGuidColumnValue(&user.ID)
	dictionary["Name"] = column.CreateStringColumnValue(user.Name)
	dictionary["IdentityType"] = column.CreateIntColumnValue(int(user.IdentityType))
	dictionary["DefaultIdentityId"] = column.CreateGuidColumnValue(&user.DefaultIdentity.ID)
	return CreateInsert(UserTableName, dictionary)
}

func getCreateUserCredentialsInsert(user *model.User) *query.Insert {
	if user.IdentityType == model.IdentityTypeDefault {
		return getCreateUserDefaultCredentialsInsert(user)
	}
	panic("not implemented")
}

func getCreateUserDefaultCredentialsInsert(user *model.User) *query.Insert {
	var dictionary = map[string]column.ColumnValue{}
	dictionary["Id"] = column.CreateGuidColumnValue(&user.DefaultIdentity.ID)
	dictionary["Login"] = column.CreateStringColumnValue(user.DefaultIdentity.Login)
	dictionary["Password"] = column.CreateStringColumnValue(user.DefaultIdentity.Password.String())
	return CreateInsert(DefaultIdentityTableName, dictionary)
}

func getUserJoins() *[]*join.TableJoin {
	return &[]*join.TableJoin {
		{
			Type:          join.InnerJoin,
			JoinTableName: DefaultIdentityTableName,
			Conditions: &condition.BinaryQueryCondition{
				ComparisonType: condition.ComparisonTypeEqual,
				LeftCondition: &condition.ColumnQueryCondition{
					QueryColumn: CreateTableColumn("DefaultIdentityId", UserTableName),
				},
				RightCondition: &condition.ColumnQueryCondition{
					QueryColumn: CreateTableColumn("Id", DefaultIdentityTableName),
				},
			},
		},
	}
}

func getUserColumns() *[]*column.TableColumn {
	return &[]*column.TableColumn {
		CreateTableColumn("Id", UserTableName),
		CreateTableColumn("Name", UserTableName),
		CreateTableColumn("IdentityType", UserTableName),
		CreateTableColumn("Id", DefaultIdentityTableName),
		CreateTableColumn("Login", DefaultIdentityTableName),
		CreateTableColumn("Password", DefaultIdentityTableName),
	}
}

	func getUserCondition(login string) condition.QueryCondition {
		return &condition.BinaryQueryCondition{
			ComparisonType: condition.ComparisonTypeEqual,
			LeftCondition: &condition.ColumnQueryCondition{
				QueryColumn: &column.TableColumn{
					ColumnName: "Login",
					TableName:  "DefaultIdentity",
				},
			},
			RightCondition: &condition.ConstantQueryCondition{
				Value: column.CreateStringColumnValue(login),
			},
		}
	}