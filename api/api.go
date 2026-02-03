package api

import (
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type API interface {
	Client() *resty.Client
	Authentification() AuthenticationService
	MonitoringServer() MonitoringServerService
	Host() HostService
	HostCategory() HostCategoryService
	HostSeverity() HostSeverityService
	HostTemplate() HostTemplateService
	HostGroup() HostGroupService
	Media() MediaService
	Command() CommandService
	TimePeriod() TimePeriodService
	Service() ServiceService
	ServiceTemplate() ServiceTemplateService
	ServiceGroup() ServiceGroupService
	ServiceCategory() ServiceCategoryService
	ServiceSeverity() ServiceSeverityService
	Downtime() DowntimeService
	Acknowledgement() AcknowledgementInterface
}

// DefaultAPI implements API
type DefaultAPI struct {
	client *resty.Client
	logger *logrus.Entry
}

// New creates a new instance of DefaultAPI
func New(client *resty.Client, logger *logrus.Entry) API {
	return &DefaultAPI{
		client: client,
		logger: logger,
	}
}

// Client returns the underlying resty client
func (h *DefaultAPI) Client() *resty.Client {
	return h.client
}

// Authentification returns the Authentication interface
func (h *DefaultAPI) Authentification() AuthenticationService {
	return NewAuthenticationService(h.client, h.logger)
}

func (h *DefaultAPI) Host() HostService {
	return NewHostService(h.client, h.logger)
}

func (h *DefaultAPI) MonitoringServer() MonitoringServerService {
	return NewMonitoringServerService(h.client, h.logger)
}

func (h *DefaultAPI) HostTemplate() HostTemplateService {
	return NewHostTemplateService(h.client, h.logger)
}

func (h *DefaultAPI) HostCategory() HostCategoryService {
	return NewHostCategoryService(h.client, h.logger)
}

func (h *DefaultAPI) HostSeverity() HostSeverityService {
	return NewHostSeverityService(h.client, h.logger)
}

func (h *DefaultAPI) HostGroup() HostGroupService {
	return NewHostGroupService(h.client, h.logger)
}

func (h *DefaultAPI) Media() MediaService {
	return NewMediaService(h.client, h.logger)
}

func (h *DefaultAPI) Command() CommandService {
	return NewCommandService(h.client, h.logger)
}

func (h *DefaultAPI) TimePeriod() TimePeriodService {
	return NewTimePeriodService(h.client, h.logger)
}

func (h *DefaultAPI) Service() ServiceService {
	return NewServiceService(h.client, h.logger)
}

func (h *DefaultAPI) ServiceTemplate() ServiceTemplateService {
	return NewServiceTemplateService(h.client, h.logger)
}

func (h *DefaultAPI) ServiceGroup() ServiceGroupService {
	return NewServiceGroupService(h.client, h.logger)
}

func (h *DefaultAPI) ServiceCategory() ServiceCategoryService {
	return NewServiceCategoryService(h.client, h.logger)
}

func (h *DefaultAPI) ServiceSeverity() ServiceSeverityService {
	return NewServiceSeverityService(h.client, h.logger)
}

func (h *DefaultAPI) Downtime() DowntimeService {
	return NewDowntimeService(h.client, h.logger)
}

func (h *DefaultAPI) Acknowledgement() AcknowledgementInterface {
	return NewAcknowledgementService(h.client, h.logger)
}
