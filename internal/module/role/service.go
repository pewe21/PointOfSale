package role

import (
	"context"
	"errors"

	"github.com/pewe21/PointOfSale/dto"
	"github.com/pewe21/PointOfSale/internal/domain"
)

type service struct {
	repository domain.RoleRepository
}

func NewService(repository domain.RoleRepository) domain.RoleService {
	return &service{repository: repository}
}

func (s service) Save(ctx context.Context, req dto.CreateRoleRequest) error {

	roles, err := s.repository.FindAll(ctx)

	if err != nil {
		return err
	}

	for _, v := range roles {
		if v.Name == req.Name {
			return errors.New("role already exist")
		}
	}

	role := domain.Role{
		Name:        req.Name,
		DisplayName: req.DisplayName,
	}

	err = s.repository.Save(ctx, &role)

	if err != nil {
		return err
	}

	return nil
}

func (s service) Update(ctx context.Context, req dto.UpdateRoleRequest, id string) error {
	data, err := s.repository.FindById(ctx, id)
	if err != nil {
		return err
	}

	if data.Id == "" {
		return errors.New("role not exist")
	}

	role := domain.Role{
		DisplayName: req.DisplayName,
	}

	if err := s.repository.Update(ctx, &role, id); err != nil {
		return err
	}

	return nil
}

func (s service) Index(ctx context.Context) ([]dto.RoleData, error) {
	var roles []dto.RoleData

	rolesData, err := s.repository.FindAll(ctx)
	if err != nil {
		return roles, err
	}

	for _, v := range rolesData {
		role := dto.RoleData{
			Id:          v.Id,
			Name:        v.Name,
			DisplayName: v.DisplayName,
		}

		roles = append(roles, role)
	}

	return roles, nil
}

func (s service) GetById(ctx context.Context, id string) (dto.RoleData, error) {
	role, err := s.repository.FindById(ctx, id)
	if err != nil {
		return dto.RoleData{}, err
	}

	if role.Id == "" {
		return dto.RoleData{}, errors.New("role not found")
	}

	newRole := dto.RoleData{
		Id:          role.Id,
		Name:        role.Name,
		DisplayName: role.DisplayName,
	}

	return newRole, nil
}

func (s service) Delete(ctx context.Context, req string) error {
	role, err := s.repository.FindById(ctx, req)

	if err != nil {
		return err
	}

	if role.Id == "" {
		return errors.New("role not found")
	}

	err = s.repository.Delete(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
