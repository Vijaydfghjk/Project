package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Invoice struct {
	Table_number   string         `json:"table_number" validate:"required"`
	Payment_method string         `json:"payment_method" validate:"required"`
	Payment_status string         `json:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	Invoice_id     string         `gorm:"primaryKey;unique;" json:"invoice_id"`
	Stored         datatypes.JSON `json:"stored"`
}

type InvoiceProcess interface {
	CreateInvoice(in Invoice) (Invoice, error)
}

type InvoiceDB struct {
	db *gorm.DB
}

func Invoice_Repo(db *gorm.DB) *InvoiceDB {

	return &InvoiceDB{db: db}
}

func (a *InvoiceDB) CreateInvoice(in Invoice) (Invoice, error) {

	in.Invoice_id = uuid.New().String()

	err := a.db.Create(&in).Error
	if err != nil {

		return in, err
	}
	return in, nil
}
