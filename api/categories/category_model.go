package categories

import (
	"database/sql"
	"time"
)

type Category struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
	Title       string
	Path        string
	DefaultUnit string
}
