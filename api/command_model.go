package api

type CommandType int

const (
	CommandTypeNotification  CommandType = 1
	CommandTypeCheck         CommandType = 2
	CommandTypeMiscellaneous CommandType = 3
	CommandTypeDiscovery     CommandType = 4
)

// CommandCreateOrUpdateRequest represents the payload to create or update a command in Centreon.
type CommandCreateOrUpdateRequest struct {
	Name            string            `json:"name" validate:"required"`
	Type            CommandType       `json:"type" validate:"required,oneof=1 2 3 4"`
	CommandLine     string            `json:"command_line" validate:"required"`
	IsShell         *bool             `json:"is_shell,omitempty"`
	ArgumentExample *string           `json:"argument_example,omitempty" validate:"omitempty"`
	Arguments       []CommandArgument `json:"arguments,omitempty" validate:"omitempty,dive"`
	Macros          []CommandMacro    `json:"macros,omitempty" validate:"omitempty,dive"`
	ConnectorId     *int64            `json:"connector_id,omitempty" validate:"omitempty,gt=0"`
	GraphTemplateId *int64            `json:"graph_template_id,omitempty" validate:"omitempty,gt=0"`
}

// CommandArgument represents an argument for a command.
type CommandArgument struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
}

// CommandMacro represents a macro associated with a command in Centreon.
type CommandMacro struct {
	Name        string           `json:"name" validate:"required"`
	Type        CommandMacroType `json:"type" validate:"required,oneof=1 2"`
	Description *string          `json:"description,omitempty"`
}

// CommandResponse represents the command in Centreon.
type CommandResponse struct {
	Id              int64             `json:"id"`
	Name            string            `json:"name"`
	Type            CommandType       `json:"type"`
	CommandLine     string            `json:"command_line"`
	IsShell         bool              `json:"is_shell,omitempty"`
	IsLocked        bool              `json:"is_locked,omitempty"`
	IsActivated     bool              `json:"is_activated,omitempty"`
	ArgumentExample *string           `json:"argument_example,omitempty"`
	Arguments       []CommandArgument `json:"arguments,omitempty"`
	Macros          []CommandMacro    `json:"macros,omitempty"`
	Connector       *IdName           `json:"connector_id,omitempty"`
	GraphTemplate   *IdName           `json:"graph_template_id,omitempty"`
}

// CommandFind represents the command found in Centreon.
type CommandFindResponse struct {
	Id          int64       `json:"id"`
	Name        string      `json:"name"`
	Type        CommandType `json:"type"`
	CommandLine string      `json:"command_line"`
	IsShell     bool        `json:"is_shell,omitempty"`
	IsLocked    bool        `json:"is_locked,omitempty"`
	IsActivated bool        `json:"is_activated,omitempty"`
}

type CommandMacroType int

const (
	HostMacro    CommandMacroType = 1
	ServiceMacro CommandMacroType = 2
)
