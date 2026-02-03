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

// ServiceSeverityService defines the interface for managing service severities.
type ServiceSeverityService interface {

	// Create creates a new service severity.
	Create(serviceSeverity *ServiceSeverityCreateOrUpdateRequest) (serviceSeverityResponse *ServiceSeverityResponse, err error)

	// Update updates an existing service severity by its ID.
	Update(id int64, serviceSeverity *ServiceSeverityCreateOrUpdateRequest) (err error)

	// Delete deletes a service severity by its ID.
	Delete(id int64) (err error)

	// Get retrieves a service severity by its ID.
	// Need PR https://github.com/centreon/centreon/issues/9458
	Get(id int64) (serviceSeverityResponse *ServiceSeverityResponse, err error)

	// GetByName retrieves a service severity by its Name.
	// It uses the List method to get the service severity.
	GetByName(name string) (serviceSeverityResponse *ServiceSeverityResponse, err error)

	// List retrieves a list of service severities based on the provided options.
	List(options *ListOptions) (serviceSeverityListResponse *ListResponse[ServiceSeverityResponse], err error)
}

type DefaultServiceSeverityService struct {
	client *resty.Client
	logger *logrus.Entry
}

// NewServiceSeverityService creates a new instance of DefaultServiceSeverityService
func NewServiceSeverityService(client *resty.Client, logger *logrus.Entry) ServiceSeverityService {
	return &DefaultServiceSeverityService{
		client: client,
		logger: logger.WithField("service", "serviceSeverity"),
	}
}

// Create creates a new service severity.
func (s *DefaultServiceSeverityService) Create(serviceSeverity *ServiceSeverityCreateOrUpdateRequest) (serviceSeverityResponse *ServiceSeverityResponse, err error) {
	s.logger.Debugf("Create service severity: %+v", serviceSeverity)

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(serviceSeverity); err != nil {
		return nil, errors.Wrap(err, "validation error on create service severity")
	}

	serviceSeverityResponse = new(ServiceSeverityResponse)

	response, err := s.client.R().
		SetBody(serviceSeverity).
		SetResult(serviceSeverityResponse).
		Post("/configuration/services/severities")

	s.logger.Debugf("Response from create service severity: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during create service severity request")
	}

	if response.IsError() {
		return nil, errors.Errorf("error response from create service severity request with status: %s and message: %s", response.Status(), response.String())
	}

	return serviceSeverityResponse, nil
}

// Update updates an existing service severity by its ID.
func (s *DefaultServiceSeverityService) Update(id int64, serviceSeverity *ServiceSeverityCreateOrUpdateRequest) (err error) {
	s.logger.Debugf("Update service severity ID %d with data: %+v", id, serviceSeverity)

	if id <= 0 {
		return errors.New("invalid service severity ID for update")
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(serviceSeverity); err != nil {
		return errors.Wrap(err, "validation error on update service severity")
	}

	response, err := s.client.R().
		SetBody(serviceSeverity).
		SetPathParam("id", fmt.Sprintf("%d", id)).
		Put("/configuration/services/severities/{id}")

	s.logger.Debugf("Response from update service severity: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during update service severity request")
	}

	if response.IsError() {
		return errors.Errorf("error response from update service severity request with status: %s and message: %s", response.Status(), response.String())
	}

	return nil
}

// Delete deletes a service severity by its ID.
func (s *DefaultServiceSeverityService) Delete(id int64) (err error) {
	s.logger.Debugf("Delete service severity ID %d", id)

	if id <= 0 {
		return errors.New("invalid service severity ID for delete")
	}

	response, err := s.client.R().
		SetPathParam("id", fmt.Sprintf("%d", id)).
		Delete("/configuration/services/severities/{id}")

	s.logger.Debugf("Response from delete service severity: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during delete service severity request")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil
		}
		return errors.Errorf("error response from delete service severity request with status: %s and message: %s", response.Status(), response.String())
	}

	return nil
}

// Get retrieves a service severity by its ID.
func (s *DefaultServiceSeverityService) Get(id int64) (serviceSeverityResponse *ServiceSeverityResponse, err error) {
	s.logger.Debugf("Get service severity ID %d", id)

	if id <= 0 {
		return nil, errors.New("invalid service severity ID for get")
	}

	serviceSeverityResponse = new(ServiceSeverityResponse)

	response, err := s.client.R().
		SetPathParam("id", fmt.Sprintf("%d", id)).
		SetResult(serviceSeverityResponse).
		Get("/configuration/services/severities/{id}")

	s.logger.Debugf("Response from get service severity: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during get service severity request")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil, nil
		}
		return nil, errors.Errorf("error response from get service severity request with status: %s and message: %s", response.Status(), response.String())
	}

	return serviceSeverityResponse, nil
}

// GetByName retrieves a service severity by its Name.
func (s *DefaultServiceSeverityService) GetByName(name string) (serviceSeverityResponse *ServiceSeverityResponse, err error) {
	s.logger.Debugf("Get service severity by Name: %s", name)

	if strings.TrimSpace(name) == "" {
		return nil, errors.New("service severity name cannot be empty")
	}

	options := &ListOptions{
		Search: map[string]any{
			"name": name,
		},
	}

	serviceSeverityListResponse, err := s.List(options)
	if err != nil {
		return nil, errors.Wrap(err, "error during list service severities request")
	}

	if serviceSeverityListResponse.Meta.Total == 1 {
		return &serviceSeverityListResponse.Result[0], nil
	}

	return nil, nil
}

// List retrieves a list of service severities based on the provided options.
func (s *DefaultServiceSeverityService) List(options *ListOptions) (serviceSeverityListResponse *ListResponse[ServiceSeverityResponse], err error) {
	if options == nil {
		options = &ListOptions{}
	}

	s.logger.Debugf("List service severities with options: %+v", options)

	serviceSeverityListResponse = new(ListResponse[ServiceSeverityResponse])

	reponse, err := s.client.R().
		SetResult(serviceSeverityListResponse).
		SetQueryParams(options.GetQueryParams()).
		Get("/configuration/services/severities")

	s.logger.Debugf("Response from list service severities: %s", reponse.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during list service severities request")
	}

	if reponse.IsError() {
		return nil, errors.Errorf("error response from list service severities request with status: %s and message: %s", reponse.Status(), reponse.String())
	}

	return serviceSeverityListResponse, nil
}
