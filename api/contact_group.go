package api

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// ContactGroup represents a contact group configuration
type ContactGroup struct {
	ID          *int64  `json:"id,omitempty"`
	Name        string  `json:"name" validate:"required"`
	Alias       string  `json:"alias" validate:"required"`
	IsActivated *bool   `json:"is_activated,omitempty"`
	Comment     string  `json:"comment,omitempty"`
	Members     []int64 `json:"members,omitempty"`

	// Relationships
	Contacts []Contact `json:"contacts,omitempty"`
}

// ContactGroupTemplate represents a contact group template
type ContactGroupTemplate struct {
	ID          *int64 `json:"id,omitempty"`
	Name        string `json:"name" validate:"required"`
	Alias       string `json:"alias" validate:"required"`
	IsActivated *bool  `json:"is_activated,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

// ACLResource represents ACL resource permissions
type ACLResource struct {
	ResourceType string   `json:"resource_type"`
	ResourceID   *int64   `json:"resource_id"`
	ResourceName string   `json:"resource_name"`
	Permissions  []string `json:"permissions"`
}

// ACLMenu represents ACL menu permissions
type ACLMenu struct {
	MenuID    *int64 `json:"menu_id"`
	MenuName  string `json:"menu_name"`
	HasAccess *bool  `json:"has_access"`
}

// ACLAction represents ACL action permissions
type ACLAction struct {
	ActionID   *int64 `json:"action_id"`
	ActionName string `json:"action_name"`
	HasAccess  *bool  `json:"has_access"`
}

// Validation functions for ContactGroup
func (cg *ContactGroup) ValidateForCreate() error {
	if cg.Name == "" || len(strings.TrimSpace(cg.Name)) == 0 {
		return fmt.Errorf("name is required for contact group creation")
	}
	if cg.Alias == "" || len(strings.TrimSpace(cg.Alias)) == 0 {
		return fmt.Errorf("alias is required for contact group creation")
	}
	return nil
}

func (cg *ContactGroup) ValidateForUpdate() error {
	if cg.ID == nil || *cg.ID <= 0 {
		return fmt.Errorf("valid contact group ID is required for update operations")
	}
	if cg.Name != "" && len(strings.TrimSpace(cg.Name)) == 0 {
		return fmt.Errorf("contact group name cannot be empty")
	}
	if cg.Alias != "" && len(strings.TrimSpace(cg.Alias)) == 0 {
		return fmt.Errorf("contact group alias cannot be empty")
	}
	return nil
}

// ContactGroupInterface defines CRUD operations for ContactGroup entities
type ContactGroupInterface interface {
	// Create creates a new contact group
	Create(ctx context.Context, group *ContactGroup) (*ContactGroup, error)

	// GetByID retrieves a contact group by its ID
	GetByID(ctx context.Context, id int64) (*ContactGroup, error)

	// GetByName retrieves a contact group by its name
	GetByName(ctx context.Context, name string) (*ContactGroup, error)

	// List retrieves contact groups with optional filtering and pagination
	List(ctx context.Context, opts *ListOptions) (*ListResponse[ContactGroup], error)

	// Update updates an existing contact group
	Update(ctx context.Context, id int64, group *ContactGroup) error

	// Delete deletes a contact group by ID
	Delete(ctx context.Context, id int64) error

	// BulkCreate creates multiple contact groups in a single operation
	BulkCreate(ctx context.Context, groups []ContactGroup) error

	// BulkUpdate updates multiple contact groups in a single operation
	BulkUpdate(ctx context.Context, groups []ContactGroup) error

	// BulkDelete deletes multiple contact groups by their IDs
	BulkDelete(ctx context.Context, ids []int64) error

	// ListMembers retrieves contacts that are members of the contact group
	ListMembers(ctx context.Context, groupID int64, opts *ListOptions) (*ListResponse[Contact], error)

	// AddMember adds a contact to the contact group
	AddMember(ctx context.Context, groupID int64, contactID int64) error

	// RemoveMember removes a contact from the contact group
	RemoveMember(ctx context.Context, groupID int64, contactID int64) error

	// SetMembers sets the complete list of members for the contact group
	SetMembers(ctx context.Context, groupID int64, contactIDs []int64) error

	// BulkAddMembers adds multiple contacts to the contact group
	BulkAddMembers(ctx context.Context, groupID int64, contactIDs []int64) error

	// BulkRemoveMembers removes multiple contacts from the contact group
	BulkRemoveMembers(ctx context.Context, groupID int64, contactIDs []int64) error

	// ListHosts retrieves hosts that use this contact group for notifications
	ListHosts(ctx context.Context, groupID int64, opts *ListOptions) (*ListResponse[Host], error)

	// ListServices retrieves services that use this contact group for notifications
	ListServices(ctx context.Context, groupID int64, opts *ListOptions) (*ListResponse[Service], error)

	// ListHostTemplates retrieves host templates that use this contact group
	ListHostTemplates(ctx context.Context, groupID int64, opts *ListOptions) (*ListResponse[HostTemplate], error)

	// ListServiceTemplates retrieves service templates that use this contact group
	ListServiceTemplates(ctx context.Context, groupID int64, opts *ListOptions) (*ListResponse[ServiceTemplate], error)

	// GetUsageStatistics retrieves statistics about contact group usage
	GetUsageStatistics(ctx context.Context, groupID int64) (*ContactGroupUsage, error)

	// Activate activates a contact group
	Activate(ctx context.Context, id int64) error

	// Deactivate deactivates a contact group
	Deactivate(ctx context.Context, id int64) error

	// Duplicate creates a copy of an existing contact group with a new name
	Duplicate(ctx context.Context, id int64, newName, newAlias string) (*ContactGroup, error)

	// Export exports contact group configuration
	Export(ctx context.Context, ids []int64) ([]byte, error)

	// Import imports contact group configuration
	Import(ctx context.Context, data []byte, overwrite bool) (*ImportResult, error)

	// Validate validates a contact group configuration
	Validate(ctx context.Context, group *ContactGroup) (*ValidationResult, error)

	// GetACLResources retrieves ACL resources associated with the contact group
	GetACLResources(ctx context.Context, groupID int64, opts *ListOptions) (*ListResponse[ACLResource], error)

	// SetACLResources sets ACL resources for the contact group
	SetACLResources(ctx context.Context, groupID int64, resources []ACLResource) error

	// GetACLMenus retrieves ACL menu permissions for the contact group
	GetACLMenus(ctx context.Context, groupID int64, opts *ListOptions) (*ListResponse[ACLMenu], error)

	// SetACLMenus sets ACL menu permissions for the contact group
	SetACLMenus(ctx context.Context, groupID int64, menus []ACLMenu) error

	// GetACLActions retrieves ACL action permissions for the contact group
	GetACLActions(ctx context.Context, groupID int64, opts *ListOptions) (*ListResponse[ACLAction], error)

	// SetACLActions sets ACL action permissions for the contact group
	SetACLActions(ctx context.Context, groupID int64, actions []ACLAction) error
}

// ContactGroupTemplateInterface defines CRUD operations for ContactGroupTemplate entities
type ContactGroupTemplateInterface interface {
	// Create creates a new contact group template
	Create(ctx context.Context, template *ContactGroupTemplate) (*ContactGroupTemplate, error)

	// GetByID retrieves a contact group template by its ID
	GetByID(ctx context.Context, id int64) (*ContactGroupTemplate, error)

	// GetByName retrieves a contact group template by its name
	GetByName(ctx context.Context, name string) (*ContactGroupTemplate, error)

	// List retrieves contact group templates with optional filtering and pagination
	List(ctx context.Context, opts *ListOptions) (*ListResponse[ContactGroupTemplate], error)

	// Update updates an existing contact group template
	Update(ctx context.Context, id int64, template *ContactGroupTemplate) error

	// Delete deletes a contact group template by ID
	Delete(ctx context.Context, id int64) error

	// Duplicate creates a copy of an existing template with a new name
	Duplicate(ctx context.Context, id int64, newName, newAlias string) (*ContactGroupTemplate, error)

	// Apply applies a template to create a new contact group
	Apply(ctx context.Context, templateID int64, name, alias string) (*ContactGroup, error)

	// ListUsages retrieves contact groups that were created from this template
	ListUsages(ctx context.Context, templateID int64, opts *ListOptions) (*ListResponse[ContactGroup], error)
}

// Supporting types for ContactGroup interface operations

// ContactGroupUsage represents usage statistics for a contact group
type ContactGroupUsage struct {
	GroupID               int64      `json:"group_id"`
	GroupName             string     `json:"group_name"`
	TotalMembers          int        `json:"total_members"`
	ActiveMembers         int        `json:"active_members"`
	HostsUsingGroup       int        `json:"hosts_using_group"`
	ServicesUsingGroup    int        `json:"services_using_group"`
	HostTemplatesUsing    int        `json:"host_templates_using"`
	ServiceTemplatesUsing int        `json:"service_templates_using"`
	LastNotificationSent  *time.Time `json:"last_notification_sent,omitempty"`
	NotificationsSent24h  int        `json:"notifications_sent_24h"`
	NotificationsSent7d   int        `json:"notifications_sent_7d"`
	NotificationsSent30d  int        `json:"notifications_sent_30d"`
}

// ServiceTemplate represents a service template (referenced in interface)
type ServiceTemplate struct {
	ID          *int64 `json:"id,omitempty"`
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	IsActivated *bool  `json:"is_activated,omitempty"`
}

// ValidationResult represents the result of validating contact group configuration
// (reused from host_template.go but defined here for completeness)
type ContactGroupValidationResult struct {
	IsValid  bool     `json:"is_valid"`
	Errors   []string `json:"errors,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}

// ImportResult represents the result of importing contact group configuration
// (reused from host_template.go but defined here for completeness)
type ContactGroupImportResult struct {
	Imported int      `json:"imported"`
	Updated  int      `json:"updated"`
	Skipped  int      `json:"skipped"`
	Errors   []string `json:"errors,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}
