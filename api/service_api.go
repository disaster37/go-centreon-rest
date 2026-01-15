package api

import (
	"fmt"
	"net/http"

	"emperror.dev/errors"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// ServiceService defines the interface for service-related operations
type ServiceService interface {

	// Create a new service
	Create(service *ServiceCreateRequest) (serviceResponse *ServiceResponse, err error)

	// Update an existing service
	// It's a partial update, so only the fields to be updated need to be provided
	Update(id int64, service *ServiceUpdateRequest) (err error)

	// Delete a service by its ID
	Delete(id int64) (err error)

	// Find services with optional filtering, pagination, and sorting
	Find(options *ListOptions) (response *ListResponse[ServiceListResponse], err error)

	// Get a service by its ID or name
	// Nead PR https://github.com/centreon/centreon/pull/9306
	Get(id int64) (response *ServiceResponse, err error)

	// GetByName retrieves a service by its name
	// It uses Find method to get the service
	GetByName(name string) (response *ServiceListResponse, err error)

	// GetServicesByHostID retrieves services associated with a specific host ID
	// It uses Find method to get the services
	GetServicesByHostID(hostID int64) (response *ListResponse[ServiceListResponse], err error)

	// GetServicesByHostName retrieves services associated with a specific host name
	// It uses Find method to get the services
	GetServicesByHostName(hostName string) (response *ListResponse[ServiceListResponse], err error)

	// GetFromRealTime retrieves real-time data for a service by its ID
	GetFromRealTime(hostId int64, serviceId int64) (response *ServiceRealTimeResponse, err error)

	// ListFromRealTime retrieves a list of services with real-time data, with optional filtering, pagination, and sorting
	ListFromRealTime(options *ListOptions) (response *ListResponse[ServiceResponse], err error)
}

// DefaultServiceService implements ServiceService
type DefaultServiceService struct {
	client *resty.Client
	logger *logrus.Entry
}

// NewServiceService creates a new instance of DefaultServiceService
func NewServiceService(client *resty.Client, logger *logrus.Entry) ServiceService {
	return &DefaultServiceService{
		client: client,
		logger: logger.WithField("service", "service"),
	}
}

// Create creates a new service
func (h *DefaultServiceService) Create(service *ServiceCreateRequest) (serviceResponse *ServiceResponse, err error) {
	h.logger.Debugf("Create Service: %+v", service)

	Validator := validator.New()
	if err := Validator.Struct(service); err != nil {
		return nil, errors.Wrap(err, "validation error on create service")
	}

	serviceResponse = new(ServiceResponse)

	response, err := h.client.R().
		SetBody(service).
		SetResult(serviceResponse).
		Post("/configuration/services")

	h.logger.Debugf("Create Service Response: %+v", response)

	if err != nil {
		return nil, errors.Wrap(err, "error during creating service request")
	}

	if response.IsError() {
		return nil, errors.Errorf("error creating service with code %s and message %s", response.Status(), response.String())
	}

	return serviceResponse, nil

}

// Update updates an existing service
func (h *DefaultServiceService) Update(id int64, service *ServiceUpdateRequest) (err error) {
	h.logger.Debugf("Update Service ID %d: %+v", id, service)

	Validator := validator.New()
	if err := Validator.Struct(service); err != nil {
		return errors.Wrap(err, "validation error on update service")
	}

	response, err := h.client.R().
		SetBody(service).
		SetPathParam("id", fmt.Sprintf("%d", id)).
		Patch("/configuration/services/{id}")

	h.logger.Debugf("Update Service Response: %+v", response)

	if err != nil {
		return errors.Wrap(err, "error during updating service request")
	}

	if response.IsError() {
		return errors.Errorf("error updating service with code %s and message %s", response.Status(), response.String())
	}

	return nil
}

// Delete deletes a service by its ID
func (h *DefaultServiceService) Delete(id int64) (err error) {
	h.logger.Debugf("Delete Service ID %d", id)

	response, err := h.client.R().
		SetPathParam("id", fmt.Sprintf("%d", id)).
		Delete("/configuration/services/{id}")

	h.logger.Debugf("Delete Service Response: %+v", response)

	if err != nil {
		return errors.Wrap(err, "error during deleting service request")
	}

	if response.IsError() {

		if response.StatusCode() == http.StatusNotFound {
			return nil
		}
		return errors.Errorf("error deleting service with code %s and message %s", response.Status(), response.String())
	}

	return nil
}

// Find retrieves a list of services based on the provided options.
func (h *DefaultServiceService) Find(options *ListOptions) (response *ListResponse[ServiceListResponse], err error) {
	if options == nil {
		options = &ListOptions{}
	}

	h.logger.Debugf("Find Services with options: %+v", options)

	response = new(ListResponse[ServiceListResponse])

	httpResponse, err := h.client.R().
		SetQueryParams(options.GetQueryParams()).
		SetResult(response).
		Get("/configuration/services")

	h.logger.Debugf("Response from find services: %s", httpResponse.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during find services request")
	}

	if httpResponse.IsError() {
		return nil, errors.Errorf("list services failed with status code %d and message %s", httpResponse.StatusCode(), httpResponse.String())
	}

	return response, nil
}

// Get retrieves a service by its ID
func (h *DefaultServiceService) Get(id int64) (serviceResponse *ServiceResponse, err error) {
	h.logger.Debugf("Get Service by ID: %d", id)

	serviceResponse = new(ServiceResponse)

	response, err := h.client.R().
		SetPathParam("id", fmt.Sprintf("%d", id)).
		SetResult(serviceResponse).
		Get("/configuration/services/{id}")

	h.logger.Debugf("Response from get service: %s", response.String())

	if err != nil {
		return nil, errors.Wrapf(err, "error getting service by ID %d", id)
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil, nil
		}
		return nil, errors.Errorf("get service failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return serviceResponse, nil
}

// GetByName retrieves a service by its name
func (h *DefaultServiceService) GetByName(name string) (response *ServiceListResponse, err error) {
	h.logger.Debugf("Get Service by Name: %s", name)

	options := &ListOptions{
		Search: map[string]interface{}{
			"name": name,
		},
	}

	servicesResponse, err := h.Find(options)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting service by Name %s", name)
	}

	if servicesResponse.Meta.Total == 1 {
		return &servicesResponse.Result[0], nil
	}

	return nil, nil
}

// GetServicesByHostID retrieves services associated with a specific host ID
func (h *DefaultServiceService) GetServicesByHostID(hostID int64) (response *ListResponse[ServiceListResponse], err error) {
	h.logger.Debugf("Get Services by Host ID: %d", hostID)

	options := &ListOptions{
		Search: map[string]interface{}{
			"host.id": hostID,
		},
	}

	servicesResponse, err := h.Find(options)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting services by Host ID %d", hostID)
	}

	return servicesResponse, nil
}

// GetServicesByHostName retrieves services associated with a specific host name
func (h *DefaultServiceService) GetServicesByHostName(hostName string) (response *ListResponse[ServiceListResponse], err error) {
	h.logger.Debugf("Get Services by Host Name: %s", hostName)

	options := &ListOptions{
		Search: map[string]interface{}{
			"host.name": hostName,
		},
	}

	servicesResponse, err := h.Find(options)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting services by Host Name %s", hostName)
	}

	return servicesResponse, nil
}

// GetFromRealTime retrieves real-time data for a service by its ID
func (h *DefaultServiceService) GetFromRealTime(hostId int64, serviceId int64) (serviceResponse *ServiceRealTimeResponse, err error) {
	h.logger.Debugf("Get Service from Real-Time by ID hostId: %d, serviceId: %d", hostId, serviceId)

	serviceResponse = new(ServiceRealTimeResponse)

	response, err := h.client.R().
		SetPathParam("hostId", fmt.Sprintf("%d", hostId)).
		SetPathParam("serviceId", fmt.Sprintf("%d", serviceId)).
		SetResult(serviceResponse).
		Get("/monitoring/hosts/{hostDd}/services/{serviceId}")

	h.logger.Debugf("Response from get service from real-time: %s", response.String())

	if err != nil {
		return nil, errors.Wrapf(err, "error during get service from real-time request")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil, nil
		}
		return nil, errors.Errorf("get service from real-time failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return serviceResponse, nil
}

// ListFromRealTime retrieves a list of services with real-time data
func (h *DefaultServiceService) ListFromRealTime(options *ListOptions) (serviceResponse *ListResponse[ServiceResponse], err error) {
	if options == nil {
		options = &ListOptions{}
	}

	h.logger.Debugf("List Services from Real-Time with options: %+v", options)

	serviceResponse = new(ListResponse[ServiceResponse])

	response, err := h.client.R().
		SetQueryParams(options.GetQueryParams()).
		SetResult(serviceResponse).
		Get("/monitoring/services")

	h.logger.Debugf("Response from list services from real-time: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during list services from real-time request")
	}

	if response.IsError() {
		return nil, errors.Errorf("list services from real-time failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return serviceResponse, nil
}
