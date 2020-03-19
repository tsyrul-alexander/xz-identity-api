package condition

import "identity-web-api/storage/column"

type ColumnQueryCondition struct {
	QueryColumn *column.TableColumn
}

func (c *ColumnQueryCondition) getType() Type  {
	return Column
}