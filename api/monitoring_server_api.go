package api

import (
	"fmt"

	"emperror.dev/errors"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// MonitoringServerService defines operations for monitoring server entities
type MonitoringServerService interface {
	// Get retrieves monitoring server details
	List(opts *ListOptions) (*ListResponse[MonitoringServerListResult], error)

	// GenerateConfiguration generates configuration for a monitoring server
	GenerateConfiguration(id int64) error

	// ReloadConfiguration reloads configuration for a monitoring server
	ReloadConfiguration(id int64) error

	// GenerateAndReloadConfiguration generates and reloads configuration for a monitoring server
	GenerateAndReloadConfiguration(id int64) error

	// GenerateConfigurationAll generates configuration for all monitoring servers
	GenerateConfigurationAll() error

	// ReloadConfigurationAll reloads configuration for all monitoring servers
	ReloadConfigurationAll() error

	// GenerateAndReloadConfigurationAll generates and reloads configuration for all monitoring servers
	GenerateAndReloadConfigurationAll() error

	// ListRealTime retrieves real-time monitoring server details
	ListFromRealTime(opts *ListOptions) (*ListResponse[MonitoringServerListRealTimeResult], error)
}

// DefaultMonitoringServerService implements MonitoringServerService
type DefaultMonitoringServerService struct {
	client *resty.Client
	logger *logrus.Entry
}

// NewMonitoringServerService creates a new instance of DefaultMonitoringServerService
func NewMonitoringServerService(client *resty.Client, logger *logrus.Entry) MonitoringServerService {
	return &DefaultMonitoringServerService{
		client: client,
		logger: logger.WithField("service", "monitoringServer"),
	}
}

// List retrieves monitoring server details
func (h *DefaultMonitoringServerService) List(opts *ListOptions) (*ListResponse[MonitoringServerListResult], error) {

	h.logger.Debugf("List monitoring servers with options: %+v", opts)

	if opts == nil {
		opts = &ListOptions{}
	}

	listResponse := new(ListResponse[MonitoringServerListResult])

	response, err := h.client.R().
		SetQueryParams(opts.GetQueryParams()).
		SetResult(listResponse).
		Get("/configuration/monitoring-servers")

	h.logger.Debugf("Response from list monitoring servers: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during list monitoring server request")
	}

	if response.IsError() {
		return nil, errors.Errorf("error response from list monitoring server request: %s", response.String())
	}

	return listResponse, nil
}

// ListRealTime retrieves real-time monitoring server details
func (h *DefaultMonitoringServerService) ListFromRealTime(opts *ListOptions) (*ListResponse[MonitoringServerListRealTimeResult], error) {

	h.logger.Debugf("List real-time monitoring servers with options: %+v", opts)

	if opts == nil {
		opts = &ListOptions{}
	}

	listResponse := new(ListResponse[MonitoringServerListRealTimeResult])

	response, err := h.client.R().
		SetQueryParams(opts.GetQueryParams()).
		SetResult(listResponse).
		Get("/monitoring/servers")

	h.logger.Debugf("Response from list real-time monitoring servers: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during list real-time monitoring server request")
	}

	if response.IsError() {
		return nil, errors.Errorf("error response from list real-time monitoring server request: %s", response.String())
	}

	return listResponse, nil
}

// GenerateConfiguration generates configuration for a monitoring server
func (h *DefaultMonitoringServerService) GenerateConfiguration(id int64) error {
	h.logger.Debugf("Generate configuration for monitoring server ID: %d", id)

	if id <= 0 {
		return errors.Errorf("invalid monitoring server id: %d", id)
	}

	response, err := h.client.R().
		SetPathParam("monitoringServerId", fmt.Sprintf("%d", id)).
		Get("/configuration/monitoring-servers/{monitoringServerId}/generate")

	h.logger.Debugf("Response from generate configuration: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during generate configuration request")
	}

	if response.IsError() {
		return errors.Errorf("error response from generate configuration request: %s", response.String())
	}

	return nil
}

// ReloadConfiguration reloads configuration for a monitoring server
func (h *DefaultMonitoringServerService) ReloadConfiguration(id int64) error {
	h.logger.Debugf("Reload configuration for monitoring server ID: %d", id)

	if id <= 0 {
		return errors.Errorf("invalid monitoring server id: %d", id)
	}

	response, err := h.client.R().
		SetPathParam("monitoringServerId", fmt.Sprintf("%d", id)).
		Get("/configuration/monitoring-servers/{monitoringServerId}/reload")

	h.logger.Debugf("Response from reload configuration: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during reload configuration request")
	}

	if response.IsError() {
		return errors.Errorf("error response from reload configuration request: %s", response.String())
	}

	return nil
}

// GenerateAndReloadConfiguration generates and reloads configuration for a monitoring server
func (h *DefaultMonitoringServerService) GenerateAndReloadConfiguration(id int64) error {
	h.logger.Debugf("Generate and reload configuration for monitoring server ID: %d", id)

	if id <= 0 {
		return errors.Errorf("invalid monitoring server id: %d", id)
	}

	response, err := h.client.R().
		SetPathParam("monitoringServerId", fmt.Sprintf("%d", id)).
		Get("/configuration/monitoring-servers/{monitoringServerId}/generate-and-reload")

	h.logger.Debugf("Response from generate and reload configuration: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during generate and reload configuration request")
	}

	if response.IsError() {
		return errors.Errorf("error response from generate and reload configuration request: %s", response.String())
	}

	return nil
}

// GenerateConfigurationAll generates configuration for all monitoring servers
func (h *DefaultMonitoringServerService) GenerateConfigurationAll() error {
	h.logger.Debug("Generate configuration for all monitoring servers")

	response, err := h.client.R().
		Get("/configuration/monitoring-servers/generate")

	h.logger.Debugf("Response from generate configuration all: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during generate configuration all request")
	}

	if response.IsError() {
		return errors.Errorf("error response from generate configuration all request: %s", response.String())
	}

	return nil
}

// ReloadConfigurationAll reloads configuration for all monitoring servers
func (h *DefaultMonitoringServerService) ReloadConfigurationAll() error {
	h.logger.Debug("Reload configuration for all monitoring servers")

	response, err := h.client.R().
		Get("/configuration/monitoring-servers/reload")

	h.logger.Debugf("Response from reload configuration all: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during reload configuration all request")
	}

	if response.IsError() {
		return errors.Errorf("error response from reload configuration all request: %s", response.String())
	}

	return nil
}

// GenerateAndReloadConfigurationAll generates and reloads configuration for all monitoring servers
func (h *DefaultMonitoringServerService) GenerateAndReloadConfigurationAll() error {
	h.logger.Debug("Generate and reload configuration for all monitoring servers")

	response, err := h.client.R().
		Get("/configuration/monitoring-servers/generate-and-reload")

	h.logger.Debugf("Response from generate and reload configuration all: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during generate and reload configuration all request")
	}

	if response.IsError() {
		return errors.Errorf("error response from generate and reload configuration all request: %s", response.String())
	}

	return nil
}
