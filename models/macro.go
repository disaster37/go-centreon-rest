package models

type Macro struct {
	Name        string `json:"macro name,omitempty"`
	Value       string `json:"macro value,omitempty"`
	IsPassword  string `json:"is_password,omitempty"`
	Description string `json:"description,omitempty"`
	Source      string `json:"source,omitempty"`
}

func (m *Macro) IsValid() bool {
	if m.Name == "" || (m.IsPassword != "0" && m.IsPassword != "1") {
		return false
	}
	return true
}
