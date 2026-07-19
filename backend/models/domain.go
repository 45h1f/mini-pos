package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseModel provides a UUID primary key and standard timestamp fields.
type BaseModel struct {
	ID        string         `gorm:"type:varchar(36);primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if base.ID == "" {
		base.ID = uuid.New().String()
	}
	return
}

type Category struct {
	BaseModel
	Name     string    `gorm:"not null" json:"name"`
	Products []Product `json:"products,omitempty"`
}

type Product struct {
	BaseModel
	CategoryID string   `gorm:"type:varchar(36);index" json:"category_id"`
	Category   Category `json:"category,omitempty"`
	Name       string   `gorm:"not null" json:"name"`
	SKU        string   `gorm:"uniqueIndex" json:"sku"`
	Barcode    string   `gorm:"index" json:"barcode"`
	Price      int64    `gorm:"not null" json:"price"` // Stored in cents/paisa
	Stock      int      `gorm:"not null;default:0" json:"stock"`
	IsActive   bool     `gorm:"default:true" json:"is_active"`
}

type Customer struct {
	BaseModel
	Name  string `gorm:"not null" json:"name"`
	Phone string `json:"phone"`
	Sales []Sale `json:"sales,omitempty"`
}

type Sale struct {
	BaseModel
	InvoiceNo     string      `gorm:"uniqueIndex;not null" json:"invoice_no"`
	CustomerID    *string     `gorm:"type:varchar(36);index" json:"customer_id"`
	Customer      *Customer   `json:"customer,omitempty"`
	Subtotal      int64       `gorm:"not null" json:"subtotal"` // in cents
	Tax           int64       `gorm:"not null" json:"tax"`      // in cents
	Discount      int64       `gorm:"not null" json:"discount"` // in cents
	Total         int64       `gorm:"not null" json:"total"`    // in cents
	Paid          int64       `gorm:"not null" json:"paid"`     // in cents
	Change        int64       `gorm:"not null" json:"change"`   // in cents
	PaymentMethod string      `gorm:"not null" json:"payment_method"`
	Items         []SaleItem  `json:"items,omitempty"`
}

type SaleItem struct {
	BaseModel
	SaleID        string  `gorm:"type:varchar(36);index;not null" json:"sale_id"`
	ProductID     string  `gorm:"type:varchar(36);index" json:"product_id"`
	Product       Product `json:"product,omitempty"`
	NameSnapshot  string  `gorm:"not null" json:"name_snapshot"`
	PriceSnapshot int64   `gorm:"not null" json:"price_snapshot"` // in cents
	Qty           int     `gorm:"not null" json:"qty"`
	LineTotal     int64   `gorm:"not null" json:"line_total"` // in cents
}

type Setting struct {
	Key   string `gorm:"primaryKey;type:varchar(100)" json:"key"`
	Value string `gorm:"type:text" json:"value"`
}
