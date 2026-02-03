package api

// ServiceTemplateCreateRequest represents the payload to create a service template in Centreon.
// https://docs-api.centreon.com/api/centreon-web/25.10/#tag/Service-template/paths/~1configuration~1services~1templates/post
type ServiceTemplateCreateRequest struct {
	Name                              string                           `json:"name" validate:"required"`
	Alias                             string                           `json:"alias" validate:"required"`
	Comment                           *string                          `json:"comment,omitempty"`
	ServiceTemplateId                 *int64                           `json:"service_template_id,omitempty"`
	CheckCommandId                    *int64                           `json:"check_command_id,omitempty"`
	CheckCommandArgs                  []string                         `json:"check_command_args,omitempty"`
	CheckTimeperiodId                 *int64                           `json:"check_timeperiod_id,omitempty"`
	MaxCheckAttempts                  *int                             `json:"max_check_attempts,omitempty"`
	NormalCheckInterval               *int                             `json:"normal_check_interval,omitempty"`
	RetryCheckInterval                *int                             `json:"retry_check_interval,omitempty"`
	ActiveCheckEnabled                *CheckState                      `json:"active_check_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	PassiveCheckEnabled               *CheckState                      `json:"passive_check_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	VolatilityEnabled                 *CheckState                      `json:"volatility_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	NotificationEnabled               *CheckState                      `json:"notification_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	IsContactAdditiveInheritance      *bool                            `json:"is_contact_additive_inheritance,omitempty"`
	IsContactGroupAdditiveInheritance *bool                            `json:"is_contact_group_additive_inheritance,omitempty"`
	NotificationInterval              *int                             `json:"notification_interval,omitempty"`
	NotificationTimeperiodId          *int64                           `json:"notification_timeperiod_id,omitempty"`
	NotificationType                  *int                             `json:"notification_type,omitempty"`
	FirstNotificationDelay            *int                             `json:"first_notification_delay,omitempty"`
	RecoveryNotificationDelay         *int                             `json:"recovery_notification_delay,omitempty"`
	AcknowledgementTimeout            *int                             `json:"acknowledgement_timeout,omitempty"`
	FreshnessChecked                  *CheckState                      `json:"freshness_checked,omitempty" validate:"omitempty,oneof=0 1 2"`
	FreshnessThreshold                *int                             `json:"freshness_threshold,omitempty"`
	FlapDetectionEnabled              *CheckState                      `json:"flap_detection_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	LowFlapThreshold                  *int                             `json:"low_flap_threshold,omitempty"`
	HighFlapThreshold                 *int                             `json:"high_flap_threshold,omitempty"`
	EventHandlerEnabled               *CheckState                      `json:"event_handler_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	EventHandlerCommandId             *int64                           `json:"event_handler_command_id,omitempty"`
	EventHandlerCommandArgs           []string                         `json:"event_handler_command_args,omitempty"`
	GraphTemplateId                   *int64                           `json:"graph_template_id,omitempty"`
	Note                              *string                          `json:"note,omitempty"`
	NoteUrl                           *string                          `json:"note_url,omitempty"`
	ActionUrl                         *string                          `json:"action_url,omitempty"`
	IconId                            *int64                           `json:"icon_id,omitempty"`
	IconAlternative                   *string                          `json:"icon_alternative,omitempty"`
	SeverityId                        *int64                           `json:"severity_id,omitempty"`
	HostTemplates                     []int64                          `json:"host_templates,omitempty"`
	ServiceCategories                 []int64                          `json:"service_categories,omitempty"`
	ServiceGroups                     []ServiceTemplateGroupAssignment `json:"service_groups,omitempty"`
	Macros                            []Macro                          `json:"macros,omitempty"`
}

// ServiceTemplateUpdateRequest represents the payload to update a service template in Centreon.
// https://docs-api.centreon.com/api/centreon-web/25.10/#tag/Service-template/paths/~1configuration~1services~1templates~1%7Bservice_template_id%7D/patch
type ServiceTemplateUpdateRequest struct {
	Name                              *string                          `json:"name,omitempty" validate:"omitempty,required"`
	Alias                             *string                          `json:"alias,omitempty" validate:"omitempty,required"`
	Comment                           *string                          `json:"comment,omitempty"`
	ServiceTemplateId                 *int64                           `json:"service_template_id,omitempty"`
	CheckCommandId                    *int64                           `json:"check_command_id,omitempty"`
	CheckCommandArgs                  []string                         `json:"check_command_args,omitempty"`
	CheckTimeperiodId                 *int64                           `json:"check_timeperiod_id,omitempty"`
	MaxCheckAttempts                  *int                             `json:"max_check_attempts,omitempty"`
	NormalCheckInterval               *int                             `json:"normal_check_interval,omitempty"`
	RetryCheckInterval                *int                             `json:"retry_check_interval,omitempty"`
	ActiveCheckEnabled                *CheckState                      `json:"active_check_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	PassiveCheckEnabled               *CheckState                      `json:"passive_check_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	VolatilityEnabled                 *CheckState                      `json:"volatility_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	NotificationEnabled               *CheckState                      `json:"notification_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	IsContactAdditiveInheritance      *bool                            `json:"is_contact_additive_inheritance,omitempty"`
	IsContactGroupAdditiveInheritance *bool                            `json:"is_contact_group_additive_inheritance,omitempty"`
	NotificationInterval              *int                             `json:"notification_interval,omitempty"`
	NotificationTimeperiodId          *int64                           `json:"notification_timeperiod_id,omitempty"`
	NotificationType                  *int                             `json:"notification_type,omitempty"`
	FirstNotificationDelay            *int                             `json:"first_notification_delay,omitempty"`
	RecoveryNotificationDelay         *int                             `json:"recovery_notification_delay,omitempty"`
	AcknowledgementTimeout            *int                             `json:"acknowledgement_timeout,omitempty"`
	FreshnessChecked                  *CheckState                      `json:"freshness_checked,omitempty" validate:"omitempty,oneof=0 1 2"`
	FreshnessThreshold                *int                             `json:"freshness_threshold,omitempty"`
	FlapDetectionEnabled              *CheckState                      `json:"flap_detection_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	LowFlapThreshold                  *int                             `json:"low_flap_threshold,omitempty"`
	HighFlapThreshold                 *int                             `json:"high_flap_threshold,omitempty"`
	EventHandlerEnabled               *CheckState                      `json:"event_handler_enabled,omitempty" validate:"omitempty,oneof=0 1 2"`
	EventHandlerCommandId             *int64                           `json:"event_handler_command_id,omitempty"`
	EventHandlerCommandArgs           []string                         `json:"event_handler_command_args,omitempty"`
	GraphTemplateId                   *int64                           `json:"graph_template_id,omitempty"`
	Note                              *string                          `json:"note,omitempty"`
	NoteUrl                           *string                          `json:"note_url,omitempty"`
	ActionUrl                         *string                          `json:"action_url,omitempty"`
	IconId                            *int64                           `json:"icon_id,omitempty"`
	IconAlternative                   *string                          `json:"icon_alternative,omitempty"`
	SeverityId                        *int64                           `json:"severity_id,omitempty"`
	HostTemplates                     []int64                          `json:"host_templates,omitempty"`
	ServiceCategories                 []int64                          `json:"service_categories,omitempty"`
	ServiceGroups                     []ServiceTemplateGroupAssignment `json:"service_groups,omitempty"`
	Macros                            []Macro                          `json:"macros,omitempty"`
}

// ServiceTemplateGroupAssignment represents a service group assignment in a service template create request.
type ServiceTemplateGroupAssignment struct {
	ServiceGroupId int64  `json:"service_group_id" validate:"required"`
	HostTemplateId *int64 `json:"host_template_id,omitempty"`
}

// ServiceTemplateListResponse represents a service template in list responses from Centreon.
type ServiceTemplateListResponse struct {
	Id                                int64                            `json:"id"`
	IsLocked                          bool                             `json:"is_locked"`
	Name                              string                           `json:"name"`
	Alias                             string                           `json:"alias"`
	Comment                           *string                          `json:"comment"`
	ServiceTemplateId                 *int64                           `json:"service_template_id"`
	CheckCommandId                    *int64                           `json:"check_command_id"`
	CheckCommandArgs                  []string                         `json:"check_command_args"`
	CheckTimeperiodId                 *int64                           `json:"check_timeperiod_id"`
	MaxCheckAttempts                  *int                             `json:"max_check_attempts"`
	NormalCheckInterval               *int                             `json:"normal_check_interval"`
	RetryCheckInterval                *int                             `json:"retry_check_interval"`
	ActiveCheckEnabled                CheckState                       `json:"active_check_enabled"`
	PassiveCheckEnabled               CheckState                       `json:"passive_check_enabled"`
	VolatilityEnabled                 CheckState                       `json:"volatility_enabled"`
	NotificationEnabled               CheckState                       `json:"notification_enabled"`
	IsContactAdditiveInheritance      bool                             `json:"is_contact_additive_inheritance"`
	IsContactGroupAdditiveInheritance bool                             `json:"is_contact_group_additive_inheritance"`
	NotificationInterval              *int                             `json:"notification_interval"`
	NotificationTimeperiodId          *int64                           `json:"notification_timeperiod_id"`
	NotificationType                  *int                             `json:"notification_type"`
	FirstNotificationDelay            *int                             `json:"first_notification_delay"`
	RecoveryNotificationDelay         *int                             `json:"recovery_notification_delay"`
	AcknowledgementTimeout            *int                             `json:"acknowledgement_timeout"`
	FreshnessChecked                  CheckState                       `json:"freshness_checked"`
	FreshnessThreshold                *int                             `json:"freshness_threshold"`
	FlapDetectionEnabled              CheckState                       `json:"flap_detection_enabled"`
	LowFlapThreshold                  *int                             `json:"low_flap_threshold"`
	HighFlapThreshold                 *int                             `json:"high_flap_threshold"`
	EventHandlerEnabled               CheckState                       `json:"event_handler_enabled"`
	EventHandlerCommandId             *int64                           `json:"event_handler_command_id"`
	EventHandlerCommandArgs           []string                         `json:"event_handler_command_args"`
	GraphTemplateId                   *int64                           `json:"graph_template_id"`
	Note                              *string                          `json:"note"`
	NoteUrl                           *string                          `json:"note_url"`
	ActionUrl                         *string                          `json:"action_url"`
	IconId                            *int64                           `json:"icon_id"`
	IconAlternative                   *string                          `json:"icon_alternative"`
	SeverityId                        *int64                           `json:"severity_id"`
	HostTemplates                     []int64                          `json:"host_templates"`
	ServiceCategories                 []int64                          `json:"service_categories"`
	ServiceGroups                     []ServiceTemplateGroupAssignment `json:"service_groups"`
	Macros                            []Macro                          `json:"macros"`
}

// ServiceTemplateResponse represents the response of a service template in Centreon.
type ServiceTemplateResponse struct {
	Id                                int64                  `json:"id"`
	IsLocked                          bool                   `json:"is_locked"`
	Name                              string                 `json:"name"`
	Alias                             string                 `json:"alias"`
	Comment                           *string                `json:"comment"`
	ServiceTemplateId                 *int64                 `json:"service_template_id"`
	CheckCommandId                    *int64                 `json:"check_command_id"`
	CheckCommandArgs                  []string               `json:"check_command_args"`
	CheckTimeperiodId                 *int64                 `json:"check_timeperiod_id"`
	MaxCheckAttempts                  *int                   `json:"max_check_attempts"`
	NormalCheckInterval               *int                   `json:"normal_check_interval"`
	RetryCheckInterval                *int                   `json:"retry_check_interval"`
	ActiveCheckEnabled                CheckState             `json:"active_check_enabled"`
	PassiveCheckEnabled               CheckState             `json:"passive_check_enabled"`
	VolatilityEnabled                 CheckState             `json:"volatility_enabled"`
	NotificationEnabled               CheckState             `json:"notification_enabled"`
	IsContactAdditiveInheritance      bool                   `json:"is_contact_additive_inheritance"`
	IsContactGroupAdditiveInheritance bool                   `json:"is_contact_group_additive_inheritance"`
	NotificationInterval              *int                   `json:"notification_interval"`
	NotificationTimeperiodId          *int64                 `json:"notification_timeperiod_id"`
	NotificationType                  *int                   `json:"notification_type"`
	FirstNotificationDelay            *int                   `json:"first_notification_delay"`
	RecoveryNotificationDelay         *int                   `json:"recovery_notification_delay"`
	AcknowledgementTimeout            *int                   `json:"acknowledgement_timeout"`
	FreshnessChecked                  CheckState             `json:"freshness_checked"`
	FreshnessThreshold                *int                   `json:"freshness_threshold"`
	FlapDetectionEnabled              CheckState             `json:"flap_detection_enabled"`
	LowFlapThreshold                  *int                   `json:"low_flap_threshold"`
	HighFlapThreshold                 *int                   `json:"high_flap_threshold"`
	EventHandlerEnabled               CheckState             `json:"event_handler_enabled"`
	EventHandlerCommandId             *int64                 `json:"event_handler_command_id"`
	EventHandlerCommandArgs           []string               `json:"event_handler_command_args"`
	GraphTemplateId                   *int64                 `json:"graph_template_id"`
	Note                              *string                `json:"note"`
	NoteUrl                           *string                `json:"note_url"`
	ActionUrl                         *string                `json:"action_url"`
	IconId                            *int64                 `json:"icon_id"`
	IconAlternative                   *string                `json:"icon_alternative"`
	SeverityId                        *int64                 `json:"severity_id"`
	HostTemplates                     []int64                `json:"host_templates"`
	Categories                        []IdName               `json:"categories"`
	Groups                            []ServiceTemplateGroup `json:"groups"`
	Macros                            []Macro                `json:"macros"`
}

// ServiceTemplateGroup represents a service group within a service template response.
type ServiceTemplateGroup struct {
	Id               int64   `json:"id"`
	Name             string  `json:"name"`
	HostTemplateId   *int64  `json:"host_template_id"`
	HostTemplateName *string `json:"host_template_name"`
}
