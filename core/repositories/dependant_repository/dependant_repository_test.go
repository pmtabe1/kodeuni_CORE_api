package dependant_repository

import (
	"testing"

	"github.com/paulmsegeya/pos/core/models/base_models"
	"github.com/paulmsegeya/pos/core/models/pos_models"
	"github.com/paulmsegeya/pos/core/models/auth_models"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {

	got := New()

	require.NotNilf(t, got, "Expected non nil but got %v instead ", got)
}

func TestAdd(t *testing.T) {

	ref := 29372112

	data := pos_models.Dependant{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "DependantFName",
			Lastname:      "DependantLName",
			Dob:           "2010-20-20",
			Mobile:        "36236236823",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		MemberID:     0,
		VirtualCard:  &pos_models.VirtualCard{},
		WorkflowList: []*pos_models.Workflow{},
		Wallet:       &pos_models.Wallet{},
	}
	data.ID = uint(ref)
	got := New().Add(&data)
	require.NotNilf(t, got.Dependant, "Expected non Nil but received %v  instead", got.DependantList)

}

func TestUpdate(t *testing.T) {

	ref := 29372112

	data := pos_models.Dependant{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "DependantFName",
			Lastname:      "DependantLName",
			Dob:           "2010-20-20",
			Mobile:        "36236236823",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		MemberID:     0,
		VirtualCard:  &pos_models.VirtualCard{},
		WorkflowList: []*pos_models.Workflow{},
		Wallet:       &pos_models.Wallet{},
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().Update(uint(ref), &data)
	require.NotNilf(t, got.Dependant, "Expected non Nil but received %v  instead", got.DependantList)
}
func TestAddOrUpdate(t *testing.T) {
	ref := 29372115

	data := pos_models.Dependant{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "DependantFName",
			Lastname:      "DependantLName",
			Dob:           "2010-20-20",
			Mobile:        "36236236823",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		MemberID:     0,
		VirtualCard:  &pos_models.VirtualCard{},
		WorkflowList: []*pos_models.Workflow{},
		Wallet:       &pos_models.Wallet{},
	}
	data.Locale = "en"
	data.ID = uint(ref)

	got := New().AddOrUpdate(uint(ref), &data)
	require.NotNilf(t, got.Dependant, "Expected non Nil but received %v  instead", got.DependantList)
}
func TestGetByID(t *testing.T) {

	ref := 29372112

	data := pos_models.Dependant{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "DependantFName",
			Lastname:      "DependantLName",
			Dob:           "2010-20-20",
			Mobile:        "36236236823",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		MemberID:     0,
		VirtualCard:  &pos_models.VirtualCard{},
		WorkflowList: []*pos_models.Workflow{},
		Wallet:       &pos_models.Wallet{},
	}
	data.ID = uint(ref)

	got := New().GetByID(data.ID)
	require.NotNilf(t, got.Dependant, "Expected non Nil but received %v  instead", got.DependantList)
}
func TestGetByName(t *testing.T) {
	ref := 29372112

	data := pos_models.Dependant{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "DependantFName",
			Lastname:      "DependantLName",
			Dob:           "2010-20-20",
			Mobile:        "36236236823",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		MemberID:     0,
		VirtualCard:  &pos_models.VirtualCard{},
		WorkflowList: []*pos_models.Workflow{},
		Wallet:       &pos_models.Wallet{},
	}
	data.ID = uint(ref)

	got := New().GetByName(data.Name)
	require.NotNilf(t, got.Dependant, "Expected non Nil but received %v  instead", got.DependantList)
}
func TestGetByStage(t *testing.T) {
	ref := 29372112

	data := pos_models.Dependant{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "DependantFName",
			Lastname:      "DependantLName",
			Dob:           "2010-20-20",
			Mobile:        "36236236823",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		MemberID:     0,
		VirtualCard:  &pos_models.VirtualCard{},
		WorkflowList: []*pos_models.Workflow{},
		Wallet:       &pos_models.Wallet{},
	}
	data.ID = uint(ref)
	data.Stage = "updated"
	got := New().GetByStage(data.Stage)
	require.NotNilf(t, got.Dependant, "Expected non Nil but received %v  instead", got.DependantList)
}
func TestGetByType(t *testing.T) {
	ref := 29372112

	data := pos_models.Dependant{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "DependantFName",
			Lastname:      "DependantLName",
			Dob:           "2010-20-20",
			Mobile:        "36236236823",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		MemberID:     0,
		VirtualCard:  &pos_models.VirtualCard{},
		WorkflowList: []*pos_models.Workflow{},
		Wallet:       &pos_models.Wallet{},
	}
	data.ID = uint(ref)

	got := New().GetByType(data.Type)
	require.NotNilf(t, got.Dependant, "Expected non Nil but received %v  instead", got.DependantList)
}
func TestGetByDate(t *testing.T) {

	ref := 29372112

	data := pos_models.Dependant{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "DependantFName",
			Lastname:      "DependantLName",
			Dob:           "2010-20-20",
			Mobile:        "36236236823",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		MemberID:     0,
		VirtualCard:  &pos_models.VirtualCard{},
		WorkflowList: []*pos_models.Workflow{},
		Wallet:       &pos_models.Wallet{},
	}
	data.ID = uint(ref)

	got := New().GetByDate("2021-11-03")
	require.NotNilf(t, got.Dependant, "Expected non Nil but received %v  instead", got.DependantList)
}
func TestGetByStatus(t *testing.T) {
	ref := 29372112

	data := pos_models.Dependant{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "DependantFName",
			Lastname:      "DependantLName",
			Dob:           "2010-20-20",
			Mobile:        "36236236823",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		MemberID:     0,
		VirtualCard:  &pos_models.VirtualCard{},
		WorkflowList: []*pos_models.Workflow{},
		Wallet:       &pos_models.Wallet{},
	}
	data.ID = uint(ref)

	got := New().GetByStatus(data.Enabled)
	require.NotNilf(t, got.Dependant, "Expected non Nil but received %v  instead", got.DependantList)
}
func TestGetByEnabled(t *testing.T) {

	ref := 29372112

	data := pos_models.Dependant{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "DependantFName",
			Lastname:      "DependantLName",
			Dob:           "2010-20-20",
			Mobile:        "36236236823",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		MemberID:     0,
		VirtualCard:  &pos_models.VirtualCard{},
		WorkflowList: []*pos_models.Workflow{},
		Wallet:       &pos_models.Wallet{},
	}
	data.ID = uint(ref)

	got := New().GetByEnabled(data.Enabled)
	require.NotNilf(t, got.Dependant, "Expected non Nil but received %v  instead", got.DependantList)
}
func TestGetByLocale(t *testing.T) {
	ref := 29372112

	data := pos_models.Dependant{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "DependantFName",
			Lastname:      "DependantLName",
			Dob:           "2010-20-20",
			Mobile:        "36236236823",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		MemberID:     0,
		VirtualCard:  &pos_models.VirtualCard{},
		WorkflowList: []*pos_models.Workflow{},
		Wallet:       &pos_models.Wallet{},
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().GetByLocate(data.Locale)
	require.NotNilf(t, got.Dependant, "Expected non Nil but received %v  instead", got.DependantList)
}
func TestCheckIFExists(t *testing.T) {
	ref := 29372112

	data := pos_models.Dependant{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "DependantFName",
			Lastname:      "DependantLName",
			Dob:           "2010-20-20",
			Mobile:        "36236236823",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		MemberID:     0,
		VirtualCard:  &pos_models.VirtualCard{},
		WorkflowList: []*pos_models.Workflow{},
		Wallet:       &pos_models.Wallet{},
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().CheckIFExists(data.ID)
	require.NotNilf(t, got.Dependant, "Expected non Nil but received %v  instead", got.DependantList)
}
func TestGetAll(t *testing.T) {
	ref := 29372112

	data := pos_models.Dependant{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "DependantFName",
			Lastname:      "DependantLName",
			Dob:           "2010-20-20",
			Mobile:        "36236236823",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		MemberID:     0,
		VirtualCard:  &pos_models.VirtualCard{},
		WorkflowList: []*pos_models.Workflow{},
		Wallet:       &pos_models.Wallet{},
	}
	data.ID = uint(ref)

	got := New().GetAll()
	require.NotNilf(t, got.Dependant, "Expected non Nil but received %v  instead", got.DependantList)
}
func TestDelete(t *testing.T) {
	ref := 1

	data := pos_models.Dependant{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "DependantFName",
			Lastname:      "DependantLName",
			Dob:           "2010-20-20",
			Mobile:        "36236236823",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		MemberID:     0,
		VirtualCard:  &pos_models.VirtualCard{},
		WorkflowList: []*pos_models.Workflow{},
		Wallet:       &pos_models.Wallet{},
	}
	data.ID = uint(ref)

	got := New().Delete(data.ID)
	require.LessOrEqual(t, len(got.DependantList), 0, "Expected non Nil but received %v  instead", got.DependantList)
}
