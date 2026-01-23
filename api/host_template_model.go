package api

// HostTemplateResponse represents a host template reponse in Centreon.
// https://docs-api.centreon.com/api/centreon-web/25.10/#tag/Host-template/paths/~1configuration~1hosts~1templates/get
type HostTemplateListResponse struct {
	Id                        int64        `json:"id"`
	Name                      string       `json:"name"`
	Alias                     string       `json:"alias"`
	SnmpVersion               *SnmpVersion `json:"snmp_version"`
	TimezoneId                *int64       `json:"timezone_id"`
	SeverityId                *int64       `json:"severity_id"`
	CheckCommandId            *int64       `json:"check_command_id"`
	CheckCommandArgs          []string     `json:"check_command_args"`
	CheckTimeperiodId         *int64       `json:"check_timeperiod_id"`
	MaxCheckAttempts          *int         `json:"max_check_attempts"`
	NormalCheckInterval       *int         `json:"normal_check_interval"`
	RetryCheckInterval        *int         `json:"retry_check_interval"`
	ActiveCheckEnabled        *CheckState  `json:"active_check_enabled"`
	PassiveCheckEnabled       *CheckState  `json:"passive_check_enabled"`
	NotificationEnabled       *CheckState  `json:"notification_enabled"`
	NotificationOptions       *int         `json:"notification_options"`
	NotificationInterval      *int         `json:"notification_interval"`
	NotificationTimeperiodId  *int64       `json:"notification_timeperiod_id"`
	AddInheritedContactGroup  bool         `json:"add_inherited_contact_group"`
	AddInheritedContact       bool         `json:"add_inherited_contact"`
	FirstNotificationDelay    *int         `json:"first_notification_delay"`
	RecoveryNotificationDelay *int         `json:"recovery_notification_delay"`
	AcknowledgementTimeout    *int         `json:"acknowledgement_timeout"`
	FreshnessChecked          *CheckState  `json:"freshness_checked"`
	FreshnessThreshold        *int         `json:"freshness_threshold"`
	FlapDetectionEnabled      *CheckState  `json:"flap_detection_enabled"`
	LowFlapThreshold          *int         `json:"low_flap_threshold"`
	HighFlapThreshold         *int         `json:"high_flap_threshold"`
	EventHandlerEnabled       *CheckState  `json:"event_handler_enabled"`
	EventHandlerCommandId     *int64       `json:"event_handler_command_id"`
	EventHandlerCommandArgs   []string     `json:"event_handler_command_args"`
	NoteUrl                   *string      `json:"note_url"`
	Note                      *string      `json:"note"`
	ActionUrl                 *string      `json:"action_url"`
	IconId                    *int64       `json:"icon_id"`
	IconAlternative           *string      `json:"icon_alternative"`
	Comment                   *string      `json:"comment"`
	IsLocked                  bool         `json:"is_locked"`
}

// HostTemplateCreateRequest represents the payload to create a host template in Centreon.
// https://docs-api.centreon.com/api/centreon-web/25.10/#tag/Host-template/paths/~1configuration~1hosts~1templates/post
type HostTemplateCreateRequest struct {
	Name                      string       `json:"name" validate:"required"`
	Alias                     string       `json:"alias" validate:"required"`
	SnmpCommunity             *string      `json:"snmp_community,omitempty"`
	SnmpVersion               *SnmpVersion `json:"snmp_version,omitempty" validate:"omitempty,oneof=1 2c 3"`
	TimezoneId                *int64       `json:"timezone_id,omitempty"`
	SeverityId                *int64       `json:"severity_id,omitempty"`
	CheckCommandId            *int64       `json:"check_command_id,omitempty"`
	CheckCommandArgs          []string     `json:"check_command_args,omitempty"`
	CheckTimeperiodId         *int64       `json:"check_timeperiod_id,omitempty"`
	MaxCheckAttempts          *int         `json:"max_check_attempts,omitempty"`
	NormalCheckInterval       *int         `json:"normal_check_interval,omitempty"`
	RetryCheckInterval        *int         `json:"retry_check_interval,omitempty"`
	ActiveCheckEnabled        *CheckState  `json:"active_check_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	PassiveCheckEnabled       *CheckState  `json:"passive_check_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	NotificationEnabled       *CheckState  `json:"notification_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	NotificationOptions       *int         `json:"notification_options,omitempty"`
	NotificationInterval      *int         `json:"notification_interval,omitempty"`
	NotificationTimeperiodId  *int64       `json:"notification_timeperiod_id,omitempty"`
	AddInheritedContactGroup  *bool        `json:"add_inherited_contact_group,omitempty"`
	AddInheritedContact       *bool        `json:"add_inherited_contact,omitempty"`
	FirstNotificationDelay    *int         `json:"first_notification_delay,omitempty"`
	RecoveryNotificationDelay *int         `json:"recovery_notification_delay,omitempty"`
	AcknowledgementTimeout    *int         `json:"acknowledgement_timeout,omitempty"`
	FreshnessChecked          *CheckState  `json:"freshness_checked,omitempty" validate:"omitempty,oneof=0 1 2"`
	FreshnessThreshold        *int         `json:"freshness_threshold,omitempty"`
	FlapDetectionEnabled      *CheckState  `json:"flap_detection_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	LowFlapThreshold          *int         `json:"low_flap_threshold,omitempty"`
	HighFlapThreshold         *int         `json:"high_flap_threshold,omitempty"`
	EventHandlerEnabled       *CheckState  `json:"event_handler_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	EventHandlerCommandId     *int64       `json:"event_handler_command_id,omitempty"`
	EventHandlerCommandArgs   []string     `json:"event_handler_command_args,omitempty"`
	NoteUrl                   *string      `json:"note_url,omitempty"`
	Note                      *string      `json:"note,omitempty"`
	ActionUrl                 *string      `json:"action_url,omitempty"`
	IconId                    *int64       `json:"icon_id,omitempty"`
	IconAlternative           *string      `json:"icon_alternative,omitempty"`
	Comment                   *string      `json:"comment,omitempty"`
	Categories                []int64      `json:"categories,omitempty"`
	Templates                 []int64      `json:"templates,omitempty"`
	Macros                    []Macro      `json:"macros,omitempty"`
}

// HostTemplateResponse represents the response of a host template in Centreon.
// https://docs-api.centreon.com/api/centreon-web/25.10/#tag/Host-template/paths/~1configuration~1hosts~1templates/post
type HostTemplateResponse struct {
	Id                        int64        `json:"id"`
	Name                      string       `json:"name"`
	Alias                     string       `json:"alias"`
	SnmpVersion               *SnmpVersion `json:"snmp_version"`
	TimezoneId                *int64       `json:"timezone_id"`
	SeverityId                *int64       `json:"severity_id"`
	CheckCommandId            *int64       `json:"check_command_id"`
	CheckCommandArgs          []string     `json:"check_command_args"`
	CheckTimeperiodId         *int64       `json:"check_timeperiod_id"`
	MaxCheckAttempts          *int         `json:"max_check_attempts"`
	NormalCheckInterval       *int         `json:"normal_check_interval"`
	RetryCheckInterval        *int         `json:"retry_check_interval"`
	ActiveCheckEnabled        CheckState   `json:"active_check_enabled"`
	PassiveCheckEnabled       CheckState   `json:"passive_check_enabled"`
	NotificationEnabled       CheckState   `json:"notification_enabled"`
	NotificationOptions       *int         `json:"notification_options"`
	NotificationInterval      *int         `json:"notification_interval"`
	NotificationTimeperiodId  *int64       `json:"notification_timeperiod_id"`
	AddInheritedContactGroup  bool         `json:"add_inherited_contact_group"`
	AddInheritedContact       bool         `json:"add_inherited_contact"`
	FirstNotificationDelay    *int         `json:"first_notification_delay"`
	RecoveryNotificationDelay *int         `json:"recovery_notification_delay"`
	AcknowledgementTimeout    *int         `json:"acknowledgement_timeout"`
	FreshnessChecked          CheckState   `json:"freshness_checked"`
	FreshnessThreshold        *int         `json:"freshness_threshold"`
	FlapDetectionEnabled      CheckState   `json:"flap_detection_enabled"`
	LowFlapThreshold          *int         `json:"low_flap_threshold"`
	HighFlapThreshold         *int         `json:"high_flap_threshold"`
	EventHandlerEnabled       CheckState   `json:"event_handler_enabled"`
	EventHandlerCommandId     *int64       `json:"event_handler_command_id"`
	EventHandlerCommandArgs   []string     `json:"event_handler_command_args"`
	NoteUrl                   *string      `json:"note_url"`
	Note                      *string      `json:"note"`
	ActionUrl                 *string      `json:"action_url"`
	IconId                    *int64       `json:"icon_id"`
	IconAlternative           *string      `json:"icon_alternative"`
	Comment                   *string      `json:"comment"`
	IsLocked                  bool         `json:"is_locked"`
	Categories                []int64      `json:"categories"`
	Templates                 []int64      `json:"templates"`
	Macros                    []Macro      `json:"macros"`
}

// HostTemplateUpdateRequest represents the payload to update a host template in Centreon.
// https://docs-api.centreon.com/api/centreon-web/25.10/#tag/Host-template/paths/~1configuration~1hosts~1templates~1%7Bhost_template_id%7D/patch
type HostTemplateUpdateRequest struct {
	Name                      *string      `json:"name,omitempty" validate:"omitempty,required"`
	Alias                     *string      `json:"alias,omitempty" validate:"omitempty,required"`
	SnmpCommunity             *string      `json:"snmp_community,omitempty"`
	SnmpVersion               *SnmpVersion `json:"snmp_version,omitempty" validate:"omitempty,oneof=1 2c 3"`
	TimezoneId                *int64       `json:"timezone_id,omitempty"`
	SeverityId                *int64       `json:"severity_id,omitempty"`
	CheckCommandId            *int64       `json:"check_command_id,omitempty"`
	CheckCommandArgs          []string     `json:"check_command_args,omitempty"`
	CheckTimeperiodId         *int64       `json:"check_timeperiod_id,omitempty"`
	MaxCheckAttempts          *int         `json:"max_check_attempts,omitempty"`
	NormalCheckInterval       *int         `json:"normal_check_interval,omitempty"`
	RetryCheckInterval        *int         `json:"retry_check_interval,omitempty"`
	ActiveCheckEnabled        *CheckState  `json:"active_check_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	PassiveCheckEnabled       *CheckState  `json:"passive_check_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	NotificationEnabled       *CheckState  `json:"notification_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	NotificationOptions       *int         `json:"notification_options,omitempty"`
	NotificationInterval      *int         `json:"notification_interval,omitempty"`
	NotificationTimeperiodId  *int64       `json:"notification_timeperiod_id,omitempty"`
	AddInheritedContactGroup  *bool        `json:"add_inherited_contact_group,omitempty"`
	AddInheritedContact       *bool        `json:"add_inherited_contact,omitempty"`
	FirstNotificationDelay    *int         `json:"first_notification_delay,omitempty"`
	RecoveryNotificationDelay *int         `json:"recovery_notification_delay,omitempty"`
	AcknowledgementTimeout    *int         `json:"acknowledgement_timeout,omitempty"`
	FreshnessChecked          *CheckState  `json:"freshness_checked,omitempty" validate:"omitempty,oneof=0 1 2"`
	FreshnessThreshold        *int         `json:"freshness_threshold,omitempty"`
	FlapDetectionEnabled      *CheckState  `json:"flap_detection_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	LowFlapThreshold          *int         `json:"low_flap_threshold,omitempty"`
	HighFlapThreshold         *int         `json:"high_flap_threshold,omitempty"`
	EventHandlerEnabled       *CheckState  `json:"event_handler_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	EventHandlerCommandId     *int64       `json:"event_handler_command_id,omitempty"`
	EventHandlerCommandArgs   []string     `json:"event_handler_command_args,omitempty"`
	NoteUrl                   *string      `json:"note_url,omitempty"`
	Note                      *string      `json:"note,omitempty"`
	ActionUrl                 *string      `json:"action_url,omitempty"`
	IconId                    *int64       `json:"icon_id,omitempty"`
	IconAlternative           *string      `json:"icon_alternative,omitempty"`
	Comment                   *string      `json:"comment,omitempty"`
	Categories                []int64      `json:"categories,omitempty"`
	Templates                 []int64      `json:"templates,omitempty"`
	Macros                    []Macro      `json:"macros,omitempty"`
}
