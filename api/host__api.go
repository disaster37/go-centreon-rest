package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"emperror.dev/errors"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// HostService defines CRUD operations for Host entities
type HostService interface {
	// Create creates a new host
	Create(host *HostUpdateRequest) (*HostCreateResponse, error)

	// Update updates an existing host
	Update(id int64, host *HostUpdateRequest) error

	// Delete deletes a host by ID
	Delete(id int64) error

	// Find retrieves hosts based on specific criteria
	Find(opts *ListOptions) (*ListResponse[HostFindResult], error)

	Get(id int64) (*HostFindResult, error)

	GetByName(name string) (*HostFindResult, error)

	// GetFromRealTime retrieves a host by its ID
	// It use monitoring API to get host details
	GetFromRealTime(id int64) (*HostResponse, error)

	// ListFromRealTime retrieves hosts with optional filtering and pagination
	// It use monitoring API to list hosts
	ListFromRealTime(opts *HostListOptions) (*ListResponse[HostResponse], error)
}

// DefaultHostService implements HostService
type DefaultHostService struct {
	client *resty.Client
	logger *logrus.Entry
}

// NewHostService creates a new instance of DefaultHostService
func NewHostService(client *resty.Client, logger *logrus.Entry) HostService {
	return &DefaultHostService{
		client: client,
		logger: logger.WithField("service", "host"),
	}
}

// Get retrieves a host by its ID
func (h *DefaultHostService) GetFromRealTime(id int64) (*HostResponse, error) {

	h.logger.Debugf("Get host with Id: %d", id)

	hostReponse := new(HostResponse)

	response, err := h.client.R().
		SetResult(hostReponse).
		SetPathParam("hostId", fmt.Sprintf("%d", id)).
		Get("/monitoring/hosts/{hostId}")

	h.logger.Debugf("Response from get host: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during get host request")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil, nil
		}
		return nil, errors.Errorf("get host failed with status code: %d", response.StatusCode())
	}

	return hostReponse, nil
}

// List retrieves hosts with optional filtering and pagination
func (h *DefaultHostService) ListFromRealTime(opts *HostListOptions) (*ListResponse[HostResponse], error) {

	if opts == nil {
		opts = &HostListOptions{}
	}

	h.logger.Debugf("List hosts with options: %+v", opts)

	listReponseHost := new(ListResponse[HostResponse])

	response, err := h.client.R().
		SetResult(listReponseHost).
		SetQueryParams(opts.GetQueryParams()).
		Get("/monitoring/hosts")

	h.logger.Debugf("Response from list hosts: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during list hosts request")
	}

	if response.IsError() {
		return nil, errors.Errorf("list hosts failed with status code: %d", response.StatusCode())
	}

	return listReponseHost, nil

}

// Create creates a new host
func (h *DefaultHostService) Create(host *HostUpdateRequest) (*HostCreateResponse, error) {

	h.logger.Debugf("Create Host: %s", host.String())

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(host); err != nil {
		return nil, errors.Wrap(err, "validation error on create host")
	}

	hostCreateResponse := new(HostCreateResponse)

	response, err := h.client.R().
		SetResult(hostCreateResponse).
		SetBody(host).
		Post("/configuration/hosts")

	h.logger.Debugf("Response from create host: %s", response.String())

	if err != nil {
		return nil, errors.Wrapf(err, "error during create host request: %s", response.String())
	}

	if response.IsError() {
		return nil, errors.Errorf("create host failed with status code: %d and message: %s", response.StatusCode(), response.String())
	}

	return hostCreateResponse, nil

}

// Update updates an existing host
func (h *DefaultHostService) Update(id int64, host *HostUpdateRequest) error {

	h.logger.Debugf("Update host with id: %d, Host: %s", id, host.String())

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(host); err != nil {
		return errors.Wrap(err, "validation error on update host")
	}

	responseMessage := new(ResponseMessage)

	response, err := h.client.R().
		SetResult(responseMessage).
		SetBody(host).
		SetPathParam("hostId", fmt.Sprintf("%d", id)).
		Patch("/configuration/hosts/{hostId}")

	h.logger.Debugf("Response from update host: %s", response.String())

	if err != nil {
		return errors.Wrapf(err, "error during update host request: %s", responseMessage.Message)
	}

	if response.IsError() {
		return errors.Errorf("update host failed with status code: %d and message: %s", responseMessage.Code, responseMessage.Message)
	}

	return nil
}

// Delete deletes a host by ID
func (h *DefaultHostService) Delete(id int64) error {

	h.logger.Debugf("Delete host with Id: %d", id)

	responseMessage := new(ResponseMessage)

	response, err := h.client.R().
		SetResult(responseMessage).
		SetPathParam("hostId", fmt.Sprintf("%d", id)).
		Delete("/configuration/hosts/{hostId}")

	h.logger.Debugf("Response from delete host: %s", response.String())

	if err != nil {
		return errors.Wrapf(err, "error during delete host request: %s", responseMessage.Message)
	}

	if response.IsError() && response.StatusCode() != http.StatusNotFound {
		return errors.Errorf("delete host failed with status code: %d and message: %s", responseMessage.Code, responseMessage.Message)
	}

	return nil
}

// Find retrieves hosts based on specific criteria
func (h *DefaultHostService) Find(opts *ListOptions) (*ListResponse[HostFindResult], error) {
	if opts == nil {
		opts = &ListOptions{}
	}

	h.logger.Debugf("Find hosts with options: %+v", opts)

	listReponseHost := new(ListResponse[HostFindResult])

	response, err := h.client.R().
		SetResult(listReponseHost).
		SetQueryParams(opts.GetQueryParams()).
		Get("/configuration/hosts")

	h.logger.Debugf("Response from list hosts: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during list hosts request")
	}

	if response.IsError() {
		return nil, errors.Errorf("list hosts failed with status code: %d", response.StatusCode())
	}

	return listReponseHost, nil
}

// Get retrieves a host by its ID
func (h *DefaultHostService) Get(id int64) (*HostFindResult, error) {
	h.logger.Debugf("Get host with Id: %d", id)

	hostListResponse, err := h.Find(&ListOptions{
		Search: map[string]interface{}{
			"id": id,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "error during find host by ID")
	}

	if hostListResponse.Meta.Total == 1 {
		return &hostListResponse.Result[0], nil
	}

	return nil, nil
}

// GetByName retrieves a host by its name
func (h *DefaultHostService) GetByName(name string) (*HostFindResult, error) {
	h.logger.Debugf("Get host with name: %s", name)

	hostListResponse, err := h.Find(&ListOptions{
		Search: map[string]interface{}{
			"name": name,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "error during find host by name")
	}

	if hostListResponse.Meta.Total == 1 {
		return &hostListResponse.Result[0], nil
	}

	return nil, nil
}

// GetQueryParams converts HostListOptions into a map of query parameters
func (h HostListOptions) GetQueryParams() map[string]string {
	params := make(map[string]string)

	if h.ShowService != nil {
		params["show_service"] = fmt.Sprintf("%t", *h.ShowService)
	}

	if h.Search != nil {
		b, err := json.Marshal(h.Search)
		if err != nil {
			panic(fmt.Sprintf("failed to marshal search parameters: %s", err.Error()))
		}

		params["search"] = string(b)
	}

	if h.Page > 0 {
		params["page"] = fmt.Sprintf("%d", h.Page)
	}

	if h.Limit > 0 {
		params["limit"] = fmt.Sprintf("%d", h.Limit)
	}

	if h.SortBy != nil {
		b, err := json.Marshal(h.SortBy)
		if err != nil {
			panic(fmt.Sprintf("failed to marshal sort_by parameters: %s", err.Error()))
		}

		params["sort_by"] = string(b)
	}

	return params
}
