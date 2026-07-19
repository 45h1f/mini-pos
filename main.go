package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"pos-mini/backend/app"
	"pos-mini/backend/database"
	"pos-mini/backend/repository"
	"pos-mini/backend/services"
	"pos-mini/backend/utils"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Initialize structured logger
	if err := utils.InitLogger(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// Initialize Database
	if err := database.InitDatabase(); err != nil {
		utils.Logger.Error("Failed to initialize database", "error", err.Error())
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Setup Dependency Injection
	productRepo := repository.NewProductRepository(database.DB)
	productService := services.NewProductService(productRepo)

	categoryRepo := repository.NewCategoryRepository(database.DB)
	categoryService := services.NewCategoryService(categoryRepo)

	customerRepo := repository.NewCustomerRepository(database.DB)
	customerService := services.NewCustomerService(customerRepo)

	saleRepo := repository.NewSaleRepository(database.DB)
	saleService := services.NewSaleService(saleRepo)

	// Create an instance of the app structure
	application := app.NewApp(productService, categoryService, customerService, saleService)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Desktop POS - Mini App",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        application.Startup,
		Bind: []interface{}{
			application,
		},
	})

	if err != nil {
		utils.Logger.Error("Application error", "error", err.Error())
	}
}
