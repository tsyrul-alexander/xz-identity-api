package pq

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/tsyrul-alexander/go-query-builder/core/column"
	"github.com/tsyrul-alexander/go-query-builder/core/condition"
	"github.com/tsyrul-alexander/go-query-builder/core/parameter"
	"github.com/tsyrul-alexander/go-query-builder/pq/builder"
	"github.com/tsyrul-alexander/go-query-builder/query"
	"github.com/tsyrul-alexander/xz-identity-api/model"
	"github.com/tsyrul-alexander/xz-identity-api/model/data"
	"github.com/tsyrul-alexander/xz-identity-api/storage"
)
//DataStorage ...
type DataStorage struct {
	Config *storage.Config
}

const (
	UserTableName = "User"
	DefaultIdentityTableName = "DefaultIdentity"
)
//CreatePQStore ...
func CreateStore(config *storage.Config) *DataStorage {
	return &DataStorage{Config: config}
}

//CreateUser ...
func (store *DataStorage) CreateUser(user *data.User, roles ...model.UserRole) error {
	var userInsert = getCreateUserInsert(user)
	var credentialsUserInsert = getCreateUserCredentialsInsert(user)

	var err = store.OpenDb(func(db *sql.DB) error {
		var tx, err = db.Begin()
		if err != nil {
			return err
		}
		if _, err = credentialsUserInsert.ExecuteTx(tx); err != nil {
			return tx.Rollback()
		}
		if _, err = userInsert.ExecuteTx(tx); err != nil {
			return tx.Rollback()
		}
		if err := store.CreateUserRole(user.ID, tx, roles...); err != nil {
			return tx.Rollback()
		}
		return tx.Commit()
	})
	return err
}

//GetUserByLogin ...
func (store *DataStorage)GetUserByLogin(login string) (*data.User, error) {
	return store.getUser("Login", login)
}

func (store *DataStorage)CreateUserRole(userId uuid.UUID, tx *sql.Tx, roles ...model.UserRole) error {
	var rolesIds, err = store.getRolesId(roles)
	if err != nil {
		return err
	}
	for _, r := range rolesIds {
		var insert = getUserRoleInsert(userId, r)
		if _, err := insert.ExecuteTx(tx); err != nil {
			return err
		}
	}
	return nil
}

func getUserRoleInsert(userId uuid.UUID, roleId uuid.UUID) *query.Insert {
	var columnValues = column.ValueList{}
	columnValues["UserId"] = parameter.CreateGuidParameter(userId)
	columnValues["RoleId"] = parameter.CreateGuidParameter(roleId)
	return builder.CreateInsert("UserRole", &columnValues)
}

func (store *DataStorage)GetUserById(id uuid.UUID) (*data.User, error) {
	return store.getUser("Id", id)
}

func (store *DataStorage)GetUserRoles(id uuid.UUID) ([]model.UserRole, error)  {
	var s = builder.CreateSelect("Role")
	s.AddTableColumn("Role", "Code").Alias = "RoleCode"
	s.AddLeftJoin("UserRole", "RoleId", "Role", "Id")
	s.AddColumnValueCondition(condition.ComparisonTypeEqual, "UserRole", "UserId", id)
	var rows, err = store.ExecuteSelect(s)
	if err != nil {
		return nil, err
	}
	var roles []model.UserRole
	for _, row := range *rows {
		roles = append(roles, model.UserRole(row.GetIntValue("RoleCode")))
	}
	return roles, nil
}

func (store *DataStorage) getUser(filterColumnName string, filterColumnValue interface{}) (*data.User, error) {
	var s = getUserQuery(filterColumnName, filterColumnValue)
	var rows, err = store.ExecuteSelect(s)
	if err != nil {
		return nil, err
	}
	if rows == nil || len(*rows) == 0 {
		return nil, nil
	}
	var row = (*rows)[0]
	var user = parseUserResponse(row)
	return user, nil
}

func (store *DataStorage) getRolesId(roles []model.UserRole) ([]uuid.UUID, error) {
	var s = builder.CreateSelect("Role")
	s.AddTableColumn("Role", "Id")
	s.AddColumnValueCondition(condition.ComparisonTypeEqual, "Role", "Code", roles)
	var rows, err = store.ExecuteSelect(s)
	if err != nil {
		return nil, err
	}
	var ids []uuid.UUID
	for _, r := range *rows {
		ids = append(ids, r.GetUuidValue("RoleId"))
	}
	return ids, nil
}

func parseUserResponse(row query.Row) *data.User {
	return &data.User{
		ID:           row.GetUuidValue("UserId"),
		Name:         row.GetStringValue("UserName"),
		IdentityType: data.IdentityType(row.GetIntValue("UserIdentityType")),
		DefaultIdentity: data.DefaultIdentity{
			ID:       row.GetUuidValue(DefaultIdentityTableName + "Id"),
			Login:    row.GetStringValue(DefaultIdentityTableName + "Login"),
			Password: model.HashPassword(row.GetStringValue(DefaultIdentityTableName + "Password")),
		},
	}
}

func getUserQuery(columnFilterName string, columnFilterValue interface{}) *query.Select {
	var s = builder.CreateSelect(UserTableName)
	s.RowCount = 1
	setUserColumns(s)
	setUserJoins(s)
	setUserCondition(s, columnFilterName, columnFilterValue)
	return s
}

func getCreateUserInsert(user *data.User) *query.Insert {
	var columnValues = column.ValueList{}
	columnValues["Id"] = parameter.CreateGuidParameter(user.ID)
	columnValues["Name"] = parameter.CreateStringParameter(user.Name)
	columnValues["IdentityType"] = parameter.CreateIntParameter(int(user.IdentityType))
	columnValues["DefaultIdentityId"] = parameter.CreateGuidParameter(user.DefaultIdentity.ID)
	return builder.CreateInsert(UserTableName, &columnValues)
}

func getCreateUserCredentialsInsert(user *data.User) *query.Insert {
	if user.IdentityType == data.IdentityTypeDefault {
		return getCreateUserDefaultCredentialsInsert(user)
	}
	panic("not implemented")
}

func getCreateUserDefaultCredentialsInsert(user *data.User) *query.Insert {
	var columnValues = column.ValueList{}
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

func setUserCondition(s *query.Select, columnFilterName string, columnFilterValue interface{}) {
	s.AddColumnValueCondition(condition.ComparisonTypeEqual, DefaultIdentityTableName, columnFilterName, columnFilterValue)
}