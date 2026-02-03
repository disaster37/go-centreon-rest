package api

import "time"

// ServiceCreateRequest represents the payload to create a service in Centreon.
// https://docs-api.centreon.com/api/centreon-web/25.10/#tag/Service/paths/~1configuration~1services/post
type ServiceCreateRequest struct {
	Name                              string      `json:"name" validate:"required,max=200"`
	HostId                            int64       `json:"host_id" validate:"required"`
	GeoCoords                         *string     `json:"geo_coords,omitempty" validate:"omitempty,max=32"`
	Comment                           *string     `json:"comment,omitempty"`
	ServiceTemplateId                 *int64      `json:"service_template_id,omitempty"`
	CheckCommandId                    *int64      `json:"check_command_id,omitempty"`
	CheckCommandArgs                  []string    `json:"check_command_args,omitempty"`
	CheckTimeperiodId                 *int64      `json:"check_timeperiod_id,omitempty"`
	MaxCheckAttempts                  *int        `json:"max_check_attempts,omitempty"`
	NormalCheckInterval               *int        `json:"normal_check_interval,omitempty"`
	RetryCheckInterval                *int        `json:"retry_check_interval,omitempty"`
	ActiveCheckEnabled                *CheckState `json:"active_check_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	PassiveCheckEnabled               *CheckState `json:"passive_check_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	VolatilityEnabled                 *CheckState `json:"volatility_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	NotificationEnabled               *CheckState `json:"notification_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	IsContactAdditiveInheritance      *bool       `json:"is_contact_additive_inheritance,omitempty"`
	IsContactGroupAdditiveInheritance *bool       `json:"is_contact_group_additive_inheritance,omitempty"`
	NotificationInterval              *int        `json:"notification_interval,omitempty"`
	NotificationTimeperiodId          *int64      `json:"notification_timeperiod_id,omitempty"`
	NotificationType                  *int        `json:"notification_type,omitempty"`
	FirstNotificationDelay            *int        `json:"first_notification_delay,omitempty"`
	RecoveryNotificationDelay         *int        `json:"recovery_notification_delay,omitempty"`
	AcknowledgementTimeout            *int        `json:"acknowledgement_timeout,omitempty"`
	FreshnessChecked                  *CheckState `json:"freshness_checked,omitempty" validate:"omitempty,oneof=0 1 2"`
	FreshnessThreshold                *int        `json:"freshness_threshold,omitempty"`
	FlapDetectionEnabled              *CheckState `json:"flap_detection_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	LowFlapThreshold                  *int        `json:"low_flap_threshold,omitempty"`
	HighFlapThreshold                 *int        `json:"high_flap_threshold,omitempty"`
	EventHandlerEnabled               *CheckState `json:"event_handler_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	EventHandlerCommandId             *int64      `json:"event_handler_command_id,omitempty"`
	EventHandlerCommandArgs           []string    `json:"event_handler_command_args,omitempty"`
	GraphTemplateId                   *int64      `json:"graph_template_id,omitempty"`
	Note                              *string     `json:"note,omitempty" validate:"omitempty,max=65535"`
	NoteUrl                           *string     `json:"note_url,omitempty" validate:"omitempty,max=65535"`
	ActionUrl                         *string     `json:"action_url,omitempty" validate:"omitempty,max=65535"`
	IconId                            *int64      `json:"icon_id,omitempty"`
	IconAlternative                   *string     `json:"icon_alternative,omitempty" validate:"omitempty,max=200"`
	SeverityId                        *int64      `json:"severity_id,omitempty"`
	IsActivated                       *bool       `json:"is_activated,omitempty"`
	ServiceCategories                 []int64     `json:"service_categories,omitempty"`
	ServiceGroups                     []int64     `json:"service_groups,omitempty"`
	Macros                            []Macro     `json:"macros,omitempty"`
}

// String returns the string representation of the ServiceCreateRequest.
func (s ServiceCreateRequest) String() string {
	return s.Name
}

// ServiceResponse represents the response after creating a service in Centreon or when get Service by Id.
// https://docs-api.centreon.com/api/centreon-web/25.10/#tag/Service/paths/~1configuration~1services/post
type ServiceResponse struct {
	Id                                int64      `json:"id"`
	Name                              string     `json:"name"`
	HostId                            int64      `json:"host_id"`
	GeoCoords                         *string    `json:"geo_coords"`
	Comment                           *string    `json:"comment"`
	ServiceTemplateId                 *int64     `json:"service_template_id"`
	CheckCommandId                    *int64     `json:"check_command_id"`
	CheckCommandArgs                  []string   `json:"check_command_args"`
	CheckTimeperiodId                 *int64     `json:"check_timeperiod_id"`
	MaxCheckAttempts                  int        `json:"max_check_attempts"`
	NormalCheckInterval               int        `json:"normal_check_interval"`
	RetryCheckInterval                int        `json:"retry_check_interval"`
	ActiveCheckEnabled                CheckState `json:"active_check_enabled"`
	PassiveCheckEnabled               CheckState `json:"passive_check_enabled"`
	VolatilityEnabled                 CheckState `json:"volatility_enabled"`
	NotificationEnabled               CheckState `json:"notification_enabled"`
	IsContactAdditiveInheritance      bool       `json:"is_contact_additive_inheritance"`
	IsContactGroupAdditiveInheritance bool       `json:"is_contact_group_additive_inheritance"`
	NotificationInterval              int        `json:"notification_interval"`
	NotificationTimeperiodId          *int64     `json:"notification_timeperiod_id"`
	NotificationType                  int        `json:"notification_type"`
	FirstNotificationDelay            int        `json:"first_notification_delay"`
	RecoveryNotificationDelay         int        `json:"recovery_notification_delay"`
	AcknowledgementTimeout            int        `json:"acknowledgement_timeout"`
	FreshnessChecked                  CheckState `json:"freshness_checked"`
	FreshnessThreshold                int        `json:"freshness_threshold"`
	FlapDetectionEnabled              CheckState `json:"flap_detection_enabled"`
	LowFlapThreshold                  int        `json:"low_flap_threshold"`
	HighFlapThreshold                 int        `json:"high_flap_threshold"`
	EventHandlerEnabled               CheckState `json:"event_handler_enabled"`
	EventHandlerCommandId             *int64     `json:"event_handler_command_id"`
	EventHandlerCommandArgs           []string   `json:"event_handler_command_args"`
	GraphTemplateId                   *int64     `json:"graph_template_id"`
	Note                              *string    `json:"note"`
	NoteUrl                           *string    `json:"note_url"`
	ActionUrl                         *string    `json:"action_url"`
	IconId                            *int64     `json:"icon_id"`
	IconAlternative                   *string    `json:"icon_alternative"`
	SeverityId                        *int64     `json:"severity_id"`
	IsActivated                       bool       `json:"is_activated"`
	Categories                        []IdName   `json:"categories"`
	Groups                            []IdName   `json:"groups"`
	Macros                            []Macro    `json:"macros"`
}

// ServiceUpdateRequest represents the payload to update a service in Centreon.
// https://docs-api.centreon.com/api/centreon-web/25.10/#tag/Service/paths/~1configuration~1services~1%7Bservice_id%7D/patch
type ServiceUpdateRequest struct {
	Name                              *string     `json:"name,omitempty" validate:"omitempty,required,max=200"`
	HostId                            *int64      `json:"host_id,omitempty" validate:"omitempty,required"`
	GeoCoords                         *string     `json:"geo_coords,omitempty" validate:"omitempty,max=32"`
	Comment                           *string     `json:"comment,omitempty"`
	ServiceTemplateId                 *int64      `json:"service_template_id,omitempty"`
	CheckCommandId                    *int64      `json:"check_command_id,omitempty"`
	CheckCommandArgs                  []string    `json:"check_command_args,omitempty"`
	CheckTimeperiodId                 *int64      `json:"check_timeperiod_id,omitempty"`
	MaxCheckAttempts                  *int        `json:"max_check_attempts,omitempty"`
	NormalCheckInterval               *int        `json:"normal_check_interval,omitempty"`
	RetryCheckInterval                *int        `json:"retry_check_interval,omitempty"`
	ActiveCheckEnabled                *CheckState `json:"active_check_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	PassiveCheckEnabled               *CheckState `json:"passive_check_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	VolatilityEnabled                 *CheckState `json:"volatility_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	NotificationEnabled               *CheckState `json:"notification_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	IsContactAdditiveInheritance      *bool       `json:"is_contact_additive_inheritance,omitempty"`
	IsContactGroupAdditiveInheritance *bool       `json:"is_contact_group_additive_inheritance,omitempty"`
	NotificationInterval              *int        `json:"notification_interval,omitempty"`
	NotificationTimeperiodId          *int64      `json:"notification_timeperiod_id,omitempty"`
	NotificationType                  *int        `json:"notification_type,omitempty"`
	FirstNotificationDelay            *int        `json:"first_notification_delay,omitempty"`
	RecoveryNotificationDelay         *int        `json:"recovery_notification_delay,omitempty"`
	AcknowledgementTimeout            *int        `json:"acknowledgement_timeout,omitempty"`
	FreshnessChecked                  *CheckState `json:"freshness_checked,omitempty" validate:"omitempty,oneof=0 1 2"`
	FreshnessThreshold                *int        `json:"freshness_threshold,omitempty"`
	FlapDetectionEnabled              *CheckState `json:"flap_detection_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	LowFlapThreshold                  *int        `json:"low_flap_threshold,omitempty"`
	HighFlapThreshold                 *int        `json:"high_flap_threshold,omitempty"`
	EventHandlerEnabled               *CheckState `json:"event_handler_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	EventHandlerCommandId             *int64      `json:"event_handler_command_id,omitempty"`
	EventHandlerCommandArgs           []string    `json:"event_handler_command_args,omitempty"`
	GraphTemplateId                   *int64      `json:"graph_template_id,omitempty"`
	Note                              *string     `json:"note,omitempty" validate:"omitempty,max=65535"`
	NoteUrl                           *string     `json:"note_url,omitempty" validate:"omitempty,max=65535"`
	ActionUrl                         *string     `json:"action_url,omitempty" validate:"omitempty,max=65535"`
	IconId                            *int64      `json:"icon_id,omitempty"`
	IconAlternative                   *string     `json:"icon_alternative,omitempty" validate:"omitempty,max=200"`
	SeverityId                        *int64      `json:"severity_id,omitempty"`
	IsActivated                       *bool       `json:"is_activated,omitempty"`
	ServiceCategories                 []int64     `json:"service_categories,omitempty"`
	ServiceGroups                     []int64     `json:"service_groups,omitempty"`
	Macros                            []Macro     `json:"macros,omitempty"`
}

// ServiceGroupHost represents a service group along with its associated host information.
type ServiceGroupHost struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	HostId   int64  `json:"host_id"`
	HostName string `json:"host_name"`
}

// ServiceListResponse represents the response structure for a service listing in Centreon.
// https://docs-api.centreon.com/api/centreon-web/25.10/#tag/Service/paths/~1configuration~1services/get
type ServiceListResponse struct {
	Id                     int64              `json:"id"`
	Name                   string             `json:"name"`
	Hosts                  []IdName           `json:"hosts"`
	ServiceTemplate        *IdName            `json:"service_template"`
	CheckTimeperiod        *IdName            `json:"check_timeperiod"`
	NotificationTimeperiod *IdName            `json:"notification_timeperiod"`
	Severity               *IdName            `json:"severity"`
	Categories             []IdName           `json:"categories"`
	Groups                 []ServiceGroupHost `json:"groups"`
	NormalCheckInterval    int                `json:"normal_check_interval"`
	RetryCheckInterval     int                `json:"retry_check_interval"`
	IsActivated            bool               `json:"is_activated"`
}

// ServiceStatus represents the status details of a service in Centreon.
type ServiceStatus struct {
	Code         ServiceCheckState `json:"code"`
	Name         string            `json:"name"`
	SeverityCode int               `json:"severity_code"`
}

// ServiceDowntime represents the downtime details for a service in Centreon.
type ServiceDowntime struct {
	Id              int64      `json:"id"`
	AuthorId        int64      `json:"author_id"`
	AuthorName      string     `json:"author_name"`
	HostId          int64      `json:"host_id"`
	Comment         *string    `json:"comment"`
	Duration        int        `json:"duration"`
	EntryTime       time.Time  `json:"entry_time"`
	StartTime       time.Time  `json:"start_time"`
	EndTime         time.Time  `json:"end_time"`
	DeletionTime    *time.Time `json:"deletion_time"`
	ActualStartTime *time.Time `json:"actual_start_time"`
	ActualEndTime   *time.Time `json:"actual_end_time"`
	IsStarted       bool       `json:"is_started"`
	IsCancelled     bool       `json:"is_cancelled"`
	IsFixed         bool       `json:"is_fixed"`
	ServiceId       int64      `json:"service_id"`
}

// ServiceAcknowledgement represents the acknowledgement details for a service in Centreon.
type ServiceAcknowledgement struct {
	Id                  int64             `json:"id"`
	AuthorId            int64             `json:"author_id"`
	AuthorName          string            `json:"author_name"`
	Comment             *string           `json:"comment"`
	DeletionTime        *time.Time        `json:"deletion_time"`
	EntryTime           time.Time         `json:"entry_time"`
	HostId              int64             `json:"host_id"`
	PollerId            int64             `json:"poller_id"`
	IsNotifyContacts    bool              `json:"is_notify_contacts"`
	IsPersistentComment bool              `json:"is_persistent_comment"`
	IsSticky            bool              `json:"is_sticky"`
	State               AcknowledgedState `json:"state"`
	ServiceId           int64             `json:"service_id"`
}

// ServiceRealTimeResponse represents the detailed response structure for a service in Centreon.
// https://docs-api.centreon.com/api/centreon-web/25.10/#tag/Service/paths/~1monitoring~1hosts~1%7Bhost_id%7D~1services~1%7Bservice_id%7D/get
type ServiceRealTimeResponse struct {
	Id                     int64                   `json:"id"`
	Description            *string                 `json:"description"`
	DisplayName            *string                 `json:"display_name"`
	State                  ServiceCheckState       `json:"state"`
	CheckAttempt           *int                    `json:"check_attempt"`
	IconImage              *string                 `json:"icon_image"`
	IconImageAlt           *string                 `json:"icon_image_alt"`
	LastCheck              *time.Time              `json:"last_check"`
	LastStateChange        *time.Time              `json:"last_state_change"`
	MaxCheckAttempts       *int                    `json:"max_check_attempts"`
	Output                 string                  `json:"output"`
	StateType              ServiceStateType        `json:"state_type"`
	Criticality            int                     `json:"criticality"`
	Status                 *ServiceStatus          `json:"status"`
	Duration               *string                 `json:"duration"`
	CheckCommand           *string                 `json:"check_command"`
	CheckInterval          *float64                `json:"check_interval"`
	CheckPeriod            *string                 `json:"check_period"`
	CheckType              ServiceCheckType        `json:"check_type"`
	CommandLine            *string                 `json:"command_line"`
	ExecutionTime          *float64                `json:"execution_time"`
	IsAcknowledged         bool                    `json:"is_acknowledged"`
	IsActiveCheck          bool                    `json:"is_active_check"`
	IsChecked              bool                    `json:"is_checked"`
	LastHardStateChange    *time.Time              `json:"last_hard_state_change"`
	LastNotification       *time.Time              `json:"last_notification"`
	LastTimeCritical       *time.Time              `json:"last_time_critical"`
	LastTimeOk             *time.Time              `json:"last_time_ok"`
	LastTimeUnknown        *time.Time              `json:"last_time_unknown"`
	LastTimeWarning        *time.Time              `json:"last_time_warning"`
	LastUpdate             *time.Time              `json:"last_update"`
	Latency                *float64                `json:"latency"`
	NextCheck              *time.Time              `json:"next_check"`
	PerformanceData        *string                 `json:"performance_data"`
	ScheduledDowntimeDepth int                     `json:"scheduled_downtime_depth"`
	Downtimes              []ServiceDowntime       `json:"downtimes"`
	Acknowledgement        *ServiceAcknowledgement `json:"acknowledgement"`
	Flapping               bool                    `json:"flapping"`
	Notify                 bool                    `json:"notify"`
}

// ServiceCheckState represents the check state of a service.
type ServiceCheckState int

const (
	ServiceCheckStateOk       ServiceCheckState = 0
	ServiceCheckStateWarning  ServiceCheckState = 1
	ServiceCheckStateCritical ServiceCheckState = 2
	ServiceCheckStateUnknown  ServiceCheckState = 3
	ServiceCheckStatePending  ServiceCheckState = 4
)

// ServiceStateType represents the state type of a service (soft or hard).
type ServiceStateType int

const (
	ServiceStateTypeSoft ServiceStateType = 0
	ServiceStateTypeHard ServiceStateType = 1
)

// ServiceCheckType represents the check type of a service (active or passive).
type ServiceCheckType int

const (
	ServiceCheckTypeActive  ServiceCheckType = 0
	ServiceCheckTypePassive ServiceCheckType = 1
)


