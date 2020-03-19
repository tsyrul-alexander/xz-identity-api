package condition

type GroupQueryCondition struct {
	LogicalOperation LogicalOperation
	QueryConditions []QueryCondition
}

func (c *GroupQueryCondition) getType() Type  {
	return Group
}
