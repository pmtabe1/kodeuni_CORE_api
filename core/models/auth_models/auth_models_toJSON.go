package auth_models

import (
	"encoding/json"
	"fmt"
)

func (model *Signup) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}
func (model *Secret) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *SignupRepositoryResponse) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *User) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *Acl) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *Permission) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *Group) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *Role) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}
func (model *UserRepositoryResponse) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *LoginRepositoryResponse) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *Login) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}
