package services

import (
	"errors"
	"pos-mini/backend/models"
	"pos-mini/backend/repository"
	"pos-mini/backend/utils"
)

type CustomerService struct {
	repo *repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) GetAllCustomers() ([]models.Customer, error) {
	return s.repo.GetAll()
}

func (s *CustomerService) CreateCustomer(customer *models.Customer) error {
	if customer.Name == "" {
		return errors.New("customer name is required")
	}
	err := s.repo.Create(customer)
	if err == nil {
		utils.Logger.Info("Created customer", "id", customer.ID)
	} else {
		utils.Logger.Error("Failed to create customer", "error", err.Error())
	}
	return err
}

func (s *CustomerService) UpdateCustomer(customer *models.Customer) error {
	if customer.ID == "" {
		return errors.New("customer ID is required")
	}
	return s.repo.Update(customer)
}

func (s *CustomerService) DeleteCustomer(id string) error {
	if id == "" {
		return errors.New("customer ID is required")
	}
	return s.repo.Delete(id)
}
