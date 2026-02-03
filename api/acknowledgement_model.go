package api

import "time"

// AcknowledgementResponse represents an acknowledgement in Centreon
type AcknowledgementResponse struct {
	// Id is the unique identifier of the acknowledgement
	Id int64 `json:"id"`

	// AuthorId is the ID of the contact who requested an acknowledgement
	AuthorId int64 `json:"author_id"`

	// AuthorName is the name of the contact who requested an acknowledgement
	AuthorName string `json:"author_name"`

	// Comment is the short description of the acknowledgement
	Comment string `json:"comment"`

	// DeletionTime is the date of the request for cancellation of the acknowledgement (ISO8601), can be null
	DeletionTime *time.Time `json:"deletion_time"`

	// EntryTime is the date of the request for acknowledgement (ISO8601)
	EntryTime time.Time `json:"entry_time"`

	// HostId is the ID of the host
	HostId int64 `json:"host_id"`

	// PollerId is the ID of the poller
	PollerId int64 `json:"poller_id"`

	// IsNotifyContacts indicates whether notification is sent to the contacts linked to the host or service
	IsNotifyContacts bool `json:"is_notify_contacts"`

	// IsPersistentComment indicates whether acknowledgement is maintained in the case of a restart of the scheduler
	IsPersistentComment bool `json:"is_persistent_comment"`

	// IsSticky indicates whether acknowledgement is maintained in the case of a change of status
	IsSticky bool `json:"is_sticky"`

	// State is the state type (1: WARNING, 2: CRITICAL, 3: UNKNOWN)
	State AcknowledgedState `json:"state"`

	// ServiceId is the ID of the service
	ServiceId int64 `json:"service_id"`
}

// AcknowledgementListHostResponse represents a host acknowledgement in list responses
type AcknowledgementListHostResponse struct {
	// Id is the unique identifier of the acknowledgement
	Id int64 `json:"id"`

	// AuthorId is the ID of the contact who requested an acknowledgement
	AuthorId int64 `json:"author_id"`

	// AuthorName is the name of the contact who requested an acknowledgement
	AuthorName string `json:"author_name"`

	// Comment is the short description of the acknowledgement
	Comment string `json:"comment"`

	// DeletionTime is the date of the request for cancellation of the acknowledgement (ISO8601), can be null
	DeletionTime *time.Time `json:"deletion_time"`

	// EntryTime is the date of the request for acknowledgement (ISO8601)
	EntryTime time.Time `json:"entry_time"`

	// HostId is the ID of the host
	HostId int64 `json:"host_id"`

	// PollerId is the ID of the poller
	PollerId int64 `json:"poller_id"`

	// IsNotifyContacts indicates whether notification is sent to the contacts linked to the host
	IsNotifyContacts bool `json:"is_notify_contacts"`

	// IsPersistentComment indicates whether acknowledgement is maintained in the case of a restart of the scheduler
	IsPersistentComment bool `json:"is_persistent_comment"`

	// IsSticky indicates whether acknowledgement is maintained in the case of a change of status
	IsSticky bool `json:"is_sticky"`

	// State is the state type (1: WARNING, 2: CRITICAL, 3: UNKNOWN)
	State AcknowledgedState `json:"state"`
}

// AcknowledgementListServiceResponse represents a service acknowledgement in list responses
type AcknowledgementListServiceResponse struct {
	// Id is the unique identifier of the acknowledgement
	Id int64 `json:"id"`

	// AuthorId is the ID of the contact who requested an acknowledgement
	AuthorId int64 `json:"author_id"`

	// AuthorName is the name of the contact who requested an acknowledgement
	AuthorName string `json:"author_name"`

	// Comment is the short description of the acknowledgement
	Comment string `json:"comment"`

	// DeletionTime is the date of the request for cancellation of the acknowledgement (ISO8601), can be null
	DeletionTime *time.Time `json:"deletion_time"`

	// EntryTime is the date of the request for acknowledgement (ISO8601)
	EntryTime time.Time `json:"entry_time"`

	// HostId is the ID of the host
	HostId int64 `json:"host_id"`

	// PollerId is the ID of the poller
	PollerId int64 `json:"poller_id"`

	// IsNotifyContacts indicates whether notification is sent to the contacts linked to the host or service
	IsNotifyContacts bool `json:"is_notify_contacts"`

	// IsPersistentComment indicates whether acknowledgement is maintained in the case of a restart of the scheduler
	IsPersistentComment bool `json:"is_persistent_comment"`

	// IsSticky indicates whether acknowledgement is maintained in the case of a change of status
	IsSticky bool `json:"is_sticky"`

	// State is the state type (1: WARNING, 2: CRITICAL, 3: UNKNOWN)
	State AcknowledgedState `json:"state"`

	// ServiceId is the ID of the service
	ServiceId int64 `json:"service_id"`
}

// AcknowledgementCreateServiceRequest represents the payload to create an acknowledgement on service in Centreon
type AcknowledgementCreateServiceRequest struct {
	// Comment is the short description of the acknowledgement
	Comment string `json:"comment" validate:"required"`

	// IsNotifyContacts indicates whether notification is sent to the contacts linked to the host or service
	IsNotifyContacts bool `json:"is_notify_contacts" validate:"required"`

	// IsPersistentComment indicates whether acknowledgement is maintained in the case of a restart of the scheduler
	IsPersistentComment bool `json:"is_persistent_comment" validate:"required"`

	// IsSticky indicates whether acknowledgement is maintained in the case of a change of status
	IsSticky bool `json:"is_sticky" validate:"required"`
}

// AcknowledgementCreateHostRequest represents the payload to create an acknowledgement on host in Centreon
type AcknowledgementCreateHostRequest struct {
	// Comment is the short description of the acknowledgement
	Comment string `json:"comment" validate:"required"`

	// IsNotifyContacts indicates whether notification is sent to the contacts linked to the host or service
	IsNotifyContacts bool `json:"is_notify_contacts" validate:"required"`

	// IsPersistentComment indicates whether acknowledgement is maintained in the case of a restart of the scheduler
	IsPersistentComment bool `json:"is_persistent_comment" validate:"required"`

	// IsSticky indicates whether acknowledgement is maintained in the case of a change of status
	IsSticky bool `json:"is_sticky" validate:"required"`

	// WithServices indicates whether we should add the acknowledgement on the host-related services
	WithServices bool `json:"with_services" validate:"required"`
}

// AcknowledgedState represents the acknowledged state of a service.
type AcknowledgedState int

const (
	AcknowledgedStateWarning  AcknowledgedState = 1
	AcknowledgedStateCritical AcknowledgedState = 2
	AcknowledgedStateUnknown  AcknowledgedState = 3
)
