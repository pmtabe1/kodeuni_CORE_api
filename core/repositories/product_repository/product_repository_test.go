package product_repository

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/paulmsegeya/pos/core/models/base_models"
	"github.com/paulmsegeya/pos/core/models/pos_models"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {

	got := New()

	require.NotNilf(t, got, "Expected non nil but got %v instead ", got)
}

func TestAdd(t *testing.T) {

	ref := 29372112

	data := pos_models.Product{
		FoundationWithDocumentBytes: base_models.FoundationWithDocumentBytes{
			Foundation: base_models.Foundation{
				Model:          gorm.Model{},
				Name:           "PRODUCT",
				Type:           "PRODUCT",
				Stage:          "received",
				Maker:          "Maker",
				Checker:        "Checker",
				Approver:       "Approver",
				Description:    "Product ....",
				Status:         "instock",
				WorkflowLevels: 0,
				SyncToken:      0,
				Version:        0,
				Enabled:        0,
				Locale:         "",
				Events:         []byte{},
			},
			Qrcode:     []byte{},
			Barcode:    []byte{},
			Attachment: []byte{},
		},
		ProductItemList: []*pos_models.ProductItem{},
		WorkflowList:    []*pos_models.Workflow{},
		Sku:             "111",
		SN:              "222111",
	}
	data.ID = uint(ref)
	got := New().Add(&data)
	require.NotNilf(t, got.Product, "Expected non Nil but received %v  instead", got.ProductList)

}

func TestUpdate(t *testing.T) {

	ref := 29372112

	data := pos_models.Product{
		FoundationWithDocumentBytes: base_models.FoundationWithDocumentBytes{
			Foundation: base_models.Foundation{
				Model:          gorm.Model{},
				Name:           "PRODUCT",
				Type:           "PRODUCT",
				Stage:          "received",
				Maker:          "Maker",
				Checker:        "Checker",
				Approver:       "Approver",
				Description:    "Product ....",
				Status:         "instock",
				WorkflowLevels: 0,
				SyncToken:      0,
				Version:        0,
				Enabled:        0,
				Locale:         "",
				Events:         []byte{},
			},
			Qrcode:     []byte{},
			Barcode:    []byte{},
			Attachment: []byte{},
		},
		ProductItemList: []*pos_models.ProductItem{},
		WorkflowList:    []*pos_models.Workflow{},
		Sku:             "111",
		SN:              "222111",
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().Update(uint(ref), &data)
	require.NotNilf(t, got.Product, "Expected non Nil but received %v  instead", got.ProductList)
}
func TestAddOrUpdate(t *testing.T) {
	ref := 29372115

	data := pos_models.Product{
		FoundationWithDocumentBytes: base_models.FoundationWithDocumentBytes{
			Foundation: base_models.Foundation{
				Model:          gorm.Model{},
				Name:           "PRODUCT",
				Type:           "PRODUCT",
				Stage:          "received",
				Maker:          "Maker",
				Checker:        "Checker",
				Approver:       "Approver",
				Description:    "Product ....",
				Status:         "instock",
				WorkflowLevels: 0,
				SyncToken:      0,
				Version:        0,
				Enabled:        0,
				Locale:         "",
				Events:         []byte{},
			},
			Qrcode:     []byte{},
			Barcode:    []byte{},
			Attachment: []byte{},
		},
		ProductItemList: []*pos_models.ProductItem{},
		WorkflowList:    []*pos_models.Workflow{},
		Sku:             "111",
		SN:              "222111",
	}
	data.Locale = "en"
	data.ID = uint(ref)

	got := New().AddOrUpdate(uint(ref), &data)
	require.NotNilf(t, got.Product, "Expected non Nil but received %v  instead", got.ProductList)
}
func TestGetByID(t *testing.T) {

	ref := 29372112

	data := pos_models.Product{
		FoundationWithDocumentBytes: base_models.FoundationWithDocumentBytes{
			Foundation: base_models.Foundation{
				Model:          gorm.Model{},
				Name:           "PRODUCT",
				Type:           "PRODUCT",
				Stage:          "received",
				Maker:          "Maker",
				Checker:        "Checker",
				Approver:       "Approver",
				Description:    "Product ....",
				Status:         "instock",
				WorkflowLevels: 0,
				SyncToken:      0,
				Version:        0,
				Enabled:        0,
				Locale:         "",
				Events:         []byte{},
			},
			Qrcode:     []byte{},
			Barcode:    []byte{},
			Attachment: []byte{},
		},
		ProductItemList: []*pos_models.ProductItem{},
		WorkflowList:    []*pos_models.Workflow{},
		Sku:             "111",
		SN:              "222111",
	}
	data.ID = uint(ref)

	got := New().GetByID(data.ID)
	require.NotNilf(t, got.Product, "Expected non Nil but received %v  instead", got.ProductList)
}
func TestGetByName(t *testing.T) {
	ref := 29372112

	data := pos_models.Product{
		FoundationWithDocumentBytes: base_models.FoundationWithDocumentBytes{
			Foundation: base_models.Foundation{
				Model:          gorm.Model{},
				Name:           "PRODUCT",
				Type:           "PRODUCT",
				Stage:          "received",
				Maker:          "Maker",
				Checker:        "Checker",
				Approver:       "Approver",
				Description:    "Product ....",
				Status:         "instock",
				WorkflowLevels: 0,
				SyncToken:      0,
				Version:        0,
				Enabled:        0,
				Locale:         "",
				Events:         []byte{},
			},
			Qrcode:     []byte{},
			Barcode:    []byte{},
			Attachment: []byte{},
		},
		ProductItemList: []*pos_models.ProductItem{},
		WorkflowList:    []*pos_models.Workflow{},
		Sku:             "111",
		SN:              "222111",
	}
	data.ID = uint(ref)

	got := New().GetByName(data.Name)
	require.NotNilf(t, got.Product, "Expected non Nil but received %v  instead", got.ProductList)
}
func TestGetByStage(t *testing.T) {
	ref := 29372112

	data := pos_models.Product{
		FoundationWithDocumentBytes: base_models.FoundationWithDocumentBytes{
			Foundation: base_models.Foundation{
				Model:          gorm.Model{},
				Name:           "PRODUCT",
				Type:           "PRODUCT",
				Stage:          "received",
				Maker:          "Maker",
				Checker:        "Checker",
				Approver:       "Approver",
				Description:    "Product ....",
				Status:         "instock",
				WorkflowLevels: 0,
				SyncToken:      0,
				Version:        0,
				Enabled:        0,
				Locale:         "",
				Events:         []byte{},
			},
			Qrcode:     []byte{},
			Barcode:    []byte{},
			Attachment: []byte{},
		},
		ProductItemList: []*pos_models.ProductItem{},
		WorkflowList:    []*pos_models.Workflow{},
		Sku:             "111",
		SN:              "222111",
	}
	data.ID = uint(ref)
	data.Stage = "updated"
	got := New().GetByStage(data.Stage)
	require.NotNilf(t, got.Product, "Expected non Nil but received %v  instead", got.ProductList)
}
func TestGetByType(t *testing.T) {
	ref := 29372112

	data := pos_models.Product{
		FoundationWithDocumentBytes: base_models.FoundationWithDocumentBytes{
			Foundation: base_models.Foundation{
				Model:          gorm.Model{},
				Name:           "PRODUCT",
				Type:           "PRODUCT",
				Stage:          "received",
				Maker:          "Maker",
				Checker:        "Checker",
				Approver:       "Approver",
				Description:    "Product ....",
				Status:         "instock",
				WorkflowLevels: 0,
				SyncToken:      0,
				Version:        0,
				Enabled:        0,
				Locale:         "",
				Events:         []byte{},
			},
			Qrcode:     []byte{},
			Barcode:    []byte{},
			Attachment: []byte{},
		},
		ProductItemList: []*pos_models.ProductItem{},
		WorkflowList:    []*pos_models.Workflow{},
		Sku:             "111",
		SN:              "222111",
	}
	data.ID = uint(ref)

	got := New().GetByType(data.Type)
	require.NotNilf(t, got.Product, "Expected non Nil but received %v  instead", got.ProductList)
}
func TestGetByDate(t *testing.T) {

	ref := 29372112

	data := pos_models.Product{
		FoundationWithDocumentBytes: base_models.FoundationWithDocumentBytes{
			Foundation: base_models.Foundation{
				Model:          gorm.Model{},
				Name:           "PRODUCT",
				Type:           "PRODUCT",
				Stage:          "received",
				Maker:          "Maker",
				Checker:        "Checker",
				Approver:       "Approver",
				Description:    "Product ....",
				Status:         "instock",
				WorkflowLevels: 0,
				SyncToken:      0,
				Version:        0,
				Enabled:        0,
				Locale:         "",
				Events:         []byte{},
			},
			Qrcode:     []byte{},
			Barcode:    []byte{},
			Attachment: []byte{},
		},
		ProductItemList: []*pos_models.ProductItem{},
		WorkflowList:    []*pos_models.Workflow{},
		Sku:             "111",
		SN:              "222111",
	}
	data.ID = uint(ref)

	got := New().GetByDate("2021-11-03")
	require.NotNilf(t, got.Product, "Expected non Nil but received %v  instead", got.ProductList)
}
func TestGetByStatus(t *testing.T) {
	ref := 29372112

	data := pos_models.Product{
		FoundationWithDocumentBytes: base_models.FoundationWithDocumentBytes{
			Foundation: base_models.Foundation{
				Model:          gorm.Model{},
				Name:           "PRODUCT",
				Type:           "PRODUCT",
				Stage:          "received",
				Maker:          "Maker",
				Checker:        "Checker",
				Approver:       "Approver",
				Description:    "Product ....",
				Status:         "instock",
				WorkflowLevels: 0,
				SyncToken:      0,
				Version:        0,
				Enabled:        0,
				Locale:         "",
				Events:         []byte{},
			},
			Qrcode:     []byte{},
			Barcode:    []byte{},
			Attachment: []byte{},
		},
		ProductItemList: []*pos_models.ProductItem{},
		WorkflowList:    []*pos_models.Workflow{},
		Sku:             "111",
		SN:              "222111",
	}
	data.ID = uint(ref)

	got := New().GetByStatus(data.Enabled)
	require.NotNilf(t, got.Product, "Expected non Nil but received %v  instead", got.ProductList)
}
func TestGetByEnabled(t *testing.T) {

	ref := 29372112

	data := pos_models.Product{
		FoundationWithDocumentBytes: base_models.FoundationWithDocumentBytes{
			Foundation: base_models.Foundation{
				Model:          gorm.Model{},
				Name:           "PRODUCT",
				Type:           "PRODUCT",
				Stage:          "received",
				Maker:          "Maker",
				Checker:        "Checker",
				Approver:       "Approver",
				Description:    "Product ....",
				Status:         "instock",
				WorkflowLevels: 0,
				SyncToken:      0,
				Version:        0,
				Enabled:        0,
				Locale:         "",
				Events:         []byte{},
			},
			Qrcode:     []byte{},
			Barcode:    []byte{},
			Attachment: []byte{},
		},
		ProductItemList: []*pos_models.ProductItem{},
		WorkflowList:    []*pos_models.Workflow{},
		Sku:             "111",
		SN:              "222111",
	}
	data.ID = uint(ref)

	got := New().GetByEnabled(data.Enabled)
	require.NotNilf(t, got.Product, "Expected non Nil but received %v  instead", got.ProductList)
}
func TestGetByLocale(t *testing.T) {
	ref := 29372112

	data := pos_models.Product{
		FoundationWithDocumentBytes: base_models.FoundationWithDocumentBytes{
			Foundation: base_models.Foundation{
				Model:          gorm.Model{},
				Name:           "PRODUCT",
				Type:           "PRODUCT",
				Stage:          "received",
				Maker:          "Maker",
				Checker:        "Checker",
				Approver:       "Approver",
				Description:    "Product ....",
				Status:         "instock",
				WorkflowLevels: 0,
				SyncToken:      0,
				Version:        0,
				Enabled:        0,
				Locale:         "",
				Events:         []byte{},
			},
			Qrcode:     []byte{},
			Barcode:    []byte{},
			Attachment: []byte{},
		},
		ProductItemList: []*pos_models.ProductItem{},
		WorkflowList:    []*pos_models.Workflow{},
		Sku:             "111",
		SN:              "222111",
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().GetByLocate(data.Locale)
	require.NotNilf(t, got.Product, "Expected non Nil but received %v  instead", got.ProductList)
}
func TestCheckIFExists(t *testing.T) {
	ref := 29372112

	data := pos_models.Product{
		FoundationWithDocumentBytes: base_models.FoundationWithDocumentBytes{
			Foundation: base_models.Foundation{
				Model:          gorm.Model{},
				Name:           "PRODUCT",
				Type:           "PRODUCT",
				Stage:          "received",
				Maker:          "Maker",
				Checker:        "Checker",
				Approver:       "Approver",
				Description:    "Product ....",
				Status:         "instock",
				WorkflowLevels: 0,
				SyncToken:      0,
				Version:        0,
				Enabled:        0,
				Locale:         "",
				Events:         []byte{},
			},
			Qrcode:     []byte{},
			Barcode:    []byte{},
			Attachment: []byte{},
		},
		ProductItemList: []*pos_models.ProductItem{},
		WorkflowList:    []*pos_models.Workflow{},
		Sku:             "111",
		SN:              "222111",
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().CheckIFExists(data.ID)
	require.NotNilf(t, got.Product, "Expected non Nil but received %v  instead", got.ProductList)
}
func TestGetAll(t *testing.T) {
	ref := 29372112

	data := pos_models.Product{
		FoundationWithDocumentBytes: base_models.FoundationWithDocumentBytes{
			Foundation: base_models.Foundation{
				Model:          gorm.Model{},
				Name:           "PRODUCT",
				Type:           "PRODUCT",
				Stage:          "received",
				Maker:          "Maker",
				Checker:        "Checker",
				Approver:       "Approver",
				Description:    "Product ....",
				Status:         "instock",
				WorkflowLevels: 0,
				SyncToken:      0,
				Version:        0,
				Enabled:        0,
				Locale:         "",
				Events:         []byte{},
			},
			Qrcode:     []byte{},
			Barcode:    []byte{},
			Attachment: []byte{},
		},
		ProductItemList: []*pos_models.ProductItem{},
		WorkflowList:    []*pos_models.Workflow{},
		Sku:             "111",
		SN:              "222111",
	}
	data.ID = uint(ref)

	got := New().GetAll()
	require.NotNilf(t, got.Product, "Expected non Nil but received %v  instead", got.ProductList)
}
func TestDelete(t *testing.T) {
	ref := 1

	data := pos_models.Product{
		FoundationWithDocumentBytes: base_models.FoundationWithDocumentBytes{
			Foundation: base_models.Foundation{
				Model:          gorm.Model{},
				Name:           "PRODUCT",
				Type:           "PRODUCT",
				Stage:          "received",
				Maker:          "Maker",
				Checker:        "Checker",
				Approver:       "Approver",
				Description:    "Product ....",
				Status:         "instock",
				WorkflowLevels: 0,
				SyncToken:      0,
				Version:        0,
				Enabled:        0,
				Locale:         "",
				Events:         []byte{},
			},
			Qrcode:     []byte{},
			Barcode:    []byte{},
			Attachment: []byte{},
		},
		ProductItemList: []*pos_models.ProductItem{},
		WorkflowList:    []*pos_models.Workflow{},
		Sku:             "111",
		SN:              "222111",
	}
	data.ID = uint(ref)

	got := New().Delete(data.ID)
	require.LessOrEqual(t, len(got.ProductList), 0, "Expected non Nil but received %v  instead", got.ProductList)
}
