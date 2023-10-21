package domain

import (
	"database/sql/driver"

	"github.com/goccy/go-json"
)

type Variable struct {
	Name        string `json:"name"`
	Data        string `json:"data"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type"`
}

type VariablesList []Variable

func (sla *VariablesList) Scan(value interface{}) error {
	var skills []Variable
	err := json.Unmarshal([]byte(value.(string)), &skills)
	if err != nil {
		return err
	}
	*sla = skills
	return nil
}

func (sla VariablesList) Value() (driver.Value, error) {
	val, err := json.Marshal(sla)
	return string(val), err
}
