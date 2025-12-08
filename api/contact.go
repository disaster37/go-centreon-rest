package api

import "context"

// Contact represents a user contact configuration in Centreon monitoring system.
// Contacts are individuals who receive notifications about monitoring events and can
// interact with the monitoring system through acknowledgements, downtimes, and administration.
type Contact struct {
	// ID is the unique identifier for the contact in the Centreon database
	ID *int64 `json:"id,omitempty"`

	// Name is the unique username for login and identification purposes
	// Used for authentication and as a reference in monitoring configurations
	Name string `json:"name" validate:"required,min=1,max=200"`

	// Alias is the human-readable display name for the contact
	// Used in the web interface and notifications for better user identification
	Alias string `json:"alias" validate:"required,min=1,max=255"`

	// Email is the contact's email address for notification delivery
	// Must be a valid email format and is used for email-based notifications
	Email string `json:"email,omitempty" validate:"omitempty,email,max=255"`

	// Password is the hashed password for local authentication
	// Minimum 8 characters required for security; stored as hash in database
	Password string `json:"password,omitempty" validate:"omitempty,min=8,max=255"`

	// Pager is the contact's pager number for SMS/pager notifications
	// Used for critical alerts that require immediate attention
	Pager string `json:"pager,omitempty" validate:"omitempty,max=255"`

	// Address1-6 are flexible contact information fields
	// Can store phone numbers, alternate emails, office locations, or other contact methods
	Address1 string `json:"address1,omitempty" validate:"omitempty,max=255"`
	Address2 string `json:"address2,omitempty" validate:"omitempty,max=255"`
	Address3 string `json:"address3,omitempty" validate:"omitempty,max=255"`
	Address4 string `json:"address4,omitempty" validate:"omitempty,max=255"`
	Address5 string `json:"address5,omitempty" validate:"omitempty,max=255"`
	Address6 string `json:"address6,omitempty" validate:"omitempty,max=255"`

	// IsActivated controls whether this contact account is active
	// Inactive contacts cannot log in and will not receive notifications
	IsActivated *bool `json:"is_activated,omitempty"`

	// IsAdmin grants administrative privileges to the contact
	// Admin contacts can modify system configuration and access all monitoring data
	IsAdmin *bool `json:"is_admin,omitempty"`

	// IsTemplate indicates whether this contact is used as a template
	// Template contacts provide default values for creating new contacts
	IsTemplate *bool `json:"is_template,omitempty"`

	// HostNotificationPeriod defines when this contact receives host notifications
	// References a time period configuration (e.g., "24x7", "workhours", "oncall")
	HostNotificationPeriod string `json:"host_notification_period,omitempty" validate:"omitempty,max=200"`

	// ServiceNotificationPeriod defines when this contact receives service notifications
	// References a time period configuration separate from host notifications
	ServiceNotificationPeriod string `json:"service_notification_period,omitempty" validate:"omitempty,max=200"`

	// HostNotificationOptions specifies which host state changes trigger notifications
	// Comma-separated values: d(own), u(nreachable), r(ecovery), f(lapping), s(cheduled downtime)
	HostNotificationOptions string `json:"host_notification_options,omitempty" validate:"omitempty,max=255"`

	// ServiceNotificationOptions specifies which service state changes trigger notifications
	// Comma-separated values: w(arning), c(ritical), u(nknown), r(ecovery), f(lapping), s(cheduled downtime)
	ServiceNotificationOptions string `json:"service_notification_options,omitempty" validate:"omitempty,max=255"`

	// HostNotificationCommands lists the notification methods for host alerts
	// Each command defines how notifications are delivered (email, SMS, etc.)
	HostNotificationCommands []string `json:"host_notification_commands,omitempty"`

	// ServiceNotificationCommands lists the notification methods for service alerts
	// Can be different from host commands to customize notification delivery per event type
	ServiceNotificationCommands []string `json:"service_notification_commands,omitempty"`

	// Language sets the preferred language for the contact's interface
	// Used for localization of the web interface and notifications (e.g., "en", "fr", "es")
	Language string `json:"language,omitempty" validate:"omitempty,max=10"`

	// AuthType specifies the authentication method for this contact
	// "local" uses Centreon's built-in authentication, "ldap" uses LDAP/Active Directory
	AuthType string `json:"auth_type,omitempty" validate:"omitempty,oneof=local ldap"`

	// LdapDN is the LDAP Distinguished Name for LDAP-authenticated contacts
	// Used to map the contact to their LDAP directory entry for authentication
	LdapDN string `json:"ldap_dn,omitempty" validate:"omitempty,max=512"`

	// Location stores the physical location or office of the contact
	// Used for geographical context and coordination of on-site activities
	Location string `json:"location,omitempty" validate:"omitempty,max=255"`

	// Timezone specifies the contact's local timezone for proper time display
	// Used for scheduling and displaying times in the contact's local context
	Timezone string `json:"timezone,omitempty" validate:"omitempty,max=100"`

	// IsLocked prevents modification of this contact configuration
	// Used to protect critical system contacts from accidental changes
	IsLocked *bool `json:"is_locked,omitempty"`
}

// ContactTemplate represents a reusable template for contact configurations.
// Templates provide default notification settings and contact information patterns,
// reducing configuration effort and ensuring consistency across similar contacts.
type ContactTemplate struct {
	// ID is the unique identifier for the contact template
	ID *int64 `json:"id,omitempty"`

	// Name is the unique template name used for identification
	// Used when applying templates to new contacts
	Name string `json:"name" validate:"required,min=1,max=200"`

	// Alias is the human-readable display name for the template
	// Provides a descriptive name for template selection and identification
	Alias string `json:"alias" validate:"required,min=1,max=255"`

	// IsActivated controls whether this template is available for use
	// Inactive templates cannot be applied to new contacts
	IsActivated *bool `json:"is_activated,omitempty"`

	// HostNotificationPeriod defines the default host notification schedule
	// New contacts using this template will inherit this notification period
	HostNotificationPeriod string `json:"host_notification_period,omitempty" validate:"omitempty,max=200"`

	// ServiceNotificationPeriod defines the default service notification schedule
	// Can be different from host notifications to provide granular control
	ServiceNotificationPeriod string `json:"service_notification_period,omitempty" validate:"omitempty,max=200"`

	// HostNotificationOptions specifies default host state change notifications
	// Template defines which host events (down, unreachable, recovery, etc.) trigger alerts
	HostNotificationOptions string `json:"host_notification_options,omitempty" validate:"omitempty,max=255"`

	// ServiceNotificationOptions specifies default service state change notifications
	// Template defines which service events (warning, critical, recovery, etc.) trigger alerts
	ServiceNotificationOptions string `json:"service_notification_options,omitempty" validate:"omitempty,max=255"`

	// HostNotificationCommands lists default notification methods for host alerts
	// New contacts inherit these delivery methods (email, SMS, etc.) for host notifications
	HostNotificationCommands []string `json:"host_notification_commands,omitempty"`

	// ServiceNotificationCommands lists default notification methods for service alerts
	// Can differ from host commands to customize delivery per event type
	ServiceNotificationCommands []string `json:"service_notification_commands,omitempty"`
}

// Permission represents a system permission that can be granted to contacts or groups.
// Permissions control access to specific Centreon features, data, and administrative functions,
// implementing role-based access control (RBAC) for security and operational governance.
type Permission struct {
	// ID is the unique identifier for the permission
	ID *int64 `json:"id,omitempty"`

	// Name is the unique permission identifier used in access control checks
	// Should be descriptive and follow naming conventions (e.g., "hosts.read", "services.write")
	Name string `json:"name" validate:"required,min=1,max=200"`

	// Description provides detailed explanation of what the permission allows
	// Helps administrators understand the scope and impact of granting this permission
	Description string `json:"description,omitempty" validate:"omitempty,max=65535"`

	// IsGranted indicates whether this permission is currently granted
	// Used in permission evaluation and access control decisions
	IsGranted *bool `json:"is_granted,omitempty"`
}

// AccessGroup represents a collection of permissions that can be assigned to contacts.
// Access groups provide a convenient way to manage permissions for multiple users with similar roles,
// implementing role-based access control patterns for easier administration.
type AccessGroup struct {
	// ID is the unique identifier for the access group
	ID *int64 `json:"id,omitempty"`

	// Name is the unique name for the access group
	// Used for identification and assignment to contacts
	Name string `json:"name" validate:"required,min=1,max=200"`

	// Description explains the purpose and scope of the access group
	// Helps administrators understand which roles or functions this group supports
	Description string `json:"description,omitempty" validate:"omitempty,max=65535"`

	// IsActivated controls whether this access group is available for assignment
	// Inactive groups cannot be assigned to new contacts but existing assignments remain
	IsActivated *bool `json:"is_activated,omitempty"`
}

// NotificationCommand represents a command used to deliver notifications to contacts.
// These commands define how alerts are sent (email, SMS, webhooks, etc.) and can be
// customized with specific parameters, recipients, and formatting options.
type NotificationCommand struct {
	// ID is the unique identifier for the notification command
	ID *int64 `json:"id,omitempty"`

	// Name is the unique command name used in contact configurations
	// Referenced by contacts to specify their preferred notification methods
	Name string `json:"name" validate:"required,min=1,max=200"`

	// CommandLine is the actual command or script executed to send notifications
	// Can include macros for dynamic content like contact info, alert details, etc.
	CommandLine string `json:"command_line" validate:"required,min=1,max=8192"`

	// Type specifies whether this command is used for host or service notifications
	// Some commands may be specialized for different types of monitoring events
	Type string `json:"type,omitempty" validate:"omitempty,oneof=host service"`
}

// ContactInterface defines CRUD operations for Contact entities
type ContactInterface interface {
	// Create creates a new contact
	Create(ctx context.Context, contact *Contact) (*Contact, error)

	// GetByID retrieves a contact by its ID
	GetByID(ctx context.Context, id int64) (*Contact, error)

	// List retrieves contacts with optional filtering and pagination
	List(ctx context.Context, opts *ListOptions) (*ListResponse[Contact], error)

	// Update updates an existing contact
	Update(ctx context.Context, id int64, contact *Contact) error

	// Delete deletes a contact by ID
	Delete(ctx context.Context, id int64) error
}

// ContactTemplateInterface defines CRUD operations for Contact Template entities
type ContactTemplateInterface interface {
	// List retrieves contact templates
	List(ctx context.Context, opts *ListOptions) (*ListResponse[ContactTemplate], error)
}

// AccessGroupInterface defines operations for access groups
type AccessGroupInterface interface {
	// List retrieves access groups
	List(ctx context.Context, opts *ListOptions) (*ListResponse[AccessGroup], error)
}
