package api

// HostGroupCreateOrUpdateRequest represents the payload to create or update a host group in Centreon.
type HostGroupCreateOrUpdateRequest struct {
	Name      string  `json:"name" validate:"required,max=200"`
	Alias     *string `json:"alias,omitempty" validate:"omitempty,required,max=200"`
	IconId    *int64  `json:"icon_id,omitempty" validate:"omitempty,required"`
	GeoCoords *string `json:"geo_coords,omitempty" validate:"omitempty,required,max=32"`
	Comment   *string `json:"comment,omitempty"`
	Hosts     []int64 `json:"hosts"`
}

// HostGroupResponse represents a host group in Centreon.
type HostGroupResponse struct {
	Id          int64    `json:"id"`
	Name        string   `json:"name"`
	Alias       *string  `json:"alias"`
	IconId      *int64   `json:"icon_id"`
	GeoCoords   *string  `json:"geo_coords"`
	Comment     *string  `json:"comment"`
	Hosts       []IdName `json:"hosts"`
	IsActivated *bool    `json:"is_activated,omitempty"`
}

// HostGroupRealTimeResponse represents a host group from real-time monitoring data.
type HostGroupRealTimeResponse struct {
	Id   int64                         `json:"id"`
	Name string                        `json:"name"`
	Host HostGroupRealTimeResponseHost `json:"host"`
}

// HostGroupRealTimeResponseHost represents a host within a host group from real-time monitoring data.
type HostGroupRealTimeResponseHost struct {
	Id          int64                                  `json:"id"`
	Name        string                                 `json:"name"`
	Alias       *string                                `json:"alias"`
	DisplayName *string                                `json:"display_name"`
	State       int64                                  `json:"state"`
	Services    []HostGroupRealTimeResponseHostService `json:"services"`
}

// HostGroupRealTimeResponseHostService represents a service of a host within a host group from real-time monitoring data.
type HostGroupRealTimeResponseHostService struct {
	Id          int64   `json:"id"`
	Description *string `json:"description"`
	State       int64   `json:"state"`
	DisplayName *string `json:"display_name"`
}
