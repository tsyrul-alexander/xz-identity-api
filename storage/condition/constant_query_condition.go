package condition

import "identity-web-api/storage/column"

type ConstantQueryCondition struct {
	Value column.ColumnValue
}

func (c *ConstantQueryCondition) getType() Type  {
	return Const
}
