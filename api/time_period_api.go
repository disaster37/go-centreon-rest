package api

import (
	"fmt"
	"net/http"

	"emperror.dev/errors"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// TimePeriodCreateOrUpdateRequest represents the payload to create or update a time period in Centreon.
type TimePeriodService interface {

	// Create creates a new time period.
	Create(timePeriod *TimePeriodCreateOrUpdateRequest) (timePeriodResponse *TimePeriodResponse, err error)

	// Update updates an existing time period by its ID.
	Update(id int64, timePeriod *TimePeriodCreateOrUpdateRequest) (err error)

	// Delete deletes a time period by its ID.
	Delete(id int64) (err error)

	// Get retrieves a time period by its ID.
	Get(id int64) (timePeriodResponse *TimePeriodResponse, err error)

	// GetByName retrieves a time period by its Name.
	// It uses the List method to get the time period.
	GetByName(name string) (timePeriodResponse *TimePeriodResponse, err error)

	// List retrieves a list of time periods based on the provided options.
	List(opts *ListOptions) (timePeriodListResponse *ListResponse[TimePeriodResponse], err error)
}

type DefaultTimePeriodService struct {
	client *resty.Client
	logger *logrus.Entry
}

// NewTimePeriodService creates a new instance of DefaultTimePeriodService
func NewTimePeriodService(client *resty.Client, logger *logrus.Entry) TimePeriodService {
	return &DefaultTimePeriodService{
		client: client,
		logger: logger.WithField("service", "timePeriod"),
	}
}

// Create creates a new time period.
func (h *DefaultTimePeriodService) Create(timePeriod *TimePeriodCreateOrUpdateRequest) (timePeriodResponse *TimePeriodResponse, err error) {
	h.logger.Debugf("Create time period: %+v", timePeriod)

	if timePeriod.Templates == nil {
		timePeriod.Templates = []int64{}
	}

	if timePeriod.Exceptions == nil {
		timePeriod.Exceptions = []TimePeriodException{}
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(timePeriod); err != nil {
		return nil, errors.Wrap(err, "validation error on create time period")
	}

	timePeriodResponse = new(TimePeriodResponse)

	response, err := h.client.R().
		SetBody(timePeriod).
		SetResult(timePeriodResponse).
		Post("/configuration/timeperiods")

	h.logger.Debugf("Response from create time period: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during create time period request")
	}

	if response.IsError() {
		return nil, errors.Errorf("create host severity failed with status code: %d and message: %s", response.StatusCode(), response.String())
	}

	return timePeriodResponse, nil
}

// Delete deletes a time period by its ID.
func (h *DefaultTimePeriodService) Delete(id int64) (err error) {
	h.logger.Debugf("Delete time period with id: %d", id)

	response, err := h.client.R().
		SetPathParam("id", fmt.Sprintf("%d", id)).
		Delete("/configuration/timeperiods/{id}")

	h.logger.Debugf("Response from delete time period: %s", response.String())

	if err != nil {
		return errors.Wrapf(err, "error during delete time period request: %s", response.String())
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil
		}
		return errors.Errorf("delete time perdio failed with status code: %d and message: %s", response.StatusCode(), response.String())
	}

	return nil
}

// Update updates an existing time period by its ID.
func (h *DefaultTimePeriodService) Update(id int64, timePeriod *TimePeriodCreateOrUpdateRequest) (err error) {
	h.logger.Debugf("Update time period with id: %d, TimePeriod: %+v", id, timePeriod)

	if timePeriod.Templates == nil {
		timePeriod.Templates = []int64{}
	}

	if timePeriod.Exceptions == nil {
		timePeriod.Exceptions = []TimePeriodException{}
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(timePeriod); err != nil {
		return errors.Wrap(err, "validation error on update time period")
	}

	response, err := h.client.R().
		SetBody(timePeriod).
		SetPathParam("id", fmt.Sprintf("%d", id)).
		Put("/configuration/timeperiods/{id}")

	h.logger.Debugf("Response from update time period: %s", response.String())

	if err != nil {
		return errors.Wrapf(err, "error during update time period request: %s", response.String())
	}

	if response.IsError() {
		return errors.Errorf("update time period failed with status code: %d and message: %s", response.StatusCode(), response.String())
	}

	return nil
}

// Get retrieves a time period by its ID.
func (h *DefaultTimePeriodService) Get(id int64) (timePeriodResponse *TimePeriodResponse, err error) {
	h.logger.Debugf("Get time period with id: %d", id)

	timePeriodResponse = new(TimePeriodResponse)

	response, err := h.client.R().
		SetResult(timePeriodResponse).
		SetPathParam("id", fmt.Sprintf("%d", id)).
		Get("/configuration/timeperiods/{id}")

	h.logger.Debugf("Response from get time period: %s", response.String())

	if err != nil {
		return nil, errors.Wrapf(err, "error during get time period request: %s", response.String())
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil, nil
		}
		return nil, errors.Errorf("get time period failed with status code: %d and message: %s", response.StatusCode(), response.String())
	}

	return timePeriodResponse, nil
}

// GetByName retrieves a time period by its Name.
func (h *DefaultTimePeriodService) GetByName(name string) (timePeriodResponse *TimePeriodResponse, err error) {
	h.logger.Debugf("Get time period with name: %s", name)

	listOptions := &ListOptions{
		Search: map[string]interface{}{
			"name": name,
		},
	}

	timePeriodListResponse, err := h.List(listOptions)
	if err != nil {
		return nil, errors.Wrap(err, "error during list time periods request")
	}

	if timePeriodListResponse.Meta.Total == 1 {
		return &timePeriodListResponse.Result[0], nil
	}

	return nil, nil
}

// List retrieves a list of time periods based on the provided options.
func (h *DefaultTimePeriodService) List(opts *ListOptions) (timePeriodListResponse *ListResponse[TimePeriodResponse], err error) {

	if opts == nil {
		opts = &ListOptions{}
	}

	h.logger.Debugf("List time periods with options: %+v", opts)

	timePeriodListResponse = new(ListResponse[TimePeriodResponse])

	response, err := h.client.R().
		SetQueryParams(opts.GetQueryParams()).
		SetResult(timePeriodListResponse).
		Get("/configuration/timeperiods")

	h.logger.Debugf("Response from list time periods: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during list time periods request")
	}

	if response.IsError() {
		return nil, errors.Errorf("list time periods failed with status code: %d and message: %s", response.StatusCode(), response.String())
	}

	return timePeriodListResponse, nil
}
