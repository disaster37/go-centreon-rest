package api

import (
	"context"
	"time"
)

// Service represents a monitored service in Centreon infrastructure monitoring.
// Services are specific checks that run on hosts to monitor particular aspects like CPU usage,
// disk space, network connectivity, or application-specific metrics. Each service belongs to a host.
type Service struct {
	// ID is the unique identifier for the service in the Centreon database
	ID *int64 `json:"id,omitempty"`

	// Description is the unique service name/description used for identification
	// This describes what the service is monitoring (e.g., "CPU Usage", "Disk /var", "HTTP Response")
	Description string `json:"description" validate:"required,min=1,max=200"`

	// DisplayName is the name shown in the Centreon web interface
	// If not set, the Description field is used for display purposes
	DisplayName string `json:"display_name,omitempty" validate:"omitempty,max=255"`

	// State represents the current monitoring state of the service
	// 0 = OK, 1 = WARNING, 2 = CRITICAL, 3 = UNKNOWN
	State *int `json:"state,omitempty" validate:"omitempty,oneof=0 1 2 3"`

	// HostID is the unique identifier of the host this service belongs to
	// Links the service to its parent host for hierarchical organization
	HostID *int64 `json:"host_id,omitempty" validate:"omitempty,gt=0"`

	// HostName is the name of the host this service belongs to
	// Provides context when viewing service information independently
	HostName string `json:"host_name,omitempty" validate:"omitempty,min=1,max=200"`

	// CheckAttempt is the current attempt number for the ongoing check
	// Increments with each failed check until MaxCheckAttempts is reached
	CheckAttempt *string `json:"check_attempt,omitempty" validate:"omitempty,max=10"`

	// IconImage is the URL or path to an icon representing this service
	// Used in the web interface to provide visual identification
	IconImage string `json:"icon_image,omitempty" validate:"omitempty,max=500"`

	// IconImageAlt is the alternative text for the service icon
	// Used for accessibility and when the icon cannot be displayed
	IconImageAlt string `json:"icon_image_alt,omitempty" validate:"omitempty,max=200"`

	// LastCheck is the timestamp of the most recent monitoring check
	LastCheck *time.Time `json:"last_check,omitempty"`

	// LastStateChange is when the service state last changed (soft or hard)
	// This includes both temporary state changes and confirmed state changes
	LastStateChange *time.Time `json:"last_state_change,omitempty"`

	// MaxCheckAttempts is the maximum number of check attempts before a hard state
	// After this many consecutive failures, the state becomes "hard" and notifications are sent
	MaxCheckAttempts *int `json:"max_check_attempts,omitempty" validate:"omitempty,gte=1,lte=10"`

	// Output contains the text output from the last monitoring check
	// This typically includes status information, metrics, and diagnostic details
	Output string `json:"output,omitempty" validate:"omitempty,max=8192"`

	// StateType indicates whether the current state is soft (0) or hard (1)
	// Soft states are temporary and may recover; hard states trigger notifications
	StateType *int `json:"state_type,omitempty" validate:"omitempty,oneof=0 1"`

	// Criticality represents the business importance of this service (0-5 scale)
	// Higher values indicate more critical service components
	Criticality *int `json:"criticality,omitempty" validate:"omitempty,gte=0,lte=5"`

	// Status contains detailed status information including severity codes
	// Provides structured status data beyond the basic State field
	Status *ResourceStatus `json:"status,omitempty"`

	// Duration represents how long the service has been in its current state
	// Useful for understanding the persistence of problems or stability
	Duration *string `json:"duration,omitempty" validate:"omitempty,max=100"`

	// CheckCommand is the command used to monitor this service
	// Defines what check plugin and parameters are used for monitoring
	CheckCommand string `json:"check_command,omitempty" validate:"omitempty,max=255"`

	// CheckInterval is the time (in seconds) between regular monitoring checks
	// Shorter intervals provide faster problem detection but increase monitoring load
	CheckInterval *float64 `json:"check_interval,omitempty" validate:"omitempty,gte=1,lte=86400"`

	// CheckPeriod defines when this service should be monitored
	// References a time period configuration (e.g., "24x7", "workhours")
	CheckPeriod string `json:"check_period,omitempty" validate:"omitempty,max=200"`

	// CheckType indicates whether checks are active (0) or passive (1)
	// Active checks are initiated by Centreon; passive checks are submitted externally
	CheckType *int `json:"check_type,omitempty" validate:"omitempty,oneof=0 1"`

	// CommandLine is the full command line executed for this service check
	// Shows the complete command with all parameters and macros resolved
	CommandLine *string `json:"command_line,omitempty" validate:"omitempty,max=8192"`

	// ExecutionTime is the duration (in seconds) of the last check execution
	// Useful for performance monitoring and identifying slow checks
	ExecutionTime *float64 `json:"execution_time,omitempty" validate:"omitempty,gte=0"`

	// IsAcknowledged indicates whether the current service problem has been acknowledged
	// When true, notifications are typically suppressed until the acknowledgement is cleared
	IsAcknowledged *bool `json:"is_acknowledged,omitempty"`

	// IsActiveCheck indicates whether this service uses active monitoring
	// Active checks are initiated by Centreon; false indicates passive-only monitoring
	IsActiveCheck *bool `json:"is_active_check,omitempty"`

	// IsChecked indicates whether the service has been checked at least once
	// False for newly configured services that haven't had their first check yet
	IsChecked *bool `json:"is_checked,omitempty"`

	// LastHardStateChange is when the service last changed to a hard state
	// Hard state changes occur after MaxCheckAttempts consecutive failures
	LastHardStateChange *time.Time `json:"last_hard_state_change,omitempty"`

	// LastNotification is when the last notification was sent for this service
	// Used to implement notification throttling and escalation policies
	LastNotification *time.Time `json:"last_notification,omitempty"`

	// LastTimeCritical is the most recent timestamp when the service was CRITICAL
	LastTimeCritical *time.Time `json:"last_time_critical,omitempty"`

	// LastTimeOK is the most recent timestamp when the service was OK
	LastTimeOK *time.Time `json:"last_time_ok,omitempty"`

	// LastTimeUnknown is the most recent timestamp when the service was UNKNOWN
	LastTimeUnknown *time.Time `json:"last_time_unknown,omitempty"`

	// LastTimeWarning is the most recent timestamp when the service was WARNING
	LastTimeWarning *time.Time `json:"last_time_warning,omitempty"`

	// LastUpdate is when the service information was last updated in the system
	LastUpdate *time.Time `json:"last_update,omitempty"`

	// Latency is the delay (in seconds) between scheduled and actual check execution
	// High latency may indicate monitoring system performance issues
	Latency *float64 `json:"latency,omitempty" validate:"omitempty,gte=0"`

	// NextCheck is the scheduled timestamp for the next monitoring check
	NextCheck *time.Time `json:"next_check,omitempty"`

	// PerformanceData contains metrics data from the monitoring check
	// Used for graphing trends, capacity planning, and performance analysis
	PerformanceData string `json:"performance_data,omitempty" validate:"omitempty,max=8192"`

	// ScheduledDowntimeDepth indicates the number of active scheduled downtimes
	// When greater than 0, the service is in a scheduled maintenance window
	ScheduledDowntimeDepth *int `json:"scheduled_downtime_depth,omitempty" validate:"omitempty,gte=0"`

	// Flapping indicates whether the service is currently in a flapping state
	// Flapping detection helps reduce notification spam from unstable services
	Flapping *bool `json:"flapping,omitempty"`

	// Notify controls whether notifications are enabled for this service
	// When false, no notifications will be sent regardless of service state
	Notify *bool `json:"notify,omitempty"`

	// ServiceTemplate is the template this service inherits configuration from
	// Templates provide default values for monitoring parameters and reduce configuration effort
	ServiceTemplate string `json:"service_template,omitempty" validate:"omitempty,max=200"`

	// IsActivated controls whether this service configuration is active
	// Inactive services are not monitored and do not appear in monitoring views
	IsActivated *bool `json:"is_activated,omitempty"`

	// Parameters contains custom macro values and extended configuration
	// Used for service-specific variables that can be referenced in commands and templates
	Parameters map[string]any `json:"parameters,omitempty"`

	// Relationships contain related monitoring objects

	// Host is the parent host that this service belongs to
	Host *Host `json:"host,omitempty"`

	// Downtimes are scheduled maintenance windows for this service
	Downtimes []Downtime `json:"downtimes,omitempty"`

	// Acknowledgement represents any active problem acknowledgement for this service
	Acknowledgement *Acknowledgement `json:"acknowledgement,omitempty"`
}

// ServiceCategory represents a logical grouping mechanism for services in Centreon.
// Categories help organize services by function, technology, or business purpose for easier
// management, reporting, and access control. Services can belong to multiple categories.
type ServiceCategory struct {
	// ID is the unique identifier for the service category
	ID *int64 `json:"id,omitempty"`

	// Name is the unique name identifier for the category
	// Used for programmatic references and API operations
	Name string `json:"name" validate:"required,min=1,max=200"`

	// Description provides detailed information about the category's purpose
	// Used to explain what types of services should belong to this category
	Description string `json:"description,omitempty" validate:"omitempty,max=65535"`

	// Level represents the hierarchical level of this category
	// Used for creating category hierarchies and nested organizational structures (0-10)
	Level *int `json:"level,omitempty" validate:"omitempty,gte=0,lte=10"`

	// IconID references an icon to visually represent this category
	// Provides visual identification in the web interface and dashboards
	IconID *int `json:"icon_id,omitempty" validate:"omitempty,gt=0"`

	// IsActivated controls whether this category is available for use
	// Inactive categories are not displayed in selection lists
	IsActivated *bool `json:"is_activated,omitempty"`
}

// Macro represents a configuration variable that can be used in monitoring commands.
// Macros provide a way to parameterize commands and templates, making configurations
// more flexible and reusable across different hosts and services.
type Macro struct {
	// Name is the macro variable name (e.g., "$_HOSTWARNING$", "$_SERVICECRITICAL$")
	// Must follow Centreon macro naming conventions with proper prefixes
	Name string `json:"name" validate:"required,min=1,max=255"`

	// Value is the actual value that will replace the macro in commands
	// Can contain static values, other macros, or dynamic expressions
	Value string `json:"value" validate:"required"`

	// IsPassword indicates whether this macro contains sensitive information
	// When set to "1", the value is masked in the web interface for security
	IsPassword *string `json:"is_password,omitempty" validate:"omitempty,oneof=0 1"`

	// Source indicates where this macro is defined (host, service, template, etc.)
	// Helps track macro inheritance and override behavior in complex configurations
	Source string `json:"source,omitempty" validate:"omitempty,max=200"`
}

// ServiceInterface defines CRUD operations for Service entities
type ServiceInterface interface {
	// Create creates a new service
	Create(ctx context.Context, service *Service) (*Service, error)

	// GetByID retrieves a service by host ID and service ID
	GetByID(ctx context.Context, hostID, serviceID int64) (*Service, error)

	// List retrieves services with optional filtering and pagination
	List(ctx context.Context, opts *ListOptions) (*ListResponse[Service], error)

	// Update updates an existing service (partial update)
	Update(ctx context.Context, hostID, serviceID int64, service *Service) error

	// Delete deletes a service
	Delete(ctx context.Context, hostID, serviceID int64) error

	// BulkDelete deletes multiple services
	BulkDelete(ctx context.Context, services []ServiceIdentifier) error
}

// ServiceCategoryInterface defines CRUD operations for Service Category entities
type ServiceCategoryInterface interface {
	// Create creates a new service category
	Create(ctx context.Context, category *ServiceCategory) (*ServiceCategory, error)

	// GetByID retrieves a service category by its ID
	GetByID(ctx context.Context, id int64) (*ServiceCategory, error)

	// List retrieves service categories with optional filtering and pagination
	List(ctx context.Context, opts *ListOptions) (*ListResponse[ServiceCategory], error)

	// Update updates an existing service category
	Update(ctx context.Context, id int64, category *ServiceCategory) error

	// Delete deletes a service category by ID
	Delete(ctx context.Context, id int64) error
}
