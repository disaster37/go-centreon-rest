package centreonapi

import (
	"encoding/json"

	"github.com/disaster37/go-centreon-rest/v21/models"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

// APIImpl implement the API interface
type APIImpl struct {
	service         ServiceAPI
	serviceTemplate ServiceTemplateAPI
	client          *resty.Client
	config          *models.Config
}

// New permit to get API handler
func New(client *resty.Client, config *models.Config) API {
	return &APIImpl{
		service:         NewService(client),
		serviceTemplate: NewServiceTemplate(client),
		client:          client,
		config:          config,
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

func (api *APIImpl) Auth() (err error) {
	resp, err := api.Client().R().
		SetFormData(map[string]string{
			"username": api.config.Username,
			"password": api.config.Password,
		}).
		SetQueryParams(map[string]string{
			"action": "authenticate",
			"object": "",
		}).
		Post("")
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 300 {
		return errors.Errorf("Error when signin: %s", resp.Body())
	}
	result := map[string]string{}
	if err = json.Unmarshal(resp.Body(), &result); err != nil {
		return err
	}
	if result["authToken"] == "" {
		return errors.New("We get an empty token...")
	}
	api.Client().SetHeader("centreon-auth-token", result["authToken"])

	return nil

}
