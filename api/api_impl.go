package centreonapi

import "github.com/go-resty/resty/v2"

// APIImpl implement the API interface
type APIImpl struct {
	service         ServiceAPI
	serviceTemplate ServiceTemplateAPI
	client          *resty.Client
}

// New permit to get API handler
func New(client *resty.Client) API {
	return &APIImpl{
		service:         NewService(client),
		serviceTemplate: NewServiceTemplate(client),
		client:          client,
	}
}

// Service permit to get service API handler
func (api *APIImpl) Service() ServiceAPI {
	return api.service
}

// ServiceTemplate permit to get service template API handler
func (api *APIImpl) ServiceTemplate() ServiceTemplateAPI {
	return api.serviceTemplate
}

// Client permit to get instance of resty client
func (api *APIImpl) Client() *resty.Client {
	return api.client
}
