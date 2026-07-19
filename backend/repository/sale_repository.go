package repository

import (
	"pos-mini/backend/models"
	"gorm.io/gorm"
)

type SaleRepository struct {
	db *gorm.DB
}

func NewSaleRepository(db *gorm.DB) *SaleRepository {
	return &SaleRepository{db: db}
}

func (r *SaleRepository) GetAll() ([]models.Sale, error) {
	var sales []models.Sale
	err := r.db.Preload("Items").Preload("Customer").Find(&sales).Error
	return sales, err
}

func (r *SaleRepository) CreateWithTransaction(sale *models.Sale) error {
	// Execute inside a transaction
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 1. Create the sale. Because of GORM, this will automatically
		// attempt to create the nested sale.Items as well.
		if err := tx.Create(sale).Error; err != nil {
			return err
		}

		// 2. Decrement stock for each item sold
		for _, item := range sale.Items {
			// Skip stock decrement if no product ID is attached (e.g. custom open-ring item)
			if item.ProductID == "" {
				continue
			}
			
			if err := tx.Model(&models.Product{}).
				Where("id = ?", item.ProductID).
				UpdateColumn("stock", gorm.Expr("stock - ?", item.Qty)).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
