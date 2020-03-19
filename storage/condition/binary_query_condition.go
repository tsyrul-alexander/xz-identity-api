package condition

type BinaryQueryCondition struct {
	ComparisonType ComparisonType
	LeftCondition  QueryCondition
	RightCondition QueryCondition
}

func (c *BinaryQueryCondition) getType() Type  {
	return Binary
}