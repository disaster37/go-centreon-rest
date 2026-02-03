package api

// ServiceCategoryResponse represents the response structure for a Service Category entity
type ServiceCategoryResponse struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Alias       *string `json:"alias,omitempty"`
	IsActivated bool    `json:"is_activated"`
	Comment     *string `json:"comment,omitempty"`
}

// ServiceCategoryCreateOrUpdateRequest is the payload for creating or updating a Service Category.
type ServiceCategoryCreateOrUpdateRequest struct {
	Name        string  `json:"name" validate:"required"`
	Alias       *string `json:"alias,omitempty"`
	IsActivated bool    `json:"is_activated"`
	Comment     *string `json:"comment,omitempty"`
}
