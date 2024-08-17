package user

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

func NewRepository(db *sql.DB) domain.UserRepository {
	return &repository{db: goqu.New("default", db)}
}

func (r repository) Save(ctx context.Context, user *domain.User) error {
	exec := r.db.Insert("users").Rows(user).Executor()

	_, err := exec.ExecContext(ctx)
	return err
}

func (r repository) Update(ctx context.Context, user *domain.User, id string) error {
	exec := r.db.Update("users").
		Set(goqu.Record{
			"name":  user.Name,
			"email": user.Email,
			"phone": user.Phone,
			"updated_at": sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		}).Where(goqu.C("id").Eq(id)).Executor()
	_, err := exec.ExecContext(ctx)
	return err
}

func (r repository) FindAll(ctx context.Context) (result []domain.User, err error) {
	dataset := r.db.From("users").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &result)

	return

}

func (r repository) FindById(ctx context.Context, id string) (result domain.User, err error) {
	dataset := r.db.From("users").Where(goqu.C("deleted_at").IsNull(), goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)

	return
}

func (r repository) Delete(ctx context.Context, id string) error {
	exec := r.db.Update("users").Set(goqu.Record{
		"deleted_at": sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}}).Where(goqu.C("id").Eq(id)).Executor()
	_, err := exec.ExecContext(ctx)
	return err
}

func (r repository) FindByEmail(ctx context.Context, email string) (result domain.User, err error) {
	dataset := r.db.From("users").Where(goqu.C("email").Eq(email)).Executor()
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}
