package subscriber_repository

import (
	"github.com/paulmsegeya/subscription/core/models/auth_models"
	"github.com/paulmsegeya/subscription/core/models/base_models"
	"github.com/paulmsegeya/subscription/core/models/subscription_models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {

	got := New()

	require.NotNilf(t, got, "Expected non nil but got %v instead ", got)
}

func TestAdd(t *testing.T) {

	ref := 29372112

	data := subscription_models.Subscriber{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "SubscriberFname",
			Lastname:      "SubscriberLname",
			Dob:           "2020-01-10",
			Mobile:        "937927329",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
	}
	data.ID = uint(ref)
	got := New().Add(&data)
	require.NotNilf(t, got.Subscriber, "Expected non Nil but received %v  instead", got.SubscriberList)

}

func TestUpdate(t *testing.T) {

	ref := 29372112

	data := subscription_models.Subscriber{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "SubscriberFname",
			Lastname:      "SubscriberLname",
			Dob:           "2020-01-10",
			Mobile:        "937927329",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().Update(uint(ref), &data)
	require.NotNilf(t, got.Subscriber, "Expected non Nil but received %v  instead", got.SubscriberList)
}
func TestAddOrUpdate(t *testing.T) {
	ref := 29372115

	data := subscription_models.Subscriber{
		User:             auth_models.User{Foundation: base_models.Foundation{}, Firstname: "SubscriberFname", Lastname: "SubscriberLname", Dob: "2020-01-10", Mobile: "937927329", RegisterID: 0, TillID: 0, UtilizationID: 0},
		SubscriptionList: []*subscription_models.Subscription{},
	}
	data.Locale = "en"
	data.ID = uint(ref)

	got := New().AddOrUpdate(uint(ref), &data)
	require.NotNilf(t, got.Subscriber, "Expected non Nil but received %v  instead", got.SubscriberList)
}
func TestGetByID(t *testing.T) {

	ref := 29372112

	data := subscription_models.Subscriber{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "SubscriberFname",
			Lastname:      "SubscriberLname",
			Dob:           "2020-01-10",
			Mobile:        "937927329",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
	}
	data.ID = uint(ref)

	got := New().GetByID(data.ID)
	require.NotNilf(t, got.Subscriber, "Expected non Nil but received %v  instead", got.SubscriberList)
}
func TestGetByName(t *testing.T) {
	ref := 29372112

	data := subscription_models.Subscriber{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "SubscriberFname",
			Lastname:      "SubscriberLname",
			Dob:           "2020-01-10",
			Mobile:        "937927329",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
	}
	data.ID = uint(ref)

	got := New().GetByName(data.Name)
	require.NotNilf(t, got.Subscriber, "Expected non Nil but received %v  instead", got.SubscriberList)
}
func TestGetByStage(t *testing.T) {
	ref := 29372112

	data := subscription_models.Subscriber{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "SubscriberFname",
			Lastname:      "SubscriberLname",
			Dob:           "2020-01-10",
			Mobile:        "937927329",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
	}
	data.ID = uint(ref)
	data.Stage = "updated"
	got := New().GetByStage(data.Stage)
	require.NotNilf(t, got.Subscriber, "Expected non Nil but received %v  instead", got.SubscriberList)
}
func TestGetByType(t *testing.T) {
	ref := 29372112

	data := subscription_models.Subscriber{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "SubscriberFname",
			Lastname:      "SubscriberLname",
			Dob:           "2020-01-10",
			Mobile:        "937927329",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
	}
	data.ID = uint(ref)

	got := New().GetByType(data.Type)
	require.NotNilf(t, got.Subscriber, "Expected non Nil but received %v  instead", got.SubscriberList)
}
func TestGetByDate(t *testing.T) {

	ref := 29372112

	data := subscription_models.Subscriber{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "SubscriberFname",
			Lastname:      "SubscriberLname",
			Dob:           "2020-01-10",
			Mobile:        "937927329",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
	}
	data.ID = uint(ref)

	got := New().GetByDate("2021-11-03")
	require.NotNilf(t, got.Subscriber, "Expected non Nil but received %v  instead", got.SubscriberList)
}
func TestGetByStatus(t *testing.T) {
	ref := 29372112

	data := subscription_models.Subscriber{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "SubscriberFname",
			Lastname:      "SubscriberLname",
			Dob:           "2020-01-10",
			Mobile:        "937927329",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
	}
	data.ID = uint(ref)

	got := New().GetByStatus(data.Enabled)
	require.NotNilf(t, got.Subscriber, "Expected non Nil but received %v  instead", got.SubscriberList)
}
func TestGetByEnabled(t *testing.T) {

	ref := 29372112

	data := subscription_models.Subscriber{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "SubscriberFname",
			Lastname:      "SubscriberLname",
			Dob:           "2020-01-10",
			Mobile:        "937927329",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
	}
	data.ID = uint(ref)

	got := New().GetByEnabled(data.Enabled)
	require.NotNilf(t, got.Subscriber, "Expected non Nil but received %v  instead", got.SubscriberList)
}
func TestGetByLocale(t *testing.T) {
	ref := 29372112

	data := subscription_models.Subscriber{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "SubscriberFname",
			Lastname:      "SubscriberLname",
			Dob:           "2020-01-10",
			Mobile:        "937927329",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().GetByLocate(data.Locale)
	require.NotNilf(t, got.Subscriber, "Expected non Nil but received %v  instead", got.SubscriberList)
}
func TestCheckIFExists(t *testing.T) {
	ref := 29372112

	data := subscription_models.Subscriber{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "SubscriberFname",
			Lastname:      "SubscriberLname",
			Dob:           "2020-01-10",
			Mobile:        "937927329",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
	}
	data.ID = uint(ref)
	data.Locale = "en"

	got := New().CheckIFExists(data.ID)
	require.NotNilf(t, got.Subscriber, "Expected non Nil but received %v  instead", got.SubscriberList)
}
func TestGetAll(t *testing.T) {
	ref := 29372112

	data := subscription_models.Subscriber{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "SubscriberFname",
			Lastname:      "SubscriberLname",
			Dob:           "2020-01-10",
			Mobile:        "937927329",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
	}
	data.ID = uint(ref)

	got := New().GetAll()
	require.NotNilf(t, got.Subscriber, "Expected non Nil but received %v  instead", got.SubscriberList)
}
func TestDelete(t *testing.T) {
	ref := 1

	data := subscription_models.Subscriber{
		User: auth_models.User{
			Foundation:    base_models.Foundation{},
			Firstname:     "SubscriberFname",
			Lastname:      "SubscriberLname",
			Dob:           "2020-01-10",
			Mobile:        "937927329",
			RegisterID:    0,
			TillID:        0,
			UtilizationID: 0,
		},
	}
	data.ID = uint(ref)

	got := New().Delete(data.ID)
	require.LessOrEqual(t, len(got.SubscriberList), 0, "Expected non Nil but received %v  instead", got.SubscriberList)
}
