package api

import "context"

// HostCategory represents a logical grouping mechanism for hosts in Centreon.
// Categories help organize hosts by type, location, or business function for easier management
// and reporting. Hosts can belong to multiple categories for flexible classification.
type HostCategory struct {
	// ID is the unique identifier for the host category
	ID *int64 `json:"id,omitempty"`

	// Name is the unique name identifier for the category
	// Used for programmatic references and API operations
	Name string `json:"name" validate:"required"`

	// Alias is the human-readable display name for the category
	// Used in the web interface and reports for better readability
	Alias string `json:"alias" validate:"required"`

	// Level represents the hierarchical level of this category
	// Used for creating category hierarchies and nested organizational structures
	Level *int `json:"level,omitempty"`

	// IconID references an icon to visually represent this category
	// Provides visual identification in the web interface and dashboards
	IconID *int `json:"icon_id,omitempty"`

	// IsActivated controls whether this category is available for use
	// Inactive categories are not displayed in selection lists
	IsActivated *bool `json:"is_activated,omitempty"`
}

// HostCategoryInterface defines CRUD operations for Host Category entities
type HostCategoryInterface interface {
	// Create creates a new host category
	Create(ctx context.Context, category *HostCategory) (*HostCategory, error)

	// GetByID retrieves a host category by its ID
	GetByID(ctx context.Context, id int64) (*HostCategory, error)

	// GetByName retrieves a host category by its name
	GetByName(ctx context.Context, name string) (*HostCategory, error)

	// List retrieves host categories with optional filtering and pagination
	List(ctx context.Context, opts *ListOptions) (*ListResponse[HostCategory], error)

	// Update updates an existing host category
	Update(ctx context.Context, id int64, category *HostCategory) error

	// Delete deletes a host category by ID
	Delete(ctx context.Context, id int64) error
}
