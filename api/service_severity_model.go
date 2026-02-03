package api

// ServiceSeverityCreateOrUpdateRequest represents the payload to create or update a service severity in Centreon.
type ServiceSeverityCreateOrUpdateRequest struct {
	Name        string `json:"name" validate:"required,max=200"`
	Alias       string `json:"alias" validate:"required,max=200"`
	Level       int    `json:"level" validate:"required"`
	IconId      int64  `json:"icon_id" validate:"required"`
	IsActivated *bool  `json:"is_activated,omitempty"`
}

// ServiceSeverityResponse represents a service severity in Centreon.
type ServiceSeverityResponse struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	Level       int    `json:"level"`
	IconId      int64  `json:"icon_id"`
	IsActivated bool   `json:"is_activated"`
}
