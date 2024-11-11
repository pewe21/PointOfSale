package brand

import (
	"context"
	"errors"

	"github.com/pewe21/PointOfSale/dto"
	"github.com/pewe21/PointOfSale/internal/domain"
)

type service struct {
	repository domain.BrandRepository
}

func NewService(repository domain.BrandRepository) domain.BrandService {
	return &service{repository: repository}
}

func (s service) Save(ctx context.Context, req dto.CreateBrandRequest) error {

	brands, _ := s.repository.FindAll(ctx)

	for _, v := range brands {
		if v.Name == req.Name {
			return errors.New("error saving brand, brand already exist")
		}
	}

	brand := domain.Brand{
		Name:        req.Name,
		Description: req.Description,
	}
	err := s.repository.Save(ctx, &brand)
	if err != nil {
		return errors.New("error saving brand")
	}

	return nil
}

func (s service) Update(ctx context.Context, req dto.UpdateBrandRequest, id string) error {
	_, err := s.repository.FindById(ctx, id)
	if err != nil {
		return errors.New("error saving brand, brand not found")
	}

	brand := domain.Brand{
		Name:        req.Name,
		Description: req.Description,
	}
	err = s.repository.Update(ctx, &brand, id)
	if err != nil {
		return errors.New("error saving brand")
	}

	return nil
}

func (s service) Index(ctx context.Context) ([]dto.BrandData, error) {
	var data []dto.BrandData
	types, err := s.repository.FindAll(ctx)
	if err != nil {
		return data, err
	}

	for _, brand := range types {
		t := dto.BrandData{
			Id:          brand.Id,
			Name:        brand.Name,
			Description: brand.Description,
		}
		data = append(data, t)
	}

	return data, nil
}

func (s service) GetById(ctx context.Context, id string) (dto.BrandData, error) {
	var data dto.BrandData
	brand, err := s.repository.FindById(ctx, id)

	if err != nil {
		return data, errors.New("error getting brand")
	}

	if brand.Id == "" {
		return data, errors.New("brand not found")
	}
	data = dto.BrandData{
		Id:          brand.Id,
		Name:        brand.Name,
		Description: brand.Description,
	}
	return data, nil
}

func (s service) Delete(ctx context.Context, req string) error {
	_, err := s.repository.FindById(ctx, req)
	if err != nil {
		return errors.New("error deleting brand, brand not found")
	}

	err = s.repository.Delete(ctx, req)
	if err != nil {
		return errors.New("error deleting brand")
	}

	return nil

}
