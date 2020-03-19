package column

type StringColumnValue struct {
	Value string
}

func (columnValue *StringColumnValue)GetValue() string {
	return "'" + columnValue.Value + "'"
}

func CreateStringColumnValue(value string) *StringColumnValue {
	return &StringColumnValue{Value:value}
}