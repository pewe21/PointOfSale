package domain

import (
	"context"
	"database/sql"
	"time"
)

type Supplier struct {
	Id        string       `json:"id" db:"id" goqu:"skipinsert"`
	Name      string       `json:"name" db:"name"`
	Email     string       `json:"email" db:"email"`
	Address   string       `json:"address" db:"address"`
	Phone     string       `json:"phone" db:"phone"`
	CreatedAt time.Time    `json:"created_at" db:"created_at" goqu:"skipinsert"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

type SupplierRepository interface {
	Save(ctx context.Context, supplier *Supplier) error
	Update(ctx context.Context, supplier *Supplier, id string) error
	FindAll(ctx context.Context) (suppliers []Supplier, err error)
	FindById(ctx context.Context, id string) (supplier Supplier, err error)
	Delete(ctx context.Context, id string) error
}

type SupplierService interface {
	Save(ctx context.Context, req *Supplier) (id string, err error)
}
