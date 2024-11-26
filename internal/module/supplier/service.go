package supplier

import (
	"context"
	"errors"

	"github.com/pewe21/PointOfSale/dto"
	"github.com/pewe21/PointOfSale/internal/domain"
)

type service struct {
	repository domain.SupplierRepository
}

func NewService(repository domain.SupplierRepository) domain.SupplierService {
	return &service{repository: repository}
}

func (s service) Save(ctx context.Context, req dto.CreateSupplierRequest) error {

	datas, err := s.repository.FindAll(ctx)

	if err != nil {
		return errors.New("error creating supplier")
	}

	for _, data := range datas {
		if data.Name == req.Name {
			return errors.New("supplier already exist")
		}
		if data.Email == req.Email {
			return errors.New("email already exist")
		}

		if data.Phone == req.Phone {
			return errors.New("phone already exist")
		}
	}

	supplier := domain.Supplier{
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
		Phone:   req.Phone,
	}
	err = s.repository.Save(ctx, &supplier)
	if err != nil {
		return errors.New("error creating supplier")
	}

	return nil
}

func (s service) Update(ctx context.Context, req dto.UpdateSupplierRequest, id string) error {
	_, err := s.repository.FindById(ctx, id)
	if err != nil {
		return errors.New("error updating supplier, supplier not found")
	}

	supplier := domain.Supplier{
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
		Phone:   req.Phone,
	}

	err = s.repository.Update(ctx, &supplier, id)
	if err != nil {
		return errors.New("error updating supplier")
	}

	return nil
}

func (s service) Index(ctx context.Context) ([]dto.SupplierData, error) {
	var data []dto.SupplierData
	suppliers, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, errors.New("error getting suppliers")
	}
	for _, supplier := range suppliers {
		s := dto.SupplierData{
			Id:      supplier.Id,
			Name:    supplier.Name,
			Email:   supplier.Email,
			Address: supplier.Address,
			Phone:   supplier.Phone,
		}
		data = append(data, s)
	}

	return data, nil
}

func (s service) GetById(ctx context.Context, id string) (dto.SupplierData, error) {
	var data dto.SupplierData
	supplier, err := s.repository.FindById(ctx, id)

	if err != nil {
		return data, errors.New("error getting data supplier")
	}

	if supplier.Id == "" {
		return data, errors.New("error getting data supplier")
	}

	data = dto.SupplierData{
		Id:      supplier.Id,
		Name:    supplier.Name,
		Email:   supplier.Email,
		Address: supplier.Address,
		Phone:   supplier.Phone,
	}
	return data, nil
}

func (s service) Delete(ctx context.Context, req string) error {

	supplier, err := s.repository.FindById(ctx, req)
	if err != nil {
		return errors.New("error deleting supplier, supplier not found")
	}

	if supplier.Id == "" {
		return errors.New("supplier not found")

	}

	err = s.repository.Delete(ctx, req)
	if err != nil {
		return errors.New("error deleting supplier")
	}

	return nil
}
