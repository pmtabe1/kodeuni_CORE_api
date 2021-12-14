package customer_repository

import (
	"testing"
	"time"

	"github.com/paulmsegeya/subscription/core/models/auth_models"
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

	data := subscription_models.Customer{
		User: auth_models.User{
			Foundation:       base_models.Foundation{},
			Firstname:        "CustomerFname",
			Lastname:         "CustomerLname",
			Username:         "CustomerUsername",
			Email:            "m@p.com",
			Password:         "232323",
			Realm:            "13133",
			SecretKey:        "2232323",
			MaxRefresh:       time.Time{},
			Timeout:          time.Time{},
			IdentityKey:      "",
			VerificationLink: "",
			Key:              []byte{},
			Dob:              "",
			Mobile:           "",
			RegisterID:       0,
			TillID:           0,
			UtilizationID:    0,
			GroupID:          0,
			RoleList:         []*auth_models.Role{},
			WorkflowList:     []*workflow_models.Workflow{},
		},
		WorkflowList: []*workflow_models.Workflow{},
	}
	data.ID = uint(ref)
	got := New().Add(&data)
	require.NotNilf(t, got.Customer, "Expected non Nil but received %v  instead", got.CustomerList)

}

func TestUpdate(t *testing.T) {

	ref := 29372112

	data := subscription_models.Customer{
		User: auth_models.User{
			Foundation:       base_models.Foundation{},
			Firstname:        "CustomerFname",
			Lastname:         "CustomerLname",
			Username:         "CustomerUsername",
			Email:            "m@p.com",
			Password:         "232323",
			Realm:            "13133",
			SecretKey:        "2232323",
			MaxRefresh:       time.Time{},
			Timeout:          time.Time{},
			IdentityKey:      "",
			VerificationLink: "",
			Key:              []byte{},
			Dob:              "",
			Mobile:           "",
			RegisterID:       0,
			TillID:           0,
			UtilizationID:    0,
			GroupID:          0,
			RoleList:         []*auth_models.Role{},
			WorkflowList:     []*workflow_models.Workflow{},
		},
		WorkflowList: []*workflow_models.Workflow{},
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().Update(uint(ref), &data)
	require.NotNilf(t, got.Customer, "Expected non Nil but received %v  instead", got.CustomerList)
}
func TestAddOrUpdate(t *testing.T) {
	ref := 29372115

	data := subscription_models.Customer{
		User: auth_models.User{
			Foundation:       base_models.Foundation{},
			Firstname:        "CustomerFname",
			Lastname:         "CustomerLname",
			Username:         "CustomerUsername",
			Email:            "m@p.com",
			Password:         "232323",
			Realm:            "13133",
			SecretKey:        "2232323",
			MaxRefresh:       time.Time{},
			Timeout:          time.Time{},
			IdentityKey:      "",
			VerificationLink: "",
			Key:              []byte{},
			Dob:              "",
			Mobile:           "",
			RegisterID:       0,
			TillID:           0,
			UtilizationID:    0,
			GroupID:          0,
			RoleList:         []*auth_models.Role{},
			WorkflowList:     []*workflow_models.Workflow{},
		},
		WorkflowList: []*workflow_models.Workflow{},
	}
	data.Locale = "en"
	data.ID = uint(ref)

	got := New().AddOrUpdate(uint(ref), &data)
	require.NotNilf(t, got.Customer, "Expected non Nil but received %v  instead", got.CustomerList)
}
func TestGetByID(t *testing.T) {

	ref := 29372112

	data := subscription_models.Customer{
		User: auth_models.User{
			Foundation:       base_models.Foundation{},
			Firstname:        "CustomerFname",
			Lastname:         "CustomerLname",
			Username:         "CustomerUsername",
			Email:            "m@p.com",
			Password:         "232323",
			Realm:            "13133",
			SecretKey:        "2232323",
			MaxRefresh:       time.Time{},
			Timeout:          time.Time{},
			IdentityKey:      "",
			VerificationLink: "",
			Key:              []byte{},
			Dob:              "",
			Mobile:           "",
			RegisterID:       0,
			TillID:           0,
			UtilizationID:    0,
			GroupID:          0,
			RoleList:         []*auth_models.Role{},
			WorkflowList:     []*workflow_models.Workflow{},
		},
		WorkflowList: []*workflow_models.Workflow{},
	}
	data.ID = uint(ref)

	got := New().GetByID(data.ID)
	require.NotNilf(t, got.Customer, "Expected non Nil but received %v  instead", got.CustomerList)
}
func TestGetByName(t *testing.T) {
	ref := 29372112

	data := subscription_models.Customer{
		User: auth_models.User{
			Foundation:       base_models.Foundation{},
			Firstname:        "CustomerFname",
			Lastname:         "CustomerLname",
			Username:         "CustomerUsername",
			Email:            "m@p.com",
			Password:         "232323",
			Realm:            "13133",
			SecretKey:        "2232323",
			MaxRefresh:       time.Time{},
			Timeout:          time.Time{},
			IdentityKey:      "",
			VerificationLink: "",
			Key:              []byte{},
			Dob:              "",
			Mobile:           "",
			RegisterID:       0,
			TillID:           0,
			UtilizationID:    0,
			GroupID:          0,
			RoleList:         []*auth_models.Role{},
			WorkflowList:     []*workflow_models.Workflow{},
		},
		WorkflowList: []*workflow_models.Workflow{},
	}
	data.ID = uint(ref)

	got := New().GetByName(data.Name)
	require.NotNilf(t, got.Customer, "Expected non Nil but received %v  instead", got.CustomerList)
}
func TestGetByStage(t *testing.T) {
	ref := 29372112

	data := subscription_models.Customer{
		User: auth_models.User{
			Foundation:       base_models.Foundation{},
			Firstname:        "CustomerFname",
			Lastname:         "CustomerLname",
			Username:         "CustomerUsername",
			Email:            "m@p.com",
			Password:         "232323",
			Realm:            "13133",
			SecretKey:        "2232323",
			MaxRefresh:       time.Time{},
			Timeout:          time.Time{},
			IdentityKey:      "",
			VerificationLink: "",
			Key:              []byte{},
			Dob:              "",
			Mobile:           "",
			RegisterID:       0,
			TillID:           0,
			UtilizationID:    0,
			GroupID:          0,
			RoleList:         []*auth_models.Role{},
			WorkflowList:     []*workflow_models.Workflow{},
		},
		WorkflowList: []*workflow_models.Workflow{},
	}
	data.ID = uint(ref)
	data.Stage = "updated"
	got := New().GetByStage(data.Stage)
	require.NotNilf(t, got.Customer, "Expected non Nil but received %v  instead", got.CustomerList)
}
func TestGetByType(t *testing.T) {
	ref := 29372112

	data := subscription_models.Customer{
		User: auth_models.User{
			Foundation:       base_models.Foundation{},
			Firstname:        "CustomerFname",
			Lastname:         "CustomerLname",
			Username:         "CustomerUsername",
			Email:            "m@p.com",
			Password:         "232323",
			Realm:            "13133",
			SecretKey:        "2232323",
			MaxRefresh:       time.Time{},
			Timeout:          time.Time{},
			IdentityKey:      "",
			VerificationLink: "",
			Key:              []byte{},
			Dob:              "",
			Mobile:           "",
			RegisterID:       0,
			TillID:           0,
			UtilizationID:    0,
			GroupID:          0,
			RoleList:         []*auth_models.Role{},
			WorkflowList:     []*workflow_models.Workflow{},
		},
		WorkflowList: []*workflow_models.Workflow{},
	}
	data.ID = uint(ref)

	got := New().GetByType(data.Type)
	require.NotNilf(t, got.Customer, "Expected non Nil but received %v  instead", got.CustomerList)
}
func TestGetByDate(t *testing.T) {

	ref := 29372112

	data := subscription_models.Customer{
		User: auth_models.User{
			Foundation:       base_models.Foundation{},
			Firstname:        "CustomerFname",
			Lastname:         "CustomerLname",
			Username:         "CustomerUsername",
			Email:            "m@p.com",
			Password:         "232323",
			Realm:            "13133",
			SecretKey:        "2232323",
			MaxRefresh:       time.Time{},
			Timeout:          time.Time{},
			IdentityKey:      "",
			VerificationLink: "",
			Key:              []byte{},
			Dob:              "",
			Mobile:           "",
			RegisterID:       0,
			TillID:           0,
			UtilizationID:    0,
			GroupID:          0,
			RoleList:         []*auth_models.Role{},
			WorkflowList:     []*workflow_models.Workflow{},
		},
		WorkflowList: []*workflow_models.Workflow{},
	}
	data.ID = uint(ref)

	got := New().GetByDate("2021-11-03")
	require.NotNilf(t, got.Customer, "Expected non Nil but received %v  instead", got.CustomerList)
}
func TestGetByStatus(t *testing.T) {
	ref := 29372112

	data := subscription_models.Customer{
		User: auth_models.User{
			Foundation:       base_models.Foundation{},
			Firstname:        "CustomerFname",
			Lastname:         "CustomerLname",
			Username:         "CustomerUsername",
			Email:            "m@p.com",
			Password:         "232323",
			Realm:            "13133",
			SecretKey:        "2232323",
			MaxRefresh:       time.Time{},
			Timeout:          time.Time{},
			IdentityKey:      "",
			VerificationLink: "",
			Key:              []byte{},
			Dob:              "",
			Mobile:           "",
			RegisterID:       0,
			TillID:           0,
			UtilizationID:    0,
			GroupID:          0,
			RoleList:         []*auth_models.Role{},
			WorkflowList:     []*workflow_models.Workflow{},
		},
		WorkflowList: []*workflow_models.Workflow{},
	}
	data.ID = uint(ref)

	got := New().GetByStatus(data.Enabled)
	require.NotNilf(t, got.Customer, "Expected non Nil but received %v  instead", got.CustomerList)
}
func TestGetByEnabled(t *testing.T) {

	ref := 29372112

	data := subscription_models.Customer{
		User: auth_models.User{
			Foundation:       base_models.Foundation{},
			Firstname:        "CustomerFname",
			Lastname:         "CustomerLname",
			Username:         "CustomerUsername",
			Email:            "m@p.com",
			Password:         "232323",
			Realm:            "13133",
			SecretKey:        "2232323",
			MaxRefresh:       time.Time{},
			Timeout:          time.Time{},
			IdentityKey:      "",
			VerificationLink: "",
			Key:              []byte{},
			Dob:              "",
			Mobile:           "",
			RegisterID:       0,
			TillID:           0,
			UtilizationID:    0,
			GroupID:          0,
			RoleList:         []*auth_models.Role{},
			WorkflowList:     []*workflow_models.Workflow{},
		},
		WorkflowList: []*workflow_models.Workflow{},
	}
	data.ID = uint(ref)

	got := New().GetByEnabled(data.Enabled)
	require.NotNilf(t, got.Customer, "Expected non Nil but received %v  instead", got.CustomerList)
}
func TestGetByLocale(t *testing.T) {
	ref := 29372112

	data := subscription_models.Customer{
		User: auth_models.User{
			Foundation:       base_models.Foundation{},
			Firstname:        "CustomerFname",
			Lastname:         "CustomerLname",
			Username:         "CustomerUsername",
			Email:            "m@p.com",
			Password:         "232323",
			Realm:            "13133",
			SecretKey:        "2232323",
			MaxRefresh:       time.Time{},
			Timeout:          time.Time{},
			IdentityKey:      "",
			VerificationLink: "",
			Key:              []byte{},
			Dob:              "",
			Mobile:           "",
			RegisterID:       0,
			TillID:           0,
			UtilizationID:    0,
			GroupID:          0,
			RoleList:         []*auth_models.Role{},
			WorkflowList:     []*workflow_models.Workflow{},
		},
		WorkflowList: []*workflow_models.Workflow{},
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().GetByLocate(data.Locale)
	require.NotNilf(t, got.Customer, "Expected non Nil but received %v  instead", got.CustomerList)
}
func TestCheckIFExists(t *testing.T) {
	ref := 29372112

	data := subscription_models.Customer{
		User: auth_models.User{
			Foundation:       base_models.Foundation{},
			Firstname:        "CustomerFname",
			Lastname:         "CustomerLname",
			Username:         "CustomerUsername",
			Email:            "m@p.com",
			Password:         "232323",
			Realm:            "13133",
			SecretKey:        "2232323",
			MaxRefresh:       time.Time{},
			Timeout:          time.Time{},
			IdentityKey:      "",
			VerificationLink: "",
			Key:              []byte{},
			Dob:              "",
			Mobile:           "",
			RegisterID:       0,
			TillID:           0,
			UtilizationID:    0,
			GroupID:          0,
			RoleList:         []*auth_models.Role{},
			WorkflowList:     []*workflow_models.Workflow{},
		},
		WorkflowList: []*workflow_models.Workflow{},
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().CheckIFExists(data.ID)
	require.NotNilf(t, got.Customer, "Expected non Nil but received %v  instead", got.CustomerList)
}
func TestGetAll(t *testing.T) {
	ref := 29372112

	data := subscription_models.Customer{
		User: auth_models.User{
			Foundation:       base_models.Foundation{},
			Firstname:        "CustomerFname",
			Lastname:         "CustomerLname",
			Username:         "CustomerUsername",
			Email:            "m@p.com",
			Password:         "232323",
			Realm:            "13133",
			SecretKey:        "2232323",
			MaxRefresh:       time.Time{},
			Timeout:          time.Time{},
			IdentityKey:      "",
			VerificationLink: "",
			Key:              []byte{},
			Dob:              "",
			Mobile:           "",
			RegisterID:       0,
			TillID:           0,
			UtilizationID:    0,
			GroupID:          0,
			RoleList:         []*auth_models.Role{},
			WorkflowList:     []*workflow_models.Workflow{},
		},
		WorkflowList: []*workflow_models.Workflow{},
	}
	data.ID = uint(ref)

	got := New().GetAll()
	require.NotNilf(t, got.Customer, "Expected non Nil but received %v  instead", got.CustomerList)
}
func TestDelete(t *testing.T) {
	ref := 1

	data := subscription_models.Customer{
		User: auth_models.User{
			Foundation:       base_models.Foundation{},
			Firstname:        "CustomerFname",
			Lastname:         "CustomerLname",
			Username:         "CustomerUsername",
			Email:            "m@p.com",
			Password:         "232323",
			Realm:            "13133",
			SecretKey:        "2232323",
			MaxRefresh:       time.Time{},
			Timeout:          time.Time{},
			IdentityKey:      "",
			VerificationLink: "",
			Key:              []byte{},
			Dob:              "",
			Mobile:           "",
			RegisterID:       0,
			TillID:           0,
			UtilizationID:    0,
			GroupID:          0,
			RoleList:         []*auth_models.Role{},
			WorkflowList:     []*workflow_models.Workflow{},
		},
		WorkflowList: []*workflow_models.Workflow{},
	}
	data.ID = uint(ref)

	got := New().Delete(data.ID)
	require.LessOrEqual(t, len(got.CustomerList), 0, "Expected non Nil but received %v  instead", got.CustomerList)
}
