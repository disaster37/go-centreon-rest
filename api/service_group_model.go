package api

// ServiceGroupResponse represents the response structure for a Service Group entity
type ServiceGroupResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Alias       *string `json:"alias,omitempty"`
	GeoCoords   *string `json:"geo_coords,omitempty"`
	Comment     *string `json:"comment,omitempty"`
	IsActivated bool    `json:"is_activated"`
}

// ServiceGroupCreateRequest represents the request structure for creating a Service Group entity
type ServiceGroupCreateOrUpdateRequest struct {
	Name        string  `json:"name" validate:"required,max=200"`
	Alias       *string `json:"alias,omitempty" validate:"omitempty,max=200"`
	GeoCoords   *string `json:"geo_coords,omitempty" validate:"omitempty,max=32"`
	Comment     *string `json:"comment,omitempty" validate:"omitempty,max=65535"`
	IsActivated *bool   `json:"is_activated,omitempty"`
}

// ServiceGroupUpdateRequest represents a partial update for a Service Group entity
type ServiceGroupUpdateRequest struct {
	Name        *string `json:"name,omitempty" validate:"omitempty,max=200"`
	Alias       *string `json:"alias,omitempty" validate:"omitempty,max=200"`
	GeoCoords   *string `json:"geo_coords,omitempty" validate:"omitempty,max=32"`
	Comment     *string `json:"comment,omitempty" validate:"omitempty,max=65535"`
	IsActivated *bool   `json:"is_activated,omitempty"`
}
