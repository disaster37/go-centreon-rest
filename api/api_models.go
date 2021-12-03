package centreonapi

import (
	"encoding/json"
	"fmt"
)

// Payload represent payload API object
type Payload struct {
	Action string `json:"action,omitempty"`
	Object string `json:"object,omitempty"`
	Values string `json:"values,omitempty"`
}

// Result represent API result object
type Result struct {
	Result json.RawMessage `json:"result,omitempty"`
}

// NewPayload permit to init payload object
func NewPayload(action, object, values string, params ...interface{}) *Payload {
	return &Payload{
		Action: action,
		Object: object,
		Values: fmt.Sprintf(values, params...),
	}
}

type ResultTest struct {
	Result interface{} `json:"result,omitempty"`
}
