package api

import "time"

// ResourceStatus represents the current operational status of a monitored resource.
// This structure provides both numeric codes and human-readable names for resource states,
// along with severity information for prioritization and alerting purposes.
type ResourceStatus struct {
	// Code is the numeric status code (0=OK/UP, 1=WARNING/DOWN, 2=CRITICAL/UNREACHABLE, 3=UNKNOWN)
	// Standardized across Centreon for programmatic status evaluation
	Code int `json:"code"`

	// Name is the human-readable status description (e.g., "OK", "WARNING", "CRITICAL", "UNKNOWN")
	// Used for display in interfaces and notifications
	Name string `json:"name"`

	// SeverityCode represents the business impact level of this status
	// Higher values typically indicate more severe conditions requiring immediate attention
	SeverityCode int `json:"severity_code"`
}

// ResourceLinks represents navigation and action links associated with a monitoring resource.
// These links provide direct access to related functionality, external systems, and detailed
// information about the resource from within the Centreon interface.
type ResourceLinks struct {
	// URIs contains web interface links for resource management and viewing
	URIs ResourceURIs `json:"uris"`

	// Endpoints contains API endpoints for programmatic resource interaction
	Endpoints ResourceEndpoints `json:"endpoints"`

	// Externals contains links to external systems and documentation
	Externals ResourceExternals `json:"externals"`
}

// ResourceURIs represents web interface navigation links for a resource.
// These URLs provide direct access to specific pages and views related to
// the resource within the Centreon web interface.
type ResourceURIs struct {
	// Configuration is the URL to the resource configuration page
	// Allows direct navigation to modify resource settings
	Configuration *string `json:"configuration"`

	// Logs is the URL to view resource-related log entries
	// Provides access to historical events and troubleshooting information
	Logs *string `json:"logs"`

	// Reporting is the URL to resource-specific reports and analytics
	// Enables access to performance trends, availability statistics, and SLA data
	Reporting *string `json:"reporting"`
}

// ResourceEndpoints represents REST API endpoints for resource operations.
// These endpoints enable programmatic interaction with resource data and actions,
// supporting automation, integration, and custom monitoring applications.
type ResourceEndpoints struct {
	// Details provides the API endpoint for retrieving detailed resource information
	Details *string `json:"details"`

	// Timeline provides the API endpoint for resource event history
	Timeline *string `json:"timeline"`

	// StatusGraph provides the API endpoint for status trend data
	StatusGraph *string `json:"status_graph"`

	// PerformanceGraph provides the API endpoint for performance metrics data
	PerformanceGraph *string `json:"performance_graph"`

	// Acknowledgement provides the API endpoint for managing resource acknowledgements
	Acknowledgement *string `json:"acknowledgement"`

	// Downtime provides the API endpoint for managing scheduled downtimes
	Downtime *string `json:"downtime"`

	// NotificationPolicy provides the API endpoint for notification configuration
	NotificationPolicy *string `json:"notification_policy"`

	// Check provides the API endpoint for triggering manual checks
	Check *string `json:"check"`

	// ForcedCheck provides the API endpoint for forcing immediate checks
	ForcedCheck *string `json:"forced_check"`

	// Metrics provides the API endpoint for retrieving performance metrics
	Metrics *string `json:"metrics"`
}

// ResourceExternals represents links to external systems and documentation.
// These links integrate Centreon with external tools, documentation systems,
// and operational procedures related to the monitored resource.
type ResourceExternals struct {
	// ActionURL is a link to external actions or tools related to this resource
	// Often points to runbooks, escalation procedures, or management interfaces
	ActionURL *string `json:"action_url"`

	// Notes contains links to external documentation or notes
	// Provides access to operational procedures, troubleshooting guides, or vendor documentation
	Notes *Notes `json:"notes"`
}

// Notes represents external documentation or reference links for a resource.
// This structure provides labeled links to external information sources that help
// operators understand and manage the monitored resource effectively.
type Notes struct {
	// URL is the link to external documentation or notes
	URL *string `json:"url"`

	// Label is the display text for the notes link
	// Provides context about what information the link contains
	Label *string `json:"label"`
}

// ResourceMain represents comprehensive information about a monitored resource in Centreon.
// This structure contains all essential data for displaying and managing resources in monitoring
// interfaces, including status, relationships, and operational metadata.
type ResourceMain struct {
	// UUID is the universally unique identifier for this resource
	// Provides a consistent reference across distributed monitoring environments
	UUID string `json:"uuid"`

	// Type specifies the resource type ("host" or "service")
	// Determines how the resource is processed and displayed
	Type string `json:"type"`

	// ShortType is an abbreviated version of the resource type
	// Used in compact displays and API responses for efficiency
	ShortType string `json:"short_type"`

	// ID is the unique numeric identifier for this resource
	ID int32 `json:"id"`

	// HostID is the identifier of the host this resource belongs to
	// For services, links to the parent host; for hosts, equals ID
	HostID int32 `json:"host_id"`

	// ServiceID is the service identifier when the resource is a service
	// Null for host resources, populated for service resources
	ServiceID *int32 `json:"service_id"`

	// Name is the primary display name for the resource
	// Host name for hosts, service description for services
	Name string `json:"name"`

	// Alias is an alternative display name for the resource
	// Provides human-readable identification when Name is technical
	Alias *string `json:"alias"`

	// FQDN is the fully qualified domain name for host resources
	// Used for network identification and DNS resolution
	FQDN *string `json:"fqdn"`

	// Links contains navigation and action URLs for this resource
	Links ResourceLinks `json:"links"`

	// MonitoringServerName identifies which monitoring server manages this resource
	// Important in distributed monitoring environments with multiple pollers
	MonitoringServerName string `json:"monitoring_server_name"`

	// Icon provides visual representation for the resource
	Icon *ResourceIcon `json:"icon"`

	// Parent contains information about the parent resource (for services)
	Parent *ResourceParent `json:"parent"`

	// Status contains current operational status information
	Status ResourceStatus `json:"status"`

	// IsInDowntime indicates if the resource is currently in scheduled maintenance
	// When true, notifications are typically suppressed
	IsInDowntime bool `json:"is_in_downtime"`

	// IsAcknowledged indicates if current problems have been acknowledged
	// Helps track issue ownership and response status
	IsAcknowledged bool `json:"is_acknowledged"`

	// Duration represents how long the resource has been in its current state
	// Useful for understanding problem persistence and stability
	Duration *string `json:"duration"`

	// LastStatusChange is when the resource last changed status
	LastStatusChange *time.Time `json:"last_status_change"`

	// LastTimeWithNoIssue is when the resource was last in a healthy state
	// Used for availability calculations and SLA reporting
	LastTimeWithNoIssue *time.Time `json:"last_time_with_no_issue"`

	// Tries indicates current/max check attempts (e.g., "3/3")
	// Shows retry attempts for failed checks
	Tries *string `json:"tries"`

	// LastCheck is a formatted string of when the resource was last checked
	LastCheck *string `json:"last_check"`

	// Information contains the latest check output or status message
	// Provides detailed status information and diagnostic data
	Information *string `json:"information"`

	// HasActiveChecksEnabled indicates if active monitoring is configured
	HasActiveChecksEnabled *bool `json:"has_active_checks_enabled"`

	// HasPassiveChecksEnabled indicates if passive check submission is allowed
	HasPassiveChecksEnabled *bool `json:"has_passive_checks_enabled"`

	// PerformanceData contains metrics from the latest check
	// Used for graphing, trending, and capacity planning
	PerformanceData *string `json:"performance_data"`

	// IsNotificationEnabled indicates if notifications are active for this resource
	IsNotificationEnabled bool `json:"is_notification_enabled"`

	// Severity contains business criticality information for the resource
	Severity *ResourceSeverity `json:"severity"`

	// Extra contains additional resource-specific data and metadata
	// Extensible field for custom attributes and specialized information
	Extra []any `json:"extra"`
}

// ResourceIcon represents visual iconography for monitored resources.
// Icons provide immediate visual identification and categorization of resources
// in monitoring dashboards, lists, and status displays.
type ResourceIcon struct {
	// ID is the unique identifier for the icon
	ID int32 `json:"id"`

	// Name is the descriptive name of the icon
	// Used for icon selection and identification
	Name string `json:"name"`

	// URL is the path or URL to the icon image file
	// Used to display the icon in web interfaces
	URL string `json:"url"`
}

// ResourceParent represents information about a parent resource in the monitoring hierarchy.
// For services, this contains details about the host they run on. This relationship
// is crucial for understanding dependencies and organizing monitoring data.
type ResourceParent struct {
	// Type specifies the parent resource type (typically "host" for services)
	Type string `json:"type"`

	// ShortType is an abbreviated version of the parent resource type
	ShortType string `json:"short_type"`

	// ID is the unique identifier of the parent resource
	ID int64 `json:"id"`

	// Name is the display name of the parent resource
	Name string `json:"name"`

	// Alias is an alternative display name for the parent
	Alias *string `json:"alias"`

	// FQDN is the fully qualified domain name of the parent host
	FQDN string `json:"fqdn"`

	// Links contains navigation and action URLs for the parent resource
	Links ResourceLinks `json:"links"`

	// Status contains the current operational status of the parent
	Status ResourceStatus `json:"status"`
}

// ResourceSeverity represents business criticality and impact classification for resources.
// Severity levels help prioritize incident response, define SLA requirements,
// and organize monitoring focus based on business importance.
type ResourceSeverity struct {
	// Type categorizes the severity classification system being used
	Type string `json:"type"`

	// ID is the unique identifier for this severity level
	ID int `json:"id"`

	// Name is the descriptive name for the severity (e.g., "Critical", "High", "Medium", "Low")
	Name string `json:"name"`

	// Icon provides visual representation for the severity level
	Icon ResourceSeverityIcon `json:"icon"`
}

// ResourceSeverityIcon represents visual indicators for resource severity levels.
// These icons provide immediate visual cues about the business criticality
// and required response urgency for monitored resources.
type ResourceSeverityIcon struct {
	// ID is the unique identifier for the severity icon
	ID int `json:"id"`

	// Name is the descriptive name of the severity icon
	Name string `json:"name"`

	// URL is the path or URL to the severity icon image file
	URL string `json:"url"`
}

// TimelineEvent represents an event in the monitoring timeline
type TimelineEvent struct {
	ID        int            `json:"id"`
	Type      string         `json:"type"`
	Date      time.Time      `json:"date"`
	StartDate *time.Time     `json:"start_date"`
	EndDate   *time.Time     `json:"end_date"`
	Content   string         `json:"content"`
	Contact   *EventContact  `json:"contact"`
	Status    ResourceStatus `json:"status"`
	Tries     *int           `json:"tries"`
}

// EventContact represents contact information for an event
type EventContact struct {
	ID   *int   `json:"id"`
	Name string `json:"name"`
}

// SubmitResultHost represents host result submission data
type SubmitResultHost struct {
	Status          int     `json:"status"`
	Output          *string `json:"output"`
	PerformanceData *string `json:"performance_data"`
}

// SubmitResultService represents service result submission data
type SubmitResultService struct {
	Status          string  `json:"status"`
	Output          *string `json:"output"`
	PerformanceData *string `json:"performance_data"`
}

// SubmitResultResources represents bulk resource result submission
type SubmitResultResources struct {
	Resources []SubmitResourceItem `json:"resources"`
}

// SubmitResourceItem represents a single resource in bulk submission
type SubmitResourceItem struct {
	Type            string                `json:"type,omitempty"`
	ID              int                   `json:"id"`
	Parent          *SubmitResourceParent `json:"parent,omitempty"`
	Status          int                   `json:"status"`
	Output          *string               `json:"output,omitempty"`
	PerformanceData *string               `json:"performance_data,omitempty"`
}

// SubmitResourceParent represents parent resource in bulk submission
type SubmitResourceParent struct {
	ID int `json:"id"`
}

// ResourceAction represents an action that can be performed on resources
type ResourceAction struct {
	Type   string                `json:"type"`
	ID     int32                 `json:"id"`
	Parent *ResourceActionParent `json:"parent,omitempty"`
}

// ResourceActionParent represents parent for resource actions
type ResourceActionParent struct {
	ID int64 `json:"id"`
}

// BulkResourceID represents resource identification for bulk operations
type BulkResourceID struct {
	ResourceID       int64  `json:"resource_id"`
	ParentResourceID *int64 `json:"parent_resource_id"`
}
