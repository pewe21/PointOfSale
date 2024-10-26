package domain

import (
	"context"
	"database/sql"
	"time"
)

type CustomerRole struct {
	Id         string       `json:"id" db:"id" goqu:"skipinsert,skipupdate"`
	RoleId     string       `json:"role_id" db:"role_id"`
	CustomerId string       `json:"customer_id" db:"customer_id"`
	CreatedAt  time.Time    `json:"created_at" db:"created_at"`
	DeletedAt  sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

type CustomerRolesRepository interface {
	Create(ctx context.Context, roleId string, customerId string) (err error)
	Update(ctx context.Context, roleId string, customerId string) (err error)
}
