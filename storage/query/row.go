package query

import "github.com/google/uuid"

type Row map[string]interface{}

func (r *Row) GetUuidValue(key string) uuid.UUID {
	return (*r)[key].(uuid.UUID)
}


