package api

import (
	"fmt"
	"strings"
	"time"
)

// Platform represents a platform in the topology
type Platform struct {
	ID                   *int64         `json:"id,omitempty"`
	PlatformName         string         `json:"platformName" validate:"required"`
	Hostname             string         `json:"hostname,omitempty"`
	Type                 string         `json:"type,omitempty"`
	Address              string         `json:"address" validate:"required"`
	ParentAddress        string         `json:"parent_address,omitempty"`
	IsRemote             *bool          `json:"isRemote,omitempty"`
	CentralServerAddress string         `json:"centralServerAddress,omitempty"`
	APIUsername          string         `json:"apiUsername,omitempty"`
	APICredentials       string         `json:"apiCredentials,omitempty"`
	APIScheme            string         `json:"apiScheme,omitempty"`
	APIPort              *int           `json:"apiPort,omitempty"`
	APIPath              string         `json:"apiPath,omitempty"`
	PeerValidation       *bool          `json:"peerValidation,omitempty"`
	Proxy                *PlatformProxy `json:"proxy,omitempty"`
	CreatedAt            *time.Time     `json:"created_at,omitempty"`
	UpdatedAt            *time.Time     `json:"updated_at,omitempty"`
}

// PlatformProxy represents proxy configuration for a platform
type PlatformProxy struct {
	Host     *string `json:"host,omitempty"`
	Port     *int    `json:"port,omitempty"`
	User     *string `json:"user,omitempty"`
	Password *string `json:"password,omitempty"`
}

// TopologyNode represents a node in the topology tree
type TopologyNode struct {
	ID               *int64         `json:"id,omitempty"`
	Name             string         `json:"name" validate:"required"`
	Type             string         `json:"type" validate:"required"`
	Address          string         `json:"address" validate:"required"`
	ParentID         *int64         `json:"parent_id,omitempty"`
	Level            *int           `json:"level,omitempty"`
	IsActive         *bool          `json:"is_active,omitempty"`
	LastHeartbeat    *time.Time     `json:"last_heartbeat,omitempty"`
	Version          string         `json:"version,omitempty"`
	Children         []TopologyNode `json:"children,omitempty"`
	ConnectionStatus string         `json:"connection_status,omitempty"`
	LastSyncTime     *time.Time     `json:"last_sync_time,omitempty"`
}

// TopologyConnection represents connection information between topology nodes
type TopologyConnection struct {
	SourceNodeID   *int64     `json:"source_node_id"`
	TargetNodeID   *int64     `json:"target_node_id"`
	ConnectionType string     `json:"connection_type"`
	Status         string     `json:"status"`
	LastTest       *time.Time `json:"last_test,omitempty"`
	ResponseTime   *float64   `json:"response_time,omitempty"`
	ErrorMessage   string     `json:"error_message,omitempty"`
}

// Validation functions for Platform
func (p *Platform) ValidateForCreate() error {
	if p.PlatformName == "" || len(strings.TrimSpace(p.PlatformName)) == 0 {
		return fmt.Errorf("platformName is required for platform creation")
	}
	if p.Address == "" || len(strings.TrimSpace(p.Address)) == 0 {
		return fmt.Errorf("address is required for platform creation")
	}
	return nil
}

func (p *Platform) ValidateForUpdate() error {
	if p.ID == nil || *p.ID <= 0 {
		return fmt.Errorf("valid platform ID is required for update operations")
	}
	if p.PlatformName != "" && len(strings.TrimSpace(p.PlatformName)) == 0 {
		return fmt.Errorf("platformName cannot be empty")
	}
	if p.Address != "" && len(strings.TrimSpace(p.Address)) == 0 {
		return fmt.Errorf("address cannot be empty")
	}
	return nil
}

func (p *Platform) ValidateForRegister() error {
	if p.PlatformName == "" || len(strings.TrimSpace(p.PlatformName)) == 0 {
		return fmt.Errorf("platformName is required for platform registration")
	}
	if p.Address == "" || len(strings.TrimSpace(p.Address)) == 0 {
		return fmt.Errorf("address is required for platform registration")
	}
	if p.Type != "" && p.Type != "central" && p.Type != "poller" && p.Type != "remote" && p.Type != "map" && p.Type != "mbi" {
		return fmt.Errorf("type must be one of: central, poller, remote, map, mbi")
	}
	return nil
}

// Validation functions for TopologyNode
func (tn *TopologyNode) ValidateForCreate() error {
	if tn.Name == "" || len(strings.TrimSpace(tn.Name)) == 0 {
		return fmt.Errorf("name is required for topology node creation")
	}
	if tn.Type == "" || len(strings.TrimSpace(tn.Type)) == 0 {
		return fmt.Errorf("type is required for topology node creation")
	}
	if tn.Address == "" || len(strings.TrimSpace(tn.Address)) == 0 {
		return fmt.Errorf("address is required for topology node creation")
	}
	return nil
}

func (tn *TopologyNode) ValidateForUpdate() error {
	if tn.ID == nil || *tn.ID <= 0 {
		return fmt.Errorf("valid topology node ID is required for update operations")
	}
	if tn.Name != "" && len(strings.TrimSpace(tn.Name)) == 0 {
		return fmt.Errorf("topology node name cannot be empty")
	}
	if tn.Address != "" && len(strings.TrimSpace(tn.Address)) == 0 {
		return fmt.Errorf("topology node address cannot be empty")
	}
	return nil
}
