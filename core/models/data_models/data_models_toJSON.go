package data_models

import (
	"encoding/json"
	"fmt"
)

func (model *Datalog) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}