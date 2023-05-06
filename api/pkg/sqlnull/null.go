package sqlnull

import (
	"database/sql"
	"time"
)

func FromInt32(i int32) sql.NullInt32 {
	return sql.NullInt32{Int32: i, Valid: true}
}
func FromInt32Ptr(i *int32) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{Int32: *i, Valid: true}
}

func FromInt32Invalidable(i int32) sql.NullInt32 {
	if i == 0 {
		return sql.NullInt32{}
	}
	return sql.NullInt32{Valid: true, Int32: i}
}

func FromTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: true}
}

func FromString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: true}
}

func FromStringInvalidatable(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}

func FromFloat64(f float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: f, Valid: true}
}
