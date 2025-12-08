package api

import (
	"fmt"
	"time"
)

// HostBase contains the common fields shared between Host and HostTemplate structs.
// This base struct reduces code duplication and ensures consistency between
// host instances and their templates in Centreon monitoring configuration.
type HostBase struct {
	// ID is the unique identifier in the Centreon database
	ID *int64 `json:"id,omitempty"`

	// Name is the unique name used for identification in Centreon
	// This is typically the hostname or a meaningful identifier
	Name string `json:"name" validate:"required,min=1,max=200"`

	// Alias is a human-readable alternative name
	// Used for display purposes and can contain spaces and special characters
	Alias string `json:"alias,omitempty" validate:"omitempty,max=255"`

	// Address is the network address (IP or FQDN) used to reach the host/template
	// This is the primary address used by monitoring plugins to connect
	Address string `json:"address,omitempty" validate:"omitempty,ip|fqdn"`

	// CheckCommand is the command used to monitor availability
	// Typically a ping-based command to verify network reachability
	CheckCommand string `json:"check_command,omitempty" validate:"omitempty,max=255"`

	// CheckInterval is the time (in seconds/minutes) between regular monitoring checks
	// Shorter intervals provide faster problem detection but increase monitoring load
	CheckInterval *float64 `json:"check_interval,omitempty" validate:"omitempty,gte=1,lte=86400"`

	// CheckPeriod defines when monitoring should occur
	// References a time period configuration (e.g., "24x7", "workhours")
	CheckPeriod string `json:"check_period,omitempty" validate:"omitempty,max=200"`

	// MaxCheckAttempts is the maximum number of check attempts before a hard state
	// After this many consecutive failures, the state becomes "hard" and notifications are sent
	MaxCheckAttempts *int64 `json:"max_check_attempts,omitempty" validate:"omitempty,gte=1,lte=10"`

	// NotificationInterval is the time (in seconds/minutes) between repeat notifications
	// Set to 0 to disable repeat notifications for persistent problems
	NotificationInterval *float64 `json:"notification_interval,omitempty" validate:"omitempty,gte=0,lte=86400"`

	// Notification options control when notifications are sent

	// NotifyOnDown enables notifications when going to DOWN state
	NotifyOnDown *bool `json:"notify_on_down,omitempty"`

	// NotifyOnUnreachable enables notifications when becoming UNREACHABLE
	NotifyOnUnreachable *bool `json:"notify_on_unreachable,omitempty"`

	// NotifyOnRecovery enables notifications when recovering to UP state
	NotifyOnRecovery *bool `json:"notify_on_recovery,omitempty"`

	// NotifyOnFlapping enables notifications when starting/stopping flapping
	// Flapping occurs when rapidly changing between UP and DOWN states
	NotifyOnFlapping *bool `json:"notify_on_flapping,omitempty"`

	// NotifyOnDowntime enables notifications when scheduled downtime starts/ends
	NotifyOnDowntime *bool `json:"notify_on_downtime,omitempty"`

	// ActiveChecks indicates whether Centreon actively monitors this host
	// When false, relies solely on passive checks from external sources
	ActiveChecks *bool `json:"active_checks,omitempty"`

	// PassiveChecks indicates whether this host accepts passive check results
	// When true, external systems can submit check results
	PassiveChecks *bool `json:"passive_checks,omitempty"`

	// CheckType indicates whether checks are active (0) or passive (1)
	// Active checks are initiated by Centreon; passive checks are submitted externally
	CheckType *int `json:"check_type,omitempty" validate:"omitempty,oneof=0 1"`

	// Notify controls whether notifications are enabled
	// When false, no notifications will be sent regardless of other notification settings
	Notify *bool `json:"notify,omitempty"`

	// Flapping indicates whether currently in a flapping state
	// Flapping detection helps reduce notification spam from unstable hosts
	Flapping *bool `json:"flapping,omitempty"`

	// PercentStateChange represents the percentage of state change for flap detection
	// Higher values indicate more frequent state changes and potential flapping
	PercentStateChange *float64 `json:"percent_state_change,omitempty" validate:"omitempty,gte=0,lte=100"`

	// IconImage is the URL or path to an icon for visual identification
	// Used in the web interface to provide visual identification
	IconImage string `json:"icon_image,omitempty" validate:"omitempty,max=500"`

	// IconImageAlt is the alternative text for the icon
	// Used for accessibility and when the icon cannot be displayed
	IconImageAlt string `json:"icon_image_alt,omitempty" validate:"omitempty,max=200"`

	// Timezone specifies the timezone for scheduling and timestamps
	// Used for time-based operations like maintenance windows and reporting
	Timezone string `json:"timezone,omitempty" validate:"omitempty,max=100"`

	// Criticality represents the business importance (0-5 scale)
	// Higher values indicate more critical infrastructure components
	Criticality *int `json:"criticality,omitempty" validate:"omitempty,gte=0,lte=5"`

	// IsActivated controls whether this configuration is active
	// Inactive configurations are not used for monitoring
	IsActivated *bool `json:"is_activated,omitempty"`

	// IsLocked prevents modification of this configuration
	// Used to protect critical configurations from accidental changes
	IsLocked *bool `json:"is_locked,omitempty"`

	// Parameters contains custom macro values and extended configuration
	// Used for variables that can be referenced in commands and templates
	Parameters map[string]any `json:"parameters,omitempty"`

	// Audit and metadata fields

	// CreatedAt is when the configuration was initially created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// UpdatedAt is when the configuration was last modified
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// CreatedBy identifies who created the configuration
	CreatedBy string `json:"created_by,omitempty"`

	// UpdatedBy identifies who last modified the configuration
	UpdatedBy string `json:"updated_by,omitempty"`

	// Version tracks configuration changes
	// Incremented with each modification for change tracking
	Version *int64 `json:"version,omitempty"`

	// Comment contains administrative notes about recent changes
	// Used for change management and audit trails
	Comment string `json:"comment,omitempty" validate:"omitempty,max=8192"`
}

// GetEffectiveCheckInterval returns the check interval, using a default if not set
func (hb *HostBase) GetEffectiveCheckInterval() float64 {
	if hb.CheckInterval != nil {
		return *hb.CheckInterval
	}
	return 300.0 // Default 5 minutes (300 seconds)
}

// GetEffectiveMaxCheckAttempts returns the max check attempts, using a default if not set
func (hb *HostBase) GetEffectiveMaxCheckAttempts() int64 {
	if hb.MaxCheckAttempts != nil {
		return *hb.MaxCheckAttempts
	}
	return 3 // Default 3 attempts
}

// GetEffectiveNotificationInterval returns the notification interval, using a default if not set
func (hb *HostBase) GetEffectiveNotificationInterval() float64 {
	if hb.NotificationInterval != nil {
		return *hb.NotificationInterval
	}
	return 0.0 // Default: no repeat notifications
}

// IsActive returns whether the configuration is activated and can be used
func (hb *HostBase) IsActive() bool {
	return hb.IsActivated == nil || *hb.IsActivated
}

// IsReadOnly returns whether the configuration is locked and cannot be modified
func (hb *HostBase) IsReadOnly() bool {
	return hb.IsLocked != nil && *hb.IsLocked
}

// ValidateBaseFields validates common fields shared between Host and HostTemplate
func (hb *HostBase) ValidateBaseFields() error {
	// Validate notification intervals
	if hb.NotificationInterval != nil && hb.CheckInterval != nil {
		if *hb.NotificationInterval < *hb.CheckInterval {
			return fmt.Errorf("notification interval should not be less than check interval")
		}
	}

	return nil
}
