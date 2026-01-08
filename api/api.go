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
	/*
		Acknowledgement() AcknowledgementInterface
		Command() CommandInterface
		ContactGroup() ContactGroupInterface
		Contact() ContactInterface
		Downtime() DowntimeInterface
		Gorgone() GorgoneInterface
		Host() HostInterface
		HostCategory() HostCategoryInterface
		//HostSeverity() HostSeverityInterface
		Media() MediaInterface
		HostTemplate() HostTemplateInterface
		Service() ServiceInterface
		ServiceCategory() ServiceCategoryInterface
		ServiceSeverity() ServiceSeverityInterface
	*/
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
