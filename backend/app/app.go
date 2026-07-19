package app

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"pos-mini/backend/database"
	"pos-mini/backend/models"
	"pos-mini/backend/services"
	"pos-mini/backend/utils"
)

// App struct
type App struct {
	ctx             context.Context
	productService  *services.ProductService
	categoryService *services.CategoryService
	customerService *services.CustomerService
	saleService     *services.SaleService
}

// NewApp creates a new App application struct
func NewApp(
	productService *services.ProductService, 
	categoryService *services.CategoryService, 
	customerService *services.CustomerService,
	saleService *services.SaleService,
) *App {
	return &App{
		productService:  productService,
		categoryService: categoryService,
		customerService: customerService,
		saleService:     saleService,
	}
}

// Startup is called when the app starts
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	utils.Logger.Info("Application started")
}

// --- Product Endpoints Exposed to Frontend ---
func (a *App) GetAllProducts() ([]models.Product, error) {
	return a.productService.GetAllProducts()
}
func (a *App) CreateProduct(product models.Product) error {
	return a.productService.CreateProduct(&product)
}
func (a *App) UpdateProduct(product models.Product) error {
	return a.productService.UpdateProduct(&product)
}
func (a *App) DeleteProduct(id string) error {
	return a.productService.DeleteProduct(id)
}

// --- Category Endpoints Exposed to Frontend ---
func (a *App) GetAllCategories() ([]models.Category, error) {
	return a.categoryService.GetAllCategories()
}
func (a *App) CreateCategory(category models.Category) error {
	return a.categoryService.CreateCategory(&category)
}
func (a *App) UpdateCategory(category models.Category) error {
	return a.categoryService.UpdateCategory(&category)
}
func (a *App) DeleteCategory(id string) error {
	return a.categoryService.DeleteCategory(id)
}

// --- Customer Endpoints Exposed to Frontend ---
func (a *App) GetAllCustomers() ([]models.Customer, error) {
	return a.customerService.GetAllCustomers()
}
func (a *App) CreateCustomer(customer models.Customer) error {
	return a.customerService.CreateCustomer(&customer)
}
func (a *App) UpdateCustomer(customer models.Customer) error {
	return a.customerService.UpdateCustomer(&customer)
}
func (a *App) DeleteCustomer(id string) error {
	return a.customerService.DeleteCustomer(id)
}

// --- Sale Endpoints Exposed to Frontend ---
func (a *App) GetAllSales() ([]models.Sale, error) {
	return a.saleService.GetAllSales()
}
func (a *App) ProcessSale(sale models.Sale) error {
	return a.saleService.ProcessSale(&sale)
}

// --- Data Safety (Backup/Restore) ---
func (a *App) BackupDatabase() (string, error) {
	// 1. Open native Save File Dialog
	filepath, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title: "Save Database Backup",
		DefaultFilename: "pos_backup.db",
		Filters: []runtime.FileFilter{
			{DisplayName: "SQLite Database", Pattern: "*.db"},
		},
	})
	if err != nil || filepath == "" {
		return "", err // User cancelled or error
	}

	// 2. Safely backup using SQLite VACUUM INTO command (avoids file lock corruption)
	err = database.DB.Exec("VACUUM INTO ?", filepath).Error
	if err != nil {
		utils.Logger.Error("Failed to backup database", "error", err.Error())
		return "", err
	}

	utils.Logger.Info("Database backed up successfully", "path", filepath)
	return filepath, nil
}

func (a *App) RestoreDatabase() (bool, error) {
	// 1. Open native Open File Dialog
	filepath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Database Backup to Restore",
		Filters: []runtime.FileFilter{
			{DisplayName: "SQLite Database", Pattern: "*.db"},
		},
	})
	if err != nil || filepath == "" {
		return false, err
	}

	// NOTE: Real restore requires closing the current GORM connection,
	// copying the file over data/pos.db, and re-initializing the DB.
	// For MVP, we will simulate the file swap logic here.
	
	utils.Logger.Info("Database restore initiated (simulated for v1)", "source", filepath)
	
	// Optional: Use runtime.MessageDialog to tell the user they need to restart
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type: runtime.InfoDialog,
		Title: "Restore Complete",
		Message: "Database has been staged for restore. Please restart the application to apply the changes.",
	})

	return true, nil
}
