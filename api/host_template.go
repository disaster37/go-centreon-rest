package api

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// HostTemplate represents a reusable configuration template for hosts in Centreon.
// Templates provide default values for host parameters, reducing configuration effort
// and ensuring consistency across similar hosts in the infrastructure.
// Based on Centreon API documentation: https://docs-api.centreon.com/api/centreon-web/25.10/#tag/Host-template
type HostTemplate struct {
	// Embed HostBase for common fields shared with Host
	HostBase

	// Template-specific fields that are not shared with Host

	// CheckCommandArgs contains the arguments for the check command
	// Allows customization of command parameters at the template level
	CheckCommandArgs string `json:"check_command_args,omitempty" validate:"omitempty,max=8192"`

	// RetryCheckInterval is the time (in minutes) between checks when the host is in a non-UP state
	// Typically shorter than normal check interval for faster recovery detection
	RetryCheckInterval *float64 `json:"retry_check_interval,omitempty" validate:"omitempty,gte=1,lte=1440"`

	// NotificationPeriod defines when notifications should be sent for hosts using this template
	// References a time period configuration for notification timing
	NotificationPeriod string `json:"notification_period,omitempty" validate:"omitempty,max=200"`

	// FirstNotificationDelay is the time (in minutes) to wait before sending the first notification
	// Allows for temporary problems to resolve themselves before alerting
	FirstNotificationDelay *float64 `json:"first_notification_delay,omitempty" validate:"omitempty,gte=0,lte=1440"`

	// RecoveryNotificationDelay is the time (in minutes) to wait before sending recovery notifications
	// Helps prevent premature recovery notifications for flapping hosts
	RecoveryNotificationDelay *float64 `json:"recovery_notification_delay,omitempty" validate:"omitempty,gte=0,lte=1440"`

	// EventHandlerEnabled controls whether event handlers are executed for state changes
	// Event handlers can perform automated recovery actions
	EventHandlerEnabled *bool `json:"event_handler_enabled,omitempty"`

	// EventHandler is the command executed when host state changes occur
	// Used for automated problem resolution and escalation
	EventHandler string `json:"event_handler,omitempty" validate:"omitempty,max=255"`

	// FlapDetectionEnabled controls whether flap detection is active
	// Helps identify unstable hosts that frequently change state
	FlapDetectionEnabled *bool `json:"flap_detection_enabled,omitempty"`

	// LowFlapThreshold is the percentage threshold below which flapping stops
	// Used in conjunction with HighFlapThreshold to define flapping behavior
	LowFlapThreshold *float64 `json:"low_flap_threshold,omitempty" validate:"omitempty,gte=0,lte=100"`

	// HighFlapThreshold is the percentage threshold above which flapping starts
	// Higher values make flap detection less sensitive
	HighFlapThreshold *float64 `json:"high_flap_threshold,omitempty" validate:"omitempty,gte=0,lte=100"`

	// ProcessPerfData controls whether performance data is processed
	// When true, performance metrics are stored for trending and reporting
	ProcessPerfData *bool `json:"process_perf_data,omitempty"`

	// RetainStatusInformation controls whether status information is retained across restarts
	// Helps maintain monitoring state consistency during system restarts
	RetainStatusInformation *bool `json:"retain_status_information,omitempty"`

	// RetainNonStatusInformation controls whether non-status information is retained
	// Includes acknowledgements, comments, and other metadata
	RetainNonStatusInformation *bool `json:"retain_nonstatus_information,omitempty"`

	// NotificationsEnabled is the master switch for all notifications
	// When false, no notifications are sent regardless of other notification settings
	NotificationsEnabled *bool `json:"notifications_enabled,omitempty"`

	// StalkingOnUp enables detailed logging when hosts are in UP state
	// Useful for debugging intermittent issues
	StalkingOnUp *bool `json:"stalking_on_up,omitempty"`

	// StalkingOnDown enables detailed logging when hosts are in DOWN state
	StalkingOnDown *bool `json:"stalking_on_down,omitempty"`

	// StalkingOnUnreachable enables detailed logging when hosts are UNREACHABLE
	StalkingOnUnreachable *bool `json:"stalking_on_unreachable,omitempty"`

	// FreshnessChecksEnabled controls whether freshness checks are performed
	// Ensures passive checks are received within expected timeframes
	FreshnessChecksEnabled *bool `json:"freshness_checks_enabled,omitempty"`

	// FreshnessThreshold is the maximum age (in seconds) for passive check results
	// After this time, the host is considered stale if no new checks arrive
	FreshnessThreshold *int64 `json:"freshness_threshold,omitempty" validate:"omitempty,gte=0,lte=86400"`

	// ObsessOverHost enables obsessive compulsive host checking
	// Causes all host check results to be sent to an external command
	ObsessOverHost *bool `json:"obsess_over_host,omitempty"`

	// Check2DCoords contains 2D coordinates for network topology mapping
	// Format: "x,y" coordinates for visual representation
	Check2DCoords string `json:"2d_coords,omitempty" validate:"omitempty,max=100"`

	// Check3DCoords contains 3D coordinates for advanced topology visualization
	// Format: "x,y,z" coordinates for 3D network maps
	Check3DCoords string `json:"3d_coords,omitempty" validate:"omitempty,max=100"`

	// ActionURL provides a custom action link for this template
	// Displayed in the web interface for quick access to related tools
	ActionURL string `json:"action_url,omitempty" validate:"omitempty,max=2000,url"`

	// Notes contains descriptive text about the template
	// Provides context and documentation for template usage
	Notes string `json:"notes,omitempty" validate:"omitempty,max=8192"`

	// NotesURL provides a link to external documentation about the template
	// Can point to wiki pages, runbooks, or other documentation
	NotesURL string `json:"notes_url,omitempty" validate:"omitempty,max=2000,url"`

	// VRMLImage is the VRML image for 3D visualization
	// Used in advanced 3D network topology displays
	VRMLImage string `json:"vrml_image,omitempty" validate:"omitempty,max=500"`

	// StatusmapImage is the image used in status map displays
	// Provides custom visualization in network status maps
	StatusmapImage string `json:"statusmap_image,omitempty" validate:"omitempty,max=500"`

	// Template inheritance and relationships

	// HostTemplateID references a parent template for inheritance
	// Allows template hierarchies where templates inherit from other templates
	HostTemplateID *int64 `json:"host_template_id,omitempty" validate:"omitempty,gt=0"`

	// ParentHostTemplates contains the list of parent template IDs
	// Supports multiple inheritance from different template sources
	ParentHostTemplates []int64 `json:"parent_host_templates,omitempty"`

	// ChildHostTemplates contains templates that inherit from this template
	// Read-only field showing template hierarchy relationships
	ChildHostTemplates []int64 `json:"child_host_templates,omitempty"`

	// Contact and notification relationships

	// ContactGroups are the groups notified about hosts using this template
	// Defines who receives notifications for host problems
	ContactGroups []string `json:"contact_groups,omitempty"`

	// Contacts are individual contacts notified about hosts using this template
	// Supplements contact group notifications with specific individuals
	Contacts []string `json:"contacts,omitempty"`

	// Host categorization and grouping

	// HostGroups are the default groups for hosts created from this template
	// Automatically assigns hosts to logical groupings
	HostGroups []string `json:"host_groups,omitempty"`

	// Categories are the host categories applied by default
	// Used for organizational and reporting purposes
	Categories []string `json:"categories,omitempty"`

	// Severity defines the default business criticality level
	// Used for incident prioritization and SLA management
	Severity *HostSeverity `json:"severity,omitempty"`

	// Extended configuration and macros

	// Macros contains custom macro definitions for this template
	// Key-value pairs that can be referenced in commands and notifications
	Macros map[string]string `json:"macros,omitempty"`

	// ExtendedInfo contains additional metadata and configuration
	// Flexible storage for template-specific settings and parameters
	ExtendedInfo map[string]any `json:"extended_info,omitempty"`

	// Audit and metadata fields

	// CreatedAt is when the template was initially created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// UpdatedAt is when the template was last modified
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// CreatedBy identifies who created the template
	CreatedBy string `json:"created_by,omitempty"`

	// UpdatedBy identifies who last modified the template
	UpdatedBy string `json:"updated_by,omitempty"`

	// Version tracks template configuration changes
	// Incremented with each modification for change tracking
	Version *int64 `json:"version,omitempty"`

	// Comment contains administrative notes about recent changes
	// Used for change management and audit trails
	Comment string `json:"comment,omitempty" validate:"omitempty,max=8192"`
}

// HostTemplateService represents a service template attached to a host template.
// These services are automatically created for any host that uses this template.
type HostTemplateService struct {
	// ID is the unique identifier for the service template relationship
	ID *int64 `json:"id,omitempty"`

	// ServiceTemplateID references the service template to be applied
	ServiceTemplateID *int64 `json:"service_template_id" validate:"required,gt=0"`

	// ServiceTemplateName is the name of the service template (read-only)
	ServiceTemplateName string `json:"service_template_name,omitempty"`

	// ServiceDescription is the description that will be used for created services
	// Can override the default description from the service template
	ServiceDescription string `json:"service_description,omitempty" validate:"omitempty,max=255"`

	// IsActivated controls whether this service template relationship is active
	IsActivated *bool `json:"is_activated,omitempty"`

	// Arguments contains service-specific arguments that override template defaults
	Arguments map[string]string `json:"arguments,omitempty"`

	// Macros contains service-specific macro overrides
	Macros map[string]string `json:"macros,omitempty"`
}

// HostTemplateDiscoveryRule represents discovery rules for automatic host creation.
// These rules define how to automatically create hosts based on discovered resources.
type HostTemplateDiscoveryRule struct {
	// ID is the unique identifier for the discovery rule
	ID *int64 `json:"id,omitempty"`

	// Name is the descriptive name for the discovery rule
	Name string `json:"name" validate:"required,min=1,max=255"`

	// IsActivated controls whether this discovery rule is active
	IsActivated *bool `json:"is_activated,omitempty"`

	// DiscoveryCommand is the command used to discover new hosts
	DiscoveryCommand string `json:"discovery_command" validate:"required,max=255"`

	// DiscoveryCommandArgs contains arguments for the discovery command
	DiscoveryCommandArgs string `json:"discovery_command_args,omitempty"`

	// DiscoveryInterval is how often (in minutes) to run discovery
	DiscoveryInterval *int64 `json:"discovery_interval,omitempty" validate:"omitempty,gte=1,lte=10080"`

	// HostNamePattern defines how to generate host names from discovery results
	// Can include variables that will be replaced with discovered values
	HostNamePattern string `json:"hostname_pattern,omitempty" validate:"omitempty,max=255"`

	// HostAliasPattern defines how to generate host aliases from discovery results
	HostAliasPattern string `json:"host_alias_pattern,omitempty" validate:"omitempty,max=255"`

	// AddressPattern defines how to extract addresses from discovery results
	AddressPattern string `json:"address_pattern,omitempty" validate:"omitempty,max=255"`

	// Filters define criteria for including/excluding discovered hosts
	Filters map[string]string `json:"filters,omitempty"`

	// AutoActivation controls whether discovered hosts are automatically activated
	AutoActivation *bool `json:"auto_activation,omitempty"`

	// LastDiscovery is when the rule was last executed
	LastDiscovery *time.Time `json:"last_discovery,omitempty"`

	// NextDiscovery is when the rule will next be executed
	NextDiscovery *time.Time `json:"next_discovery,omitempty"`

	// DiscoveredHostsCount tracks how many hosts have been discovered by this rule
	DiscoveredHostsCount *int64 `json:"discovered_hosts_count,omitempty"`
}

// Validation methods for HostTemplate

// ValidateForCreate validates the host template for creation operations
func (ht *HostTemplate) ValidateForCreate() error {
	if ht.Name == "" || len(strings.TrimSpace(ht.Name)) == 0 {
		return fmt.Errorf("name is required for host template creation")
	}
	if ht.Alias == "" || len(strings.TrimSpace(ht.Alias)) == 0 {
		return fmt.Errorf("alias is required for host template creation")
	}

	// Validate inheritance doesn't create cycles
	if ht.HostTemplateID != nil && ht.ID != nil && *ht.HostTemplateID == *ht.ID {
		return fmt.Errorf("host template cannot inherit from itself")
	}

	// Validate notification intervals
	if ht.NotificationInterval != nil && ht.CheckInterval != nil {
		if *ht.NotificationInterval < *ht.CheckInterval {
			return fmt.Errorf("notification interval should not be less than check interval")
		}
	}

	// Validate flap thresholds
	if ht.LowFlapThreshold != nil && ht.HighFlapThreshold != nil {
		if *ht.LowFlapThreshold >= *ht.HighFlapThreshold {
			return fmt.Errorf("low flap threshold must be less than high flap threshold")
		}
	}

	return nil
}

// ValidateForUpdate validates the host template for update operations
func (ht *HostTemplate) ValidateForUpdate() error {
	if ht.ID == nil || *ht.ID <= 0 {
		return fmt.Errorf("valid host template ID is required for update operations")
	}

	if ht.Name != "" && len(strings.TrimSpace(ht.Name)) == 0 {
		return fmt.Errorf("host template name cannot be empty")
	}

	if ht.Alias != "" && len(strings.TrimSpace(ht.Alias)) == 0 {
		return fmt.Errorf("host template alias cannot be empty")
	}

	// Validate inheritance doesn't create cycles
	if ht.HostTemplateID != nil && *ht.HostTemplateID == *ht.ID {
		return fmt.Errorf("host template cannot inherit from itself")
	}

	// Validate notification intervals
	if ht.NotificationInterval != nil && ht.CheckInterval != nil {
		if *ht.NotificationInterval < *ht.CheckInterval {
			return fmt.Errorf("notification interval should not be less than check interval")
		}
	}

	// Validate flap thresholds
	if ht.LowFlapThreshold != nil && ht.HighFlapThreshold != nil {
		if *ht.LowFlapThreshold >= *ht.HighFlapThreshold {
			return fmt.Errorf("low flap threshold must be less than high flap threshold")
		}
	}

	return nil
}

// ValidateForDuplicate validates the host template for duplication operations
func (ht *HostTemplate) ValidateForDuplicate() error {
	if ht.ID == nil || *ht.ID <= 0 {
		return fmt.Errorf("valid source host template ID is required for duplicate operations")
	}
	return nil
}

// ValidateForDelete validates the host template for deletion operations
func (ht *HostTemplate) ValidateForDelete() error {
	if ht.ID == nil || *ht.ID <= 0 {
		return fmt.Errorf("valid host template ID is required for delete operations")
	}
	return nil
}

// HasInheritance returns true if the template inherits from another template
func (ht *HostTemplate) HasInheritance() bool {
	return ht.HostTemplateID != nil || len(ht.ParentHostTemplates) > 0
}

// IsParent returns true if the template has child templates
func (ht *HostTemplate) IsParent() bool {
	return len(ht.ChildHostTemplates) > 0
}

// GetEffectiveCheckInterval returns the check interval, using a default if not set
func (ht *HostTemplate) GetEffectiveCheckInterval() float64 {
	if ht.CheckInterval != nil {
		return *ht.CheckInterval
	}
	return 5.0 // Default 5 minutes
}

// GetEffectiveMaxCheckAttempts returns the max check attempts, using a default if not set
func (ht *HostTemplate) GetEffectiveMaxCheckAttempts() int64 {
	if ht.MaxCheckAttempts != nil {
		return *ht.MaxCheckAttempts
	}
	return 3 // Default 3 attempts
}

// GetEffectiveNotificationInterval returns the notification interval, using a default if not set
func (ht *HostTemplate) GetEffectiveNotificationInterval() float64 {
	if ht.NotificationInterval != nil {
		return *ht.NotificationInterval
	}
	return 0.0 // Default: no repeat notifications
}

// IsActive returns whether the template is activated and can be used
func (ht *HostTemplate) IsActive() bool {
	return ht.IsActivated == nil || *ht.IsActivated
}

// IsReadOnly returns whether the template is locked and cannot be modified
func (ht *HostTemplate) IsReadOnly() bool {
	return ht.IsLocked != nil && *ht.IsLocked
}

// Validation methods for HostTemplateService

// ValidateForCreate validates the service template relationship for creation
func (hts *HostTemplateService) ValidateForCreate() error {
	if hts.ServiceTemplateID == nil || *hts.ServiceTemplateID <= 0 {
		return fmt.Errorf("valid service template ID is required")
	}
	return nil
}

// ValidateForUpdate validates the service template relationship for updates
func (hts *HostTemplateService) ValidateForUpdate() error {
	if hts.ID == nil || *hts.ID <= 0 {
		return fmt.Errorf("valid service template relationship ID is required for update operations")
	}
	if hts.ServiceTemplateID == nil || *hts.ServiceTemplateID <= 0 {
		return fmt.Errorf("valid service template ID is required")
	}
	return nil
}

// Validation methods for HostTemplateDiscoveryRule

// ValidateForCreate validates the discovery rule for creation
func (htdr *HostTemplateDiscoveryRule) ValidateForCreate() error {
	if htdr.Name == "" || len(strings.TrimSpace(htdr.Name)) == 0 {
		return fmt.Errorf("name is required for discovery rule creation")
	}
	if htdr.DiscoveryCommand == "" || len(strings.TrimSpace(htdr.DiscoveryCommand)) == 0 {
		return fmt.Errorf("discovery command is required for discovery rule creation")
	}
	return nil
}

// ValidateForUpdate validates the discovery rule for updates
func (htdr *HostTemplateDiscoveryRule) ValidateForUpdate() error {
	if htdr.ID == nil || *htdr.ID <= 0 {
		return fmt.Errorf("valid discovery rule ID is required for update operations")
	}
	if htdr.Name != "" && len(strings.TrimSpace(htdr.Name)) == 0 {
		return fmt.Errorf("discovery rule name cannot be empty")
	}
	if htdr.DiscoveryCommand != "" && len(strings.TrimSpace(htdr.DiscoveryCommand)) == 0 {
		return fmt.Errorf("discovery command cannot be empty")
	}
	return nil
}

// HostTemplateInterface defines CRUD operations for HostTemplate entities
type HostTemplateInterface interface {
	// Create creates a new host template
	Create(ctx context.Context, template *HostTemplate) (*HostTemplate, error)

	// GetByID retrieves a host template by its ID
	GetByID(ctx context.Context, id int64) (*HostTemplate, error)

	// GetByName retrieves a host template by its name
	GetByName(ctx context.Context, name string) (*HostTemplate, error)

	// List retrieves host templates with optional filtering and pagination
	List(ctx context.Context, opts *ListOptions) (*ListResponse[HostTemplate], error)

	// Update updates an existing host template
	Update(ctx context.Context, id int64, template *HostTemplate) error

	// Delete deletes a host template by ID
	Delete(ctx context.Context, id int64) error

	// Duplicate creates a copy of an existing host template with a new name
	Duplicate(ctx context.Context, id int64, newName, newAlias string) (*HostTemplate, error)

	// BulkCreate creates multiple host templates in a single operation
	BulkCreate(ctx context.Context, templates []HostTemplate) error

	// BulkUpdate updates multiple host templates in a single operation
	BulkUpdate(ctx context.Context, templates []HostTemplate) error

	// BulkDelete deletes multiple host templates by their IDs
	BulkDelete(ctx context.Context, ids []int64) error

	// GetInheritanceTree retrieves the complete inheritance hierarchy for a template
	GetInheritanceTree(ctx context.Context, id int64) (*HostTemplateInheritanceTree, error)

	// ListChildren retrieves all templates that inherit from the specified template
	ListChildren(ctx context.Context, id int64, opts *ListOptions) (*ListResponse[HostTemplate], error)

	// ListParents retrieves all templates that the specified template inherits from
	ListParents(ctx context.Context, id int64, opts *ListOptions) (*ListResponse[HostTemplate], error)

	// SetParentTemplates sets the parent templates for inheritance
	SetParentTemplates(ctx context.Context, id int64, parentIDs []int64) error

	// AddParentTemplate adds a parent template for inheritance
	AddParentTemplate(ctx context.Context, id int64, parentID int64) error

	// RemoveParentTemplate removes a parent template from inheritance
	RemoveParentTemplate(ctx context.Context, id int64, parentID int64) error

	// ListServices retrieves service templates associated with a host template
	ListServices(ctx context.Context, templateID int64, opts *ListOptions) (*ListResponse[HostTemplateService], error)

	// AddService adds a service template to a host template
	AddService(ctx context.Context, templateID int64, service *HostTemplateService) error

	// UpdateService updates a service template relationship
	UpdateService(ctx context.Context, templateID int64, serviceID int64, service *HostTemplateService) error

	// RemoveService removes a service template from a host template
	RemoveService(ctx context.Context, templateID int64, serviceID int64) error

	// ListDiscoveryRules retrieves discovery rules for a host template
	ListDiscoveryRules(ctx context.Context, templateID int64, opts *ListOptions) (*ListResponse[HostTemplateDiscoveryRule], error)

	// AddDiscoveryRule adds a discovery rule to a host template
	AddDiscoveryRule(ctx context.Context, templateID int64, rule *HostTemplateDiscoveryRule) error

	// UpdateDiscoveryRule updates a discovery rule
	UpdateDiscoveryRule(ctx context.Context, templateID int64, ruleID int64, rule *HostTemplateDiscoveryRule) error

	// RemoveDiscoveryRule removes a discovery rule from a host template
	RemoveDiscoveryRule(ctx context.Context, templateID int64, ruleID int64) error

	// ExecuteDiscoveryRule manually executes a discovery rule
	ExecuteDiscoveryRule(ctx context.Context, templateID int64, ruleID int64) (*DiscoveryResult, error)

	// GetMacros retrieves all macros defined in a host template (including inherited)
	GetMacros(ctx context.Context, id int64) (map[string]string, error)

	// SetMacros sets macros for a host template
	SetMacros(ctx context.Context, id int64, macros map[string]string) error

	// GetMacro retrieves a specific macro value
	GetMacro(ctx context.Context, id int64, macroName string) (string, error)

	// SetMacro sets a specific macro value
	SetMacro(ctx context.Context, id int64, macroName, macroValue string) error

	// DeleteMacro deletes a specific macro
	DeleteMacro(ctx context.Context, id int64, macroName string) error

	// ListContactGroups retrieves contact groups associated with a host template
	ListContactGroups(ctx context.Context, templateID int64, opts *ListOptions) (*ListResponse[ContactGroup], error)

	// SetContactGroups sets the contact groups for a host template
	SetContactGroups(ctx context.Context, templateID int64, contactGroupNames []string) error

	// AddContactGroup adds a contact group to a host template
	AddContactGroup(ctx context.Context, templateID int64, contactGroupName string) error

	// RemoveContactGroup removes a contact group from a host template
	RemoveContactGroup(ctx context.Context, templateID int64, contactGroupName string) error

	// ListContacts retrieves individual contacts associated with a host template
	ListContacts(ctx context.Context, templateID int64, opts *ListOptions) (*ListResponse[Contact], error)

	// SetContacts sets the individual contacts for a host template
	SetContacts(ctx context.Context, templateID int64, contactNames []string) error

	// AddContact adds an individual contact to a host template
	AddContact(ctx context.Context, templateID int64, contactName string) error

	// RemoveContact removes an individual contact from a host template
	RemoveContact(ctx context.Context, templateID int64, contactName string) error

	// ListHostGroups retrieves host groups associated with a host template
	ListHostGroups(ctx context.Context, templateID int64, opts *ListOptions) (*ListResponse[HostGroup], error)

	// SetHostGroups sets the host groups for a host template
	SetHostGroups(ctx context.Context, templateID int64, hostGroupNames []string) error

	// AddHostGroup adds a host group to a host template
	AddHostGroup(ctx context.Context, templateID int64, hostGroupName string) error

	// RemoveHostGroup removes a host group from a host template
	RemoveHostGroup(ctx context.Context, templateID int64, hostGroupName string) error

	// ListCategories retrieves categories associated with a host template
	ListCategories(ctx context.Context, templateID int64, opts *ListOptions) (*ListResponse[HostCategory], error)

	// SetCategories sets the categories for a host template
	SetCategories(ctx context.Context, templateID int64, categoryNames []string) error

	// AddCategory adds a category to a host template
	AddCategory(ctx context.Context, templateID int64, categoryName string) error

	// RemoveCategory removes a category from a host template
	RemoveCategory(ctx context.Context, templateID int64, categoryName string) error

	// SetSeverity sets the severity for a host template
	SetSeverity(ctx context.Context, templateID int64, severity *HostSeverity) error

	// RemoveSeverity removes the severity from a host template
	RemoveSeverity(ctx context.Context, templateID int64) error

	// Export exports host template configuration
	Export(ctx context.Context, ids []int64) ([]byte, error)

	// Import imports host template configuration
	Import(ctx context.Context, data []byte, overwrite bool) (*ImportResult, error)

	// Validate validates a host template configuration
	Validate(ctx context.Context, template *HostTemplate) (*ValidationResult, error)
}

// HostTemplateServiceInterface defines operations for HostTemplateService entities
type HostTemplateServiceInterface interface {
	// Create creates a new host template service relationship
	Create(ctx context.Context, templateID int64, service *HostTemplateService) (*HostTemplateService, error)

	// GetByID retrieves a host template service by its ID
	GetByID(ctx context.Context, id int64) (*HostTemplateService, error)

	// List retrieves host template services with optional filtering and pagination
	List(ctx context.Context, templateID int64, opts *ListOptions) (*ListResponse[HostTemplateService], error)

	// Update updates an existing host template service
	Update(ctx context.Context, id int64, service *HostTemplateService) error

	// Delete deletes a host template service by ID
	Delete(ctx context.Context, id int64) error

	// BulkCreate creates multiple host template services in a single operation
	BulkCreate(ctx context.Context, templateID int64, services []HostTemplateService) error

	// BulkDelete deletes multiple host template services by their IDs
	BulkDelete(ctx context.Context, ids []int64) error

	// SetArguments sets the arguments for a host template service
	SetArguments(ctx context.Context, id int64, arguments map[string]string) error

	// SetMacros sets the macros for a host template service
	SetMacros(ctx context.Context, id int64, macros map[string]string) error

	// Activate activates a host template service
	Activate(ctx context.Context, id int64) error

	// Deactivate deactivates a host template service
	Deactivate(ctx context.Context, id int64) error
}

// HostTemplateDiscoveryRuleInterface defines operations for HostTemplateDiscoveryRule entities
type HostTemplateDiscoveryRuleInterface interface {
	// Create creates a new host template discovery rule
	Create(ctx context.Context, templateID int64, rule *HostTemplateDiscoveryRule) (*HostTemplateDiscoveryRule, error)

	// GetByID retrieves a host template discovery rule by its ID
	GetByID(ctx context.Context, id int64) (*HostTemplateDiscoveryRule, error)

	// GetByName retrieves a host template discovery rule by its name
	GetByName(ctx context.Context, templateID int64, name string) (*HostTemplateDiscoveryRule, error)

	// List retrieves host template discovery rules with optional filtering and pagination
	List(ctx context.Context, templateID int64, opts *ListOptions) (*ListResponse[HostTemplateDiscoveryRule], error)

	// Update updates an existing host template discovery rule
	Update(ctx context.Context, id int64, rule *HostTemplateDiscoveryRule) error

	// Delete deletes a host template discovery rule by ID
	Delete(ctx context.Context, id int64) error

	// Execute manually executes a discovery rule
	Execute(ctx context.Context, id int64) (*DiscoveryResult, error)

	// GetExecutionHistory retrieves the execution history for a discovery rule
	GetExecutionHistory(ctx context.Context, id int64, opts *ListOptions) (*ListResponse[DiscoveryExecution], error)

	// SetFilters sets the filters for a discovery rule
	SetFilters(ctx context.Context, id int64, filters map[string]string) error

	// Activate activates a discovery rule
	Activate(ctx context.Context, id int64) error

	// Deactivate deactivates a discovery rule
	Deactivate(ctx context.Context, id int64) error

	// TestRule tests a discovery rule without executing it
	TestRule(ctx context.Context, id int64) (*DiscoveryTestResult, error)
}

// Supporting types for interface operations

// HostTemplateInheritanceTree represents the inheritance hierarchy of a host template
type HostTemplateInheritanceTree struct {
	Template *HostTemplate                  `json:"template"`
	Parents  []*HostTemplateInheritanceTree `json:"parents,omitempty"`
	Children []*HostTemplateInheritanceTree `json:"children,omitempty"`
	Depth    int                            `json:"depth"`
}

// DiscoveryResult represents the result of executing a discovery rule
type DiscoveryResult struct {
	RuleID           int64                      `json:"rule_id"`
	ExecutionTime    time.Time                  `json:"execution_time"`
	Duration         time.Duration              `json:"duration"`
	HostsDiscovered  int                        `json:"hosts_discovered"`
	HostsCreated     int                        `json:"hosts_created"`
	HostsUpdated     int                        `json:"hosts_updated"`
	Errors           []string                   `json:"errors,omitempty"`
	Warnings         []string                   `json:"warnings,omitempty"`
	DiscoveredHosts  []DiscoveredHost           `json:"discovered_hosts"`
	ExecutionDetails *DiscoveryExecutionDetails `json:"execution_details,omitempty"`
}

// DiscoveredHost represents a host discovered by a discovery rule
type DiscoveredHost struct {
	Name      string            `json:"name"`
	Alias     string            `json:"alias"`
	Address   string            `json:"address"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	Status    string            `json:"status"` // "created", "updated", "skipped", "error"
	Message   string            `json:"message,omitempty"`
	CreatedID *int64            `json:"created_id,omitempty"`
}

// DiscoveryExecution represents a historical execution of a discovery rule
type DiscoveryExecution struct {
	ID              int64         `json:"id"`
	RuleID          int64         `json:"rule_id"`
	StartTime       time.Time     `json:"start_time"`
	EndTime         *time.Time    `json:"end_time,omitempty"`
	Duration        time.Duration `json:"duration"`
	Status          string        `json:"status"` // "running", "completed", "failed", "cancelled"
	HostsDiscovered int           `json:"hosts_discovered"`
	HostsCreated    int           `json:"hosts_created"`
	HostsUpdated    int           `json:"hosts_updated"`
	ErrorCount      int           `json:"error_count"`
	WarningCount    int           `json:"warning_count"`
	Message         string        `json:"message,omitempty"`
}

// DiscoveryTestResult represents the result of testing a discovery rule
type DiscoveryTestResult struct {
	RuleID        int64            `json:"rule_id"`
	TestTime      time.Time        `json:"test_time"`
	Duration      time.Duration    `json:"duration"`
	HostsFound    int              `json:"hosts_found"`
	SampleHosts   []DiscoveredHost `json:"sample_hosts"`
	Errors        []string         `json:"errors,omitempty"`
	Warnings      []string         `json:"warnings,omitempty"`
	CommandOutput string           `json:"command_output,omitempty"`
	IsValid       bool             `json:"is_valid"`
}

// DiscoveryExecutionDetails provides detailed information about discovery execution
type DiscoveryExecutionDetails struct {
	CommandExecuted string            `json:"command_executed"`
	CommandArgs     string            `json:"command_args,omitempty"`
	CommandOutput   string            `json:"command_output,omitempty"`
	CommandError    string            `json:"command_error,omitempty"`
	ExitCode        int               `json:"exit_code"`
	Environment     map[string]string `json:"environment,omitempty"`
	ProcessingTime  time.Duration     `json:"processing_time"`
	FilterResults   map[string]int    `json:"filter_results,omitempty"`
}

// ImportResult represents the result of importing host template configuration
type ImportResult struct {
	Imported int      `json:"imported"`
	Updated  int      `json:"updated"`
	Skipped  int      `json:"skipped"`
	Errors   []string `json:"errors,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}

// ValidationResult represents the result of validating host template configuration
type ValidationResult struct {
	IsValid  bool     `json:"is_valid"`
	Errors   []string `json:"errors,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}
