package customer

import (
	"context"
	"errors"
	"github.com/pewe21/PointOfSale/dto"
	"github.com/pewe21/PointOfSale/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	domain.TypeRepository
}

func NewService(customerRepository domain.TypeRepository) domain.CustomerService {
	return &service{TypeRepository: customerRepository}
}

func (s service) Save(ctx context.Context, req dto.CreateCustomerRequest) error {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	customer := domain.Customer{
		Name:     req.Name,
		Username: req.Username,
		Password: string(password),
		Email:    req.Email,
	}
	err = s.TypeRepository.Save(ctx, &customer)
	if err != nil {
		return errors.New("error creating customer")
	}
	return nil
}

func (s service) Update(ctx context.Context, req dto.UpdateCustomerRequest, id string) error {
	customer := domain.Customer{
		Name:     req.Name,
		Username: req.Username,
		Phone:    req.Phone,
		Address:  req.Address,
	}
	err := s.TypeRepository.Update(ctx, &customer, id)
	if err != nil {
		return errors.New("error updating customer")
	}

	return nil
}

func (s service) Index(ctx context.Context) ([]dto.CustomerData, error) {
	var data []dto.CustomerData
	customers, err := s.TypeRepository.FindAll(ctx)
	if err != nil {
		return nil, errors.New("error getting customers")
	}

	for _, customer := range customers {
		c := dto.CustomerData{
			Name:     customer.Name,
			Username: customer.Username,
			Email:    customer.Email,
			Phone:    customer.Phone,
		}

		data = append(data, c)
	}

	return data, nil
}

func (s service) GetById(ctx context.Context, id string) (customer domain.Customer, err error) {
	customer, err = s.TypeRepository.FindById(ctx, id)
	if err != nil {
		return domain.Customer{}, errors.New("customer not found")
	}

	return customer, nil
}

func (s service) GetByUsername(ctx context.Context, username string) (customer domain.Customer, err error) {
	customer, err = s.TypeRepository.FindByUsername(ctx, username)
	if err != nil {
		return domain.Customer{}, errors.New("customer not found")
	}

	return customer, nil
}

func (s service) GetByEmail(ctx context.Context, email string) (customer domain.Customer, err error) {
	customer, err = s.TypeRepository.FindByEmail(ctx, email)
	if err != nil {
		return domain.Customer{}, errors.New("customer not found")
	}

	return customer, nil
}

func (s service) Delete(ctx context.Context, req string) error {
	_, err := s.TypeRepository.FindById(ctx, req)
	if err != nil {
		return err
	}

	err = s.TypeRepository.Delete(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
