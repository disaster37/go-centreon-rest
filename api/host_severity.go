package api

import "context"

// HostSeverity represents the business criticality classification for hosts.
// Severity levels help prioritize incident response and define SLA requirements.
// Higher severity hosts typically require faster response times and more attention.
type HostSeverity struct {
	// ID is the unique identifier for the host severity level
	ID *int64 `json:"id,omitempty"`

	// Name is the descriptive name for the severity level (e.g., "Critical", "High", "Medium", "Low")
	Name string `json:"name" validate:"required"`

	// Level is the numeric priority level for this severity
	// Lower numbers typically indicate higher priority/more critical systems
	Level *int64 `json:"level,omitempty"`

	// Type categorizes the kind of severity classification
	// Used to distinguish between different severity frameworks or standards
	Type string `json:"type,omitempty"`

	// Icon provides visual representation for the severity level
	// Helps users quickly identify severity in dashboards and reports
	Icon *SeverityIcon `json:"icon,omitempty"`
}

// SeverityIcon represents visual iconography for severity levels.
// Icons provide immediate visual cues about the criticality of monitored resources.
type SeverityIcon struct {
	// Name is the identifier for the icon
	Name string `json:"name"`

	// URL is the path or URL to the icon image file
	URL string `json:"url"`
}

// HostSeverityInterface defines CRUD operations for Host Severity entities
type HostSeverityInterface interface {
	// Create creates a new host severity
	Create(ctx context.Context, severity *HostSeverity) (*HostSeverity, error)

	// GetByID retrieves a host severity by its ID
	GetByID(ctx context.Context, id int64) (*HostSeverity, error)

	// List retrieves host severities with optional filtering and pagination
	List(ctx context.Context, opts *ListOptions) (*ListResponse[HostSeverity], error)

	// Update updates an existing host severity
	Update(ctx context.Context, id int64, severity *HostSeverity) error

	// Delete deletes a host severity by ID
	Delete(ctx context.Context, id int64) error
}
