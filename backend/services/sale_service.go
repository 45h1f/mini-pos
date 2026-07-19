package services

import (
	"errors"
	"pos-mini/backend/models"
	"pos-mini/backend/repository"
	"pos-mini/backend/utils"
)

type SaleService struct {
	repo *repository.SaleRepository
}

func NewSaleService(repo *repository.SaleRepository) *SaleService {
	return &SaleService{repo: repo}
}

func (s *SaleService) GetAllSales() ([]models.Sale, error) {
	return s.repo.GetAll()
}

func (s *SaleService) ProcessSale(sale *models.Sale) error {
	if len(sale.Items) == 0 {
		return errors.New("sale must have at least one item")
	}
	
	if sale.Total < 0 {
		return errors.New("sale total cannot be negative")
	}

	err := s.repo.CreateWithTransaction(sale)
	if err == nil {
		utils.Logger.Info("Processed sale successfully", "invoice_no", sale.InvoiceNo)
	} else {
		utils.Logger.Error("Failed to process sale", "error", err.Error())
	}
	
	return err
}
