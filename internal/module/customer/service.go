package customer

import (
	"context"
	"errors"
	"github.com/pewe21/PointOfSale/dto"
	"github.com/pewe21/PointOfSale/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repository        domain.CustomerRepository
	roleRepository    domain.RoleRepository
	cusRoleRepository domain.CustomerRolesRepository
}

func NewService(customerRepository domain.CustomerRepository, roleRepository domain.RoleRepository, cusRoleRepository domain.CustomerRolesRepository) domain.CustomerService {
	return &service{repository: customerRepository, roleRepository: roleRepository, cusRoleRepository: cusRoleRepository}
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
	err = s.repository.Save(ctx, &customer)
	if err != nil {
		return errors.New("error creating customer")
	}
	return nil
}

func (s service) Update(ctx context.Context, req dto.UpdateCustomerRequest, id string) error {
	_, err := s.repository.FindById(ctx, id)
	if err != nil {
		return errors.New("error updating customer, customer not found")
	}

	customer := domain.Customer{
		Name:     req.Name,
		Username: req.Username,
		Phone:    req.Phone,
		Address:  req.Address,
	}
	err = s.repository.Update(ctx, &customer, id)
	if err != nil {
		return errors.New("error updating customer")
	}

	return nil
}

func (s service) Index(ctx context.Context) ([]dto.CustomerData, error) {
	var data []dto.CustomerData
	customers, err := s.repository.FindAll(ctx)
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
	customer, err = s.repository.FindById(ctx, id)
	if err != nil {
		return domain.Customer{}, errors.New("customer not found")
	}

	return customer, nil
}

func (s service) GetByUsername(ctx context.Context, username string) (customer domain.Customer, err error) {
	customer, err = s.repository.FindByUsername(ctx, username)
	if err != nil {
		return domain.Customer{}, errors.New("customer not found")
	}

	return customer, nil
}

func (s service) GetByEmail(ctx context.Context, email string) (customer domain.Customer, err error) {
	customer, err = s.repository.FindByEmail(ctx, email)
	if err != nil {
		return domain.Customer{}, errors.New("customer not found")
	}

	return customer, nil
}

func (s service) Delete(ctx context.Context, req string) error {
	_, err := s.repository.FindById(ctx, req)
	if err != nil {
		return err
	}

	err = s.repository.Delete(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s service) AddRole(ctx context.Context, req dto.AddCustomerRoleRequest) error {
	customer, err := s.repository.FindById(ctx, req.CustomerId)
	if customer.Id == "" {
		return errors.New("customer not found")
	}

	if err != nil {
		return err
	}

	role, err := s.roleRepository.FindById(ctx, req.RoleId)

	if role.Id == "" {
		return errors.New("role not found")
	}

	if err != nil {
		return err
	}

	err = s.cusRoleRepository.Create(ctx, req.RoleId, req.CustomerId)
	if err != nil {
		return errors.New("error adding customer role")
	}

	return nil
}

func (s service) ChangeRole(ctx context.Context, req dto.UpdateCustomerRoleRequest, customerId string) error {
	customer, err := s.repository.FindById(ctx, customerId)
	if customer.Id == "" {
		return errors.New("customer not found")
	}

	if err != nil {
		return err
	}

	role, err := s.roleRepository.FindById(ctx, req.RoleId)

	if role.Id == "" {
		return errors.New("role not found")
	}

	if err != nil {
		return err
	}

	err = s.cusRoleRepository.Update(ctx, req.RoleId, customerId)
	if err != nil {
		return errors.New("error changing customer role")
	}

	return nil
}
