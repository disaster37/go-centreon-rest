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

type ServiceGroupService interface {
	// Create creates a new service group.
	Create(serviceGroup *ServiceGroupCreateOrUpdateRequest) (serviceGroupResponse *ServiceGroupResponse, err error)

	// Update updates an existing service group by its ID.
	// Need PR https://github.com/centreon/centreon/pull/9463
	Update(id int64, serviceGroup *ServiceGroupCreateOrUpdateRequest) (err error)

	// Delete deletes a service group by its ID.
	Delete(id int64) (err error)

	// Get retrieves a service group by its ID.
	// Need PR https://github.com/centreon/centreon/pull/9463
	Get(id int64) (serviceGroupResponse *ServiceGroupResponse, err error)

	// GetByName retrieves a service group by its Name.
	// It uses the List method to get the service group.
	GetByName(name string) (serviceGroupResponse *ServiceGroupResponse, err error)

	// List retrieves a list of service groups based on the provided options.
	List(options *ListOptions) (serviceGroupListResponse *ListResponse[ServiceGroupResponse], err error)
}

type DefaultServiceGroupService struct {
	client *resty.Client
	logger *logrus.Entry
}

// NewServiceGroupService creates a new instance of DefaultServiceGroupService
func NewServiceGroupService(client *resty.Client, logger *logrus.Entry) ServiceGroupService {
	return &DefaultServiceGroupService{
		client: client,
		logger: logger.WithField("service", "serviceGroup"),
	}
}

// Create creates a new service group.
func (s *DefaultServiceGroupService) Create(serviceGroup *ServiceGroupCreateOrUpdateRequest) (serviceGroupResponse *ServiceGroupResponse, err error) {
	s.logger.Debugf("Create service group: %+v", serviceGroup)

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(serviceGroup); err != nil {
		return nil, errors.Wrap(err, "validation error on create service group")
	}

	serviceGroupResponse = new(ServiceGroupResponse)

	response, err := s.client.R().
		SetBody(serviceGroup).
		SetResult(serviceGroupResponse).
		Post("/configuration/services/groups")

	s.logger.Debugf("Response from create service group: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during create service group request")
	}

	if response.IsError() {
		return nil, errors.Errorf("failed to create service group, status code: %d, response: %s", response.StatusCode(), response.String())
	}

	return serviceGroupResponse, nil
}

// Update updates an existing service group by its ID.
func (s *DefaultServiceGroupService) Update(id int64, serviceGroup *ServiceGroupCreateOrUpdateRequest) (err error) {
	s.logger.Debugf("Update service group ID %d: %+v", id, serviceGroup)

	if id <= 0 {
		return errors.New("invalid service group ID for update")
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(serviceGroup); err != nil {
		return errors.Wrap(err, "validation error on update service group")
	}

	response, err := s.client.R().
		SetBody(serviceGroup).
		SetPathParam("id", fmt.Sprintf("%d", id)).
		Put("/configuration/services/groups/{id}")

	s.logger.Debugf("Response from update service group: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during update service group request")
	}

	if response.IsError() {
		return errors.Errorf("failed to update service group, status code: %d, response: %s", response.StatusCode(), response.String())
	}

	return nil
}

// Delete deletes a service group by its ID.
func (s *DefaultServiceGroupService) Delete(id int64) (err error) {
	s.logger.Debugf("Delete service group with Id: %d", id)

	if id <= 0 {
		return errors.New("invalid service group id: 0")
	}

	response, err := s.client.R().
		SetPathParam("id", fmt.Sprintf("%d", id)).
		Delete("/configuration/services/groups/{id}")

	s.logger.Debugf("Response from delete service group: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during delete service group request")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil
		}
		return errors.Errorf("failed to delete service group, status code: %d, response: %s", response.StatusCode(), response.String())
	}

	return nil
}

// Get retrieves a service group by its ID.
func (s *DefaultServiceGroupService) Get(id int64) (serviceGroupResponse *ServiceGroupResponse, err error) {
	s.logger.Debugf("Get service group with Id: %d", id)

	if id <= 0 {
		return nil, errors.New("invalid service group id: 0")
	}

	serviceGroupResponse = new(ServiceGroupResponse)

	response, err := s.client.R().
		SetPathParam("id", fmt.Sprintf("%d", id)).
		SetResult(serviceGroupResponse).
		Get("/configuration/services/groups/{id}")

	s.logger.Debugf("Response from get service group: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during get service group request")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil, nil
		}
		return nil, errors.Errorf("failed to get service group, status code: %d, response: %s", response.StatusCode(), response.String())
	}

	return serviceGroupResponse, nil
}

// GetByName retrieves a service group by its Name.
// It uses the List method to get the service group.
func (s *DefaultServiceGroupService) GetByName(name string) (serviceGroupResponse *ServiceGroupResponse, err error) {
	s.logger.Debugf("Get service group with Name: %s", name)

	if strings.TrimSpace(name) == "" {
		return nil, errors.New("invalid service group name: empty")
	}

	serviceGroupListResponse, err := s.List(&ListOptions{
		Search: map[string]interface{}{
			"name": name,
		},
	})

	if err != nil {
		return nil, errors.Wrapf(err, "error during list service groups to get by name %s", name)
	}

	if serviceGroupListResponse.Meta.Total == 1 {
		return &serviceGroupListResponse.Result[0], nil
	}

	return nil, nil
}

// List retrieves a list of service groups based on the provided options.
func (s *DefaultServiceGroupService) List(options *ListOptions) (serviceGroupListResponse *ListResponse[ServiceGroupResponse], err error) {

	if options == nil {
		options = &ListOptions{}
	}

	s.logger.Debugf("List service groups with options: %+v", options)

	serviceGroupListResponse = new(ListResponse[ServiceGroupResponse])

	reponse, err := s.client.R().
		SetResult(serviceGroupListResponse).
		SetQueryParams(options.GetQueryParams()).
		Get("/configuration/services/groups")

	s.logger.Debugf("Response from list service groups: %s", reponse.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during list service groups request")
	}

	if reponse.IsError() {
		return nil, errors.Errorf("failed to list service groups, status code: %d, response: %s", reponse.StatusCode(), reponse.String())
	}

	return serviceGroupListResponse, nil
}
