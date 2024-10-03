package product

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

func NewRepository(db *sql.DB) domain.ProductRepository {
	return &repository{db: goqu.New("default", db)}
}

func (r repository) FindAll(ctx context.Context) (products []domain.ProductWithDetail, err error) {
	dataset := r.db.From("products").LeftJoin(goqu.T("suppliers"), goqu.On(goqu.Ex{
		"products.supplier_id": goqu.L("suppliers.id"),
	})).LeftJoin(goqu.T("brands"), goqu.On(goqu.Ex{
		"products.brand_id": goqu.L("brands.id"),
	})).Where(goqu.L("products.deleted_at").IsNull()).Select(
		goqu.L("products.id").As("id"),
		goqu.L("products.name").As("name"),
		goqu.L("products.sku").As("sku"),
		goqu.L("products.stock").As("stock"),
		goqu.L("products.brand_id").As("brand_id"),
		goqu.L("brands.name").As("brand_name"),
		goqu.L("products.supplier_id").As("supplier_id"),
		goqu.L("suppliers.name").As("supplier_name"),
		//goqu.L("products.created_at").As("created_at"),
		//goqu.L("products.updated_at").As("updated_at"),
		//goqu.L("products.deleted_at").As("deleted_at"),
	).Executor()
	err = dataset.ScanStructsContext(ctx, &products)
	return

}

func (r repository) FindById(ctx context.Context, id string) (product domain.ProductWithDetail, err error) {
	dataset := r.db.From("products").LeftJoin(goqu.T("suppliers"), goqu.On(goqu.Ex{
		"products.supplier_id": goqu.L("suppliers.id"),
	})).LeftJoin(goqu.T("brands"), goqu.On(goqu.Ex{
		"products.brand_id": goqu.L("brands.id"),
	})).Where(goqu.L("products.deleted_at").IsNull()).Where(goqu.L("products.id").Eq(id)).Select(
		goqu.L("products.id").As("id"),
		goqu.L("products.name").As("name"),
		goqu.L("products.sku").As("sku"),
		goqu.L("products.stock").As("stock"),
		goqu.L("products.brand_id").As("brand_id"),
		goqu.L("brands.name").As("brand_name"),
		goqu.L("products.supplier_id").As("supplier_id"),
		goqu.L("suppliers.name").As("supplier_name"),
		//goqu.L("products.created_at").As("created_at"),
		//goqu.L("products.updated_at").As("updated_at"),
		//goqu.L("products.deleted_at").As("deleted_at"),
	).Executor()
	_, err = dataset.ScanStructContext(ctx, &product)
	return
}

func (r repository) Save(ctx context.Context, product *domain.Product) error {
	executor := r.db.Insert("products").Rows(product).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (r repository) Update(ctx context.Context, product *domain.Product, id string) error {
	product.UpdatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	executor := r.db.Update("products").Set(product).Where(goqu.C("id").Eq(id)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (r repository) Delete(ctx context.Context, id string) error {
	executor := r.db.Update("products").Set(goqu.Record{
		"deleted_at": sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}).Where(goqu.C("id").Eq(id)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
