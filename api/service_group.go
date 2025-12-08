package api

import (
	"fmt"
	"strings"
)

// ServiceGroup represents a service group configuration
type ServiceGroup struct {
	ID           *int64 `json:"id,omitempty"`
	Name         string `json:"name" validate:"required"`
	Alias        string `json:"alias" validate:"required"`
	Notes        string `json:"notes,omitempty"`
	NotesURL     string `json:"notes_url,omitempty"`
	ActionURL    string `json:"action_url,omitempty"`
	IconImage    string `json:"icon_image,omitempty"`
	MapIconImage string `json:"map_icon_image,omitempty"`
	IsActivated  *bool  `json:"is_activated,omitempty"`
	Comment      string `json:"comment,omitempty"`

	// Relationships
	Services []Service `json:"services,omitempty"`
}

// ServiceGroupMember represents service group member for operations
type ServiceGroupMember struct {
	HostID    *int64 `json:"host_id" validate:"required"`
	ServiceID *int64 `json:"service_id" validate:"required"`
}

// Validation functions for ServiceGroup
func (sg *ServiceGroup) ValidateForCreate() error {
	if sg.Name == "" || len(strings.TrimSpace(sg.Name)) == 0 {
		return fmt.Errorf("name is required for service group creation")
	}
	if sg.Alias == "" || len(strings.TrimSpace(sg.Alias)) == 0 {
		return fmt.Errorf("alias is required for service group creation")
	}
	return nil
}

func (sg *ServiceGroup) ValidateForUpdate() error {
	if sg.ID == nil || *sg.ID <= 0 {
		return fmt.Errorf("valid service group ID is required for update operations")
	}
	if sg.Name != "" && len(strings.TrimSpace(sg.Name)) == 0 {
		return fmt.Errorf("service group name cannot be empty")
	}
	if sg.Alias != "" && len(strings.TrimSpace(sg.Alias)) == 0 {
		return fmt.Errorf("service group alias cannot be empty")
	}
	return nil
}

func (sg *ServiceGroup) ValidateForDuplicate() error {
	if sg.ID == nil || *sg.ID <= 0 {
		return fmt.Errorf("valid source service group ID is required for duplicate operations")
	}
	return nil
}
