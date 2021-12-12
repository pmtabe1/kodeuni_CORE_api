package reference_repository

import (
	"strconv"
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
	refString := strconv.Itoa(ref)

	data := pos_models.Reference{
		Foundation:       base_models.Foundation{Model: gorm.Model{}, Name: "REFERENCE", Type: "Reference", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "REFERENCE", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
		ReferenceNumber:  uint(ref),
		Reference:        refString,
		BankReference:    refString,
		ChequeReference:  refString,
		ChannelReference: refString,
		ReconReference:   refString,
	}
	data.ID = uint(ref)
	data.GlobalCounter = uint(ref)
	data.DailyCounter = uint(ref)
	got := New().Add(&data)
	require.NotNilf(t, got.Reference, "Expected non Nil but received %v  instead", got.ReferenceList)

}

func TestUpdate(t *testing.T) {

	ref :=1 //29372111
	refString := strconv.Itoa(ref)

	data := pos_models.Reference{
		Foundation:       base_models.Foundation{Model: gorm.Model{}, Name: "REFERENCE", Type: "Reference", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "REFERENCE", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
		ReferenceNumber:  uint(ref),
		Reference:        refString,
		BankReference:    refString,
		ChequeReference:  refString,
		ChannelReference: refString,
		ReconReference:   refString,
	}
	data.ID = uint(ref)
	data.Locale="en"
	data.GlobalCounter = uint(ref)
	data.DailyCounter = uint(ref)
	got := New().Update(uint(ref), &data)
	require.NotNilf(t, got.Reference, "Expected non Nil but received %v  instead", got.ReferenceList)
}
func TestAddOrUpdate(t *testing.T) {
	ref := 29372115
	refString := strconv.Itoa(ref)

	data := pos_models.Reference{
		Foundation:       base_models.Foundation{Model: gorm.Model{}, Name: "REFERENCE", Type: "Reference", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "REFERENCE", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
		ReferenceNumber:  uint(ref),
		Reference:        refString,
		BankReference:    refString,
		ChequeReference:  refString,
		ChannelReference: refString,
		ReconReference:   refString,
	}
	data.ID = uint(ref)
	data.GlobalCounter = uint(ref)
	data.DailyCounter = uint(ref)
	got := New().AddOrUpdate(uint(ref), &data)
	require.NotNilf(t, got.Reference, "Expected non Nil but received %v  instead", got.ReferenceList)
}
func TestGetByID(t *testing.T) {

	ref := 29372112
	refString := strconv.Itoa(ref)

	data := pos_models.Reference{
		Foundation:       base_models.Foundation{Model: gorm.Model{}, Name: "REFERENCE", Type: "Reference", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "REFERENCE", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
		ReferenceNumber:  uint(ref),
		Reference:        refString,
		BankReference:    refString,
		ChequeReference:  refString,
		ChannelReference: refString,
		ReconReference:   refString,
	}
	data.ID = uint(ref)
	data.GlobalCounter = uint(ref)
	data.DailyCounter = uint(ref)
	got := New().GetByID(data.ID)
	require.NotNilf(t, got.Reference, "Expected non Nil but received %v  instead", got.ReferenceList)
}
func TestGetByName(t *testing.T) {
	ref := 29372112
	refString := strconv.Itoa(ref)

	data := pos_models.Reference{
		Foundation:       base_models.Foundation{Model: gorm.Model{}, Name: "REFERENCE", Type: "Reference", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "REFERENCE", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
		ReferenceNumber:  uint(ref),
		Reference:        refString,
		BankReference:    refString,
		ChequeReference:  refString,
		ChannelReference: refString,
		ReconReference:   refString,
	}
	data.ID = uint(ref)
	data.GlobalCounter = uint(ref)
	data.DailyCounter = uint(ref)
	got := New().GetByName(data.Name)
	require.NotNilf(t, got.Reference, "Expected non Nil but received %v  instead", got.ReferenceList)
}
func TestGetByStage(t *testing.T) {
	ref := 29372112
	refString := strconv.Itoa(ref)

	data := pos_models.Reference{
		Foundation:       base_models.Foundation{Model: gorm.Model{}, Name: "REFERENCE", Type: "Reference", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "REFERENCE", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
		ReferenceNumber:  uint(ref),
		Reference:        refString,
		BankReference:    refString,
		ChequeReference:  refString,
		ChannelReference: refString,
		ReconReference:   refString,
	}
	data.ID = uint(ref)
	data.GlobalCounter = uint(ref)
	data.DailyCounter = uint(ref)
	got := New().GetByStage(data.Stage)
	require.NotNilf(t, got.Reference, "Expected non Nil but received %v  instead", got.ReferenceList)
}
func TestGetByType(t *testing.T) {
	ref := 29372112
	refString := strconv.Itoa(ref)

	data := pos_models.Reference{
		Foundation:       base_models.Foundation{Model: gorm.Model{}, Name: "REFERENCE", Type: "Reference", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "REFERENCE", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
		ReferenceNumber:  uint(ref),
		Reference:        refString,
		BankReference:    refString,
		ChequeReference:  refString,
		ChannelReference: refString,
		ReconReference:   refString,
	}
	data.ID = uint(ref)
	data.GlobalCounter = uint(ref)
	data.DailyCounter = uint(ref)
	got := New().GetByType(data.Type)
	require.NotNilf(t, got.Reference, "Expected non Nil but received %v  instead", got.ReferenceList)
}
func TestGetByDate(t *testing.T) {

	ref := 29372112
	refString := strconv.Itoa(ref)

	data := pos_models.Reference{
		Foundation:       base_models.Foundation{Model: gorm.Model{}, Name: "REFERENCE", Type: "Reference", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "REFERENCE", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
		ReferenceNumber:  uint(ref),
		Reference:        refString,
		BankReference:    refString,
		ChequeReference:  refString,
		ChannelReference: refString,
		ReconReference:   refString,
	}
	data.ID = uint(ref)
	data.GlobalCounter = uint(ref)
	data.DailyCounter = uint(ref)
	got := New().GetByDate("2021-11-03")
	require.NotNilf(t, got.Reference, "Expected non Nil but received %v  instead", got.ReferenceList)
}
func TestGetByStatus(t *testing.T) {
	ref := 29372112
	refString := strconv.Itoa(ref)

	data := pos_models.Reference{
		Foundation:       base_models.Foundation{Model: gorm.Model{}, Name: "REFERENCE", Type: "Reference", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "REFERENCE", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
		ReferenceNumber:  uint(ref),
		Reference:        refString,
		BankReference:    refString,
		ChequeReference:  refString,
		ChannelReference: refString,
		ReconReference:   refString,
	}
	data.ID = uint(ref)
	data.GlobalCounter = uint(ref)
	data.DailyCounter = uint(ref)
	got := New().GetByStatus(data.Enabled)
	require.NotNilf(t, got.Reference, "Expected non Nil but received %v  instead", got.ReferenceList)
}
func TestGetByEnabled(t *testing.T) {

	ref := 29372112
	refString := strconv.Itoa(ref)

	data := pos_models.Reference{
		Foundation:       base_models.Foundation{Model: gorm.Model{}, Name: "REFERENCE", Type: "Reference", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "REFERENCE", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
		ReferenceNumber:  uint(ref),
		Reference:        refString,
		BankReference:    refString,
		ChequeReference:  refString,
		ChannelReference: refString,
		ReconReference:   refString,
	}
	data.ID = uint(ref)
	data.GlobalCounter = uint(ref)
	data.DailyCounter = uint(ref)
	got := New().GetByEnabled(data.Enabled)
	require.NotNilf(t, got.Reference, "Expected non Nil but received %v  instead", got.ReferenceList)
}
func TestGetByLocale(t *testing.T) {
	ref := 29372112
	refString := strconv.Itoa(ref)

	data := pos_models.Reference{
		Foundation:       base_models.Foundation{Model: gorm.Model{}, Name: "REFERENCE", Type: "Reference", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "REFERENCE", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
		ReferenceNumber:  uint(ref),
		Reference:        refString,
		BankReference:    refString,
		ChequeReference:  refString,
		ChannelReference: refString,
		ReconReference:   refString,
	}
	data.ID = uint(ref)
	data.Locale="en"
	data.GlobalCounter = uint(ref)
	data.DailyCounter = uint(ref)
	got := New().GetByLocate(data.Locale)
	require.NotNilf(t, got.Reference, "Expected non Nil but received %v  instead", got.ReferenceList)
}
func TestCheckIFExists(t *testing.T) {
	ref := 29372112
	refString := strconv.Itoa(ref)

	data := pos_models.Reference{
		Foundation:       base_models.Foundation{Model: gorm.Model{}, Name: "REFERENCE", Type: "Reference", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "REFERENCE", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
		ReferenceNumber:  uint(ref),
		Reference:        refString,
		BankReference:    refString,
		ChequeReference:  refString,
		ChannelReference: refString,
		ReconReference:   refString,
	}
	data.ID = uint(ref)
	data.Locale="en"
	data.GlobalCounter = uint(ref)
	data.DailyCounter = uint(ref)
	got := New().CheckIFExists(data.ID)
	require.NotNilf(t, got.Reference, "Expected non Nil but received %v  instead", got.ReferenceList)
}
func TestGetAll(t *testing.T) {
	ref := 29372112
	refString := strconv.Itoa(ref)

	data := pos_models.Reference{
		Foundation:       base_models.Foundation{Model: gorm.Model{}, Name: "REFERENCE", Type: "Reference", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "REFERENCE", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
		ReferenceNumber:  uint(ref),
		Reference:        refString,
		BankReference:    refString,
		ChequeReference:  refString,
		ChannelReference: refString,
		ReconReference:   refString,
	}
	data.ID = uint(ref)
	data.GlobalCounter = uint(ref)
	data.DailyCounter = uint(ref)
	got := New().GetAll()
	require.NotNilf(t, got.Reference, "Expected non Nil but received %v  instead", got.ReferenceList)
}
func TestDelete(t *testing.T) {
	ref := 1
	refString := strconv.Itoa(ref)

	data := pos_models.Reference{
		Foundation:       base_models.Foundation{Model: gorm.Model{}, Name: "REFERENCE", Type: "Reference", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "REFERENCE", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
		ReferenceNumber:  uint(ref),
		Reference:        refString,
		BankReference:    refString,
		ChequeReference:  refString,
		ChannelReference: refString,
		ReconReference:   refString,
	}
	data.ID = uint(ref)
	data.GlobalCounter = uint(ref)
	data.DailyCounter = uint(ref)
	got := New().Delete(data.ID)
	require.LessOrEqual(t,len(got.ReferenceList),0, "Expected non Nil but received %v  instead", got.ReferenceList)
}
