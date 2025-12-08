package api

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// GorgoneResponse represents a response from Gorgone
type GorgoneResponse struct {
	Message string            `json:"message"`
	Token   string            `json:"token"`
	Data    []GorgoneDataItem `json:"data,omitempty"`
}

// GorgoneDataItem represents a single data item in Gorgone response
type GorgoneDataItem struct {
	CreationTime time.Time `json:"creation_time"`
	EventTime    time.Time `json:"event_time"`
	ID           *int      `json:"id,omitempty"`
	Token        string    `json:"token"`
	Code         *int      `json:"code,omitempty"`
	Data         string    `json:"data"`
}

// GorgoneCommand represents a command to be sent to Gorgone
type GorgoneCommand struct {
	Action  string         `json:"action" validate:"required"`
	Target  string         `json:"target,omitempty"`
	Content map[string]any `json:"content,omitempty"`
	Token   string         `json:"token,omitempty"`
}

// GorgoneActionRequest represents an action request to Gorgone
type GorgoneActionRequest struct {
	Action     string         `json:"action" validate:"required"`
	Parameters map[string]any `json:"parameters,omitempty"`
	Target     string         `json:"target,omitempty"`
}

// GorgoneLog represents a log entry from Gorgone
type GorgoneLog struct {
	ID           *int      `json:"id,omitempty"`
	Token        string    `json:"token"`
	CreationTime time.Time `json:"creation_time"`
	EventTime    time.Time `json:"event_time"`
	Code         *int      `json:"code,omitempty"`
	Etime        *float64  `json:"etime,omitempty"`
	Data         string    `json:"data"`
}

// GorgoneExecuteCommand represents a command execution request
type GorgoneExecuteCommand struct {
	Command     string            `json:"command" validate:"required"`
	Target      string            `json:"target,omitempty"`
	Timeout     *int              `json:"timeout,omitempty"`
	Environment map[string]string `json:"environment,omitempty"`
	Metadata    map[string]any    `json:"metadata,omitempty"`
}

// GorgoneFileTransfer represents a file transfer operation
type GorgoneFileTransfer struct {
	Source      string                  `json:"source" validate:"required"`
	Destination string                  `json:"destination" validate:"required"`
	Target      string                  `json:"target,omitempty"`
	Options     *GorgoneTransferOptions `json:"options,omitempty"`
}

// GorgoneTransferOptions represents options for file transfer
type GorgoneTransferOptions struct {
	Timeout    *int  `json:"timeout,omitempty"`
	Recursive  *bool `json:"recursive,omitempty"`
	Preserve   *bool `json:"preserve,omitempty"`
	Overwrite  *bool `json:"overwrite,omitempty"`
	CreateDirs *bool `json:"create_dirs,omitempty"`
}

// GorgoneModule represents a module in Gorgone
type GorgoneModule struct {
	Name       string         `json:"name" validate:"required"`
	Package    string         `json:"package" validate:"required"`
	Enable     *bool          `json:"enable,omitempty"`
	Status     string         `json:"status,omitempty"`
	Version    string         `json:"version,omitempty"`
	Config     map[string]any `json:"config,omitempty"`
	LastUpdate *time.Time     `json:"last_update,omitempty"`
}

// GorgoneConfiguration represents Gorgone configuration
type GorgoneConfiguration struct {
	Name        string             `json:"name" validate:"required"`
	Description string             `json:"description,omitempty"`
	Gorgonecore *GorgoneCoreConfig `json:"gorgonecore,omitempty"`
	Modules     []GorgoneModule    `json:"modules,omitempty"`
}

// GorgoneCoreConfig represents core Gorgone configuration
type GorgoneCoreConfig struct {
	ID            *int                      `json:"id,omitempty"`
	ExternalCom   string                    `json:"external_com,omitempty"`
	Hostname      string                    `json:"hostname,omitempty"`
	GorgoneVarLib string                    `json:"gorgone_var_lib,omitempty"`
	LogLevel      string                    `json:"log_level,omitempty"`
	LogFile       string                    `json:"log_file,omitempty"`
	PidFile       string                    `json:"pid_file,omitempty"`
	InternalCom   *GorgoneInternalComConfig `json:"internal_com,omitempty"`
}

// GorgoneInternalComConfig represents internal communication configuration
type GorgoneInternalComConfig struct {
	Type       string         `json:"type"`
	Path       string         `json:"path,omitempty"`
	Address    string         `json:"address,omitempty"`
	Port       *int           `json:"port,omitempty"`
	Encryption *bool          `json:"encryption,omitempty"`
	KeyFile    string         `json:"key_file,omitempty"`
	CertFile   string         `json:"cert_file,omitempty"`
	CAFile     string         `json:"ca_file,omitempty"`
	Options    map[string]any `json:"options,omitempty"`
}

// Validation functions for GorgoneCommand
func (gc *GorgoneCommand) ValidateForExecution() error {
	if gc.Action == "" || len(strings.TrimSpace(gc.Action)) == 0 {
		return fmt.Errorf("action is required for gorgone command execution")
	}
	return nil
}

// Validation functions for GorgoneExecuteCommand
func (gec *GorgoneExecuteCommand) ValidateForExecution() error {
	if gec.Command == "" || len(strings.TrimSpace(gec.Command)) == 0 {
		return fmt.Errorf("command is required for gorgone command execution")
	}
	if gec.Timeout != nil && *gec.Timeout <= 0 {
		return fmt.Errorf("timeout must be a positive integer")
	}
	return nil
}

// Validation functions for GorgoneFileTransfer
func (gft *GorgoneFileTransfer) ValidateForTransfer() error {
	if gft.Source == "" || len(strings.TrimSpace(gft.Source)) == 0 {
		return fmt.Errorf("source is required for gorgone file transfer")
	}
	if gft.Destination == "" || len(strings.TrimSpace(gft.Destination)) == 0 {
		return fmt.Errorf("destination is required for gorgone file transfer")
	}
	return nil
}

// GorgoneInterface defines operations for Gorgone distributed task execution
type GorgoneInterface interface {
	// Command execution operations

	// ExecuteCommand executes a command on a target node
	ExecuteCommand(ctx context.Context, command *GorgoneExecuteCommand) (*GorgoneResponse, error)

	// ExecuteCommandWithToken executes a command and returns a tracking token
	ExecuteCommandWithToken(ctx context.Context, command *GorgoneExecuteCommand) (string, error)

	// GetCommandResult retrieves the result of a command execution by token
	GetCommandResult(ctx context.Context, token string) (*GorgoneDataItem, error)

	// WaitForCommand waits for a command to complete and returns the result
	WaitForCommand(ctx context.Context, token string, timeout time.Duration) (*GorgoneDataItem, error)

	// CancelCommand cancels a running command by token
	CancelCommand(ctx context.Context, token string) error

	// ListActiveCommands retrieves currently executing commands
	ListActiveCommands(ctx context.Context, target string, opts *ListOptions) (*ListResponse[GorgoneDataItem], error)

	// File transfer operations

	// TransferFile transfers a file between nodes
	TransferFile(ctx context.Context, transfer *GorgoneFileTransfer) (*GorgoneResponse, error)

	// TransferFileWithToken transfers a file and returns a tracking token
	TransferFileWithToken(ctx context.Context, transfer *GorgoneFileTransfer) (string, error)

	// GetTransferResult retrieves the result of a file transfer by token
	GetTransferResult(ctx context.Context, token string) (*GorgoneDataItem, error)

	// ListActiveTransfers retrieves currently active file transfers
	ListActiveTransfers(ctx context.Context, target string, opts *ListOptions) (*ListResponse[GorgoneDataItem], error)

	// Node and target management

	// ListNodes retrieves available Gorgone nodes
	ListNodes(ctx context.Context, opts *ListOptions) (*ListResponse[GorgoneNode], error)

	// GetNodeStatus retrieves the status of a specific node
	GetNodeStatus(ctx context.Context, nodeID string) (*GorgoneNodeStatus, error)

	// GetNodeInfo retrieves detailed information about a node
	GetNodeInfo(ctx context.Context, nodeID string) (*GorgoneNodeInfo, error)

	// PingNode tests connectivity to a node
	PingNode(ctx context.Context, nodeID string) (*GorgonePingResult, error)

	// RestartNode restarts a Gorgone daemon on a node
	RestartNode(ctx context.Context, nodeID string) (*GorgoneResponse, error)

	// StopNode stops a Gorgone daemon on a node
	StopNode(ctx context.Context, nodeID string) (*GorgoneResponse, error)

	// StartNode starts a Gorgone daemon on a node
	StartNode(ctx context.Context, nodeID string) (*GorgoneResponse, error)

	// Module management

	// ListModules retrieves available modules on a node
	ListModules(ctx context.Context, nodeID string, opts *ListOptions) (*ListResponse[GorgoneModule], error)

	// GetModule retrieves information about a specific module
	GetModule(ctx context.Context, nodeID, moduleName string) (*GorgoneModule, error)

	// EnableModule enables a module on a node
	EnableModule(ctx context.Context, nodeID, moduleName string) error

	// DisableModule disables a module on a node
	DisableModule(ctx context.Context, nodeID, moduleName string) error

	// RestartModule restarts a module on a node
	RestartModule(ctx context.Context, nodeID, moduleName string) error

	// GetModuleConfiguration retrieves module configuration
	GetModuleConfiguration(ctx context.Context, nodeID, moduleName string) (map[string]any, error)

	// SetModuleConfiguration updates module configuration
	SetModuleConfiguration(ctx context.Context, nodeID, moduleName string, config map[string]any) error

	// Configuration management

	// GetConfiguration retrieves Gorgone configuration for a node
	GetConfiguration(ctx context.Context, nodeID string) (*GorgoneConfiguration, error)

	// UpdateConfiguration updates Gorgone configuration for a node
	UpdateConfiguration(ctx context.Context, nodeID string, config *GorgoneConfiguration) error

	// ReloadConfiguration reloads configuration on a node
	ReloadConfiguration(ctx context.Context, nodeID string) error

	// ValidateConfiguration validates a configuration before applying
	ValidateConfiguration(ctx context.Context, config *GorgoneConfiguration) (*ValidationResult, error)

	// BackupConfiguration creates a backup of current configuration
	BackupConfiguration(ctx context.Context, nodeID string) ([]byte, error)

	// RestoreConfiguration restores configuration from backup
	RestoreConfiguration(ctx context.Context, nodeID string, backup []byte) error

	// Logging and monitoring

	// GetLogs retrieves logs from a node
	GetLogs(ctx context.Context, nodeID string, opts *GorgoneLogOptions) (*ListResponse[GorgoneLog], error)

	// GetLogsByToken retrieves logs for a specific token
	GetLogsByToken(ctx context.Context, token string, opts *ListOptions) (*ListResponse[GorgoneLog], error)

	// StreamLogs streams real-time logs from a node
	StreamLogs(ctx context.Context, nodeID string, opts *GorgoneLogOptions) (<-chan GorgoneLog, error)

	// GetMetrics retrieves performance metrics from a node
	GetMetrics(ctx context.Context, nodeID string) (*GorgoneMetrics, error)

	// GetStatistics retrieves execution statistics
	GetStatistics(ctx context.Context, nodeID string, timeRange *GorgoneTimeRange) (*GorgoneStatistics, error)

	// Action and workflow management

	// SendAction sends a custom action to a node
	SendAction(ctx context.Context, action *GorgoneActionRequest) (*GorgoneResponse, error)

	// SendActionWithToken sends a custom action and returns a tracking token
	SendActionWithToken(ctx context.Context, action *GorgoneActionRequest) (string, error)

	// ListActions retrieves available actions on a node
	ListActions(ctx context.Context, nodeID string, opts *ListOptions) (*ListResponse[GorgoneAction], error)

	// GetActionResult retrieves the result of an action by token
	GetActionResult(ctx context.Context, token string) (*GorgoneDataItem, error)

	// Workflow operations

	// CreateWorkflow creates a new workflow
	CreateWorkflow(ctx context.Context, workflow *GorgoneWorkflow) (*GorgoneWorkflow, error)

	// ExecuteWorkflow executes a workflow
	ExecuteWorkflow(ctx context.Context, workflowID string) (*GorgoneResponse, error)

	// GetWorkflowStatus retrieves the status of a workflow execution
	GetWorkflowStatus(ctx context.Context, workflowID string) (*GorgoneWorkflowStatus, error)

	// ListWorkflows retrieves available workflows
	ListWorkflows(ctx context.Context, opts *ListOptions) (*ListResponse[GorgoneWorkflow], error)

	// DeleteWorkflow deletes a workflow
	DeleteWorkflow(ctx context.Context, workflowID string) error

	// Utility operations

	// GetVersion retrieves Gorgone version information
	GetVersion(ctx context.Context, nodeID string) (*GorgoneVersion, error)

	// GetHealth performs health check on a node
	GetHealth(ctx context.Context, nodeID string) (*GorgoneHealthStatus, error)

	// Cleanup performs cleanup operations on a node
	Cleanup(ctx context.Context, nodeID string, opts *GorgoneCleanupOptions) (*GorgoneResponse, error)
}

// Supporting types for Gorgone interface operations

// GorgoneNode represents a Gorgone node in the distributed system
type GorgoneNode struct {
	ID            string                `json:"id"`
	Name          string                `json:"name"`
	Type          string                `json:"type"` // "central", "poller", "remote"
	Address       string                `json:"address"`
	Port          *int                  `json:"port,omitempty"`
	Status        string                `json:"status"` // "connected", "disconnected", "error"
	Version       string                `json:"version,omitempty"`
	LastSeen      *time.Time            `json:"last_seen,omitempty"`
	Configuration *GorgoneConfiguration `json:"configuration,omitempty"`
	Capabilities  []string              `json:"capabilities,omitempty"`
	Metadata      map[string]any        `json:"metadata,omitempty"`
}

// GorgoneNodeStatus represents the current status of a Gorgone node
type GorgoneNodeStatus struct {
	NodeID           string         `json:"node_id"`
	Status           string         `json:"status"`
	Uptime           *time.Duration `json:"uptime,omitempty"`
	LastConnected    *time.Time     `json:"last_connected,omitempty"`
	LastDisconnected *time.Time     `json:"last_disconnected,omitempty"`
	ActiveTasks      int            `json:"active_tasks"`
	QueuedTasks      int            `json:"queued_tasks"`
	CompletedTasks   int            `json:"completed_tasks"`
	FailedTasks      int            `json:"failed_tasks"`
	LoadAverage      *float64       `json:"load_average,omitempty"`
	MemoryUsage      *float64       `json:"memory_usage,omitempty"`
	CPUUsage         *float64       `json:"cpu_usage,omitempty"`
}

// GorgoneNodeInfo represents detailed information about a Gorgone node
type GorgoneNodeInfo struct {
	Node          *GorgoneNode          `json:"node"`
	Status        *GorgoneNodeStatus    `json:"status"`
	Modules       []GorgoneModule       `json:"modules"`
	Configuration *GorgoneConfiguration `json:"configuration"`
	SystemInfo    *GorgoneSystemInfo    `json:"system_info,omitempty"`
}

// GorgoneSystemInfo represents system information for a node
type GorgoneSystemInfo struct {
	OS                string `json:"os"`
	OSVersion         string `json:"os_version"`
	Architecture      string `json:"architecture"`
	Hostname          string `json:"hostname"`
	CPUCount          int    `json:"cpu_count"`
	TotalMemoryMB     int64  `json:"total_memory_mb"`
	AvailableMemoryMB int64  `json:"available_memory_mb"`
	DiskSpaceGB       int64  `json:"disk_space_gb"`
	AvailableDiskGB   int64  `json:"available_disk_gb"`
}

// GorgonePingResult represents the result of pinging a node
type GorgonePingResult struct {
	NodeID       string        `json:"node_id"`
	Success      bool          `json:"success"`
	ResponseTime time.Duration `json:"response_time"`
	Message      string        `json:"message,omitempty"`
	Timestamp    time.Time     `json:"timestamp"`
}

// GorgoneLogOptions represents options for retrieving logs
type GorgoneLogOptions struct {
	StartTime  *time.Time `json:"start_time,omitempty"`
	EndTime    *time.Time `json:"end_time,omitempty"`
	Level      string     `json:"level,omitempty"` // "debug", "info", "warning", "error"
	Module     string     `json:"module,omitempty"`
	SearchText string     `json:"search_text,omitempty"`
	Follow     *bool      `json:"follow,omitempty"`
	MaxLines   *int       `json:"max_lines,omitempty"`
}

// GorgoneMetrics represents performance metrics from a node
type GorgoneMetrics struct {
	NodeID               string                       `json:"node_id"`
	Timestamp            time.Time                    `json:"timestamp"`
	SystemMetrics        *GorgoneSystemMetrics        `json:"system_metrics"`
	TaskMetrics          *GorgoneTaskMetrics          `json:"task_metrics"`
	CommunicationMetrics *GorgoneCommunicationMetrics `json:"communication_metrics"`
}

// GorgoneSystemMetrics represents system-level metrics
type GorgoneSystemMetrics struct {
	CPUUsagePercent    float64 `json:"cpu_usage_percent"`
	MemoryUsagePercent float64 `json:"memory_usage_percent"`
	DiskUsagePercent   float64 `json:"disk_usage_percent"`
	LoadAverage1Min    float64 `json:"load_average_1min"`
	LoadAverage5Min    float64 `json:"load_average_5min"`
	LoadAverage15Min   float64 `json:"load_average_15min"`
	NetworkBytesIn     int64   `json:"network_bytes_in"`
	NetworkBytesOut    int64   `json:"network_bytes_out"`
}

// GorgoneTaskMetrics represents task execution metrics
type GorgoneTaskMetrics struct {
	ActiveTasks           int     `json:"active_tasks"`
	QueuedTasks           int     `json:"queued_tasks"`
	CompletedTasks24h     int     `json:"completed_tasks_24h"`
	FailedTasks24h        int     `json:"failed_tasks_24h"`
	AverageExecutionTime  float64 `json:"average_execution_time_seconds"`
	TaskThroughputPerHour int     `json:"task_throughput_per_hour"`
}

// GorgoneCommunicationMetrics represents communication metrics
type GorgoneCommunicationMetrics struct {
	MessagesReceived24h int     `json:"messages_received_24h"`
	MessagesSent24h     int     `json:"messages_sent_24h"`
	ConnectionErrors24h int     `json:"connection_errors_24h"`
	AverageResponseTime float64 `json:"average_response_time_ms"`
	ActiveConnections   int     `json:"active_connections"`
}

// GorgoneStatistics represents execution statistics
type GorgoneStatistics struct {
	NodeID      string               `json:"node_id"`
	TimeRange   *GorgoneTimeRange    `json:"time_range"`
	TaskStats   *GorgoneTaskStats    `json:"task_stats"`
	ModuleStats []GorgoneModuleStats `json:"module_stats"`
	ErrorStats  *GorgoneErrorStats   `json:"error_stats"`
}

// GorgoneTimeRange represents a time range for statistics
type GorgoneTimeRange struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// GorgoneTaskStats represents task execution statistics
type GorgoneTaskStats struct {
	TotalTasks           int     `json:"total_tasks"`
	CompletedTasks       int     `json:"completed_tasks"`
	FailedTasks          int     `json:"failed_tasks"`
	CancelledTasks       int     `json:"cancelled_tasks"`
	SuccessRate          float64 `json:"success_rate"`
	AverageExecutionTime float64 `json:"average_execution_time"`
	MinExecutionTime     float64 `json:"min_execution_time"`
	MaxExecutionTime     float64 `json:"max_execution_time"`
}

// GorgoneModuleStats represents module-specific statistics
type GorgoneModuleStats struct {
	ModuleName     string  `json:"module_name"`
	TasksProcessed int     `json:"tasks_processed"`
	TasksSucceeded int     `json:"tasks_succeeded"`
	TasksFailed    int     `json:"tasks_failed"`
	AverageTime    float64 `json:"average_time"`
}

// GorgoneErrorStats represents error statistics
type GorgoneErrorStats struct {
	ConnectionErrors     int `json:"connection_errors"`
	TimeoutErrors        int `json:"timeout_errors"`
	AuthenticationErrors int `json:"authentication_errors"`
	ValidationErrors     int `json:"validation_errors"`
	ExecutionErrors      int `json:"execution_errors"`
	OtherErrors          int `json:"other_errors"`
}

// GorgoneAction represents an available action on a node
type GorgoneAction struct {
	Name        string                   `json:"name"`
	Description string                   `json:"description,omitempty"`
	Module      string                   `json:"module"`
	Parameters  []GorgoneActionParameter `json:"parameters,omitempty"`
	Examples    []string                 `json:"examples,omitempty"`
}

// GorgoneActionParameter represents a parameter for an action
type GorgoneActionParameter struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Required    bool   `json:"required"`
	Description string `json:"description,omitempty"`
	Default     any    `json:"default,omitempty"`
}

// GorgoneWorkflow represents a workflow definition
type GorgoneWorkflow struct {
	ID          string                `json:"id,omitempty"`
	Name        string                `json:"name" validate:"required"`
	Description string                `json:"description,omitempty"`
	Steps       []GorgoneWorkflowStep `json:"steps" validate:"required,min=1"`
	Variables   map[string]any        `json:"variables,omitempty"`
	CreatedAt   *time.Time            `json:"created_at,omitempty"`
	UpdatedAt   *time.Time            `json:"updated_at,omitempty"`
	CreatedBy   string                `json:"created_by,omitempty"`
}

// GorgoneWorkflowStep represents a step in a workflow
type GorgoneWorkflowStep struct {
	ID              string         `json:"id" validate:"required"`
	Name            string         `json:"name" validate:"required"`
	Action          string         `json:"action" validate:"required"`
	Target          string         `json:"target,omitempty"`
	Parameters      map[string]any `json:"parameters,omitempty"`
	DependsOn       []string       `json:"depends_on,omitempty"`
	ContinueOnError bool           `json:"continue_on_error,omitempty"`
	Timeout         *int           `json:"timeout,omitempty"`
	RetryCount      *int           `json:"retry_count,omitempty"`
	RetryDelay      *int           `json:"retry_delay,omitempty"`
}

// GorgoneWorkflowStatus represents the status of a workflow execution
type GorgoneWorkflowStatus struct {
	WorkflowID   string                      `json:"workflow_id"`
	Status       string                      `json:"status"` // "running", "completed", "failed", "cancelled"
	StartTime    *time.Time                  `json:"start_time,omitempty"`
	EndTime      *time.Time                  `json:"end_time,omitempty"`
	StepStatuses []GorgoneWorkflowStepStatus `json:"step_statuses"`
	ErrorMessage string                      `json:"error_message,omitempty"`
	Variables    map[string]any              `json:"variables,omitempty"`
}

// GorgoneWorkflowStepStatus represents the status of a workflow step
type GorgoneWorkflowStepStatus struct {
	StepID       string     `json:"step_id"`
	Status       string     `json:"status"` // "pending", "running", "completed", "failed", "skipped"
	StartTime    *time.Time `json:"start_time,omitempty"`
	EndTime      *time.Time `json:"end_time,omitempty"`
	ErrorMessage string     `json:"error_message,omitempty"`
	Output       string     `json:"output,omitempty"`
	RetryCount   int        `json:"retry_count"`
}

// GorgoneVersion represents version information
type GorgoneVersion struct {
	NodeID    string `json:"node_id"`
	Version   string `json:"version"`
	BuildDate string `json:"build_date,omitempty"`
	GitCommit string `json:"git_commit,omitempty"`
	GoVersion string `json:"go_version,omitempty"`
}

// GorgoneHealthStatus represents health check results
type GorgoneHealthStatus struct {
	NodeID       string               `json:"node_id"`
	Status       string               `json:"status"` // "healthy", "degraded", "unhealthy"
	Timestamp    time.Time            `json:"timestamp"`
	Checks       []GorgoneHealthCheck `json:"checks"`
	OverallScore float64              `json:"overall_score"` // 0-100
}

// GorgoneHealthCheck represents an individual health check
type GorgoneHealthCheck struct {
	Name     string        `json:"name"`
	Status   string        `json:"status"` // "pass", "warn", "fail"
	Message  string        `json:"message,omitempty"`
	Duration time.Duration `json:"duration"`
	Score    float64       `json:"score"` // 0-100
}

// GorgoneCleanupOptions represents options for cleanup operations
type GorgoneCleanupOptions struct {
	OlderThan       *time.Duration `json:"older_than,omitempty"`
	MaxLogs         *int           `json:"max_logs,omitempty"`
	MaxTempFiles    *int           `json:"max_temp_files,omitempty"`
	CleanLogs       *bool          `json:"clean_logs,omitempty"`
	CleanTempFiles  *bool          `json:"clean_temp_files,omitempty"`
	CleanCacheFiles *bool          `json:"clean_cache_files,omitempty"`
	DryRun          *bool          `json:"dry_run,omitempty"`
}
