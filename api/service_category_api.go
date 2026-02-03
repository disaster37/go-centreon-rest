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

type ServiceCategoryService interface {
	// Create creates a new service category
	Create(serviceCategory *ServiceCategoryCreateOrUpdateRequest) (serviceCategoryResponse *ServiceCategoryResponse, err error)

	// Update updates an existing service category by its ID
	// Need PR https://github.com/centreon/centreon/pull/9472
	Update(id int64, serviceCategory *ServiceCategoryCreateOrUpdateRequest) (err error)

	// Delete deletes a service category by its ID
	Delete(id int64) (err error)

	// Get retrieves a service category by its ID
	// Need PR https://github.com/centreon/centreon/pull/9472
	Get(id int64) (serviceCategoryResponse *ServiceCategoryResponse, err error)

	// GetByName retrieves a service category by its Name
	// It uses the List method to get the service category
	GetByName(name string) (serviceCategoryResponse *ServiceCategoryResponse, err error)

	// List retrieves a list of service categories based on the provided options
	List(options *ListOptions) (serviceCategoryListResponse *ListResponse[ServiceCategoryResponse], err error)
}

type DefaultServiceCategoryService struct {
	client *resty.Client
	logger *logrus.Entry
}

// NewServiceCategoryService creates a new instance of DefaultServiceCategoryService
func NewServiceCategoryService(client *resty.Client, logger *logrus.Entry) ServiceCategoryService {
	return &DefaultServiceCategoryService{
		client: client,
		logger: logger.WithField("service", "serviceCategory"),
	}
}

// Create creates a new service category
func (s *DefaultServiceCategoryService) Create(serviceCategory *ServiceCategoryCreateOrUpdateRequest) (serviceCategoryResponse *ServiceCategoryResponse, err error) {
	s.logger.Debugf("Create service category: %+v", serviceCategory)

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(serviceCategory); err != nil {
		return nil, errors.Wrap(err, "validation error on create service category")
	}

	serviceCategoryResponse = new(ServiceCategoryResponse)

	response, err := s.client.R().
		SetBody(serviceCategory).
		SetResult(serviceCategoryResponse).
		Post("/configuration/services/categories")

	s.logger.Debugf("Response from create service category: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during create service category request")
	}

	if response.IsError() {
		return nil, errors.Errorf("error response from create service category request: %s", response.String())
	}

	return serviceCategoryResponse, nil
}

// Update updates an existing service category by its ID
func (s *DefaultServiceCategoryService) Update(id int64, serviceCategory *ServiceCategoryCreateOrUpdateRequest) (err error) {
	s.logger.Debugf("Update service category ID %d: %+v", id, serviceCategory)

	if id <= 0 {
		return errors.New("invalid service category ID for update")
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(serviceCategory); err != nil {
		return errors.Wrap(err, "validation error on update service category")
	}

	response, err := s.client.R().
		SetBody(serviceCategory).
		SetPathParam("id", fmt.Sprintf("%d", id)).
		Put("/configuration/services/categories/{id}")

	s.logger.Debugf("Response from update service category: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during update service category request")
	}

	if response.IsError() {
		return errors.Errorf("error response from update service category request: %s", response.String())
	}

	return nil
}

// Delete deletes a service category by its ID
func (s *DefaultServiceCategoryService) Delete(id int64) (err error) {
	s.logger.Debugf("Delete service category ID %d", id)

	if id <= 0 {
		return errors.New("invalid service category ID for delete")
	}

	response, err := s.client.R().
		SetPathParam("id", fmt.Sprintf("%d", id)).
		Delete("/configuration/services/categories/{id}")

	s.logger.Debugf("Response from delete service category: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during delete service category request")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil
		}
		return errors.Errorf("error response from delete service category request: %s", response.String())
	}

	return nil
}

// Get retrieves a service category by its ID
func (s *DefaultServiceCategoryService) Get(id int64) (serviceCategoryResponse *ServiceCategoryResponse, err error) {
	s.logger.Debugf("Get service category ID %d", id)

	if id <= 0 {
		return nil, errors.New("invalid service category ID for get")
	}

	serviceCategoryResponse = new(ServiceCategoryResponse)

	response, err := s.client.R().
		SetPathParam("id", fmt.Sprintf("%d", id)).
		SetResult(serviceCategoryResponse).
		Get("/configuration/services/categories/{id}")

	s.logger.Debugf("Response from get service category: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during get service category request")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil, nil
		}
		return nil, errors.Errorf("error response from get service category request: %s", response.String())
	}

	return serviceCategoryResponse, nil
}

// GetByName retrieves a service category by its Name
func (s *DefaultServiceCategoryService) GetByName(name string) (serviceCategoryResponse *ServiceCategoryResponse, err error) {
	s.logger.Debugf("Get service category by Name: %s", name)

	if strings.TrimSpace(name) == "" {
		return nil, errors.New("service category name cannot be empty")
	}

	listResponse, err := s.List(&ListOptions{
		Search: map[string]interface{}{
			"name": name,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list service categories for name %s", name)
	}

	if listResponse.Meta.Total == 1 {
		return &listResponse.Result[0], nil
	}

	return nil, nil
}

// List retrieves a list of service categories based on the provided options
func (s *DefaultServiceCategoryService) List(options *ListOptions) (serviceCategoryListResponse *ListResponse[ServiceCategoryResponse], err error) {

	if options == nil {
		options = &ListOptions{}
	}

	s.logger.Debugf("List service categories with options: %+v", options)

	serviceCategoryListResponse = new(ListResponse[ServiceCategoryResponse])

	reponse, err := s.client.R().
		SetResult(serviceCategoryListResponse).
		SetQueryParams(options.GetQueryParams()).
		Get("/configuration/services/categories")

	s.logger.Debugf("Response from list service categories: %s", reponse.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during list service categories request")
	}

	if reponse.IsError() {
		return nil, errors.Errorf("error response from list service categories request: %s", reponse.String())
	}

	return serviceCategoryListResponse, nil
}
