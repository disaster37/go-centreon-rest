package api

import "github.com/go-resty/resty/v2"

type API interface {
	Authentification() AuthenticationInterface
	/*
		Acknowledgement() AcknowledgementInterface
		Command() CommandInterface
		ContactGroup() ContactGroupInterface
		Contact() ContactInterface
		Downtime() DowntimeInterface
		Gorgone() GorgoneInterface
		Host() HostInterface
		HostCategory() HostCategoryInterface
		HostSeverity() HostSeverityInterface
		Media() MediaInterface
		HostTemplate() HostTemplateInterface
		Service() ServiceInterface
		ServiceCategory() ServiceCategoryInterface
		ServiceSeverity() ServiceSeverityInterface
	*/
}

type APIImpl struct {
	client *resty.Client
}

func New(client *resty.Client) API {
	return &APIImpl{
		client: client,
	}
}

func (api *APIImpl) Authentification() AuthenticationInterface {
	return NewAuthentication(api.client)
}
