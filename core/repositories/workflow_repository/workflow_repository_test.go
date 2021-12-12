package workflow_repository

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

	ref := 29372114

	data := pos_models.Workflow{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Workflow", Type: "Workflow", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Workflow", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)
	got := New().Add(&data)
	require.NotNilf(t, got.Workflow, "Expected non Nil but received %v  instead", got.WorkflowList)

}

func TestUpdate(t *testing.T) {

	ref := 29372112

	data := pos_models.Workflow{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Workflow", Type: "Workflow", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Workflow", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().Update(uint(ref), &data)
	require.NotNilf(t, got.Workflow, "Expected non Nil but received %v  instead", got.WorkflowList)
}
func TestAddOrUpdate(t *testing.T) {
	ref := 29372115

	data := pos_models.Workflow{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Workflow", Type: "Workflow", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Workflow", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.Locale = "en"
	data.ID = uint(ref)

	got := New().AddOrUpdate(uint(ref), &data)
	require.NotNilf(t, got.Workflow, "Expected non Nil but received %v  instead", got.WorkflowList)
}
func TestGetByID(t *testing.T) {

	ref := 29372112

	data := pos_models.Workflow{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Workflow", Type: "Workflow", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Workflow", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetByID(data.ID)
	require.NotNilf(t, got.Workflow, "Expected non Nil but received %v  instead", got.WorkflowList)
}
func TestGetByName(t *testing.T) {
	ref := 29372112

	data := pos_models.Workflow{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Workflow", Type: "Workflow", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Workflow", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetByName(data.Name)
	require.NotNilf(t, got.Workflow, "Expected non Nil but received %v  instead", got.WorkflowList)
}
func TestGetByStage(t *testing.T) {
	ref := 29372112

	data := pos_models.Workflow{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Workflow", Type: "Workflow", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Workflow", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)
	data.Stage = "updated"
	got := New().GetByStage(data.Stage)
	require.NotNilf(t, got.Workflow, "Expected non Nil but received %v  instead", got.WorkflowList)
}
func TestGetByType(t *testing.T) {
	ref := 29372112

	data := pos_models.Workflow{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Workflow", Type: "Workflow", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Workflow", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetByType(data.Type)
	require.NotNilf(t, got.Workflow, "Expected non Nil but received %v  instead", got.WorkflowList)
}
func TestGetByDate(t *testing.T) {

	ref := 29372112

	data := pos_models.Workflow{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Workflow", Type: "Workflow", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Workflow", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetByDate("2021-11-03")
	require.NotNilf(t, got.Workflow, "Expected non Nil but received %v  instead", got.WorkflowList)
}
func TestGetByStatus(t *testing.T) {
	ref := 29372112

	data := pos_models.Workflow{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Workflow", Type: "Workflow", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Workflow", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetByStatus(data.Enabled)
	require.NotNilf(t, got.Workflow, "Expected non Nil but received %v  instead", got.WorkflowList)
}
func TestGetByEnabled(t *testing.T) {

	ref := 29372112

	data := pos_models.Workflow{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Workflow", Type: "Workflow", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Workflow", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetByEnabled(data.Enabled)
	require.NotNilf(t, got.Workflow, "Expected non Nil but received %v  instead", got.WorkflowList)
}
func TestGetByLocale(t *testing.T) {
	ref := 29372112

	data := pos_models.Workflow{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Workflow", Type: "Workflow", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Workflow", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().GetByLocate(data.Locale)
	require.NotNilf(t, got.Workflow, "Expected non Nil but received %v  instead", got.WorkflowList)
}
func TestCheckIFExists(t *testing.T) {
	ref := 29372112

	data := pos_models.Workflow{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Workflow", Type: "Workflow", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Workflow", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().CheckIFExists(data.ID)
	require.NotNilf(t, got.Workflow, "Expected non Nil but received %v  instead", got.WorkflowList)
}
func TestGetAll(t *testing.T) {
	ref := 29372112

	data := pos_models.Workflow{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Workflow", Type: "Workflow", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Workflow", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetAll()
	require.NotNilf(t, got.Workflow, "Expected non Nil but received %v  instead", got.WorkflowList)
}
func TestDelete(t *testing.T) {
	ref := 1

	data := pos_models.Workflow{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Workflow", Type: "Workflow", Stage: "added", Maker: "maker", Checker: "checker", Approver: "approver", Description: "Workflow", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().Delete(data.ID)
	require.LessOrEqual(t, len(got.WorkflowList), 0, "Expected non Nil but received %v  instead", got.WorkflowList)
}
