package api

import (
	"context"
	"time"
)

// Acknowledgement represents acknowledgement information for hosts and services in Centreon monitoring.
// Acknowledgements are used to indicate that an administrator is aware of a problem and is working on it,
// which helps prevent duplicate notifications and provides accountability for issue resolution.
type Acknowledgement struct {
	// ID is the unique identifier for the acknowledgement record
	ID *int64 `json:"id,omitempty"`

	// EntryTime is the timestamp when the acknowledgement was created in the system
	EntryTime *time.Time `json:"entry_time,omitempty"`

	// Author is the username of the person who created the acknowledgement
	// This field helps track who acknowledged the issue for accountability purposes
	Author string `json:"author,omitempty" validate:"omitempty,min=1,max=255"`

	// Comment is a required description explaining the acknowledgement reason or planned actions
	// This provides context about why the issue was acknowledged and what steps are being taken
	Comment string `json:"comment" validate:"required,min=1,max=65535"`

	// DeletionTime is the timestamp when the acknowledgement was removed from the system
	// This is set when the acknowledgement is manually deleted or automatically cleared
	DeletionTime *time.Time `json:"deletion_time,omitempty"`

	// IsSticky indicates whether the acknowledgement persists through state changes
	// When true, the acknowledgement remains active even if the resource changes state
	// When false, the acknowledgement is automatically cleared on state changes
	IsSticky *bool `json:"is_sticky,omitempty"`

	// IsPersistent determines if the acknowledgement survives Centreon engine restarts
	// When true, the acknowledgement is preserved across monitoring engine restarts
	// When false, the acknowledgement is lost if the monitoring engine is restarted
	IsPersistent *bool `json:"is_persistent,omitempty"`

	// NotifyContacts specifies whether to send notifications to contacts when acknowledging
	// When true, all contacts associated with the resource will be notified of the acknowledgement
	// This helps inform the team that someone is working on the issue
	NotifyContacts *bool `json:"notify_contacts,omitempty"`

	// WithServices indicates whether to acknowledge associated services when acknowledging a host
	// Only applies to host acknowledgements - when true, all services of the host are also acknowledged
	// This prevents service alerts when the underlying host issue is being addressed
	WithServices *bool `json:"with_services,omitempty"`

	// ResourceType specifies the type of resource being acknowledged ("host" or "service")
	// This determines how the acknowledgement is applied and what validation rules are used
	ResourceType string `json:"resource_type,omitempty" validate:"omitempty,resource_type"`

	// ResourceID is the unique identifier of the resource (host or service) being acknowledged
	// For services, this represents the service ID; for hosts, this represents the host ID
	ResourceID *int64 `json:"resource_id,omitempty" validate:"omitempty,gt=0"`

	// ParentResourceID is the host ID when acknowledging a service
	// For service acknowledgements, this links the service to its parent host
	// For host acknowledgements, this field is typically not used
	ParentResourceID *int64 `json:"parent_resource_id,omitempty" validate:"omitempty,gt=0"`

	// ResourceName is the display name of the resource being acknowledged
	// For hosts, this is the hostname; for services, this is the service description
	ResourceName string `json:"resource_name,omitempty" validate:"omitempty,min=1,max=200"`

	// ParentName is the hostname when acknowledging a service
	// This provides the host context for service acknowledgements
	// For host acknowledgements, this field is typically not used
	ParentName *string `json:"parent_name,omitempty" validate:"omitempty,min=1,max=200"`
}

// AcknowledgementInterface defines operations for Acknowledgement entities
type AcknowledgementInterface interface {
	// GetByID retrieves an acknowledgement by its ID
	GetByID(ctx context.Context, id int64) (*Acknowledgement, error)

	// List retrieves acknowledgements with optional filtering and pagination
	List(ctx context.Context, opts *ListOptions) (*ListResponse[Acknowledgement], error)

	// ListForHost retrieves acknowledgements for a specific host
	ListForHost(ctx context.Context, hostID int64, opts *ListOptions) (*ListResponse[Acknowledgement], error)

	// ListForService retrieves acknowledgements for a specific service
	ListForService(ctx context.Context, hostID, serviceID int64, opts *ListOptions) (*ListResponse[Acknowledgement], error)

	// AcknowledgeHost acknowledges a host
	AcknowledgeHost(ctx context.Context, hostID int64, ack *AcknowledgementRequest) error

	// AcknowledgeService acknowledges a service
	AcknowledgeService(ctx context.Context, hostID, serviceID int64, ack *AcknowledgementRequest) error

	// DisacknowledgeHost removes acknowledgement from a host
	DisacknowledgeHost(ctx context.Context, hostID int64) error

	// DisacknowledgeService removes acknowledgement from a service
	DisacknowledgeService(ctx context.Context, hostID, serviceID int64) error
}
