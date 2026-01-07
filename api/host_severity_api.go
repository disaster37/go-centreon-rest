package api

import (
	"fmt"
	"net/http"

	"emperror.dev/errors"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// HostSeverityService defines the interface for managing host severities in Centreon.
type HostSeverityService interface {
	// Create creates a new host severity.
	Create(hostSeverity *HostSeverityUpdateRequest) (*HostSeverityResponse, error)

	// Update updates an existing host severity by its ID.
	Update(id int64, hostSeverity *HostSeverityUpdateRequest) error

	//Delete deletes a host severity by its ID.
	Delete(id int64) error

	// Get retrieves a host severity by its ID.
	Get(id int64) (*HostSeverityResponse, error)

	// GetByName retrieves a host severity by its Name.
	GetByName(name string) (*HostSeverityResponse, error)

	// List retrieves a list of host severities based on the provided options.
	List(opts *ListOptions) (*ListResponse[HostSeverityResponse], error)
}

type DefaultHostSeverityService struct {
	client *resty.Client
	logger *logrus.Entry
}

// NewHostSeverityService creates a new instance of DefaultHostSeverityService
func NewHostSeverityService(client *resty.Client, logger *logrus.Entry) HostSeverityService {
	return &DefaultHostSeverityService{
		client: client,
		logger: logger.WithField("service", "hostSeverity"),
	}
}

// Create creates a new host severity.
func (s *DefaultHostSeverityService) Create(hostSeverity *HostSeverityUpdateRequest) (*HostSeverityResponse, error) {
	s.logger.Debugf("Create host severity: %+v", hostSeverity)

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(hostSeverity); err != nil {
		return nil, errors.Wrap(err, "validation error on create host severity")
	}

	hostSeverityResponse := new(HostSeverityResponse)

	response, err := s.client.R().
		SetBody(hostSeverity).
		SetResult(hostSeverityResponse).
		Post("/configuration/hosts/severities")

	s.logger.Debugf("Response from create host severity: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during create host severity request")
	}

	if response.IsError() {
		return nil, errors.Errorf("failed to create host severity, status code: %d, response: %s", response.StatusCode(), response.String())
	}

	return hostSeverityResponse, nil
}

// Delete deletes a host severity by its ID.
func (s *DefaultHostSeverityService) Delete(id int64) error {
	s.logger.Debugf("Delete host severity with Id: %d", id)

	response, err := s.client.R().
		SetPathParam("severityId", fmt.Sprintf("%d", id)).
		Delete("/configuration/hosts/severities/{severityId}")

	s.logger.Debugf("Response from delete host severity: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during delete host severity request")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil
		}
		return errors.Errorf("failed to delete host severity, status code: %d, response: %s", response.StatusCode(), response.String())
	}

	return nil
}

// Get retrieves a host severity by its ID.
func (s *DefaultHostSeverityService) Get(id int64) (*HostSeverityResponse, error) {
	s.logger.Debugf("Get host severity with Id: %d", id)

	responseStruct := new(HostSeverityResponse)

	response, err := s.client.R().
		SetResult(responseStruct).
		SetPathParam("severityId", fmt.Sprintf("%d", id)).
		Get("/configuration/hosts/severities/{severityId}")

	s.logger.Debugf("Response from get host severity: %s", response.String())

	if err != nil {
		return nil, errors.Wrapf(err, "failed to get host severity with id %d", id)
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil, nil
		}
		return nil, errors.Errorf("get host severity failed with status code: %d", response.StatusCode())
	}

	return responseStruct, nil
}

// GetByName retrieves a host severity by its Name.
func (s *DefaultHostSeverityService) GetByName(name string) (*HostSeverityResponse, error) {
	s.logger.Debugf("Get host severity with name: %s", name)

	listResponse, err := s.List(&ListOptions{
		Search: map[string]interface{}{
			"name": name,
		},
	})

	if err != nil {
		return nil, errors.Wrapf(err, "failed to get host severity with name %s", name)
	}

	if listResponse.Meta.Total == 1 {
		return &listResponse.Result[0], nil
	}

	return nil, nil
}

// List retrieves a list of host severities based on the provided options.
func (s *DefaultHostSeverityService) List(opts *ListOptions) (*ListResponse[HostSeverityResponse], error) {
	if opts == nil {
		opts = &ListOptions{}
	}

	s.logger.Debugf("List host severities with options: %+v", opts)

	listResponse := new(ListResponse[HostSeverityResponse])

	response, err := s.client.R().
		SetQueryParams(opts.GetQueryParams()).
		SetResult(listResponse).
		Get("/configuration/hosts/severities")

	s.logger.Debugf("Response from list host severities: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during list host severities request")
	}

	if response.IsError() {
		return nil, errors.Errorf("failed to list host severities, status code: %d, response: %s", response.StatusCode(), response.String())
	}

	return listResponse, nil
}

// Update updates an existing host severity by its ID.
func (h *DefaultHostSeverityService) Update(id int64, hostSeverity *HostSeverityUpdateRequest) error {
	h.logger.Debugf("Update host severity with Id: %d, Data: %+v", id, hostSeverity)

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(hostSeverity); err != nil {
		return errors.Wrap(err, "validation error on update host severity")
	}

	response, err := h.client.R().
		SetBody(hostSeverity).
		SetPathParam("severityId", fmt.Sprintf("%d", id)).
		Put("/configuration/hosts/severities/{severityId}")

	h.logger.Debugf("Response from update host severity: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during update host severity request")
	}

	if response.IsError() {
		return errors.Errorf("failed to update host severity, status code: %d, response: %s", response.StatusCode(), response.String())
	}

	return nil
}
