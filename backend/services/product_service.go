package services

import (
	"errors"
	"pos-mini/backend/models"
	"pos-mini/backend/repository"
	"pos-mini/backend/utils"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	utils.Logger.Info("Fetching all products")
	return s.repo.GetAll()
}

func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	if id == "" {
		return nil, errors.New("product ID is required")
	}
	return s.repo.GetByID(id)
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price < 0 {
		return errors.New("product price cannot be negative")
	}
	
	err := s.repo.Create(product)
	if err != nil {
		utils.Logger.Error("Failed to create product", "error", err.Error())
		return err
	}
	
	utils.Logger.Info("Created new product", "id", product.ID)
	return nil
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	if product.ID == "" {
		return errors.New("product ID is required for update")
	}
	
	err := s.repo.Update(product)
	if err != nil {
		utils.Logger.Error("Failed to update product", "id", product.ID, "error", err.Error())
		return err
	}
	return nil
}

func (s *ProductService) DeleteProduct(id string) error {
	if id == "" {
		return errors.New("product ID is required for deletion")
	}
	
	err := s.repo.Delete(id)
	if err != nil {
		utils.Logger.Error("Failed to delete product", "id", id, "error", err.Error())
		return err
	}
	return nil
}
