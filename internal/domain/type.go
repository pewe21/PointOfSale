package domain

import (
	"database/sql"
	"time"
)

type Type struct {
	Id          string       `json:"id" db:"id" goqu:"skipinsert"`
	Name        string       `json:"name" db:"name"`
	Description string       `json:"description" db:"description"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at" goqu:"skipinsert"`
	UpdatedAt   sql.NullTime `json:"updated_at" db:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at" db:"deleted_at"`
}
