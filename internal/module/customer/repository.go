package customer

import (
	"context"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"github.com/pewe21/PointOfSale/internal/domain"
	"time"
)

type repository struct {
	db *goqu.Database
}

func NewRepository(db *sql.DB) domain.CustomerRepository {
	return &repository{db: goqu.New("default", db)}
}

func (r repository) Save(ctx context.Context, customer *domain.Customer) error {
	exec := r.db.Insert("customers").Rows(customer).Executor()
	_, err := exec.ExecContext(ctx)
	return err
}

func (r repository) Update(ctx context.Context, customer *domain.Customer, id string) error {
	customer.UpdatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	executor := r.db.Update("customers").Set(customer).Where(goqu.C("id").Eq(id)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (r repository) FindById(ctx context.Context, id string) (customer domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.C("id").Eq(id)).Where(goqu.C("deleted_at").IsNull()).Executor()
	_, err = dataset.ScanStructContext(ctx, &customer)
	return
}

func (r repository) FindByUsername(ctx context.Context, username string) (customer domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.C("username").Eq(username)).Executor()
	_, err = dataset.ScanStructContext(ctx, &customer)
	return
}

func (r repository) FindByEmail(ctx context.Context, email string) (customer domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.C("email").Eq(email)).Executor()
	_, err = dataset.ScanStructContext(ctx, &customer)
	return
}

func (r repository) FindAll(ctx context.Context) (customers []domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.C("deleted_at").IsNull()).Executor()
	err = dataset.ScanStructsContext(ctx, &customers)
	return
}

func (r repository) Delete(ctx context.Context, id string) error {
	executor := r.db.Update("customers").Set(goqu.Record{
		"deleted_at": sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}).Where(goqu.C("id").Eq(id)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
