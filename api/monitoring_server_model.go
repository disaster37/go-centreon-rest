package api

import "time"

// MonitoringServerListResult represents the monitoring server details
type MonitoringServerListResult struct {
	Id                       int64      `json:"id"`
	Name                     string     `json:"name"`
	Address                  string     `json:"address"`
	IsLocalhost              bool       `json:"is_localhost"`
	IsDefault                bool       `json:"is_default"`
	SshPort                  int        `json:"ssh_port"`
	LastRestart              *time.Time `json:"last_restart,omitempty"`
	EngineStartCommand       *string    `json:"engine_start_command"`
	EngineStopCommand        *string    `json:"engine_stop_command"`
	EngineRestartCommand     *string    `json:"engine_restart_command"`
	EngineReloadCommand      *string    `json:"engine_reload_command"`
	NagiosBin                *string    `json:"nagios_bin"`
	NagiostatsBin            *string    `json:"nagiostats_bin"`
	BrokerReloadCommand      *string    `json:"broker_reload_command"`
	CentreonBrokerCfgPath    *string    `json:"centreonbroker_cfg_path"`
	CentreonBrokerModulePath *string    `json:"centreonbroker_module_path"`
	CentreonBrokerLogsPath   *string    `json:"centreonbroker_logs_path"`
	CentreonConnectorPath    *string    `json:"centreonconnector_path"`
	InitScriptCentreontrapd  *string    `json:"init_script_centreontrapd"`
	SnmpTrapdPathConf        *string    `json:"snmp_trapd_path_conf"`
	RemoteId                 *int64     `json:"remote_id"`
	RemoteServerUseAsProxy   bool       `json:"remote_server_use_as_proxy"`
	IsUpdated                bool       `json:"is_updated"`
	IsActivate               bool       `json:"is_activate"`
}

// MonitoringServerListRealTimeResult represents the real-time monitoring server details
type MonitoringServerListRealTimeResult struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Description *string   `json:"description,omitempty"`
	IsRunning   bool      `json:"is_running"`
	LastAlive   Timestamp `json:"last_alive"`
	Version     string    `json:"version,omitempty"`
}
