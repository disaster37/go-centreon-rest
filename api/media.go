package api

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// Media represents a media file configuration
type Media struct {
	ID          *int64    `json:"id,omitempty"`
	Name        string    `json:"name" validate:"required"`
	Filename    string    `json:"filename,omitempty"`
	Directory   string    `json:"directory,omitempty"`
	Comment     string    `json:"comment,omitempty"`
	Size        *int64    `json:"size,omitempty"`
	ContentType string    `json:"content_type,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	URL         string    `json:"url,omitempty"`
	FileContent []byte    `json:"file_content,omitempty"`
}

// MediaFolder represents a media folder configuration
type MediaFolder struct {
	ID         *int64        `json:"id,omitempty"`
	Name       string        `json:"name" validate:"required"`
	Path       string        `json:"path" validate:"required"`
	ParentID   *int64        `json:"parent_id,omitempty"`
	Comment    string        `json:"comment,omitempty"`
	CreatedAt  time.Time     `json:"created_at,omitempty"`
	UpdatedAt  time.Time     `json:"updated_at,omitempty"`
	MediaCount *int          `json:"media_count,omitempty"`
	SubFolders []MediaFolder `json:"sub_folders,omitempty"`
}

// Validation functions for Media
func (m *Media) ValidateForCreate() error {
	if m.Name == "" || len(strings.TrimSpace(m.Name)) == 0 {
		return fmt.Errorf("name is required for media creation")
	}
	if len(m.FileContent) == 0 {
		return fmt.Errorf("file_content is required for media creation")
	}
	return nil
}

func (m *Media) ValidateForUpdate() error {
	if m.ID == nil || *m.ID <= 0 {
		return fmt.Errorf("valid media ID is required for update operations")
	}
	if m.Name != "" && len(strings.TrimSpace(m.Name)) == 0 {
		return fmt.Errorf("media name cannot be empty")
	}
	return nil
}

// Validation functions for MediaFolder
func (mf *MediaFolder) ValidateForCreate() error {
	if mf.Name == "" || len(strings.TrimSpace(mf.Name)) == 0 {
		return fmt.Errorf("name is required for media folder creation")
	}
	if mf.Path == "" || len(strings.TrimSpace(mf.Path)) == 0 {
		return fmt.Errorf("path is required for media folder creation")
	}
	return nil
}

// MediaInterface defines CRUD operations for Media entities
type MediaInterface interface {
	// Create creates a new media file
	Create(ctx context.Context, media *Media) (*Media, error)

	// GetByID retrieves a media file by its ID
	GetByID(ctx context.Context, id int64) (*Media, error)

	// GetByName retrieves a media file by its name
	GetByName(ctx context.Context, name string) (*Media, error)

	// List retrieves media files with optional filtering and pagination
	List(ctx context.Context, opts *ListOptions) (*ListResponse[Media], error)

	// Update updates an existing media file
	Update(ctx context.Context, id int64, media *Media) error

	// Delete deletes a media file by ID
	Delete(ctx context.Context, id int64) error

	// Upload uploads a new media file with content
	Upload(ctx context.Context, media *Media) (*Media, error)

	// Download retrieves the content of a media file
	Download(ctx context.Context, id int64) ([]byte, error)

	// ListByFolder retrieves media files within a specific folder
	ListByFolder(ctx context.Context, folderID int64, opts *ListOptions) (*ListResponse[Media], error)

	// Move moves a media file to a different folder
	Move(ctx context.Context, mediaID, targetFolderID int64) error
}

// MediaFolderInterface defines CRUD operations for MediaFolder entities
type MediaFolderInterface interface {
	// Create creates a new media folder
	Create(ctx context.Context, folder *MediaFolder) (*MediaFolder, error)

	// GetByID retrieves a media folder by its ID
	GetByID(ctx context.Context, id int64) (*MediaFolder, error)

	// GetByPath retrieves a media folder by its path
	GetByPath(ctx context.Context, path string) (*MediaFolder, error)

	// List retrieves media folders with optional filtering and pagination
	List(ctx context.Context, opts *ListOptions) (*ListResponse[MediaFolder], error)

	// Update updates an existing media folder
	Update(ctx context.Context, id int64, folder *MediaFolder) error

	// Delete deletes a media folder by ID (must be empty)
	Delete(ctx context.Context, id int64) error

	// ListSubFolders retrieves subfolders of a parent folder
	ListSubFolders(ctx context.Context, parentID int64, opts *ListOptions) (*ListResponse[MediaFolder], error)

	// Move moves a folder to a different parent folder
	Move(ctx context.Context, folderID, newParentID int64) error

	// GetMediaCount returns the number of media files in a folder
	GetMediaCount(ctx context.Context, folderID int64) (int, error)
}

func (mf *MediaFolder) ValidateForUpdate() error {
	if mf.ID == nil || *mf.ID <= 0 {
		return fmt.Errorf("valid media folder ID is required for update operations")
	}
	if mf.Name != "" && len(strings.TrimSpace(mf.Name)) == 0 {
		return fmt.Errorf("media folder name cannot be empty")
	}
	if mf.Path != "" && len(strings.TrimSpace(mf.Path)) == 0 {
		return fmt.Errorf("media folder path cannot be empty")
	}
	return nil
}
