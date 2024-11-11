package customer_roles

import (
	"context"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"github.com/pewe21/PointOfSale/internal/domain"
)

type customerRolesRepository struct {
	db *goqu.Database
}

func NewCustomerRolesRepository(db *sql.DB) domain.CustomerRolesRepository {
	return &customerRolesRepository{db: goqu.New("default", db)}
}

func (r customerRolesRepository) Create(ctx context.Context, roleId string, customerId string) (err error) {
	executor := r.db.Insert("customer_roles").Rows(goqu.Ex{
		"customer_id": customerId,
		"role_id":     roleId,
	}).Executor()
	_, err = executor.ExecContext(ctx)
	return
}

func (r customerRolesRepository) Update(ctx context.Context, roleId string, customerId string) (err error) {
	executor := r.db.Update("customer_roles").Set(goqu.Ex{
		"role_id": roleId,
	}).Where(goqu.C("customer_id").Eq(customerId)).Executor()
	_, err = executor.ExecContext(ctx)
	return
}
