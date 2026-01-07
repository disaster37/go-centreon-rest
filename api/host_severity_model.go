package api

// HostSeverityUpdateRequest represents the payload to create or update a host severity.
type HostSeverityUpdateRequest struct {
	Name        string  `json:"name" validate:"required,max=200"`
	Alias       string  `json:"alias" validate:"required,max=200"`
	Level       int     `json:"level" validate:"required"`
	IconId      int64   `json:"icon_id" validate:"required"`
	IsActivated *bool   `json:"is_activated,omitempty"`
	Comment     *string `json:"comment,omitempty"`
}

// HostSeverityResponse represents a host severity in Centreon.
type HostSeverityResponse struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	Level       int    `json:"level"`
	IconId      int64  `json:"icon_id"`
	IsActivated bool   `json:"is_activated"`
	Comment     string `json:"comment"`
}
