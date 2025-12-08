package api

import (
	"context"
	"time"
)

// Host represents a monitored host in Centreon infrastructure monitoring.
// Hosts are the fundamental entities representing servers, network devices, or any infrastructure
// components that need to be monitored. Each host can have multiple services associated with it.
type Host struct {
	// Embed HostBase for common fields shared with HostTemplate
	HostBase

	// Host-specific fields that are not shared with HostTemplate

	// DisplayName is the name shown in the Centreon web interface
	// If not set, the Name field is used for display purposes
	DisplayName string `json:"display_name,omitempty" validate:"omitempty,max=255"`

	// AddressIP is the specific IP address of the host
	// Used when the Address field contains an FQDN but you need to store the resolved IP
	AddressIP string `json:"address_ip,omitempty" validate:"omitempty,ip"`

	// State represents the current monitoring state of the host
	// 0 = UP, 1 = DOWN, 2 = UNREACHABLE
	State *int `json:"state,omitempty" validate:"omitempty,oneof=0 1 2"`

	// PollerID identifies which monitoring poller is responsible for checking this host
	// Pollers are distributed monitoring engines that perform checks on behalf of the central server
	PollerID *int64 `json:"poller_id,omitempty" validate:"omitempty,gt=0"`

	// Acknowledged indicates whether the current host problem has been acknowledged
	// When true, notifications are typically suppressed until the acknowledgement is cleared
	Acknowledged *bool `json:"acknowledged,omitempty"`

	// CheckAttempt is the current attempt number for the ongoing check
	// Increments with each failed check until MaxCheckAttempts is reached
	CheckAttempt *int `json:"check_attempt,omitempty" validate:"omitempty,gte=1"`

	// Checked indicates whether the host has been checked at least once
	// False for newly configured hosts that haven't had their first check yet
	Checked *bool `json:"checked,omitempty"`

	// ExecutionTime is the duration (in seconds) of the last check execution
	// Useful for performance monitoring and identifying slow checks
	ExecutionTime *float64 `json:"execution_time,omitempty" validate:"omitempty,gte=0"`

	// Runtime state and history fields

	// LastCheck is the timestamp of the most recent monitoring check
	LastCheck *time.Time `json:"last_check,omitempty"`

	// LastHardStateChange is when the host last changed to a hard state
	// Hard state changes occur after MaxCheckAttempts consecutive failures
	LastHardStateChange *time.Time `json:"last_hard_state_change,omitempty"`

	// LastStateChange is when the host state last changed (soft or hard)
	// This includes both temporary state changes and confirmed state changes
	LastStateChange *time.Time `json:"last_state_change,omitempty"`

	// LastTimeDown is the most recent timestamp when the host was in DOWN state
	LastTimeDown *time.Time `json:"last_time_down,omitempty"`

	// LastTimeUnreachable is the most recent timestamp when the host was UNREACHABLE
	LastTimeUnreachable *time.Time `json:"last_time_unreachable,omitempty"`

	// LastTimeUp is the most recent timestamp when the host was in UP state
	LastTimeUp *time.Time `json:"last_time_up,omitempty"`

	// LastUpdate is when the host information was last updated in the system
	LastUpdate *time.Time `json:"last_update,omitempty"`

	// Output contains the text output from the last monitoring check
	// This typically includes status information and performance metrics
	Output string `json:"output,omitempty" validate:"omitempty,max=8192"`

	// StateType indicates whether the current state is soft (0) or hard (1)
	// Soft states are temporary and may recover; hard states trigger notifications
	StateType *int `json:"state_type,omitempty" validate:"omitempty,oneof=0 1"`

	// ScheduledDowntimeDepth indicates the number of active scheduled downtimes
	// When greater than 0, the host is in a scheduled maintenance window
	ScheduledDowntimeDepth *int `json:"scheduled_downtime_depth,omitempty" validate:"omitempty,gte=0"`

	// LastHardState is the timestamp of the last hard state change
	// Used for calculating availability metrics and SLA reporting
	LastHardState *time.Time `json:"last_hard_state,omitempty"`

	// LastNotification is when the last notification was sent for this host
	// Used to implement notification throttling and escalation policies
	LastNotification *time.Time `json:"last_notification,omitempty"`

	// Latency is the delay (in seconds) between scheduled and actual check execution
	// High latency may indicate monitoring system performance issues
	Latency *float64 `json:"latency,omitempty" validate:"omitempty,gte=0"`

	// NextCheck is the scheduled timestamp for the next monitoring check
	NextCheck *time.Time `json:"next_check,omitempty"`

	// NextHostNotification is the sequence number for the next notification
	// Used by notification escalation and throttling mechanisms
	NextHostNotification *int `json:"next_host_notification,omitempty" validate:"omitempty,gte=0"`

	// NotificationNumber tracks how many notifications have been sent
	// Increments with each notification for escalation and throttling purposes
	NotificationNumber *int `json:"notification_number,omitempty" validate:"omitempty,gte=0"`

	// HostTemplate is the template this host inherits configuration from
	// Templates provide default values for monitoring parameters and reduce configuration effort
	HostTemplate string `json:"host_template,omitempty" validate:"omitempty,max=200"`

	// MonitoringServer identifies which monitoring server manages this host
	// Used in distributed monitoring environments with multiple Centreon instances
	MonitoringServer string `json:"monitoring_server,omitempty" validate:"omitempty,max=200"`

	// Relationships contain related monitoring objects

	// Services are the individual service checks running on this host
	Services []Service `json:"services,omitempty"`

	// Downtimes are scheduled maintenance windows for this host
	Downtimes []Downtime `json:"downtimes,omitempty"`

	// Acknowledgement represents any active problem acknowledgement for this host
	Acknowledgement *Acknowledgement `json:"acknowledgement,omitempty"`
}

// HostInterface defines CRUD operations for Host entities
type HostInterface interface {
	// Create creates a new host
	Create(ctx context.Context, host *Host) (*Host, error)

	// GetByID retrieves a host by its ID
	GetByID(ctx context.Context, id int64) (*Host, error)

	// GetByName retrieves a host by its name
	GetByName(ctx context.Context, name string) (*Host, error)

	// List retrieves hosts with optional filtering and pagination
	List(ctx context.Context, opts *ListOptions) (*ListResponse[Host], error)

	// Update updates an existing host
	Update(ctx context.Context, id int64, host *Host) error

	// Delete deletes a host by ID
	Delete(ctx context.Context, id int64) error

	// BulkCreate creates multiple hosts in a single operation
	BulkCreate(ctx context.Context, hosts []Host) error

	// BulkUpdate updates multiple hosts in a single operation
	BulkUpdate(ctx context.Context, hosts []Host) error

	// BulkDelete deletes multiple hosts by their IDs
	BulkDelete(ctx context.Context, ids []int64) error

	// ListServices retrieves services for a host
	ListServices(ctx context.Context, hostID int64, opts *ListOptions) (*ListResponse[Service], error)

	// AddService adds a service to a host
	AddService(ctx context.Context, hostID int64, service *Service) error

	// RemoveService removes a service from a host
	RemoveService(ctx context.Context, hostID, serviceID int64) error
}
