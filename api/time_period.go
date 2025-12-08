package api

import (
	"fmt"
	"strings"
)

// TimePeriod represents a time period configuration
type TimePeriod struct {
	ID          *int64                `json:"id,omitempty"`
	Name        string                `json:"name" validate:"required"`
	Alias       string                `json:"alias" validate:"required"`
	IsActivated *bool                 `json:"is_activated,omitempty"`
	Sunday      string                `json:"sunday,omitempty"`
	Monday      string                `json:"monday,omitempty"`
	Tuesday     string                `json:"tuesday,omitempty"`
	Wednesday   string                `json:"wednesday,omitempty"`
	Thursday    string                `json:"thursday,omitempty"`
	Friday      string                `json:"friday,omitempty"`
	Saturday    string                `json:"saturday,omitempty"`
	Exceptions  []TimePeriodException `json:"exceptions,omitempty"`
	Templates   []int64               `json:"templates,omitempty"`
}

// TimePeriodException represents an exception in a time period
type TimePeriodException struct {
	ID           *int64 `json:"id,omitempty"`
	TimePeriodID *int64 `json:"time_period_id,omitempty"`
	Days         string `json:"days" validate:"required"`
	TimeRange    string `json:"time_range" validate:"required"`
}

// TimePeriodTemplate represents a time period template
type TimePeriodTemplate struct {
	ID   *int64 `json:"id,omitempty"`
	Name string `json:"name" validate:"required"`
}

// Validation functions for TimePeriod
func (tp *TimePeriod) ValidateForCreate() error {
	if tp.Name == "" || len(strings.TrimSpace(tp.Name)) == 0 {
		return fmt.Errorf("name is required for time period creation")
	}
	if tp.Alias == "" || len(strings.TrimSpace(tp.Alias)) == 0 {
		return fmt.Errorf("alias is required for time period creation")
	}
	// At least one day should be defined
	if tp.Sunday == "" && tp.Monday == "" && tp.Tuesday == "" && tp.Wednesday == "" &&
		tp.Thursday == "" && tp.Friday == "" && tp.Saturday == "" && len(tp.Exceptions) == 0 {
		return fmt.Errorf("at least one day definition or exception is required")
	}
	return nil
}

func (tp *TimePeriod) ValidateForUpdate() error {
	if tp.ID == nil || *tp.ID <= 0 {
		return fmt.Errorf("valid time period ID is required for update operations")
	}
	if tp.Name != "" && len(strings.TrimSpace(tp.Name)) == 0 {
		return fmt.Errorf("time period name cannot be empty")
	}
	if tp.Alias != "" && len(strings.TrimSpace(tp.Alias)) == 0 {
		return fmt.Errorf("time period alias cannot be empty")
	}
	return nil
}

func (tp *TimePeriod) ValidateForDuplicate() error {
	if tp.ID == nil || *tp.ID <= 0 {
		return fmt.Errorf("valid source time period ID is required for duplicate operations")
	}
	return nil
}
