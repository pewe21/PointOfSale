package customer

import (
	"context"
	"errors"
	"github.com/pewe21/PointOfSale/dto"
	"github.com/pewe21/PointOfSale/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	domain.CustomerRepository
}

func NewService(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &service{CustomerRepository: customerRepository}
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
	err = s.CustomerRepository.Save(ctx, &customer)
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
	err := s.CustomerRepository.Update(ctx, &customer, id)
	if err != nil {
		return errors.New("error updating customer")
	}

	return nil
}

func (s service) Index(ctx context.Context) ([]dto.CustomerData, error) {
	var data []dto.CustomerData
	customers, err := s.CustomerRepository.FindAll(ctx)
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
	customer, err = s.CustomerRepository.FindById(ctx, id)
	if err != nil {
		return domain.Customer{}, errors.New("customer not found")
	}

	return customer, nil
}

func (s service) GetByUsername(ctx context.Context, username string) (customer domain.Customer, err error) {
	customer, err = s.CustomerRepository.FindByUsername(ctx, username)
	if err != nil {
		return domain.Customer{}, errors.New("customer not found")
	}

	return customer, nil
}

func (s service) GetByEmail(ctx context.Context, email string) (customer domain.Customer, err error) {
	customer, err = s.CustomerRepository.FindByEmail(ctx, email)
	if err != nil {
		return domain.Customer{}, errors.New("customer not found")
	}

	return customer, nil
}

func (s service) Delete(ctx context.Context, req string) error {
	_, err := s.CustomerRepository.FindById(ctx, req)
	if err != nil {
		return err
	}

	err = s.CustomerRepository.Delete(ctx, req)
	if err != nil {
		return err
	}

	return nil
}