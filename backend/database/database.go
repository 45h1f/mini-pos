package database

import (
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"pos-mini/backend/models"
	"pos-mini/backend/utils"
)

var DB *gorm.DB

// InitDatabase initializes the SQLite connection and runs auto-migration.
func InitDatabase() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = "."
	}
	dataDir := filepath.Join(configDir, "pos-mini", "data")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return err
	}

	dbPath := filepath.Join(dataDir, "pos.db")
	// Use _pragma=foreign_keys(1) to enable foreign key constraints in SQLite
	dsn := dbPath + "?_pragma=foreign_keys(1)"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn), // Avoid too much noise in production
	})
	if err != nil {
		return err
	}

	DB = db
	utils.Logger.Info("Database connection established")

	// Run AutoMigrate for all models
	err = db.AutoMigrate(
		&models.Category{},
		&models.Product{},
		&models.Customer{},
		&models.Sale{},
		&models.SaleItem{},
		&models.Setting{},
	)
	if err != nil {
		utils.Logger.Error("Failed to auto-migrate database", "error", err.Error())
		return err
	}

	utils.Logger.Info("Database auto-migration completed successfully")
	return nil
}
