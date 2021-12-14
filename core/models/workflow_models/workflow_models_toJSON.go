package workflow_models

import (
	"encoding/json"
	"fmt"
)

func (model *Workflow) ToJSON() string {
	jsonBytes, _ := json.Marshal(model)
	x := fmt.Sprintf("%v", string(jsonBytes))
	return x
}
