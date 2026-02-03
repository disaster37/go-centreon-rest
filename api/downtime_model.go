package api

import "time"

// DowntimeResponse represents a downtime in Centreon
type DowntimeResponse struct {
	// Id is the unique identifier of the downtime
	Id int `json:"id"`

	// AuthorId is the ID of the contact who requested the downtime
	AuthorId int `json:"author_id"`

	// AuthorName is the name of the contact who requested the downtime
	AuthorName string `json:"author_name"`

	// HostId is the ID of the host on which the downtime is set
	HostId int `json:"host_id"`

	// Comment is the comment of the downtime
	Comment string `json:"comment"`

	// Duration is the downtime duration in seconds
	Duration int `json:"duration"`

	// EntryTime is the date of the request to create the downtime (ISO8601)
	EntryTime time.Time `json:"entry_time"`

	// StartTime is the scheduled start date of the downtime (ISO8601)
	StartTime time.Time `json:"start_time"`

	// EndTime is the scheduled end date of the downtime (ISO8601)
	EndTime time.Time `json:"end_time"`

	// DeletionTime is the date of cancellation of downtime (ISO8601), can be null
	DeletionTime *time.Time `json:"deletion_time"`

	// ActualStartTime is the start date of the downtime (ISO8601)
	ActualStartTime time.Time `json:"actual_start_time"`

	// ActualEndTime is the end date of the downtime (ISO8601), can be null
	ActualEndTime *time.Time `json:"actual_end_time"`

	// IsStarted indicates whether the downtime has started
	IsStarted bool `json:"is_started"`

	// IsCancelled indicates whether the downtime has been cancelled
	IsCancelled bool `json:"is_cancelled"`

	// IsFixed indicates whether the downtime is fixed
	IsFixed bool `json:"is_fixed"`

	// ServiceId is the ID of the service on which the downtime is set
	ServiceId int `json:"service_id"`
}

// DowntimeCreateHostRequest represents the payload to create a downtime on host in Centreon
type DowntimeCreateHostRequest struct {
	// StartTime is the scheduled start date of the downtime (ISO8601)
	StartTime time.Time `json:"start_time" validate:"required"`

	// EndTime is the scheduled end date of the downtime (ISO8601)
	EndTime time.Time `json:"end_time" validate:"required"`

	// IsFixed indicates whether the downtime is fixed
	IsFixed bool `json:"is_fixed" validate:"required"`

	// Duration is the downtime duration in seconds
	Duration int `json:"duration" validate:"required,gte=0"`

	// Comment is the comment of the downtime
	Comment string `json:"comment" validate:"required"`

	// WithServices indicates whether we should add the downtime on the host-related services
	WithServices bool `json:"with_services" validate:"required"`
}

// DowntimeCreateServiceRequest represents the payload to create a downtime on service in Centreon
type DowntimeCreateServiceRequest struct {
	// StartTime is the scheduled start date of the downtime (ISO8601)
	StartTime time.Time `json:"start_time" validate:"required"`

	// EndTime is the scheduled end date of the downtime (ISO8601)
	EndTime time.Time `json:"end_time" validate:"required"`

	// IsFixed indicates whether the downtime is fixed
	IsFixed bool `json:"is_fixed" validate:"required"`

	// Duration is the downtime duration in seconds
	Duration int `json:"duration" validate:"required,gte=0"`

	// Comment is the comment of the downtime
	Comment string `json:"comment" validate:"required"`
}
