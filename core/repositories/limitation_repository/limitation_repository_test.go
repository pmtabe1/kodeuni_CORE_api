package limitation_repository

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/paulmsegeya/subscription/core/models/base_models"
	"github.com/paulmsegeya/subscription/core/models/subscription_models"
	"github.com/paulmsegeya/subscription/core/models/workflow_models"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {

	got := New()

	require.NotNilf(t, got, "Expected non nil but got %v instead ", got)
}

func TestAdd(t *testing.T) {

	ref := 29372112

	data := subscription_models.Limitation{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Limitation", Type: "Limitation", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Limitation", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)
	got := New().Add(&data)
	require.NotNilf(t, got.Limitation, "Expected non Nil but received %v  instead", got.LimitationList)

}

func TestUpdate(t *testing.T) {

	ref := 29372112

	data := subscription_models.Limitation{
		Foundation:   base_models.Foundation{Model: gorm.Model{}, Name: "Limitation", Type: "Limitation", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Limitation", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
		ProductID:    0,
		WorkflowList: []*workflow_models.Workflow{},
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().Update(uint(ref), &data)
	require.NotNilf(t, got.Limitation, "Expected non Nil but received %v  instead", got.LimitationList)
}
func TestAddOrUpdate(t *testing.T) {
	ref := 29372115

	data := subscription_models.Limitation{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Limitation", Type: "Limitation", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Limitation", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.Locale = "en"
	data.ID = uint(ref)

	got := New().AddOrUpdate(uint(ref), &data)
	require.NotNilf(t, got.Limitation, "Expected non Nil but received %v  instead", got.LimitationList)
}
func TestGetByID(t *testing.T) {

	ref := 29372112

	data := subscription_models.Limitation{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Limitation", Type: "Limitation", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Limitation", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetByID(data.ID)
	require.NotNilf(t, got.Limitation, "Expected non Nil but received %v  instead", got.LimitationList)
}
func TestGetByName(t *testing.T) {
	ref := 29372112

	data := subscription_models.Limitation{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Limitation", Type: "Limitation", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Limitation", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetByName(data.Name)
	require.NotNilf(t, got.Limitation, "Expected non Nil but received %v  instead", got.LimitationList)
}
func TestGetByStage(t *testing.T) {
	ref := 29372112

	data := subscription_models.Limitation{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Limitation", Type: "Limitation", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Limitation", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)
	data.Stage = "updated"
	got := New().GetByStage(data.Stage)
	require.NotNilf(t, got.Limitation, "Expected non Nil but received %v  instead", got.LimitationList)
}
func TestGetByType(t *testing.T) {
	ref := 29372112

	data := subscription_models.Limitation{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Limitation", Type: "Limitation", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Limitation", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetByType(data.Type)
	require.NotNilf(t, got.Limitation, "Expected non Nil but received %v  instead", got.LimitationList)
}
func TestGetByDate(t *testing.T) {

	ref := 29372112

	data := subscription_models.Limitation{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Limitation", Type: "Limitation", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Limitation", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetByDate("2021-11-03")
	require.NotNilf(t, got.Limitation, "Expected non Nil but received %v  instead", got.LimitationList)
}
func TestGetByStatus(t *testing.T) {
	ref := 29372112

	data := subscription_models.Limitation{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Limitation", Type: "Limitation", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Limitation", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetByStatus(data.Enabled)
	require.NotNilf(t, got.Limitation, "Expected non Nil but received %v  instead", got.LimitationList)
}
func TestGetByEnabled(t *testing.T) {

	ref := 29372112

	data := subscription_models.Limitation{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Limitation", Type: "Limitation", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Limitation", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetByEnabled(data.Enabled)
	require.NotNilf(t, got.Limitation, "Expected non Nil but received %v  instead", got.LimitationList)
}
func TestGetByLocale(t *testing.T) {
	ref := 29372112

	data := subscription_models.Limitation{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Limitation", Type: "Limitation", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Limitation", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().GetByLocate(data.Locale)
	require.NotNilf(t, got.Limitation, "Expected non Nil but received %v  instead", got.LimitationList)
}
func TestCheckIFExists(t *testing.T) {
	ref := 29372112

	data := subscription_models.Limitation{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Limitation", Type: "Limitation", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Limitation", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().CheckIFExists(data.ID)
	require.NotNilf(t, got.Limitation, "Expected non Nil but received %v  instead", got.LimitationList)
}
func TestGetAll(t *testing.T) {
	ref := 29372112

	data := subscription_models.Limitation{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Limitation", Type: "Limitation", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Limitation", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetAll()
	require.NotNilf(t, got.Limitation, "Expected non Nil but received %v  instead", got.LimitationList)
}
func TestDelete(t *testing.T) {
	ref := 1

	data := subscription_models.Limitation{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Limitation", Type: "Limitation", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Limitation", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().Delete(data.ID)
	require.LessOrEqual(t, len(got.LimitationList), 0, "Expected non Nil but received %v  instead", got.LimitationList)
}
