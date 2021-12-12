package base_models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Foundation struct {
	gorm.Model     `gorm:"model"`
	Name           string
	Type           string
	Stage          string
	Maker          string
	Checker        string
	Approver       string
	Description    string
	Status         string
	WorkflowLevels int
	SyncToken      uint
	Version        uint
	Enabled        int
	Locale         string
	Events         []byte
}

type FoundationWithDocumentBytes struct {
	Foundation
	Qrcode     []byte
	Barcode    []byte
	Attachment []byte
}

type FoundationWithTransactionType struct {
	Foundation
	TransactionType string
}
type BaseModel struct {
	ID      string
	Version int
	At      time.Time
}
