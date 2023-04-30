package sqlnull

import "database/sql"

func FromInt32Ptr(i *int32) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{}
	} else {
		return sql.NullInt32{Int32: *i, Valid: true}
	}
}

func FromString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: true}
}
