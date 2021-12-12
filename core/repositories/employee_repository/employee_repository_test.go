package employee_repository

import (
	"testing"
	"time"

	"github.com/paulmsegeya/pos/core/models/auth_models"
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

	data := pos_models.Employee{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "EmployeeFname",
			Lastname:      "EmployeeLname",
			Dob:           "1980-02-01",
			Mobile:        "686868686",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		DepartmentID: 0,
		Salary:       0,
		SN:           "5555",
		SSN:          "5555-4",
		JoinedOn:     time.Time{},
		ResignedOn:   time.Time{},
		LeftOn:       time.Time{},
		ReportsTo:    "manager1",
	}
	data.ID = uint(ref)
	got := New().Add(&data)
	require.NotNilf(t, got.Employee, "Expected non Nil but received %v  instead", got.EmployeeList)

}

func TestUpdate(t *testing.T) {

	ref := 29372112

	data := pos_models.Employee{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "EmployeeFname",
			Lastname:      "EmployeeLname",
			Dob:           "1980-02-01",
			Mobile:        "686868686",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		DepartmentID: 0,
		Salary:       0,
		SN:           "5555",
		SSN:          "5555-4",
		JoinedOn:     time.Time{},
		ResignedOn:   time.Time{},
		LeftOn:       time.Time{},
		ReportsTo:    "manager1",
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().Update(uint(ref), &data)
	require.NotNilf(t, got.Employee, "Expected non Nil but received %v  instead", got.EmployeeList)
}
func TestAddOrUpdate(t *testing.T) {
	ref := 29372115

	data := pos_models.Employee{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "EmployeeFname",
			Lastname:      "EmployeeLname",
			Dob:           "1980-02-01",
			Mobile:        "686868686",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		DepartmentID: 0,
		Salary:       0,
		SN:           "5555",
		SSN:          "5555-4",
		JoinedOn:     time.Time{},
		ResignedOn:   time.Time{},
		LeftOn:       time.Time{},
		ReportsTo:    "manager1",
	}
	data.Locale = "en"
	data.ID = uint(ref)

	got := New().AddOrUpdate(uint(ref), &data)
	require.NotNilf(t, got.Employee, "Expected non Nil but received %v  instead", got.EmployeeList)
}
func TestGetByID(t *testing.T) {

	ref := 29372112

	data := pos_models.Employee{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "EmployeeFname",
			Lastname:      "EmployeeLname",
			Dob:           "1980-02-01",
			Mobile:        "686868686",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		DepartmentID: 0,
		Salary:       0,
		SN:           "5555",
		SSN:          "5555-4",
		JoinedOn:     time.Time{},
		ResignedOn:   time.Time{},
		LeftOn:       time.Time{},
		ReportsTo:    "manager1",
	}
	data.ID = uint(ref)

	got := New().GetByID(data.ID)
	require.NotNilf(t, got.Employee, "Expected non Nil but received %v  instead", got.EmployeeList)
}
func TestGetByName(t *testing.T) {
	ref := 29372112

	data := pos_models.Employee{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "EmployeeFname",
			Lastname:      "EmployeeLname",
			Dob:           "1980-02-01",
			Mobile:        "686868686",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		DepartmentID: 0,
		Salary:       0,
		SN:           "5555",
		SSN:          "5555-4",
		JoinedOn:     time.Time{},
		ResignedOn:   time.Time{},
		LeftOn:       time.Time{},
		ReportsTo:    "manager1",
	}
	data.ID = uint(ref)

	got := New().GetByName(data.Name)
	require.NotNilf(t, got.Employee, "Expected non Nil but received %v  instead", got.EmployeeList)
}
func TestGetByStage(t *testing.T) {
	ref := 29372112

	data := pos_models.Employee{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "EmployeeFname",
			Lastname:      "EmployeeLname",
			Dob:           "1980-02-01",
			Mobile:        "686868686",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		DepartmentID: 0,
		Salary:       0,
		SN:           "5555",
		SSN:          "5555-4",
		JoinedOn:     time.Time{},
		ResignedOn:   time.Time{},
		LeftOn:       time.Time{},
		ReportsTo:    "manager1",
	}
	data.ID = uint(ref)
	data.Stage = "updated"
	got := New().GetByStage(data.Stage)
	require.NotNilf(t, got.Employee, "Expected non Nil but received %v  instead", got.EmployeeList)
}
func TestGetByType(t *testing.T) {
	ref := 29372112

	data := pos_models.Employee{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "EmployeeFname",
			Lastname:      "EmployeeLname",
			Dob:           "1980-02-01",
			Mobile:        "686868686",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		DepartmentID: 0,
		Salary:       0,
		SN:           "5555",
		SSN:          "5555-4",
		JoinedOn:     time.Time{},
		ResignedOn:   time.Time{},
		LeftOn:       time.Time{},
		ReportsTo:    "manager1",
	}
	data.ID = uint(ref)

	got := New().GetByType(data.Type)
	require.NotNilf(t, got.Employee, "Expected non Nil but received %v  instead", got.EmployeeList)
}
func TestGetByDate(t *testing.T) {

	ref := 29372112

	data := pos_models.Employee{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "EmployeeFname",
			Lastname:      "EmployeeLname",
			Dob:           "1980-02-01",
			Mobile:        "686868686",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		DepartmentID: 0,
		Salary:       0,
		SN:           "5555",
		SSN:          "5555-4",
		JoinedOn:     time.Time{},
		ResignedOn:   time.Time{},
		LeftOn:       time.Time{},
		ReportsTo:    "manager1",
	}
	data.ID = uint(ref)

	got := New().GetByDate("2021-11-03")
	require.NotNilf(t, got.Employee, "Expected non Nil but received %v  instead", got.EmployeeList)
}
func TestGetByStatus(t *testing.T) {
	ref := 29372112

	data := pos_models.Employee{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "EmployeeFname",
			Lastname:      "EmployeeLname",
			Dob:           "1980-02-01",
			Mobile:        "686868686",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		DepartmentID: 0,
		Salary:       0,
		SN:           "5555",
		SSN:          "5555-4",
		JoinedOn:     time.Time{},
		ResignedOn:   time.Time{},
		LeftOn:       time.Time{},
		ReportsTo:    "manager1",
	}
	data.ID = uint(ref)

	got := New().GetByStatus(data.Enabled)
	require.NotNilf(t, got.Employee, "Expected non Nil but received %v  instead", got.EmployeeList)
}
func TestGetByEnabled(t *testing.T) {

	ref := 29372112

	data := pos_models.Employee{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "EmployeeFname",
			Lastname:      "EmployeeLname",
			Dob:           "1980-02-01",
			Mobile:        "686868686",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		DepartmentID: 0,
		Salary:       0,
		SN:           "5555",
		SSN:          "5555-4",
		JoinedOn:     time.Time{},
		ResignedOn:   time.Time{},
		LeftOn:       time.Time{},
		ReportsTo:    "manager1",
	}
	data.ID = uint(ref)

	got := New().GetByEnabled(data.Enabled)
	require.NotNilf(t, got.Employee, "Expected non Nil but received %v  instead", got.EmployeeList)
}
func TestGetByLocale(t *testing.T) {
	ref := 29372112

	data := pos_models.Employee{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "EmployeeFname",
			Lastname:      "EmployeeLname",
			Dob:           "1980-02-01",
			Mobile:        "686868686",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		DepartmentID: 0,
		Salary:       0,
		SN:           "5555",
		SSN:          "5555-4",
		JoinedOn:     time.Time{},
		ResignedOn:   time.Time{},
		LeftOn:       time.Time{},
		ReportsTo:    "manager1",
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().GetByLocate(data.Locale)
	require.NotNilf(t, got.Employee, "Expected non Nil but received %v  instead", got.EmployeeList)
}
func TestCheckIFExists(t *testing.T) {
	ref := 29372112

	data := pos_models.Employee{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "EmployeeFname",
			Lastname:      "EmployeeLname",
			Dob:           "1980-02-01",
			Mobile:        "686868686",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		DepartmentID: 0,
		Salary:       0,
		SN:           "5555",
		SSN:          "5555-4",
		JoinedOn:     time.Time{},
		ResignedOn:   time.Time{},
		LeftOn:       time.Time{},
		ReportsTo:    "manager1",
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().CheckIFExists(data.ID)
	require.NotNilf(t, got.Employee, "Expected non Nil but received %v  instead", got.EmployeeList)
}
func TestGetAll(t *testing.T) {
	ref := 29372112

	data := pos_models.Employee{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "EmployeeFname",
			Lastname:      "EmployeeLname",
			Dob:           "1980-02-01",
			Mobile:        "686868686",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		DepartmentID: 0,
		Salary:       0,
		SN:           "5555",
		SSN:          "5555-4",
		JoinedOn:     time.Time{},
		ResignedOn:   time.Time{},
		LeftOn:       time.Time{},
		ReportsTo:    "manager1",
	}
	data.ID = uint(ref)

	got := New().GetAll()
	require.NotNilf(t, got.Employee, "Expected non Nil but received %v  instead", got.EmployeeList)
}
func TestDelete(t *testing.T) {
	ref := 1

	data := pos_models.Employee{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "EmployeeFname",
			Lastname:      "EmployeeLname",
			Dob:           "1980-02-01",
			Mobile:        "686868686",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
		DepartmentID: 0,
		Salary:       0,
		SN:           "5555",
		SSN:          "5555-4",
		JoinedOn:     time.Time{},
		ResignedOn:   time.Time{},
		LeftOn:       time.Time{},
		ReportsTo:    "manager1",
	}
	data.ID = uint(ref)

	got := New().Delete(data.ID)
	require.LessOrEqual(t, len(got.EmployeeList), 0, "Expected non Nil but received %v  instead", got.EmployeeList)
}
