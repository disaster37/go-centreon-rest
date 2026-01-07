package api

import (
	"fmt"
	"net/http"

	"emperror.dev/errors"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// HostTemplateService defines CRUD operations for Host Template entities
type HostTemplateService interface {
	Get(id int64) (*HostTemplateResponse, error)
	GetByName(name string) (*HostTemplateResponse, error)
	Find(opts *ListOptions) (*ListResponse[HostTemplateResponse], error)
	Create(template *HostTemplateUpdateRequest) (*HostTemplateCreateResponse, error)
	Update(id int64, template *HostTemplateUpdateRequest) error
	Delete(id int64) error
}

// DefaultHostTemplateService implements HostTemplateService
type DefaultHostTemplateService struct {
	client *resty.Client
	logger *logrus.Entry
}

// NewHostTemplateService creates a new instance of DefaultHostTemplateService
func NewHostTemplateService(client *resty.Client, logger *logrus.Entry) HostTemplateService {
	return &DefaultHostTemplateService{
		client: client,
		logger: logger.WithField("service", "hostTemplate"),
	}
}

// Get retrieves a host template by its ID
func (h *DefaultHostTemplateService) Get(id int64) (*HostTemplateResponse, error) {

	h.logger.Debugf("Get host template with Id: %d", id)

	listResponse, err := h.Find(&ListOptions{
		Search: map[string]interface{}{
			"id": id,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find host template with id %d", id)
	}

	if listResponse.Meta.Total == 1 {
		return &listResponse.Result[0], nil
	}

	return nil, nil
}

// GetByName retrieves a host template by its Name
func (h *DefaultHostTemplateService) GetByName(name string) (*HostTemplateResponse, error) {

	h.logger.Debugf("Get host template with Name: %s", name)

	listResponse, err := h.Find(&ListOptions{
		Search: map[string]interface{}{
			"name": name,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find host template with name %s", name)
	}

	if listResponse.Meta.Total == 1 {
		return &listResponse.Result[0], nil
	}

	return nil, nil
}

// Find retrieves host templates based on specific criteria
func (h *DefaultHostTemplateService) Find(opts *ListOptions) (*ListResponse[HostTemplateResponse], error) {

	if opts == nil {
		opts = &ListOptions{}
	}

	h.logger.Debugf("Find host templates with options: %+v", opts)

	listResponse := new(ListResponse[HostTemplateResponse])

	response, err := h.client.R().
		SetQueryParams(opts.GetQueryParams()).
		SetResult(listResponse).
		Get("/configuration/hosts/templates")

	h.logger.Debugf("Response from find host templates: %s", response.String())

	if err != nil {
		h.logger.Errorf("Error while finding host templates: %v", err)
		return nil, errors.Wrap(err, "error during find host templates request")
	}

	if response.IsError() {
		return nil, errors.Errorf("list hosts templates failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return listResponse, nil
}

// Create adds a new host template
func (h *DefaultHostTemplateService) Create(template *HostTemplateUpdateRequest) (*HostTemplateCreateResponse, error) {

	h.logger.Debugf("Create host template with data: %+v", template)

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(template); err != nil {
		return nil, errors.Wrap(err, "validation error on create host template")
	}

	createResponse := new(HostTemplateCreateResponse)

	response, err := h.client.R().
		SetBody(template).
		SetResult(createResponse).
		Post("/configuration/hosts/templates")

	h.logger.Debugf("Response from create host template: %s", response.String())

	if err != nil {
		h.logger.Errorf("Error while creating host template: %v", err)
		return nil, errors.Wrap(err, "error during create host template request")
	}

	if response.IsError() {
		return nil, errors.Errorf("create host template failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return createResponse, nil
}

// Update modifies an existing host template
func (h *DefaultHostTemplateService) Update(id int64, template *HostTemplateUpdateRequest) error {

	h.logger.Debugf("Update host template with Id %d and data: %+v", id, template)

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(template); err != nil {
		return errors.Wrap(err, "validation error on update host template")
	}

	response, err := h.client.R().
		SetBody(template).
		SetPathParam("hostTemplateId", fmt.Sprintf("%d", id)).
		Patch("/configuration/hosts/templates/{hostTemplateId}")

	h.logger.Debugf("Response from update host template: %s", response.String())

	if err != nil {
		h.logger.Errorf("Error while updating host template: %v", err)
		return errors.Wrap(err, "error during update host template request")
	}

	if response.IsError() {
		return errors.Errorf("update host template failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return nil
}

// Delete removes a host template by its ID
func (h *DefaultHostTemplateService) Delete(id int64) error {

	h.logger.Debugf("Delete host template with Id: %d", id)

	response, err := h.client.R().
		SetPathParam("hostTemplateId", fmt.Sprintf("%d", id)).
		Delete("/configuration/hosts/templates/{hostTemplateId}")

	h.logger.Debugf("Response from delete host template: %s", response.String())

	if err != nil {
		h.logger.Errorf("Error while deleting host template: %v", err)
		return errors.Wrap(err, "error during delete host template request")
	}

	if response.IsError() {

		if response.StatusCode() == http.StatusNotFound {
			return nil
		}
		return errors.Errorf("delete host template failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return nil

}
