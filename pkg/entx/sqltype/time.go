package sqltype

import "entgo.io/ent/dialect/sql"

func NewNullTime() *sql.NullTime {
	return &sql.NullTime{}
}
