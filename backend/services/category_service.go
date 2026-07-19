package services

import (
	"errors"
	"pos-mini/backend/models"
	"pos-mini/backend/repository"
	"pos-mini/backend/utils"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) CreateCategory(category *models.Category) error {
	if category.Name == "" {
		return errors.New("category name is required")
	}
	err := s.repo.Create(category)
	if err == nil {
		utils.Logger.Info("Created category", "id", category.ID)
	} else {
		utils.Logger.Error("Failed to create category", "error", err.Error())
	}
	return err
}

func (s *CategoryService) UpdateCategory(category *models.Category) error {
	if category.ID == "" {
		return errors.New("category ID is required")
	}
	return s.repo.Update(category)
}

func (s *CategoryService) DeleteCategory(id string) error {
	if id == "" {
		return errors.New("category ID is required")
	}
	return s.repo.Delete(id)
}
