package condition

func GetConditionSql(c *QueryCondition) string {
	switch t := (*c).(type) {
	case *GroupQueryCondition:
		return getGroupQueryConditionSql(t)
	case *BinaryQueryCondition:
		return getBinaryQueryConditionSql(t)
	case *ColumnQueryCondition:
		return getColumnQueryConditionSql(t)
	case *ConstantQueryCondition:
		return getConstantQueryConditionSql(t)
	default:
		panic("not implemented")
	}
}

func getConstantQueryConditionSql(c *ConstantQueryCondition) string {
	return c.Value.GetValue()
}

func getColumnQueryConditionSql(c *ColumnQueryCondition) string {
	return "\"" + c.QueryColumn.TableName + "\"." + "\"" + c.QueryColumn.ColumnName + "\""
}

func getGroupQueryConditionSql(c *GroupQueryCondition) string {
	var logicalOperatorSql = getLogicalOperatorSql(c.LogicalOperation)
	var sqlText = "("
	for index, condition := range c.QueryConditions {
		sqlText += GetConditionSql(&condition)
		if index == (len(c.QueryConditions) - 1) {
			sqlText += logicalOperatorSql
		}
	}
	sqlText += ")"
	return sqlText
}

func getBinaryQueryConditionSql(c *BinaryQueryCondition) string {
	var leftExpressionSql = GetConditionSql(&c.LeftCondition)
	var rightExpressionSql = GetConditionSql(&c.RightCondition)
	return getComparisonTypeSql(c.ComparisonType, leftExpressionSql, rightExpressionSql)
}

func getComparisonTypeSql(c ComparisonType, leftExpression string, rightExpression string) string {
	switch c {
	case ComparisonTypeEqual:
		return leftExpression + " = " + rightExpression
	case ComparisonTypeNotEqual:
		return leftExpression + " != " + rightExpression
	default:
		panic("not implemented")
	}
}

func getLogicalOperatorSql(l LogicalOperation) string {
	switch l {
	case And:
		return "AND"
	case Or:
		return "OR"
	default:
		panic("not implemented")
	}
}
