package services

import (
	"testing"

	"log/slog"
	
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"pos-mini/backend/models"
	"pos-mini/backend/repository"
	"pos-mini/backend/utils"
)

func setupTestDB(t *testing.T) *gorm.DB {
	// Initialize logger for testing
	utils.Logger = slog.Default()
	
	// Use pure-Go SQLite with in-memory database for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to memory db: %v", err)
	}
	
	// Migrate models
	err = db.AutoMigrate(&models.Category{}, &models.Product{})
	if err != nil {
		t.Fatalf("failed to migrate models: %v", err)
	}
	
	return db
}

func TestProductService(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewProductRepository(db)
	service := NewProductService(repo)

	// 1. Test Create
	product := models.Product{
		Name:  "Test Product",
		Price: 1500, // 15.00
		Stock: 10,
	}

	err := service.CreateProduct(&product)
	if err != nil {
		t.Errorf("expected no error on create, got %v", err)
	}
	if product.ID == "" {
		t.Errorf("expected UUID to be generated on create")
	}

	// 2. Test Read
	products, err := service.GetAllProducts()
	if err != nil {
		t.Errorf("expected no error on read all, got %v", err)
	}
	if len(products) != 1 {
		t.Errorf("expected exactly 1 product, got %d", len(products))
	}

	fetchedProduct, err := service.GetProductByID(product.ID)
	if err != nil || fetchedProduct.Name != "Test Product" {
		t.Errorf("failed to fetch specific product")
	}

	// 3. Test Update
	product.Price = 2000
	err = service.UpdateProduct(&product)
	if err != nil {
		t.Errorf("expected no error on update, got %v", err)
	}
	updatedProduct, _ := service.GetProductByID(product.ID)
	if updatedProduct.Price != 2000 {
		t.Errorf("expected updated price to be 2000, got %d", updatedProduct.Price)
	}

	// 4. Test Delete
	err = service.DeleteProduct(product.ID)
	if err != nil {
		t.Errorf("expected no error on delete, got %v", err)
	}
	
	// Verify Deletion
	_, err = service.GetProductByID(product.ID)
	if err == nil {
		t.Errorf("expected error fetching deleted product, got none")
	}
}
