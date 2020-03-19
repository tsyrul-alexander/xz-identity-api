package column

import (
	"strconv"
)

//IntColumnValue ...
type IntColumnValue struct {
	Value int
}

func (columnValue *IntColumnValue)GetValue() string {
	return "'" + strconv.Itoa(columnValue.Value) + "'"
}
func CreateIntColumnValue(value int) *IntColumnValue {
	return &IntColumnValue{Value:value}
}
