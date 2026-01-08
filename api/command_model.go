package api

type CommandType int

const (
	CommandTypeNotification  CommandType = 1
	CommandTypeCheck         CommandType = 2
	CommandTypeMiscellaneous CommandType = 3
	CommandTypeDiscovery     CommandType = 4
)

// CommandCreateRequest represents the payload to create a new command in Centreon.
type CommandCreateRequest struct {
	Name            string            `json:"name" validate:"required"`
	Type            CommandType       `json:"type" validate:"required,oneof=1 2 3 4"`
	CommandLine     string            `json:"command_line" validate:"required"`
	IsShell         *bool             `json:"is_shell,omitempty"`
	ArgumentExample *string           `json:"argument_example,omitempty"`
	Arguments       []CommandArgument `json:"arguments,omitempty"`
	Macros          []CommandMacro    `json:"macros,omitempty"`
	ConnectorId     *int64            `json:"connector_id,omitempty"`
	GraphTemplateId *int64            `json:"graph_template_id,omitempty"`
}

// CommandResponse represents a command in Centreon.
type CommandArgument struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// CommandMacro represents a macro associated with a command in Centreon.
type CommandMacro struct {
	Name        string `json:"name" validate:"required"`
	Type        string `json:"type" validate:"required,oneof=1 2"`
	Description string `json:"description" validate:"required"`
}

// CommandCreateResponse represents the response after creating a command in Centreon.
type CommandCreateResponse struct {
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

// CommandResponse represents a command in Centreon.
type CommandResponse struct {
	Id          int64       `json:"id"`
	Name        string      `json:"name"`
	Type        CommandType `json:"type"`
	CommandLine string      `json:"command_line"`
	IsShell     bool        `json:"is_shell,omitempty"`
	IsLocked    bool        `json:"is_locked,omitempty"`
	IsActivated bool        `json:"is_activated,omitempty"`
}
