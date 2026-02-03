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

// HostSeverityService defines the interface for managing host severities in Centreon.
type HostSeverityService interface {
	// Create creates a new host severity.
	Create(hostSeverity *HostSeverityCreateOrUpdateRequest) (hostSeverityResponse *HostSeverityResponse, err error)

	// Update updates an existing host severity by its ID.
	Update(id int64, hostSeverity *HostSeverityCreateOrUpdateRequest) (err error)

	//Delete deletes a host severity by its ID.
	Delete(id int64) (err error)

	// Get retrieves a host severity by its ID.
	Get(id int64) (hostSeverityResponse *HostSeverityResponse, err error)

	// GetByName retrieves a host severity by its Name.
	GetByName(name string) (hostSeverityResponse *HostSeverityResponse, err error)

	// List retrieves a list of host severities based on the provided options.
	List(opts *ListOptions) (hostSeverityListResponse *ListResponse[HostSeverityResponse], err error)
}

// DefaultHostSeverityService implements HostSeverityService
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
func (s *DefaultHostSeverityService) Create(hostSeverity *HostSeverityCreateOrUpdateRequest) (hostSeverityResponse *HostSeverityResponse, err error) {
	s.logger.Debugf("Create host severity: %+v", hostSeverity)

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(hostSeverity); err != nil {
		return nil, errors.Wrap(err, "validation error on create host severity")
	}

	hostSeverityResponse = new(HostSeverityResponse)

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
func (s *DefaultHostSeverityService) Delete(id int64) (err error) {
	s.logger.Debugf("Delete host severity with Id: %d", id)

	if id <= 0 {
		return errors.New("invalid host severity id: 0")
	}

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
func (s *DefaultHostSeverityService) Get(id int64) (hostSeverityResponse *HostSeverityResponse, err error) {
	s.logger.Debugf("Get host severity with Id: %d", id)

	if id <= 0 {
		return nil, errors.New("invalid host severity id: 0")
	}

	hostSeverityResponse = new(HostSeverityResponse)

	response, err := s.client.R().
		SetResult(hostSeverityResponse).
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

	return hostSeverityResponse, nil
}

// GetByName retrieves a host severity by its Name.
func (s *DefaultHostSeverityService) GetByName(name string) (hostSeverityResponse *HostSeverityResponse, err error) {
	s.logger.Debugf("Get host severity with name: %s", name)

	if strings.TrimSpace(name) == "" {
		return nil, errors.New("host severity name cannot be empty")
	}

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
func (s *DefaultHostSeverityService) List(opts *ListOptions) (hostSeverityListResponse *ListResponse[HostSeverityResponse], err error) {
	if opts == nil {
		opts = &ListOptions{}
	}

	s.logger.Debugf("List host severities with options: %+v", opts)

	hostSeverityListResponse = new(ListResponse[HostSeverityResponse])

	response, err := s.client.R().
		SetQueryParams(opts.GetQueryParams()).
		SetResult(hostSeverityListResponse).
		Get("/configuration/hosts/severities")

	s.logger.Debugf("Response from list host severities: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during list host severities request")
	}

	if response.IsError() {
		return nil, errors.Errorf("failed to list host severities, status code: %d, response: %s", response.StatusCode(), response.String())
	}

	return hostSeverityListResponse, nil
}

// Update updates an existing host severity by its ID.
func (h *DefaultHostSeverityService) Update(id int64, hostSeverity *HostSeverityCreateOrUpdateRequest) (err error) {
	h.logger.Debugf("Update host severity with Id: %d, Data: %+v", id, hostSeverity)

	if id <= 0 {
		return errors.New("invalid host severity id: 0")
	}

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
