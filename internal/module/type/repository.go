package _type

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

func NewRepository(db *sql.DB) domain.TypeRepository {
	return &repository{db: goqu.New("default", db)}
}

func (r repository) Save(ctx context.Context, _type *domain.Type) error {
	executor := r.db.Insert("types").Rows(_type).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (r repository) Update(ctx context.Context, _type *domain.Type, id string) error {
	_type.UpdatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	executor := r.db.Update("types").Set(_type).Where(goqu.C("id").Eq(id)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (r repository) FindById(ctx context.Context, id string) (_type domain.Type, err error) {
	dataset := r.db.From("types").Where(goqu.C("id").Eq(id)).Executor()
	_, err = dataset.ScanStructContext(ctx, &_type)
	return
}

func (r repository) FindAll(ctx context.Context) (types []domain.Type, err error) {
	dataset := r.db.From("types").Where(goqu.C("deleted_at").IsNull()).Executor()
	err = dataset.ScanStructsContext(ctx, &types)
	return
}

func (r repository) Delete(ctx context.Context, id string) error {
	executor := r.db.Update("types").Set(goqu.Record{
		"deleted_at": sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}).Where(goqu.C("id").Eq(id)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
