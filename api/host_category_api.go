package api

import (
	"fmt"
	"net/http"

	"emperror.dev/errors"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type HostCategoryService interface {
	// Create creates a new host category
	Create(hostCategory *HostCategoryUpdateRequest) (*HostCategoryResponse, error)

	// Update updates an existing host category
	Update(id int64, hostCategory *HostCategoryUpdateRequest) error

	// Delete deletes a host category by ID
	Delete(id int64) error

	// List retrieves host categories with optional filtering and pagination
	List(opts *ListOptions) (*ListResponse[HostCategoryResponse], error)

	Get(id int64) (*HostCategoryResponse, error)

	GetByName(name string) (*HostCategoryResponse, error)

	ListFromRealTime(opts *ListOptions) (*ListResponse[HostCategoryResponse], error)
}

type DefaultHostCategoryService struct {
	client *resty.Client
	logger *logrus.Entry
}

// NewHostCategoryService creates a new instance of DefaultHostCategoryService
func NewHostCategoryService(client *resty.Client, logger *logrus.Entry) HostCategoryService {
	return &DefaultHostCategoryService{
		client: client,
		logger: logger.WithField("service", "hostCategory"),
	}
}

// Create creates a new host category
func (h *DefaultHostCategoryService) Create(hostCategory *HostCategoryUpdateRequest) (*HostCategoryResponse, error) {
	h.logger.Debugf("Create host category: %+v", hostCategory)

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(hostCategory); err != nil {
		return nil, errors.Wrap(err, "validation error on create host category")
	}

	hostCategoryResponse := new(HostCategoryResponse)

	response, err := h.client.R().
		SetBody(hostCategory).
		SetResult(hostCategoryResponse).
		Post("/configuration/hosts/categories")

	h.logger.Debugf("Response from create host category: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during create host category request")
	}

	if response.IsError() {
		return nil, errors.Errorf("failed to create host category, status code: %d, response: %s", response.StatusCode(), response.String())
	}

	return hostCategoryResponse, nil
}

// Update updates an existing host category
func (h *DefaultHostCategoryService) Update(id int64, hostCategory *HostCategoryUpdateRequest) error {
	h.logger.Debugf("Update host category with Id %d: %+v", id, hostCategory)

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(hostCategory); err != nil {
		return errors.Wrap(err, "validation error on update host category")
	}

	response, err := h.client.R().
		SetBody(hostCategory).
		SetPathParam("categoryId", fmt.Sprintf("%d", id)).
		Put("/configuration/hosts/categories/{categoryId}")

	h.logger.Debugf("Response from update host category: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during update host category request")
	}

	if response.IsError() {
		return errors.Errorf("failed to update host category, status code: %d, response: %s", response.StatusCode(), response.String())
	}

	return nil
}

// Delete deletes a host category by ID
func (h *DefaultHostCategoryService) Delete(id int64) error {
	h.logger.Debugf("Delete host category with Id: %d", id)

	response, err := h.client.R().
		SetPathParam("categoryId", fmt.Sprintf("%d", id)).
		Delete("/configuration/hosts/categories/{categoryId}")

	h.logger.Debugf("Response from delete host category: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during delete host category request")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil
		}
		return errors.Errorf("failed to delete host category, status code: %d, response: %s", response.StatusCode(), response.String())
	}

	return nil
}

// List retrieves host categories with optional filtering and pagination
func (h *DefaultHostCategoryService) List(opts *ListOptions) (*ListResponse[HostCategoryResponse], error) {
	if opts == nil {
		opts = &ListOptions{}
	}

	h.logger.Debugf("List host categories with optionq: %+v", opts)

	listResponse := new(ListResponse[HostCategoryResponse])

	response, err := h.client.R().
		SetQueryParams(opts.GetQueryParams()).
		SetResult(listResponse).
		Get("/configuration/hosts/categories")

	h.logger.Debugf("Response from list host categories: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during list host categories request")
	}

	if response.IsError() {
		return nil, errors.Errorf("failed to list host categories, status code: %d, response: %s", response.StatusCode(), response.String())
	}

	return listResponse, nil
}

// Get retrieves a host category by its ID
func (h *DefaultHostCategoryService) Get(id int64) (*HostCategoryResponse, error) {
	h.logger.Debugf("Get host category with Id: %d", id)

	hostCategoryReponse := new(HostCategoryResponse)

	reponse, err := h.client.R().
		SetResult(hostCategoryReponse).
		SetPathParam("categoryId", fmt.Sprintf("%d", id)).
		Get("/configuration/hosts/categories/{categoryId}")

	h.logger.Debugf("Response from get host category: %s", reponse.String())

	if err != nil {
		return nil, errors.Wrapf(err, "failed to get host category with id %d", id)
	}

	if reponse.IsError() {
		if reponse.StatusCode() == http.StatusNotFound {
			return nil, nil
		}
		return nil, errors.Errorf("get host category failed with status code: %d", reponse.StatusCode())
	}

	return hostCategoryReponse, nil
}

// GetByName retrieves a host category by its Name
func (h *DefaultHostCategoryService) GetByName(name string) (*HostCategoryResponse, error) {
	h.logger.Debugf("Get host category with name: %s", name)

	listResponse, err := h.List(&ListOptions{
		Search: map[string]interface{}{
			"name": name,
		},
	})

	if err != nil {
		return nil, errors.Wrapf(err, "failed to get host category with name %s", name)
	}

	if listResponse.Meta.Total == 1 {
		return &listResponse.Result[0], nil
	}

	return nil, nil
}

// ListFromRealTime retrieves host categories from real-time monitoring data
func (h *DefaultHostCategoryService) ListFromRealTime(opts *ListOptions) (*ListResponse[HostCategoryResponse], error) {
	if opts == nil {
		opts = &ListOptions{}
	}

	h.logger.Debugf("List real-time host categories with options: %+v", opts)

	listResponse := new(ListResponse[HostCategoryResponse])

	response, err := h.client.R().
		SetQueryParams(opts.GetQueryParams()).
		SetResult(listResponse).
		Get("/monitoring/hosts/categories")

	h.logger.Debugf("Response from list real-time host categories: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during list real-time host categories request")
	}

	if response.IsError() {
		return nil, errors.Errorf("failed to list real-time host categories, status code: %d, response: %s", response.StatusCode(), response.String())
	}

	return listResponse, nil
}
