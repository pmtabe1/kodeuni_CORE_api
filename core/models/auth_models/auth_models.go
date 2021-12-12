package auth_models

import (
	"time"

	"github.com/paulmsegeya/pos/core/models/base_models"
)

type Signup struct {
	Username    string `json:"username" bson:"username"`
	Mobile      string `json:"mobile" bson:"mobile"`
	Firstname   string `json:"firstname" bson:"firstname"`
	Lastname    string `json:"lastname" bson:"lastname"`
	Email       string `json:"email" bson:"email"`
	DeviceToken string `json:"deviceToken" bson:"deviceToken"`
	Password    string `json:"password" bson:"password"`
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type OnboardingResponse struct {
	StatusCode int
	Data       struct {
		Username  string
		Firstname string
		Lastname  string
		Role      string
		Token     string
		DeviceID  string
		SessionID string
	}
}
type Onboarding struct {
	StatusCode int
	Data       struct {
		Username  string
		Firstname string
		Lastname  string
		Role      string
		Token     string
		DeviceID  string
		SessionID string
	}
}

type User struct {
	base_models.Foundation
	Firstname        string
	Lastname         string
	Username         string
	Email            string
	Password         string
	Realm            string
	SecretKey        string
	MaxRefresh       time.Time
	Timeout          time.Time
	IdentityKey      string
	VerificationLink string
	Key              []byte
	Dob              string
	Mobile           string
	RegisterID       uint
	TillID           uint
	UtilizationID    uint
	GroupID          uint
	RoleList         []*Role
}

type Role struct {
	base_models.Foundation
	PermissionList []*Permission
	UserID         uint
}

type Permission struct {
	base_models.Foundation
	AclList []*Acl
	RoleID  uint
}
type Group struct {
	base_models.Foundation
	UserList []*User
}

type Acl struct {
	base_models.Foundation
	PermissionID uint
}

type Profile struct {
	base_models.Foundation
}
