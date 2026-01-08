package api

import (
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
	Create(host *HostCreateRequest) (hostCreateResponse *HostCreateResponse, err error)

	// Update updates an existing host
	Update(id int64, host *HostUpdateRequest) (err error)

	// Delete deletes a host by ID
	Delete(id int64) (err error)

	// Find retrieves hosts based on specific criteria
	Find(opts *ListOptions) (hostListResponse *ListResponse[HostFindResult], err error)

	// Get retrieves a host by its ID
	Get(id int64) (hostResponse *HostFindResult, err error)

	// GetByName retrieves a host by its Name
	GetByName(name string) (hostResponse *HostFindResult, err error)

	// GetFromRealTime retrieves a host by its ID
	// It use monitoring API to get host details
	GetFromRealTime(id int64) (hostResponse *HostResponse, err error)

	// GetByNameFromRealTime retrieves a host by its name
	// It use monitoring API to get host details
	GetByNameFromRealTime(name string) (hostResponse *HostResponse, err error)

	// ListFromRealTime retrieves hosts with optional filtering and pagination
	// It use monitoring API to list hosts
	ListFromRealTime(opts *HostListOptions) (hostListResponse *ListResponse[HostResponse], err error)

	// CountHostsByStatusFromRealTime retrieves count of hosts by their status
	// It use monitoring API to get host status counts
	CountHostsByStatusFromRealTime() (countHostResponse *CountHostStatusResponse, err error)
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

// Create creates a new host
func (h *DefaultHostService) Create(host *HostCreateRequest) (hostCreateResponse *HostCreateResponse, err error) {

	h.logger.Debugf("Create Host: %+v", host)

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(host); err != nil {
		return nil, errors.Wrap(err, "validation error on create host")
	}

	hostCreateResponse = new(HostCreateResponse)

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
// It will partial update the host with the provided fields
func (h *DefaultHostService) Update(id int64, host *HostUpdateRequest) (err error) {

	h.logger.Debugf("Update host with id: %d, Host: %+v", id, host)

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(host); err != nil {
		return errors.Wrap(err, "validation error on update host")
	}

	response, err := h.client.R().
		SetBody(host).
		SetPathParam("hostId", fmt.Sprintf("%d", id)).
		Patch("/configuration/hosts/{hostId}")

	h.logger.Debugf("Response from update host: %s", response.String())

	if err != nil {
		return errors.Wrapf(err, "error during update host request: %s", response.String())
	}

	if response.IsError() {
		return errors.Errorf("update host failed with status code: %d and message: %s", response.StatusCode(), response.String())
	}

	return nil
}

// Delete deletes a host by ID
func (h *DefaultHostService) Delete(id int64) (err error) {

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
func (h *DefaultHostService) Find(opts *ListOptions) (hostListResponse *ListResponse[HostFindResult], err error) {
	if opts == nil {
		opts = &ListOptions{}
	}

	h.logger.Debugf("Find hosts with options: %+v", opts)

	hostListResponse = new(ListResponse[HostFindResult])

	response, err := h.client.R().
		SetResult(hostListResponse).
		SetQueryParams(opts.GetQueryParams()).
		Get("/configuration/hosts")

	h.logger.Debugf("Response from list hosts: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during list hosts request")
	}

	if response.IsError() {
		return nil, errors.Errorf("list hosts failed with status code: %d", response.StatusCode())
	}

	return hostListResponse, nil
}

// Get retrieves a host by its ID
func (h *DefaultHostService) Get(id int64) (hostResponse *HostFindResult, err error) {
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
func (h *DefaultHostService) GetByName(name string) (hostResponse *HostFindResult, err error) {
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

// GetFromRealTime retrieves a host by its ID
func (h *DefaultHostService) GetFromRealTime(id int64) (hostResponse *HostResponse, err error) {

	h.logger.Debugf("Get host with Id: %d", id)

	hostResponse = new(HostResponse)

	response, err := h.client.R().
		SetResult(hostResponse).
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

	return hostResponse, nil
}

// GetByNameFromRealTime retrieves a host by its name
func (h *DefaultHostService) GetByNameFromRealTime(name string) (hostResponse *HostResponse, err error) {

	h.logger.Debugf("Get host with name: %s", name)

	hostListResponse, err := h.ListFromRealTime(&HostListOptions{
		ListOptions: ListOptions{
			Search: map[string]interface{}{
				"host.name": name,
			},
		},
	})

	if err != nil {
		return nil, errors.Wrap(err, "error during get host by name request")
	}

	if hostListResponse.Meta.Total == 1 {
		return &hostListResponse.Result[0], nil
	}

	return nil, nil
}

// List retrieves hosts with optional filtering and pagination
func (h *DefaultHostService) ListFromRealTime(opts *HostListOptions) (*ListResponse[HostResponse], error) {

	if opts == nil {
		opts = &HostListOptions{}
	}

	h.logger.Debugf("List hosts with options: %+v", opts)

	hostListResponse := new(ListResponse[HostResponse])

	response, err := h.client.R().
		SetResult(hostListResponse).
		SetQueryParams(opts.GetQueryParams()).
		Get("/monitoring/hosts")

	h.logger.Debugf("Response from list hosts: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during list hosts request")
	}

	if response.IsError() {
		return nil, errors.Errorf("list hosts failed with status code: %d", response.StatusCode())
	}

	return hostListResponse, nil

}

// CountHostsByStatusFromRealTime retrieves count of hosts by their status
func (h *DefaultHostService) CountHostsByStatusFromRealTime() (countHostResponse *CountHostStatusResponse, err error) {
	h.logger.Debugf("Count hosts by status from real-time")

	countHostResponse = new(CountHostStatusResponse)

	response, err := h.client.R().
		SetResult(countHostResponse).
		Get("monitoring/hosts/status")

	h.logger.Debugf("Response from count hosts by status: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during count hosts by status request")
	}

	if response.IsError() {
		return nil, errors.Errorf("count hosts by status failed with status code: %d and message: %s", response.StatusCode(), response.String())
	}

	return countHostResponse, nil
}
