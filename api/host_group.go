package api

import (
	"fmt"
	"strings"
)

// HostGroup represents a host group configuration
type HostGroup struct {
	ID           *int64  `json:"id,omitempty"`
	Name         string  `json:"name" validate:"required"`
	Alias        string  `json:"alias" validate:"required"`
	Notes        string  `json:"notes,omitempty"`
	NotesURL     string  `json:"notes_url,omitempty"`
	ActionURL    string  `json:"action_url,omitempty"`
	IconImage    string  `json:"icon_image,omitempty"`
	MapIconImage string  `json:"map_icon_image,omitempty"`
	RRDRetention string  `json:"rrd_retention,omitempty"`
	IsActivated  *bool   `json:"is_activated,omitempty"`
	Comment      string  `json:"comment,omitempty"`
	GeoCoords    string  `json:"geo_coords,omitempty"`
	Members      []int64 `json:"members,omitempty"`

	// Relationships
	Hosts []Host `json:"hosts,omitempty"`
}

// Validation functions for HostGroup
func (hg *HostGroup) ValidateForCreate() error {
	if hg.Name == "" || len(strings.TrimSpace(hg.Name)) == 0 {
		return fmt.Errorf("name is required for host group creation")
	}
	if hg.Alias == "" || len(strings.TrimSpace(hg.Alias)) == 0 {
		return fmt.Errorf("alias is required for host group creation")
	}
	return nil
}

func (hg *HostGroup) ValidateForUpdate() error {
	if hg.ID == nil || *hg.ID <= 0 {
		return fmt.Errorf("valid host group ID is required for update operations")
	}
	if hg.Name != "" && len(strings.TrimSpace(hg.Name)) == 0 {
		return fmt.Errorf("host group name cannot be empty")
	}
	if hg.Alias != "" && len(strings.TrimSpace(hg.Alias)) == 0 {
		return fmt.Errorf("host group alias cannot be empty")
	}
	return nil
}

func (hg *HostGroup) ValidateForDuplicate() error {
	if hg.ID == nil || *hg.ID <= 0 {
		return fmt.Errorf("valid source host group ID is required for duplicate operations")
	}
	return nil
}
