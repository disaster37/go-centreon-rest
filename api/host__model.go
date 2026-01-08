package api

import (
	"encoding/json"
	"fmt"
	"time"
)

// HostCreateRequest represents the payload to create a host in Centreon.
// https://docs-api.centreon.com/api/centreon-web/25.10/#tag/Host/paths/~1configuration~1hosts/post
type HostCreateRequest struct {
	Name                      string      `json:"name" validate:"required,max=200"`
	Alias                     string      `json:"alias,omitempty" validate:"omitempty,max=200"`
	Address                   string      `json:"address" validate:"required,ip|fqdn"`
	MonitoringServerId        int64       `json:"monitoring_server_id" validate:"required"`
	Templates                 []int64     `json:"templates,omitempty"`
	SnmpCommunity             *string     `json:"snmp_community,omitempty"`
	SnmpVersion               *string     `json:"snmp_version,omitempty" validate:"omitempty,oneof=1 2c 3"`
	GeoCoords                 *string     `json:"geo_coords,omitempty" validate:"omitempty,max=32"`
	TimezoneId                *int64      `json:"timezone_id,omitempty"`
	SeverityId                *int64      `json:"severity_id,omitempty"`
	CheckCommandId            *int64      `json:"check_command_id,omitempty"`
	CheckCommandArgs          []string    `json:"check_command_args,omitempty"`
	CheckTimeperiodId         *int64      `json:"check_timeperiod_id,omitempty"`
	MaxCheckAttempts          *int        `json:"max_check_attempts,omitempty"`
	CheckInterval             *int        `json:"check_interval,omitempty"`
	RetryCheckInterval        *int        `json:"retry_check_interval,omitempty"`
	ActiveCheckEnabled        *int        `json:"active_check_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	PassiveCheckEnabled       *int        `json:"passive_check_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	NotificationEnabled       *int        `json:"notification_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	NotificationOptions       *int        `json:"notification_options,omitempty"`
	NotificationInterval      *int        `json:"notification_interval,omitempty"`
	NotificationTimeperiodId  *int64      `json:"notification_timeperiod_id,omitempty"`
	AddInheritedContactGroup  *bool       `json:"add_inherited_contact_group,omitempty"`
	AddInheritedContact       *bool       `json:"add_inherited_contact,omitempty"`
	FirstNotificationDelay    *int        `json:"first_notification_delay,omitempty"`
	RecoveryNotificationDelay *int        `json:"recovery_notification_delay,omitempty"`
	AcknowledgementTimeout    *int        `json:"acknowledgement_timeout,omitempty"`
	FreshnessChecked          *int        `json:"freshness_checked,omitempty" validate:"omitempty,oneof=0 1 2"`
	FreshnessThreshold        *int        `json:"freshness_threshold,omitempty"`
	FlapDetectionEnabled      *int        `json:"flap_detection_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	LowFlapThreshold          *int        `json:"low_flap_threshold,omitempty"`
	HighFlapThreshold         *int        `json:"high_flap_threshold,omitempty"`
	EventHandlerEnabled       *int        `json:"event_handler_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	EventHandlerCommandId     *int64      `json:"event_handler_command_id,omitempty"`
	EventHandlerCommandArgs   []string    `json:"event_handler_command_args,omitempty"`
	NoteUrl                   *string     `json:"note_url,omitempty" validate:"omitempty,max=65535"`
	Note                      *string     `json:"note,omitempty" validate:"omitempty,max=65535"`
	ActionUrl                 *string     `json:"action_url,omitempty" validate:"omitempty,max=65535"`
	IconId                    *int64      `json:"icon_id,omitempty"`
	IconAlternative           *string     `json:"icon_alternative,omitempty" validate:"omitempty,max=200"`
	Comment                   *string     `json:"comment,omitempty"`
	IsActivated               *bool       `json:"is_activated,omitempty"`
	Categories                []int64     `json:"categories,omitempty"`
	Groups                    []int64     `json:"groups,omitempty"`
	Macros                    []HostMacro `json:"macros,omitempty"`
}

// String returns the string representation of the HostCreateRequest.
func (h HostCreateRequest) String() string {
	return h.Name
}

// HostUpdateRequest represents the payload to updatecreate a host in Centreon.
// https://docs-api.centreon.com/api/centreon-web/25.10/#tag/Host/paths/~1configuration~1hosts~1%7Bhost_id%7D/patch
type HostUpdateRequest struct {
	Name                      *string     `json:"name,omitempty" validate:"omitempty,required,max=200"`
	Alias                     *string     `json:"alias,omitempty" validate:"omitempty,max=200"`
	Address                   *string     `json:"address,omitempty" validate:"omitempty,required,ip|fqdn"`
	MonitoringServerId        *int64      `json:"monitoring_server_id,omitempty" validate:"omitempty,required"`
	Templates                 []int64     `json:"templates,omitempty"`
	SnmpCommunity             *string     `json:"snmp_community,omitempty"`
	SnmpVersion               *string     `json:"snmp_version,omitempty" validate:"omitempty,oneof=1 2c 3"`
	GeoCoords                 *string     `json:"geo_coords,omitempty" validate:"omitempty,max=32"`
	TimezoneId                *int64      `json:"timezone_id,omitempty"`
	SeverityId                *int64      `json:"severity_id,omitempty"`
	CheckCommandId            *int64      `json:"check_command_id,omitempty"`
	CheckCommandArgs          []string    `json:"check_command_args,omitempty"`
	CheckTimeperiodId         *int64      `json:"check_timeperiod_id,omitempty"`
	MaxCheckAttempts          *int        `json:"max_check_attempts,omitempty"`
	CheckInterval             *int        `json:"check_interval,omitempty"`
	RetryCheckInterval        *int        `json:"retry_check_interval,omitempty"`
	ActiveCheckEnabled        *int        `json:"active_check_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	PassiveCheckEnabled       *int        `json:"passive_check_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	NotificationEnabled       *int        `json:"notification_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	NotificationOptions       *int        `json:"notification_options,omitempty"`
	NotificationInterval      *int        `json:"notification_interval,omitempty"`
	NotificationTimeperiodId  *int64      `json:"notification_timeperiod_id,omitempty"`
	AddInheritedContactGroup  *bool       `json:"add_inherited_contact_group,omitempty"`
	AddInheritedContact       *bool       `json:"add_inherited_contact,omitempty"`
	FirstNotificationDelay    *int        `json:"first_notification_delay,omitempty"`
	RecoveryNotificationDelay *int        `json:"recovery_notification_delay,omitempty"`
	AcknowledgementTimeout    *int        `json:"acknowledgement_timeout,omitempty"`
	FreshnessChecked          *int        `json:"freshness_checked,omitempty" validate:"omitempty,oneof=0 1 2"`
	FreshnessThreshold        *int        `json:"freshness_threshold,omitempty"`
	FlapDetectionEnabled      *int        `json:"flap_detection_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	LowFlapThreshold          *int        `json:"low_flap_threshold,omitempty"`
	HighFlapThreshold         *int        `json:"high_flap_threshold,omitempty"`
	EventHandlerEnabled       *int        `json:"event_handler_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	EventHandlerCommandId     *int64      `json:"event_handler_command_id,omitempty"`
	EventHandlerCommandArgs   []string    `json:"event_handler_command_args,omitempty"`
	NoteUrl                   *string     `json:"note_url,omitempty" validate:"omitempty,max=65535"`
	Note                      *string     `json:"note,omitempty" validate:"omitempty,max=65535"`
	ActionUrl                 *string     `json:"action_url,omitempty" validate:"omitempty,max=65535"`
	IconId                    *int64      `json:"icon_id,omitempty"`
	IconAlternative           *string     `json:"icon_alternative,omitempty" validate:"omitempty,max=200"`
	Comment                   *string     `json:"comment,omitempty"`
	IsActivated               *bool       `json:"is_activated,omitempty"`
	Categories                []int64     `json:"categories,omitempty"`
	Groups                    []int64     `json:"groups,omitempty"`
	Macros                    []HostMacro `json:"macros,omitempty"`
}

// HostMacros represents a macro definition for a host in Centreon.
type HostMacro struct {
	Name        string  `json:"name" validate:"required"`
	Value       *string `json:"value,omitempty"`
	IsPassword  *bool   `json:"is_password,omitempty"`
	Description *string `json:"description,omitempty"`
}

// HostUpdateResponse represents the response after creating or updating a host in Centreon.
type HostCreateResponse struct {
	Id                        int64       `json:"id"`
	Name                      string      `json:"name"`
	Alias                     string      `json:"alias,omitempty"`
	Address                   string      `json:"address"`
	MonitoringServerId        int64       `json:"monitoring_server_id"`
	SnmpCommunity             *string     `json:"snmp_community,omitempty"`
	SnmpVersion               *string     `json:"snmp_version,omitempty"`
	GeoCoords                 *string     `json:"geo_coords,omitempty"`
	TimezoneId                *int64      `json:"timezone_id,omitempty"`
	SeverityId                *int64      `json:"severity_id,omitempty"`
	CheckCommandId            *int64      `json:"check_command_id,omitempty"`
	CheckCommandArgs          []string    `json:"check_command_args,omitempty"`
	CheckTimeperiodId         *int64      `json:"check_timeperiod_id,omitempty"`
	MaxCheckAttempts          *int        `json:"max_check_attempts,omitempty"`
	NormalCheckInterval       *int        `json:"normal_check_interval,omitempty"`
	RetryCheckInterval        *int        `json:"retry_check_interval,omitempty"`
	ActiveCheckEnabled        *int        `json:"active_check_enabled,omitempty"`
	PassiveCheckEnabled       *int        `json:"passive_check_enabled,omitempty"`
	NotificationEnabled       *int        `json:"notification_enabled,omitempty"`
	NotificationOptions       *int        `json:"notification_options,omitempty"`
	NotificationInterval      *int        `json:"notification_interval,omitempty"`
	NotificationTimeperiodId  *int64      `json:"notification_timeperiod_id,omitempty"`
	AddInheritedContactGroup  *bool       `json:"add_inherited_contact_group,omitempty"`
	AddInheritedContact       *bool       `json:"add_inherited_contact,omitempty"`
	FirstNotificationDelay    *int        `json:"first_notification_delay,omitempty"`
	RecoveryNotificationDelay *int        `json:"recovery_notification_delay,omitempty"`
	AcknowledgementTimeout    *int        `json:"acknowledgement_timeout,omitempty"`
	FreshnessChecked          *int        `json:"freshness_checked,omitempty"`
	FreshnessThreshold        *int        `json:"freshness_threshold,omitempty"`
	FlapDetectionEnabled      *int        `json:"flap_detection_enabled,omitempty"`
	LowFlapThreshold          *int        `json:"low_flap_threshold,omitempty"`
	HighFlapThreshold         *int        `json:"high_flap_threshold,omitempty"`
	EventHandlerEnabled       *int        `json:"event_handler_enabled,omitempty"`
	EventHandlerCommandId     *int64      `json:"event_handler_command_id,omitempty"`
	EventHandlerCommandArgs   []string    `json:"event_handler_command_args,omitempty"`
	NoteUrl                   *string     `json:"note_url,omitempty"`
	Note                      *string     `json:"note,omitempty"`
	ActionUrl                 *string     `json:"action_url,omitempty"`
	IconId                    *int64      `json:"icon_id,omitempty"`
	IconAlternative           *string     `json:"icon_alternative,omitempty"`
	Comment                   *string     `json:"comment,omitempty"`
	IsActivated               *bool       `json:"is_activated,omitempty"`
	Categories                []IdName    `json:"categories,omitempty"`
	Groups                    []IdName    `json:"groups,omitempty"`
	Templates                 []IdName    `json:"templates,omitempty"`
	Macros                    []HostMacro `json:"macros,omitempty"`
}

// Host represents a monitored host in Centreon infrastructure monitoring.
// https://docs-api.centreon.com/api/centreon-web/25.10/#tag/Host/paths/~1monitoring~1hosts~1%7Bhost_id%7D/get
type HostResponse struct {
	Id                     int64                       `json:"id"`
	Alias                  string                      `json:"alias"`
	DisplayName            string                      `json:"display_name"`
	Name                   string                      `json:"name"`
	State                  int                         `json:"state"`
	Services               []HostServiceResponse       `json:"services"`
	PollerID               int64                       `json:"poller_id"`
	Acknowledged           bool                        `json:"acknowledged"`
	AddressIP              string                      `json:"address_ip"`
	CheckAttempt           int                         `json:"check_attempt"`
	Checked                bool                        `json:"checked"`
	ExecutionTime          *float64                    `json:"execution_time"`
	IconImage              string                      `json:"icon_image"`
	IconImageAlt           string                      `json:"icon_image_alt"`
	LastCheck              *time.Time                  `json:"last_check"`
	LastHardStateChange    *time.Time                  `json:"last_hard_state_change"`
	LastStateChange        *time.Time                  `json:"last_state_change"`
	LastTimeDown           *time.Time                  `json:"last_time_down"`
	LastTimeUnreachable    *time.Time                  `json:"last_time_unreachable"`
	LastTimeUp             *time.Time                  `json:"last_time_up"`
	LastUpdate             *time.Time                  `json:"last_update"`
	MaxCheckAttempts       int                         `json:"max_check_attempts"`
	Output                 string                      `json:"output"`
	PassiveChecks          bool                        `json:"passive_checks"`
	StateType              int                         `json:"state_type"`
	Timezone               string                      `json:"timezone"`
	ScheduledDowntimeDepth int                         `json:"scheduled_downtime_depth"`
	Criticality            *int                        `json:"criticality"`
	ActiveChecks           bool                        `json:"active_checks"`
	CheckCommand           string                      `json:"check_command"`
	CheckInterval          float64                     `json:"check_interval"`
	CheckPeriod            string                      `json:"check_period"`
	CheckType              int                         `json:"check_type"`
	LastHardState          *int64                      `json:"last_hard_state"`
	LastNotification       *time.Time                  `json:"last_notification"`
	Latency                float64                     `json:"latency"`
	NextCheck              *time.Time                  `json:"next_check"`
	NextHostNotification   *int64                      `json:"next_host_notification"`
	NotificationInterval   float64                     `json:"notification_interval"`
	NotificationNumber     int                         `json:"notification_number"`
	Notify                 bool                        `json:"notify"`
	NotifyOnDown           bool                        `json:"notify_on_down"`
	NotifyOnDowntime       bool                        `json:"notify_on_downtime"`
	NotifyOnFlapping       bool                        `json:"notify_on_flapping"`
	NotifyOnRecovery       bool                        `json:"notify_on_recovery"`
	NotifyOnUnreachable    bool                        `json:"notify_on_unreachable"`
	Flapping               bool                        `json:"flapping"`
	PercentStateChange     float64                     `json:"percent_state_change"`
	Downtimes              []HostDowntimeResponse      `json:"downtimes"`
	Acknowledgement        *HostAcknowledgementWrapper `json:"acknowledgement"`
}

// HostServiceResponse represents a service associated with a host in Centreon.

type HostServiceResponse struct {
	Id          int64  `json:"id"`
	Description string `json:"description"`
	DisplayName string `json:"display_name"`
	State       int    `json:"state"`
}

// HostDowntimeResponse represents a scheduled downtime for a host in Centreon.
type HostDowntimeResponse struct {
	Id              int64      `json:"id"`
	AuthorID        int64      `json:"author_id"`
	AuthorName      string     `json:"author_name"`
	HostID          int64      `json:"host_id"`
	Comment         string     `json:"comment"`
	Duration        int        `json:"duration"`
	EntryTime       time.Time  `json:"entry_time"`
	StartTime       time.Time  `json:"start_time"`
	EndTime         time.Time  `json:"end_time"`
	DeletionTime    *time.Time `json:"deletion_time"`
	ActualStartTime time.Time  `json:"actual_start_time"`
	ActualEndTime   *time.Time `json:"actual_end_time"`
	IsStarted       bool       `json:"is_started"`
	IsCancelled     bool       `json:"is_cancelled"`
	IsFixed         bool       `json:"is_fixed"`
}

// HostAcknowledgementWrapper wraps the HostAcknowledgementResponse for JSON structure.
type HostAcknowledgementWrapper struct {
	Host *HostAcknowledgementResponse `json:"host"`
}

// HostAcknowledgementResponse represents an active problem acknowledgement for a host in Centreon.
type HostAcknowledgementResponse struct {
	Id                  int64      `json:"id"`
	AuthorID            int64      `json:"author_id"`
	AuthorName          string     `json:"author_name"`
	Comment             string     `json:"comment"`
	DeletionTime        *time.Time `json:"deletion_time"`
	EntryTime           time.Time  `json:"entry_time"`
	HostID              int64      `json:"host_id"`
	PollerID            int64      `json:"poller_id"`
	IsNotifyContacts    bool       `json:"is_notify_contacts"`
	IsPersistentComment bool       `json:"is_persistent_comment"`
	IsSticky            bool       `json:"is_sticky"`
	State               int        `json:"state"`
}

// HostListOptions represents options for listing hosts with optional filtering and pagination.
type HostListOptions struct {
	ListOptions
	ShowService *bool `json:"show_service,omitempty"`
}

// GetQueryParams converts HostListOptions into a map of query parameters
func (h HostListOptions) GetQueryParams() map[string]string {
	params := make(map[string]string)

	if h.ShowService != nil {
		params["show_service"] = fmt.Sprintf("%t", *h.ShowService)
	}

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

// HostFindResult represents a host found via the HostService Find method.
type HostFindResult struct {
	Id                     int64    `json:"id"`
	Name                   string   `json:"name"`
	Alias                  string   `json:"alias"`
	Address                string   `json:"address"`
	MonitoringServer       IdName   `json:"monitoring_server"`
	Templates              []IdName `json:"templates"`
	NormalCheckInterval    *float64 `json:"normal_check_interval,omitempty"`
	RetryCheckInterval     *float64 `json:"retry_check_interval,omitempty"`
	NotificationTimeperiod *IdName  `json:"notification_timeperiod,omitempty"`
	CheckTimeperiod        *IdName  `json:"check_timeperiod,omitempty"`
	Severity               *IdName  `json:"severity,omitempty"`
	Categories             []IdName `json:"categories,omitempty"`
	Groups                 []IdName `json:"groups,omitempty"`
	IsActivated            bool     `json:"is_activated"`
}

// CountHostStatusResponse represents the count of hosts by their status from the monitoring API.
type CountHostStatusResponse struct {
	Up          CountHostStatusResponseTotal `json:"up"`
	Down        CountHostStatusResponseTotal `json:"down"`
	Unreachable CountHostStatusResponseTotal `json:"unreachable"`
	Pending     CountHostStatusResponseTotal `json:"pending"`
	Total       int64                        `json:"total"`
}

// CountHostStatusResponseTotal represents the total count for a specific host status.
type CountHostStatusResponseTotal struct {
	Total int64 `json:"total"`
}
