package api

import (
	"fmt"
	"net/http"
	"strings"

	"emperror.dev/errors"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// ServiceTemplateService defines CRUD operations for Service Template entities
type ServiceTemplateService interface {
	// Get retrieves a service template by its ID
	// Need PR https://github.com/centreon/centreon/pull/9451
	Get(id int64) (serviceTemplateResponse *ServiceTemplateResponse, err error)

	// GetByName retrieves a service template by its Name
	// It use Find method to get the service template
	GetByName(name string) (serviceTemplateResponse *ServiceTemplateListResponse, err error)

	// Find retrieves service templates based on specific criteria
	Find(opts *ListOptions) (serviceTemplateListResponse *ListResponse[ServiceTemplateListResponse], err error)

	// Create adds a new service template
	Create(template *ServiceTemplateCreateRequest) (serviceTemplateResponse *ServiceTemplateResponse, err error)

	// Update modifies an existing service template
	// It is a partial update
	Update(id int64, template *ServiceTemplateUpdateRequest) (err error)

	// Delete removes a service template by its ID
	Delete(id int64) (err error)
}

// DefaultServiceTemplateService implements ServiceTemplateService
type DefaultServiceTemplateService struct {
	client *resty.Client
	logger *logrus.Entry
}

// NewServiceTemplateService creates a new instance of DefaultServiceTemplateService
func NewServiceTemplateService(client *resty.Client, logger *logrus.Entry) ServiceTemplateService {
	return &DefaultServiceTemplateService{
		client: client,
		logger: logger.WithField("service", "serviceTemplate"),
	}
}

// Get retrieves a service template by its ID
func (h *DefaultServiceTemplateService) Get(id int64) (serviceTemplateResponse *ServiceTemplateResponse, err error) {

	h.logger.Debugf("Get service template with Id: %d", id)

	if id <= 0 {
		return nil, errors.Errorf("invalid service template id: %d", id)
	}

	serviceTemplateResponse = new(ServiceTemplateResponse)

	response, err := h.client.R().
		SetPathParam("serviceTemplateId", fmt.Sprintf("%d", id)).
		SetResult(serviceTemplateResponse).
		Get("/configuration/services/templates/{serviceTemplateId}")

	h.logger.Debugf("Response from get service template: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during get service template request")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil, nil
		}
		return nil, errors.Errorf("error getting service template with code %s and message %s", response.Status(), response.String())
	}

	return serviceTemplateResponse, nil
}

// GetByName retrieves a service template by its Name
func (h *DefaultServiceTemplateService) GetByName(name string) (response *ServiceTemplateListResponse, err error) {
	h.logger.Debugf("Get Service Template by Name: %s", name)

	if strings.TrimSpace(name) == "" {
		return nil, errors.New("service template name cannot be empty")
	}

	options := &ListOptions{
		Search: map[string]any{
			"name": name,
		},
	}

	listResponse, err := h.Find(options)
	if err != nil {
		return nil, errors.Wrapf(err, "error finding service template by name %s", name)
	}

	if listResponse.Meta.Total == 1 {
		return &listResponse.Result[0], nil
	}

	return nil, nil
}

// Find retrieves service templates based on specific criteria
func (h *DefaultServiceTemplateService) Find(opts *ListOptions) (serviceTemplateListResponse *ListResponse[ServiceTemplateListResponse], err error) {

	if opts == nil {
		opts = &ListOptions{}
	}

	h.logger.Debugf("Find Service Templates with options: %+v", opts)

	serviceTemplateListResponse = new(ListResponse[ServiceTemplateListResponse])

	reponse, err := h.client.R().
		SetResult(serviceTemplateListResponse).
		SetQueryParams(opts.GetQueryParams()).
		Get("/configuration/services/templates")

	h.logger.Debugf("Response from find service templates: %s", reponse.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during find service templates request")
	}

	if reponse.IsError() {
		return nil, errors.Errorf("list service templates failed with status code %d and message %s", reponse.StatusCode(), reponse.String())
	}

	return serviceTemplateListResponse, nil
}

// Create adds a new service template
func (h *DefaultServiceTemplateService) Create(template *ServiceTemplateCreateRequest) (serviceTemplateResponse *ServiceTemplateResponse, err error) {

	h.logger.Debugf("Create Service Template with data: %+v", template)

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(template); err != nil {
		return nil, errors.Wrap(err, "validation error on create service template")
	}

	serviceTemplateResponse = new(ServiceTemplateResponse)

	response, err := h.client.R().
		SetBody(template).
		SetResult(serviceTemplateResponse).
		Post("/configuration/services/templates")

	h.logger.Debugf("Response from create service template: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during create service template request")
	}

	if response.IsError() {
		return nil, errors.Errorf("create service template failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return serviceTemplateResponse, nil
}

// Update modifies an existing service template
func (h *DefaultServiceTemplateService) Update(id int64, template *ServiceTemplateUpdateRequest) (err error) {
	h.logger.Debugf("Update Service Template ID %d with data: %+v", id, template)

	if id <= 0 {
		return errors.Errorf("invalid service template id: %d", id)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(template); err != nil {
		return errors.Wrap(err, "validation error on update service template")
	}

	response, err := h.client.R().
		SetPathParam("serviceTemplateId", fmt.Sprintf("%d", id)).
		SetBody(template).
		Patch("/configuration/services/templates/{serviceTemplateId}")

	h.logger.Debugf("Update Service Template Response: %+v", response)

	if err != nil {
		return errors.Wrap(err, "error during updating service template request")
	}

	if response.IsError() {
		return errors.Errorf("error updating service template with code %s and message %s", response.Status(), response.String())
	}

	return nil
}

// Delete removes a service template by its ID
func (h *DefaultServiceTemplateService) Delete(id int64) (err error) {
	h.logger.Debugf("Delete Service Template with Id: %d", id)

	if id <= 0 {
		return errors.Errorf("invalid service template id: %d", id)
	}

	response, err := h.client.R().
		SetPathParam("serviceTemplateId", fmt.Sprintf("%d", id)).
		Delete("/configuration/services/templates/{serviceTemplateId}")

	h.logger.Debugf("Response from delete service template: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during delete service template request")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil
		}
		return errors.Errorf("delete service template failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return nil
}
