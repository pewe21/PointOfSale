package supplier

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

func NewRepository(db *sql.DB) domain.SupplierRepository {
	return &repository{db: goqu.New("default", db)}
}

func (r repository) Save(ctx context.Context, supplier *domain.Supplier) error {
	executor := r.db.Insert("suppliers").Rows(supplier).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (r repository) Update(ctx context.Context, supplier *domain.Supplier, id string) error {
	executor := r.db.Update("suppliers").Set(supplier).Set(goqu.Record{
		"updated_at": sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}).Where(goqu.C("id").Eq(id)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (r repository) FindById(ctx context.Context, id string) (supplier domain.Supplier, err error) {
	dataset := r.db.From("suppliers").Where(goqu.C("id").Eq(id)).Executor()
	_, err = dataset.ScanStructContext(ctx, &supplier)
	return
}

func (r repository) FindAll(ctx context.Context) (suppliers []domain.Supplier, err error) {
	dataset := r.db.From("suppliers").Where(goqu.C("deleted_at").IsNull()).Executor()
	err = dataset.ScanStructsContext(ctx, &suppliers)
	return
}

func (r repository) Delete(ctx context.Context, id string) error {
	executor := r.db.Update("suppliers").Set(goqu.Record{
		"deleted_at": sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}).Where(goqu.C("id").Eq(id)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
