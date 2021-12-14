package subscription_models

import (
	"encoding/json"
	"fmt"
)

func (model *Customer) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *Subscriber) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *Subscription) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *Team) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *Staff) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *Service) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}
func (model *Schedule) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *Support) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *Licence) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}
func (model *Agreement) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}
func (model *Limitation) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *Contract) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *Notification) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *Department) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *Product) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}

func (model *Contact) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}
