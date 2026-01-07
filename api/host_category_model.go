package api

type HostCategoryUpdateRequest struct {
	Name        string  `json:"name" validate:"required,max=200"`
	Alias       string  `json:"alias" validate:"required,max=200"`
	IsActivated *bool   `json:"is_activated,omitempty"`
	Comment     *string `json:"comment,omitempty"`
}

// HostCategory represents a host category in Centreon.
type HostCategoryResponse struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	IsActivated bool   `json:"is_activated"`
	Comment     string `json:"comment"`
}
