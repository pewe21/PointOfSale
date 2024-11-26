package brand

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

func NewRepository(db *sql.DB) domain.BrandRepository {

	return &repository{db: goqu.New("default", db)}
}

func (r repository) Save(ctx context.Context, brand *domain.Brand) error {
	executor := r.db.Insert("brands").Rows(brand).Executor()
	_, err := executor.ExecContext(ctx)

	return err
}

func (r repository) Update(ctx context.Context, brand *domain.Brand, id string) error {
	brand.UpdatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	executor := r.db.Update("brands").Set(brand).Where(goqu.C("id").Eq(id)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (r repository) FindById(ctx context.Context, id string) (brand domain.Brand, err error) {
	dataset := r.db.From("brands").Where(goqu.C("id").Eq(id)).Where(goqu.C("deleted_at").IsNull()).Executor()
	_, err = dataset.ScanStructContext(ctx, &brand)

	return
}

func (r repository) FindAll(ctx context.Context) (brands []domain.Brand, err error) {
	dataset := r.db.From("brands").Where(goqu.C("deleted_at").IsNull()).Executor()
	err = dataset.ScanStructsContext(ctx, &brands)
	return
}

func (r repository) Delete(ctx context.Context, id string) error {
	executor := r.db.Update("brands").Set(goqu.Record{
		"deleted_at": sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}).Where(goqu.C("id").Eq(id)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
