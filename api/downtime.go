package api

import (
	"context"
	"time"
)

// Downtime represents scheduled maintenance windows for hosts and services in Centreon.
// Downtimes are used to suppress notifications and indicate planned maintenance activities,
// preventing false alerts during known maintenance periods or system upgrades.
type Downtime struct {
	// ID is the unique identifier for the downtime record
	ID *int64 `json:"id,omitempty"`

	// EntryTime is the timestamp when the downtime was created in the system
	// This records when the downtime was scheduled, not when it becomes active
	EntryTime *time.Time `json:"entry_time,omitempty"`

	// StartTime is the scheduled start time for the downtime period
	// Monitoring notifications will be suppressed starting from this time
	StartTime time.Time `json:"start_time" validate:"required"`

	// EndTime is the scheduled end time for the downtime period
	// Monitoring notifications will resume after this time (must be after StartTime)
	EndTime time.Time `json:"end_time" validate:"required,gtfield=StartTime"`

	// Duration specifies the length of the downtime in seconds for flexible downtimes
	// Used for flexible downtimes that start when a problem occurs during the time window
	Duration *int64 `json:"duration,omitempty" validate:"omitempty,gte=0"`

	// Fixed indicates whether this is a fixed or flexible downtime
	// Fixed downtimes run for the entire scheduled period regardless of resource state
	// Flexible downtimes only start when the resource enters a problem state
	Fixed *bool `json:"fixed,omitempty"`

	// ActualStartTime is when the downtime actually became active
	// For fixed downtimes, this matches StartTime; for flexible downtimes, this is when problems began
	ActualStartTime *time.Time `json:"actual_start_time,omitempty"`

	// ActualEndTime is when the downtime actually ended
	// May differ from EndTime if the downtime was cancelled early or extended
	ActualEndTime *time.Time `json:"actual_end_time,omitempty"`

	// IsStarted indicates whether the downtime period has begun
	// True when the current time is within the downtime window
	IsStarted *bool `json:"is_started,omitempty"`

	// IsCancelled indicates whether the downtime was manually cancelled before completion
	// Cancelled downtimes stop suppressing notifications immediately
	IsCancelled *bool `json:"is_cancelled,omitempty"`

	// Comment is a required description explaining the reason for the downtime
	// Should describe the maintenance activity or reason for the scheduled downtime
	Comment string `json:"comment" validate:"required,min=1,max=65535"`

	// DeletionTime is the timestamp when the downtime was removed from the system
	// Set when the downtime is manually deleted or automatically cleaned up
	DeletionTime *time.Time `json:"deletion_time,omitempty"`

	// Author is the username of the person who scheduled the downtime
	// Provides accountability and contact information for the maintenance activity
	Author string `json:"author,omitempty" validate:"omitempty,min=1,max=255"`

	// ResourceType specifies the type of resource ("host" or "service")
	// Determines how the downtime is applied and what validation rules are used
	ResourceType string `json:"resource_type,omitempty" validate:"omitempty,resource_type"`

	// ResourceID is the unique identifier of the resource (host or service) in downtime
	// For services, this represents the service ID; for hosts, this represents the host ID
	ResourceID *int64 `json:"resource_id,omitempty" validate:"omitempty,gt=0"`

	// ParentResourceID is the host ID when scheduling downtime for a service
	// For service downtimes, this links the service to its parent host
	// For host downtimes, this field is typically not used
	ParentResourceID *int64 `json:"parent_resource_id,omitempty" validate:"omitempty,gt=0"`

	// ResourceName is the display name of the resource in downtime
	// For hosts, this is the hostname; for services, this is the service description
	ResourceName string `json:"resource_name,omitempty" validate:"omitempty,min=1,max=200"`

	// ParentName is the hostname when scheduling downtime for a service
	// This provides the host context for service downtimes
	// For host downtimes, this field is typically not used
	ParentName *string `json:"parent_name,omitempty" validate:"omitempty,min=1,max=200"`

	// WithServices indicates whether to include associated services when scheduling host downtime
	// Only applies to host downtimes - when true, all services of the host are also put in downtime
	// This prevents service alerts when the underlying host is undergoing maintenance
	WithServices *bool `json:"with_services,omitempty"`
}

// ListMeta represents metadata for paginated API responses in Centreon.
// This structure provides information about the current page, total results,
// and search/sort parameters for efficient data navigation and filtering.
type ListMeta struct {
	// Page is the current page number in the paginated results (starts from 1)
	Page int `json:"page" validate:"gte=1"`

	// Limit is the maximum number of items returned per page (1-1000)
	// Controls the page size for performance and usability optimization
	Limit int `json:"limit" validate:"gte=1,lte=1000"`

	// Search contains the search criteria applied to filter results
	// Can be a string, object, or array depending on the search complexity
	Search any `json:"search"`

	// Sort contains the sorting parameters applied to order results
	// Defines field names and sort directions (ascending/descending)
	Sort any `json:"sort"`

	// Total is the total number of items matching the search criteria
	// Used for calculating pagination information and progress indicators
	Total int `json:"total" validate:"gte=0"`
}

// DowntimeInterface defines CRUD operations for Downtime entities
type DowntimeInterface interface {
	// Create creates a new downtime for a host or service
	Create(ctx context.Context, downtime *Downtime) (*Downtime, error)

	// CreateForHost creates a downtime for a specific host
	CreateForHost(ctx context.Context, hostID int64, request *DowntimeRequest) (*Downtime, error)

	// CreateForService creates a downtime for a specific service
	CreateForService(ctx context.Context, hostID, serviceID int64, request *DowntimeRequest) (*Downtime, error)

	// CreateForHostsByName creates downtimes for hosts by their names
	CreateForHostsByName(ctx context.Context, hostNames []string, request *DowntimeRequest) ([]Downtime, error)

	// CreateForServicesByName creates downtimes for services by their names
	CreateForServicesByName(ctx context.Context, serviceIdentifiers []ServiceIdentifier, request *DowntimeRequest) ([]Downtime, error)

	// GetByID retrieves a downtime by its ID
	GetByID(ctx context.Context, id int64) (*Downtime, error)

	// List retrieves downtimes with optional filtering and pagination
	List(ctx context.Context, opts *ListOptions) (*ListResponse[Downtime], error)

	// ListActive retrieves currently active downtimes
	ListActive(ctx context.Context, opts *ListOptions) (*ListResponse[Downtime], error)

	// ListScheduled retrieves scheduled (future) downtimes
	ListScheduled(ctx context.Context, opts *ListOptions) (*ListResponse[Downtime], error)

	// ListForHost retrieves downtimes for a specific host
	ListForHost(ctx context.Context, hostID int64, opts *ListOptions) (*ListResponse[Downtime], error)

	// ListForService retrieves downtimes for a specific service
	ListForService(ctx context.Context, hostID, serviceID int64, opts *ListOptions) (*ListResponse[Downtime], error)

	// ListByTimeRange retrieves downtimes within a specific time range
	ListByTimeRange(ctx context.Context, startTime, endTime time.Time, opts *ListOptions) (*ListResponse[Downtime], error)

	// Update updates an existing downtime
	Update(ctx context.Context, id int64, downtime *Downtime) error

	// Cancel cancels an active or scheduled downtime
	Cancel(ctx context.Context, id int64) error

	// CancelByComment cancels downtimes matching a comment pattern
	CancelByComment(ctx context.Context, commentPattern string) ([]int64, error)

	// Delete deletes a downtime by ID
	Delete(ctx context.Context, id int64) error

	// BulkCreate creates multiple downtimes in a single operation
	BulkCreate(ctx context.Context, downtimes []Downtime) ([]Downtime, error)

	// BulkCancel cancels multiple downtimes by their IDs
	BulkCancel(ctx context.Context, ids []int64) error

	// BulkDelete deletes multiple downtimes by their IDs
	BulkDelete(ctx context.Context, ids []int64) error

	// GetStatistics retrieves downtime statistics for a time period
	GetStatistics(ctx context.Context, startTime, endTime time.Time) (*DowntimeStatistics, error)

	// GetUpcomingExpiration retrieves downtimes expiring within a specified duration
	GetUpcomingExpiration(ctx context.Context, duration time.Duration) (*ListResponse[Downtime], error)

	// ExtendDowntime extends the end time of an existing downtime
	ExtendDowntime(ctx context.Context, id int64, newEndTime time.Time) error

	// GetRecurringDowntimes retrieves recurring downtime templates
	GetRecurringDowntimes(ctx context.Context, opts *ListOptions) (*ListResponse[RecurringDowntime], error)

	// CreateRecurringDowntime creates a new recurring downtime template
	CreateRecurringDowntime(ctx context.Context, template *RecurringDowntime) (*RecurringDowntime, error)

	// UpdateRecurringDowntime updates a recurring downtime template
	UpdateRecurringDowntime(ctx context.Context, id int64, template *RecurringDowntime) error

	// DeleteRecurringDowntime deletes a recurring downtime template
	DeleteRecurringDowntime(ctx context.Context, id int64) error

	// GenerateRecurringDowntimes generates downtime instances from a template for a date range
	GenerateRecurringDowntimes(ctx context.Context, templateID int64, startDate, endDate time.Time) ([]Downtime, error)

	// GetDowntimeHistory retrieves historical downtime information
	GetDowntimeHistory(ctx context.Context, resourceType string, resourceID int64, opts *ListOptions) (*ListResponse[DowntimeHistoryEntry], error)

	// ExportDowntimes exports downtime configuration
	Export(ctx context.Context, ids []int64) ([]byte, error)

	// ImportDowntimes imports downtime configuration
	Import(ctx context.Context, data []byte, overwrite bool) (*ImportResult, error)

	// Validate validates a downtime configuration
	Validate(ctx context.Context, downtime *Downtime) (*ValidationResult, error)

	// GetConflictingDowntimes retrieves downtimes that overlap with a proposed downtime
	GetConflictingDowntimes(ctx context.Context, resourceType string, resourceID int64, startTime, endTime time.Time) ([]Downtime, error)

	// GetDowntimeImpact analyzes the impact of a proposed downtime
	GetDowntimeImpact(ctx context.Context, request *DowntimeImpactRequest) (*DowntimeImpact, error)

	// NotifyDowntimeStart sends notifications when a downtime starts
	NotifyDowntimeStart(ctx context.Context, id int64) error

	// NotifyDowntimeEnd sends notifications when a downtime ends
	NotifyDowntimeEnd(ctx context.Context, id int64) error
}

// Supporting types for Downtime interface operations

// DowntimeRequest represents a request to create a downtime (reused from common.go)
type DowntimeRequestExtended struct {
	// Comment is a required description for the downtime
	Comment string `json:"comment" validate:"required,min=1"`

	// StartTime is the scheduled start time for the downtime
	StartTime time.Time `json:"start_time" validate:"required"`

	// EndTime is the scheduled end time for the downtime
	EndTime time.Time `json:"end_time" validate:"required,gtfield=StartTime"`

	// Duration specifies the length for flexible downtimes (in seconds)
	Duration *int64 `json:"duration,omitempty" validate:"omitempty,gte=0"`

	// Fixed indicates whether this is a fixed (true) or flexible (false) downtime
	Fixed *bool `json:"fixed,omitempty"`

	// WithServices indicates whether to include services when scheduling host downtime
	WithServices *bool `json:"with_services,omitempty"`

	// Author is the username creating the downtime
	Author string `json:"author,omitempty"`

	// NotifyContacts indicates whether to notify contacts about the downtime
	NotifyContacts *bool `json:"notify_contacts,omitempty"`
}

// RecurringDowntime represents a template for creating recurring downtimes
type RecurringDowntime struct {
	// ID is the unique identifier for the recurring downtime template
	ID *int64 `json:"id,omitempty"`

	// Name is a descriptive name for the recurring downtime
	Name string `json:"name" validate:"required,min=1,max=255"`

	// Description provides additional details about the recurring downtime
	Description string `json:"description,omitempty"`

	// ResourceType specifies the type of resource ("host" or "service")
	ResourceType string `json:"resource_type" validate:"required,resource_type"`

	// ResourceID is the ID of the resource for the recurring downtime
	ResourceID *int64 `json:"resource_id" validate:"required,gt=0"`

	// ParentResourceID is the host ID for service recurring downtimes
	ParentResourceID *int64 `json:"parent_resource_id,omitempty"`

	// Comment template for generated downtimes
	Comment string `json:"comment" validate:"required,min=1"`

	// Duration of each downtime instance (in seconds)
	Duration int64 `json:"duration" validate:"required,gte=0"`

	// Fixed indicates whether generated downtimes are fixed or flexible
	Fixed *bool `json:"fixed,omitempty"`

	// WithServices indicates whether to include services for host downtimes
	WithServices *bool `json:"with_services,omitempty"`

	// RecurrencePattern defines when downtimes should be created
	RecurrencePattern *RecurrencePattern `json:"recurrence_pattern" validate:"required"`

	// IsActivated controls whether this template is active
	IsActivated *bool `json:"is_activated,omitempty"`

	// Author of the recurring downtime template
	Author string `json:"author,omitempty"`

	// CreatedAt is when the template was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// UpdatedAt is when the template was last modified
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// RecurrencePattern defines how recurring downtimes are scheduled
type RecurrencePattern struct {
	// Type of recurrence: "daily", "weekly", "monthly", "yearly"
	Type string `json:"type" validate:"required,oneof=daily weekly monthly yearly"`

	// Interval specifies every nth occurrence (e.g., every 2 weeks)
	Interval int `json:"interval" validate:"gte=1"`

	// StartTime is the time of day for the downtime (HH:MM format)
	StartTime string `json:"start_time" validate:"required,datetime=15:04"`

	// DaysOfWeek specifies which days for weekly recurrence (0=Sunday, 6=Saturday)
	DaysOfWeek []int `json:"days_of_week,omitempty" validate:"omitempty,dive,gte=0,lte=6"`

	// DayOfMonth specifies the day for monthly recurrence (1-31)
	DayOfMonth *int `json:"day_of_month,omitempty" validate:"omitempty,gte=1,lte=31"`

	// WeekOfMonth specifies the week for monthly recurrence (1-5, -1 for last)
	WeekOfMonth *int `json:"week_of_month,omitempty" validate:"omitempty,oneof=1 2 3 4 5 -1"`

	// MonthsOfYear specifies which months for yearly recurrence (1-12)
	MonthsOfYear []int `json:"months_of_year,omitempty" validate:"omitempty,dive,gte=1,lte=12"`

	// EndDate is the last date for recurring downtimes (optional)
	EndDate *time.Time `json:"end_date,omitempty"`

	// MaxOccurrences limits the number of generated downtimes (optional)
	MaxOccurrences *int `json:"max_occurrences,omitempty" validate:"omitempty,gte=1"`
}

// DowntimeStatistics provides statistical information about downtimes
type DowntimeStatistics struct {
	// Period covered by these statistics
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`

	// Total counts
	TotalDowntimes     int `json:"total_downtimes"`
	ActiveDowntimes    int `json:"active_downtimes"`
	ScheduledDowntimes int `json:"scheduled_downtimes"`
	CompletedDowntimes int `json:"completed_downtimes"`
	CancelledDowntimes int `json:"cancelled_downtimes"`

	// Breakdown by resource type
	HostDowntimes    int `json:"host_downtimes"`
	ServiceDowntimes int `json:"service_downtimes"`

	// Breakdown by downtime type
	FixedDowntimes    int `json:"fixed_downtimes"`
	FlexibleDowntimes int `json:"flexible_downtimes"`

	// Time-based statistics
	TotalDowntimeHours      float64        `json:"total_downtime_hours"`
	AverageDowntimeDuration float64        `json:"average_downtime_duration"`
	LongestDowntime         *time.Duration `json:"longest_downtime,omitempty"`
	ShortestDowntime        *time.Duration `json:"shortest_downtime,omitempty"`

	// Top contributors
	TopAuthors   []AuthorStatistic   `json:"top_authors"`
	TopResources []ResourceStatistic `json:"top_resources"`
	TopReasons   []ReasonStatistic   `json:"top_reasons"`
}

// DowntimeHistoryEntry represents a historical downtime record
type DowntimeHistoryEntry struct {
	// Downtime information
	Downtime *Downtime `json:"downtime"`

	// Impact metrics
	NotificationsSuppressed int `json:"notifications_suppressed"`
	AlertsSuppressed        int `json:"alerts_suppressed"`

	// Effectiveness metrics
	ActualDuration     *time.Duration `json:"actual_duration,omitempty"`
	PlannedDuration    time.Duration  `json:"planned_duration"`
	EffectivenessScore float64        `json:"effectiveness_score"` // 0-100
}

// DowntimeImpactRequest represents a request to analyze downtime impact
type DowntimeImpactRequest struct {
	ResourceType     string    `json:"resource_type" validate:"required,resource_type"`
	ResourceID       int64     `json:"resource_id" validate:"required,gt=0"`
	ParentResourceID *int64    `json:"parent_resource_id,omitempty"`
	StartTime        time.Time `json:"start_time" validate:"required"`
	EndTime          time.Time `json:"end_time" validate:"required,gtfield=StartTime"`
	WithServices     *bool     `json:"with_services,omitempty"`
}

// DowntimeImpact represents the analyzed impact of a downtime
type DowntimeImpact struct {
	// Resources affected
	AffectedHosts    []string `json:"affected_hosts"`
	AffectedServices []string `json:"affected_services"`
	TotalResources   int      `json:"total_resources"`

	// Notification impact
	EstimatedNotificationsSuppressed int      `json:"estimated_notifications_suppressed"`
	ContactsAffected                 []string `json:"contacts_affected"`
	ContactGroupsAffected            []string `json:"contact_groups_affected"`

	// Dependencies
	DependentResources []string `json:"dependent_resources"`
	ParentResources    []string `json:"parent_resources"`

	// Conflicting downtimes
	ConflictingDowntimes []Downtime `json:"conflicting_downtimes"`

	// Business impact
	BusinessCriticality string     `json:"business_criticality,omitempty"`
	SLAImpact           *SLAImpact `json:"sla_impact,omitempty"`
}

// Supporting statistical types
type AuthorStatistic struct {
	Author     string  `json:"author"`
	Count      int     `json:"count"`
	TotalHours float64 `json:"total_hours"`
}

type ResourceStatistic struct {
	ResourceType string  `json:"resource_type"`
	ResourceID   int64   `json:"resource_id"`
	ResourceName string  `json:"resource_name"`
	Count        int     `json:"count"`
	TotalHours   float64 `json:"total_hours"`
}

type ReasonStatistic struct {
	Reason     string  `json:"reason"`
	Count      int     `json:"count"`
	TotalHours float64 `json:"total_hours"`
}

// SLAImpact represents the impact on Service Level Agreements
type SLAImpact struct {
	SLAName              string  `json:"sla_name"`
	CurrentAvailability  float64 `json:"current_availability"`
	ImpactedAvailability float64 `json:"impacted_availability"`
	AvailabilityDelta    float64 `json:"availability_delta"`
	IsCritical           bool    `json:"is_critical"`
}
