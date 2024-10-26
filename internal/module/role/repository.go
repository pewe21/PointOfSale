package role

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

func NewRepository(db *sql.DB) domain.RoleRepository {
	return &repository{db: goqu.New("default", db)}
}

func (r repository) Save(ctx context.Context, role *domain.Role) error {
	executor := r.db.Insert("roles").Rows(role).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (r repository) Update(ctx context.Context, role *domain.Role, id string) error {
	role.UpdatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	executor := r.db.Update("roles").Set(goqu.Ex{
		"display_name": role.DisplayName,
	}).Where(goqu.C("id").Eq(id)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (r repository) FindById(ctx context.Context, id string) (role domain.Role, err error) {
	dataset := r.db.From("roles").Where(goqu.C("id").Eq(id)).Executor()
	_, err = dataset.ScanStructContext(ctx, &role)
	return

}

func (r repository) FindAll(ctx context.Context) (roles []domain.Role, err error) {
	dataset := r.db.From("roles").Executor()
	err = dataset.ScanStructsContext(ctx, &roles)
	return
}

func (r repository) Delete(ctx context.Context, id string) error {
	executor := r.db.Update("roles").Set(goqu.Record{
		"deleted_at": sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}).Where(goqu.C("id").Eq(id)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
