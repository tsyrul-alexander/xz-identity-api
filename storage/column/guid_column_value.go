package column

import "github.com/google/uuid"

type GuidColumnValue struct {
	Value *uuid.UUID
}

func (columnValue *GuidColumnValue)GetValue() string {
	return "'" + columnValue.Value.String() + "'"
}
func CreateGuidColumnValue(value *uuid.UUID) *GuidColumnValue {
	return &GuidColumnValue{Value:value}
}
