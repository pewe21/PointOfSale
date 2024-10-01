package product

import (
	"context"
	"errors"
	"github.com/pewe21/PointOfSale/dto"
	"github.com/pewe21/PointOfSale/internal/domain"
)

type service struct {
	repository domain.ProductRepository
}

func NewService(repository domain.ProductRepository) domain.ProductService {
	return &service{repository: repository}

}

func (s service) IndexNew(ctx context.Context) (productsx []dto.ProductxDto, err error) {
	prod, err := s.repository.FindAll(ctx)

	for _, v := range prod {
		var product dto.ProductxDto
		product = dto.ProductxDto{
			ID:   v.Id,
			Name: v.Name,
			Supplier: dto.Supplierx{
				ID:   v.SupplierId,
				Name: v.SupplierName,
			},
		}
		productsx = append(productsx, product)
	}
	return
}

func (s service) Index(ctx context.Context) (products []domain.ProductWithDetail, err error) {
	products, err = s.repository.FindAll(ctx)
	return
}

func (s service) GetById(ctx context.Context, id string) (product dto.ProductxDto, err error) {
	prd, errs := s.repository.FindById(ctx, id)

	if errs != nil {
		return product, errs
	}

	if prd.Id == "" {
		return product, errors.New("product not found")
	}

	product = dto.ProductxDto{
		ID:   prd.Id,
		Name: prd.Name,
		Supplier: dto.Supplierx{
			ID:   prd.SupplierId,
			Name: prd.SupplierName,
		},
	}
	return
}

func (s service) Create(ctx context.Context, product *domain.Product) error {
	err := s.repository.Save(ctx, product)
	return err
}

func (s service) Update(ctx context.Context, product *domain.Product, id string) error {
	_, err := s.repository.FindById(ctx, id)
	if err != nil {
		return errors.New("Product not found")
	}

	err = s.repository.Update(ctx, product, id)
	return err
}

func (s service) Delete(ctx context.Context, id string) error {
	_, err := s.repository.FindById(ctx, id)
	if err != nil {
		return errors.New("Product not found")
	}

	err = s.repository.Delete(ctx, id)
	return err
}
