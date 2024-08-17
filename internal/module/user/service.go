package user

import (
	"context"
	"github.com/pewe21/PointOfSale/dto"
	"github.com/pewe21/PointOfSale/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	domain.UserRepository
}

func NewService(userRepository domain.UserRepository) domain.UserService {
	return &service{UserRepository: userRepository}
}

func (s service) Index(ctx context.Context) ([]dto.UserData, error) {

	users, err := s.UserRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var res []dto.UserData
	for _, u := range users {
		user := dto.UserData{
			Id:    u.Id,
			Name:  u.Name,
			Email: u.Email,
			Phone: u.Phone,
		}

		res = append(res, user)
	}

	return res, nil
}

func (s service) Save(ctx context.Context, req dto.CreateUserRequest) error {

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	user := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(password),
		Phone:    req.Phone,
	}
	err = s.UserRepository.Save(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}

func (s service) Update(ctx context.Context, req dto.UpdateUserRequest, id string) error {
	data := domain.User{
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}
	err := s.UserRepository.Update(ctx, &data, id)
	if err != nil {
		return err
	}
	return nil
}

func (s service) Delete(ctx context.Context, req string) error {
	err := s.UserRepository.Delete(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
