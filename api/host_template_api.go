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
	// Get retrieves a host template by its ID
	// It use Find method to get the host template
	Get(id int64) (hostTemplateResponse *HostTemplateResponse, err error)

	// GetByName retrieves a host template by its Name
	// It use Find method to get the host template
	GetByName(name string) (hostTemplateResponse *HostTemplateResponse, err error)

	// Find retrieves host templates based on specific criteria
	Find(opts *ListOptions) (hostTemplateListResponse *ListResponse[HostTemplateResponse], err error)

	// Create adds a new host template
	Create(template *HostTemplateCreateRequest) (hostTemplateResponse *HostTemplateCreateResponse, err error)

	// Update modifies an existing host template
	Update(id int64, template *HostTemplateUpdateRequest) (err error)

	// Delete removes a host template by its ID
	Delete(id int64) (err error)
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
func (h *DefaultHostTemplateService) Get(id int64) (hostTemplateResponse *HostTemplateResponse, err error) {

	h.logger.Debugf("Get host template with Id: %d", id)

	hostTemplateListResponse, err := h.Find(&ListOptions{
		Search: map[string]interface{}{
			"id": id,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find host template with id %d", id)
	}

	if hostTemplateListResponse.Meta.Total == 1 {
		return &hostTemplateListResponse.Result[0], nil
	}

	return nil, nil
}

// GetByName retrieves a host template by its Name
func (h *DefaultHostTemplateService) GetByName(name string) (hostTemplateResponse *HostTemplateResponse, err error) {

	h.logger.Debugf("Get host template with Name: %s", name)

	hostTemplateListResponse, err := h.Find(&ListOptions{
		Search: map[string]interface{}{
			"name": name,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find host template with name %s", name)
	}

	if hostTemplateListResponse.Meta.Total == 1 {
		return &hostTemplateListResponse.Result[0], nil
	}

	return nil, nil
}

// Find retrieves host templates based on specific criteria
func (h *DefaultHostTemplateService) Find(opts *ListOptions) (hostTemplateListResponse *ListResponse[HostTemplateResponse], err error) {

	if opts == nil {
		opts = &ListOptions{}
	}

	h.logger.Debugf("Find host templates with options: %+v", opts)

	hostTemplateListResponse = new(ListResponse[HostTemplateResponse])

	response, err := h.client.R().
		SetQueryParams(opts.GetQueryParams()).
		SetResult(hostTemplateListResponse).
		Get("/configuration/hosts/templates")

	h.logger.Debugf("Response from find host templates: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during find host templates request")
	}

	if response.IsError() {
		return nil, errors.Errorf("list hosts templates failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return hostTemplateListResponse, nil
}

// Create adds a new host template
func (h *DefaultHostTemplateService) Create(template *HostTemplateCreateRequest) (hostTemplateResponse *HostTemplateCreateResponse, err error) {

	h.logger.Debugf("Create host template with data: %+v", template)

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(template); err != nil {
		return nil, errors.Wrap(err, "validation error on create host template")
	}

	hostTemplateResponse = new(HostTemplateCreateResponse)

	response, err := h.client.R().
		SetBody(template).
		SetResult(hostTemplateResponse).
		Post("/configuration/hosts/templates")

	h.logger.Debugf("Response from create host template: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during create host template request")
	}

	if response.IsError() {
		return nil, errors.Errorf("create host template failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return hostTemplateResponse, nil
}

// Update modifies an existing host template
// It will partial update the host template with the provided fields
func (h *DefaultHostTemplateService) Update(id int64, template *HostTemplateUpdateRequest) (err error) {

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
		return errors.Wrap(err, "error during update host template request")
	}

	if response.IsError() {
		return errors.Errorf("update host template failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return nil
}

// Delete removes a host template by its ID
func (h *DefaultHostTemplateService) Delete(id int64) (err error) {

	h.logger.Debugf("Delete host template with Id: %d", id)

	response, err := h.client.R().
		SetPathParam("hostTemplateId", fmt.Sprintf("%d", id)).
		Delete("/configuration/hosts/templates/{hostTemplateId}")

	h.logger.Debugf("Response from delete host template: %s", response.String())

	if err != nil {
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
