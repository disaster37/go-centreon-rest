package api

import (
	"fmt"
	"net/http"

	"emperror.dev/errors"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// DowntimeService defines the interface for downtime-related operations
type DowntimeService interface {

	// ListHost retrieves a list of host downtimes
	ListHost() (downtimeResponse *ListResponse[DowntimeResponse], err error)

	// CreateHost creates a downtime for a host
	CreateHost(id int64, downtime *DowntimeCreateHostRequest) (err error)

	// Delete deletes a downtime by its ID
	Delete(downtimeID int64) (err error)

	// Get retrieves a downtime by its ID
	Get(downtimeID int64) (downtimeResponse *DowntimeResponse, err error)

	// ListService retrieves a list of service downtimes
	ListService() (downtimeResponse *ListResponse[DowntimeResponse], err error)

	// CreateService creates a downtime for a service
	CreateService(id int64, downtime *DowntimeCreateServiceRequest) (err error)
}

type DefaultDowntimeService struct {
	client *resty.Client
	logger *logrus.Entry
}

// NewDowntimeService creates a new instance of DefaultDowntimeService
func NewDowntimeService(client *resty.Client, logger *logrus.Entry) DowntimeService {
	return &DefaultDowntimeService{
		client: client,
		logger: logger.WithField("service", "downtime"),
	}
}

// ListHost retrieves a list of host downtimes
func (d *DefaultDowntimeService) ListHost() (downtimeResponse *ListResponse[DowntimeResponse], err error) {
	d.logger.Debug("List host downtimes")

	downtimeResponse = new(ListResponse[DowntimeResponse])
	response, err := d.client.R().
		SetResult(downtimeResponse).
		Get("/monitoring/hosts/downtimes")

	d.logger.Debugf("Response from list host downtimes: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error on list host downtimes")
	}

	if response.IsError() {
		return nil, errors.Errorf("list host failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return downtimeResponse, nil
}

// CreateHost creates a downtime for a host
func (d *DefaultDowntimeService) CreateHost(id int64, downtime *DowntimeCreateHostRequest) (err error) {

	d.logger.Debugf("Create host downtime on %d: %+v", id, downtime)

	if id <= 0 {
		return errors.New("invalid host ID")
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(downtime); err != nil {
		return errors.Wrap(err, "validation error on create host downtime")
	}

	response, err := d.client.R().
		SetBody(downtime).
		SetPathParam("id", fmt.Sprintf("%d", id)).
		Post("/monitoring/hosts/{id}/downtimes")

	d.logger.Debugf("Response from create host downtime: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error when create host downtime")
	}

	if response.IsError() {
		return errors.Errorf("create host downtime failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return nil
}

// Delete deletes a downtime by its ID
func (d *DefaultDowntimeService) Delete(id int64) (err error) {
	d.logger.Debugf("Delete downtime with ID: %d", id)

	if id <= 0 {
		return errors.New("invalid downtime ID")
	}

	response, err := d.client.R().
		SetPathParam("downtime_id", fmt.Sprintf("%d", id)).
		Delete("/monitoring/downtimes/{downtime_id}")

	d.logger.Debugf("Response from delete downtime: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error when delete downtime")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil
		}
		return errors.Errorf("delete downtime failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return nil
}

// Get retrieves a downtime by its ID
func (d *DefaultDowntimeService) Get(id int64) (downtimeResponse *DowntimeResponse, err error) {
	d.logger.Debugf("Get downtime with ID: %d", id)

	if id <= 0 {
		return nil, errors.New("invalid downtime ID")
	}

	downtimeResponse = new(DowntimeResponse)
	response, err := d.client.R().
		SetResult(downtimeResponse).
		SetPathParam("downtime_id", fmt.Sprintf("%d", id)).
		Get("/monitoring/downtimes/{downtime_id}")

	d.logger.Debugf("Response from get downtime: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error when get downtime")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil, nil
		}
		return nil, errors.Errorf("get downtime failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return downtimeResponse, nil
}

// ListService retrieves a list of service downtimes
func (d *DefaultDowntimeService) ListService() (downtimeResponse *ListResponse[DowntimeResponse], err error) {
	d.logger.Debug("List service downtimes")

	downtimeResponse = new(ListResponse[DowntimeResponse])
	response, err := d.client.R().
		SetResult(downtimeResponse).
		Get("/monitoring/services/downtimes")

	d.logger.Debugf("Response from list service downtimes: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error when list service downtime")
	}

	if response.IsError() {
		return nil, errors.Errorf("list service failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return downtimeResponse, nil
}

// CreateService creates a downtime for a service
func (d *DefaultDowntimeService) CreateService(id int64, downtime *DowntimeCreateServiceRequest) (err error) {
	d.logger.Debugf("Create service downtime on %d: %+v", id, downtime)

	if id <= 0 {
		return errors.New("invalid service ID")
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(downtime); err != nil {
		return errors.Wrap(err, "validation error on create service downtime")
	}

	response, err := d.client.R().
		SetBody(downtime).
		SetPathParam("id", fmt.Sprintf("%d", id)).
		Post("/monitoring/services/{id}/downtimes")

	d.logger.Debugf("Response from create service downtime: %s", response.String())

	if err != nil {
		return err
	}

	if response.IsError() {
		return errors.Errorf("create service downtime failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return nil
}
