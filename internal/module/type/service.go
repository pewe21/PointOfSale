package _type

import (
	"context"
	"errors"
	"github.com/pewe21/PointOfSale/dto"
	"github.com/pewe21/PointOfSale/internal/domain"
)

type service struct {
	repository domain.TypeRepository
}

func NewService(repository domain.TypeRepository) domain.TypeService {
	return &service{repository: repository}
}

func (s service) Save(ctx context.Context, req dto.CreateTypeRequest) error {
	_type := domain.Type{
		Name:        req.Name,
		Description: req.Description,
	}
	err := s.repository.Save(ctx, &_type)
	if err != nil {
		return errors.New("error saving type")
	}

	return nil
}

func (s service) Update(ctx context.Context, req dto.UpdateTypeRequest, id string) error {
	_, err := s.repository.FindById(ctx, id)
	if err != nil {
		return errors.New("error saving type, type not found")
	}

	_type := domain.Type{
		Name:        req.Name,
		Description: req.Description,
	}
	err = s.repository.Update(ctx, &_type, id)
	if err != nil {
		return errors.New("error saving type")
	}

	return nil
}

func (s service) Index(ctx context.Context) ([]dto.TypeData, error) {
	var data []dto.TypeData
	types, err := s.repository.FindAll(ctx)
	if err != nil {
		return data, err
	}

	for _, _type := range types {
		t := dto.TypeData{
			Name:        _type.Name,
			Description: _type.Description,
		}
		data = append(data, t)
	}

	return data, nil
}

func (s service) GetById(ctx context.Context, id string) (dto.TypeData, error) {
	var data dto.TypeData
	_type, err := s.repository.FindById(ctx, id)
	if err != nil {
		return data, errors.New("error getting type")
	}
	data = dto.TypeData{
		Name:        _type.Name,
		Description: _type.Description,
	}
	return data, nil
}

func (s service) Delete(ctx context.Context, req string) error {
	_, err := s.repository.FindById(ctx, req)
	if err != nil {
		return errors.New("error deleting type, type not found")
	}

	err = s.repository.Delete(ctx, req)
	if err != nil {
		return errors.New("error deleting type")
	}

	return nil

}
