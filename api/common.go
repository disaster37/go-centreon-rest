package api

import (
	"encoding/json"
	"fmt"
	"time"
)

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

type ResponseMessage struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// IdName represents a simple structure with ID and Name
type IdName struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (h ListOptions) GetQueryParams() map[string]string {
	params := make(map[string]string)

	if h.Search != nil {
		b, err := json.Marshal(h.Search)
		if err != nil {
			panic(fmt.Sprintf("failed to marshal search parameters: %s", err.Error()))
		}

		params["search"] = string(b)
	}

	if h.Page > 0 {
		params["page"] = fmt.Sprintf("%d", h.Page)
	}

	if h.Limit > 0 {
		params["limit"] = fmt.Sprintf("%d", h.Limit)
	}

	if h.SortBy != nil {
		b, err := json.Marshal(h.SortBy)
		if err != nil {
			panic(fmt.Sprintf("failed to marshal sort_by parameters: %s", err.Error()))
		}

		params["sort_by"] = string(b)
	}

	return params
}

type Timestamp struct {
	time.Time
}

// UnmarshalJSON decodes an int64 timestamp into a time.Time object
func (p *Timestamp) UnmarshalJSON(bytes []byte) error {
	// 1. Decode the bytes into an int64
	var raw int64
	err := json.Unmarshal(bytes, &raw)

	if err != nil {
		fmt.Printf("error decoding timestamp: %s\n", err)
		return err
	}

	// 2. Parse the unix timestamp
	p.Time = time.Unix(raw, 0)
	return nil
}
