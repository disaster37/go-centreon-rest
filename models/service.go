package models

type ServiceBaseGet struct {
	HostId              string `json:"host id,omitempty"`
	HostName            string `json:"host name,omitempty"`
	ID                  string `json:"id,omitempty"`
	Name                string `json:"description,omitempty"`
	CheckCommand        string `json:"check command,omitempty"`
	CheckCommandArgs    string `json:"check command arg,omitempty"`
	NormalCheckInterval string `json:"normal check interval,omitempty"`
	RetryCheckInterval  string `json:"retry check interval,omitempty"`
	MaxCheckAttempts    string `json:"max check attempts,omitempty"`
	ActiveCheckEnabled  string `json:"active checks enabled,omitempty"`
	PassiveCheckEnabled string `json:"passive checks enabled,omitempty"`
	Activated           string `json:"activate,omitempty"`
}

type ServiceGet struct {
	*ServiceBaseGet
}

type ServiceTemplateGet struct {
	*ServiceBaseGet
}
