package join

import (
	"identity-web-api/storage/condition"
)

type TableJoin struct {
	JoinTableName string
	Type          Type
	Conditions    condition.QueryCondition
}
