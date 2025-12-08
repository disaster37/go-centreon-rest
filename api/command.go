package api

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// Command represents a monitoring command
type Command struct {
	ID             *int64 `json:"id,omitempty"`
	Name           string `json:"name" validate:"required"`
	CommandLine    string `json:"command_line" validate:"required"`
	CommandType    *int   `json:"command_type,omitempty"`
	IsActivated    *bool  `json:"is_activated,omitempty"`
	IsLocked       *bool  `json:"is_locked,omitempty"`
	CommandExample string `json:"command_example,omitempty"`
	CommandComment string `json:"command_comment,omitempty"`
	GraphID        *int64 `json:"graph_id,omitempty"`
	ConnectorID    *int64 `json:"connector_id,omitempty"`
}

// CommandArgument represents a command argument
type CommandArgument struct {
	ID                  *int64 `json:"id,omitempty"`
	CommandID           *int64 `json:"command_id,omitempty"`
	ArgumentName        string `json:"argument_name" validate:"required"`
	ArgumentValue       string `json:"argument_value,omitempty"`
	ArgumentExample     string `json:"argument_example,omitempty"`
	ArgumentDescription string `json:"argument_description,omitempty"`
	IsMacro             *bool  `json:"is_macro,omitempty"`
}

// CheckCommand represents a check command execution
type CheckCommand struct {
	ResourceType     string     `json:"resource_type" validate:"required,oneof=host service"`
	ResourceID       *int64     `json:"resource_id" validate:"required"`
	ParentResourceID *int64     `json:"parent_resource_id,omitempty"`
	CheckTime        *time.Time `json:"check_time,omitempty"`
	IsForced         *bool      `json:"is_forced,omitempty"`
}

// CommandExecutionResult represents the result of command execution
type CommandExecutionResult struct {
	CommandID     *int64    `json:"command_id,omitempty"`
	ResourceType  string    `json:"resource_type,omitempty"`
	ResourceID    *int64    `json:"resource_id,omitempty"`
	ExecutionTime time.Time `json:"execution_time"`
	ExitCode      *int      `json:"exit_code,omitempty"`
	Output        string    `json:"output,omitempty"`
	ErrorOutput   string    `json:"error_output,omitempty"`
}

// Connector represents a connector used by commands
type Connector struct {
	ID          *int64 `json:"id,omitempty"`
	Name        string `json:"name" validate:"required"`
	CommandLine string `json:"command_line" validate:"required"`
	IsActivated *bool  `json:"is_activated,omitempty"`
	Description string `json:"description,omitempty"`
}

// Validation functions for Command
func (c *Command) ValidateForCreate() error {
	if c.Name == "" || len(strings.TrimSpace(c.Name)) == 0 {
		return fmt.Errorf("name is required for command creation")
	}
	if c.CommandLine == "" || len(strings.TrimSpace(c.CommandLine)) == 0 {
		return fmt.Errorf("command_line is required for command creation")
	}
	// Validate command type if provided
	if c.CommandType != nil && (*c.CommandType < 1 || *c.CommandType > 3) {
		return fmt.Errorf("command_type must be 1 (check), 2 (notification), or 3 (miscellaneous)")
	}
	return nil
}

func (c *Command) ValidateForUpdate() error {
	if c.ID == nil || *c.ID <= 0 {
		return fmt.Errorf("valid command ID is required for update operations")
	}
	if c.Name != "" && len(strings.TrimSpace(c.Name)) == 0 {
		return fmt.Errorf("command name cannot be empty")
	}
	if c.CommandLine != "" && len(strings.TrimSpace(c.CommandLine)) == 0 {
		return fmt.Errorf("command_line cannot be empty")
	}
	return nil
}

// CommandInterface defines CRUD operations for Command entities
type CommandInterface interface {
	// Create creates a new command
	Create(ctx context.Context, command *Command) (*Command, error)

	// GetByID retrieves a command by its ID
	GetByID(ctx context.Context, id int64) (*Command, error)

	// List retrieves commands with optional filtering and pagination
	List(ctx context.Context, opts *ListOptions) (*ListResponse[Command], error)

	// Update updates an existing command
	Update(ctx context.Context, id int64, command *Command) error

	// Delete deletes a command by ID
	Delete(ctx context.Context, id int64) error
}

func (c *Command) ValidateForExecution() error {
	if c.CommandLine == "" || len(strings.TrimSpace(c.CommandLine)) == 0 {
		return fmt.Errorf("command_line is required for command execution")
	}
	return nil
}
