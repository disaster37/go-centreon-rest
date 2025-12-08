package api

// Common response structures for pagination
type ListResponse[T any] struct {
	Result []T  `json:"result"`
	Meta   Meta `json:"meta"`
}

type Meta struct {
	Page   int         `json:"page"`
	Limit  int         `json:"limit"`
	Search interface{} `json:"search,omitempty"`
	SortBy interface{} `json:"sort_by,omitempty"`
	Total  int         `json:"total"`
}

// Common filter options
type ListOptions struct {
	Page   int                    `json:"page,omitempty"`
	Limit  int                    `json:"limit,omitempty"`
	Search map[string]interface{} `json:"search,omitempty"`
	SortBy map[string]string      `json:"sort_by,omitempty"`
}

// Supporting types for interface operations
type ServiceIdentifier struct {
	HostID    int64 `json:"host_id"`
	ServiceID int64 `json:"service_id"`
}

type ResourceIdentifier struct {
	Type     string `json:"type"` // "host" or "service"
	ID       int64  `json:"id"`
	ParentID *int64 `json:"parent_id,omitempty"` // For services, the host ID
}

type AcknowledgementRequest struct {
	Comment        string `json:"comment"`
	IsSticky       *bool  `json:"is_sticky,omitempty"`
	IsPersistent   *bool  `json:"is_persistent,omitempty"`
	NotifyContacts *bool  `json:"notify_contacts,omitempty"`
	WithServices   *bool  `json:"with_services,omitempty"`
}

type DowntimeRequest struct {
	Comment      string `json:"comment"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	IsFixed      *bool  `json:"is_fixed,omitempty"`
	Duration     *int   `json:"duration,omitempty"`
	WithServices *bool  `json:"with_services,omitempty"`
}
