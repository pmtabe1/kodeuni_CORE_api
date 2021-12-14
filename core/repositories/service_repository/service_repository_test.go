package service_repository

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

	ref := 29372113

	data := subscription_models.Service{
		Foundation:     base_models.Foundation{Model: gorm.Model{}, Name: "Service", Type: "Service", Stage: "added", Maker: "maker", Checker: "checker", Description: "Service", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
		SubscriptionID: 0,
		LimitationList: []*subscription_models.Limitation{},
		LicenceList:    []*subscription_models.Licence{},
		WorkflowList:   []*workflow_models.Workflow{},
	}
	data.ID = uint(ref)
	got := New().Add(&data)
	require.NotNilf(t, got.Service, "Expected non Nil but received %v  instead", got.ServiceList)

}

func TestUpdate(t *testing.T) {

	ref := 29372112

	data := subscription_models.Service{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Service", Type: "Service", Stage: "added", Maker: "maker", Checker: "checker", Description: "Service", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().Update(uint(ref), &data)
	require.NotNilf(t, got.Service, "Expected non Nil but received %v  instead", got.ServiceList)
}
func TestAddOrUpdate(t *testing.T) {
	ref := 29372115

	data := subscription_models.Service{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Service", Type: "Service", Stage: "added", Maker: "maker", Checker: "checker", Description: "Service", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.Locale = "en"
	data.ID = uint(ref)

	got := New().AddOrUpdate(uint(ref), &data)
	require.NotNilf(t, got.Service, "Expected non Nil but received %v  instead", got.ServiceList)
}
func TestGetByID(t *testing.T) {

	ref := 29372112

	data := subscription_models.Service{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Service", Type: "Service", Stage: "added", Maker: "maker", Checker: "checker", Description: "Service", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetByID(data.ID)
	require.NotNilf(t, got.Service, "Expected non Nil but received %v  instead", got.ServiceList)
}
func TestGetByName(t *testing.T) {
	ref := 29372112

	data := subscription_models.Service{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Service", Type: "Service", Stage: "added", Maker: "maker", Checker: "checker", Description: "Service", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetByName(data.Name)
	require.NotNilf(t, got.Service, "Expected non Nil but received %v  instead", got.ServiceList)
}
func TestGetByStage(t *testing.T) {
	ref := 29372112

	data := subscription_models.Service{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Service", Type: "Service", Stage: "added", Maker: "maker", Checker: "checker", Description: "Service", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)
	data.Stage = "updated"
	got := New().GetByStage(data.Stage)
	require.NotNilf(t, got.Service, "Expected non Nil but received %v  instead", got.ServiceList)
}
func TestGetByType(t *testing.T) {
	ref := 29372112

	data := subscription_models.Service{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Service", Type: "Service", Stage: "added", Maker: "maker", Checker: "checker", Description: "Service", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetByType(data.Type)
	require.NotNilf(t, got.Service, "Expected non Nil but received %v  instead", got.ServiceList)
}
func TestGetByDate(t *testing.T) {

	ref := 29372112

	data := subscription_models.Service{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Service", Type: "Service", Stage: "added", Maker: "maker", Checker: "checker", Description: "Service", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetByDate("2021-11-03")
	require.NotNilf(t, got.Service, "Expected non Nil but received %v  instead", got.ServiceList)
}
func TestGetByStatus(t *testing.T) {
	ref := 29372112

	data := subscription_models.Service{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Service", Type: "Service", Stage: "added", Maker: "maker", Checker: "checker", Description: "Service", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetByStatus(data.Enabled)
	require.NotNilf(t, got.Service, "Expected non Nil but received %v  instead", got.ServiceList)
}
func TestGetByEnabled(t *testing.T) {

	ref := 29372112

	data := subscription_models.Service{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Service", Type: "Service", Stage: "added", Maker: "maker", Checker: "checker", Description: "Service", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetByEnabled(data.Enabled)
	require.NotNilf(t, got.Service, "Expected non Nil but received %v  instead", got.ServiceList)
}
func TestGetByLocale(t *testing.T) {
	ref := 29372112

	data := subscription_models.Service{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Service", Type: "Service", Stage: "added", Maker: "maker", Checker: "checker", Description: "Service", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().GetByLocate(data.Locale)
	require.NotNilf(t, got.Service, "Expected non Nil but received %v  instead", got.ServiceList)
}
func TestCheckIFExists(t *testing.T) {
	ref := 29372112

	data := subscription_models.Service{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Service", Type: "Service", Stage: "added", Maker: "maker", Checker: "checker", Description: "Service", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().CheckIFExists(data.ID)
	require.NotNilf(t, got.Service, "Expected non Nil but received %v  instead", got.ServiceList)
}
func TestGetAll(t *testing.T) {
	ref := 29372112

	data := subscription_models.Service{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Service", Type: "Service", Stage: "added", Maker: "maker", Checker: "checker", Description: "Service", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().GetAll()
	require.NotNilf(t, got.Service, "Expected non Nil but received %v  instead", got.ServiceList)
}
func TestDelete(t *testing.T) {
	ref := 1

	data := subscription_models.Service{
		Foundation: base_models.Foundation{Model: gorm.Model{}, Name: "Service", Type: "Service", Stage: "added", Maker: "maker", Checker: "checker", Description: "Service", Status: "", WorkflowLevels: 0, SyncToken: 0, Version: 0, Enabled: 0, Locale: "", Events: []byte{}},
	}
	data.ID = uint(ref)

	got := New().Delete(data.ID)
	require.LessOrEqual(t, len(got.ServiceList), 0, "Expected non Nil but received %v  instead", got.ServiceList)
}
