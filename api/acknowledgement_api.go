package api

import (
	"fmt"
	"net/http"

	"emperror.dev/errors"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// AcknowledgementInterface defines the interface for acknowledgement-related operations
type AcknowledgementInterface interface {
	// ListHost retrieves a list of host acknowledgements
	ListHost() (acknowledgementResponse *ListResponse[AcknowledgementListHostResponse], err error)

	// CreateHost creates an acknowledgement for a host
	CreateHost(id int64, acknowledgement *AcknowledgementCreateHostRequest) (err error)

	// Delete deletes an acknowledgement for host by its ID
	DeleteHost(hostId int64) (err error)

	// ListService retrieves a list of service acknowledgements
	ListService() (acknowledgementResponse *ListResponse[AcknowledgementListServiceResponse], err error)

	// CreateService creates an acknowledgement for a service
	CreateService(id int64, acknowledgement *AcknowledgementCreateServiceRequest) (err error)

	// DeleteService deletes an acknowledgement for service by its ID
	DeleteService(serviceId int64) (err error)

	// Get retrieves an acknowledgement by its ID
	Get(id int64) (acknowledgementResponse *AcknowledgementResponse, err error)
}

// DefaultAckneledgementService implements AcknowledgementInterface
type DefaultAckneledgementService struct {
	client *resty.Client
	logger *logrus.Entry
}

// NewAcknowledgementService creates a new instance of DefaultAckneledgementService
func NewAcknowledgementService(client *resty.Client, logger *logrus.Entry) AcknowledgementInterface {
	return &DefaultAckneledgementService{
		client: client,
		logger: logger.WithField("service", "acknowledgement"),
	}
}

// Get retrieves an acknowledgement by its ID
func (h *DefaultAckneledgementService) Get(id int64) (acknowledgementResponse *AcknowledgementResponse, err error) {

	h.logger.Debugf("Get acknowledgement with Id: %d", id)

	if id <= 0 {
		return nil, errors.Errorf("invalid acknowledgement id: %d", id)
	}

	acknowledgementResponse = new(AcknowledgementResponse)
	response, err := h.client.R().
		SetResult(acknowledgementResponse).
		SetPathParam("id", fmt.Sprintf("%d", id)).
		Get("/monitoring/acknowledgements/{id}")

	h.logger.Debugf("Response from get acknowledgement: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error on get acknowledgement")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil, nil
		}
		return nil, errors.Errorf("get acknowledgement failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return acknowledgementResponse, nil
}

// ListHost retrieves a list of host acknowledgements
func (h *DefaultAckneledgementService) ListHost() (acknowledgementResponse *ListResponse[AcknowledgementListHostResponse], err error) {
	h.logger.Debug("List host acknowledgements")

	acknowledgementResponse = new(ListResponse[AcknowledgementListHostResponse])
	response, err := h.client.R().
		SetResult(acknowledgementResponse).
		Get("/monitoring/hosts/acknowledgements")

	h.logger.Debugf("Response from list host acknowledgements: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error on list host acknowledgements")
	}

	if response.IsError() {
		return nil, errors.Errorf("list host acknowledgements failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return acknowledgementResponse, nil
}

// CreateHost creates an acknowledgement for a host
func (h *DefaultAckneledgementService) CreateHost(id int64, acknowledgement *AcknowledgementCreateHostRequest) (err error) {

	h.logger.Debugf("Create host acknowledgement on %d: %+v", id, acknowledgement)

	if id <= 0 {
		return errors.New("invalid host ID")
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(acknowledgement); err != nil {
		return errors.Wrap(err, "validation error on create host acknowledgement")
	}

	response, err := h.client.R().
		SetBody(acknowledgement).
		SetPathParam("id", fmt.Sprintf("%d", id)).
		Post("/monitoring/hosts/{id}/acknowledgements")

	h.logger.Debugf("Response from create host acknowledgement: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error when create host acknowledgement")
	}

	if response.IsError() {
		return errors.Errorf("create host acknowledgement failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return nil
}

// DeleteHost deletes an acknowledgement for host by its ID
func (h *DefaultAckneledgementService) DeleteHost(hostId int64) (err error) {
	h.logger.Debugf("Delete host acknowledgement for host Id: %d", hostId)

	if hostId <= 0 {
		return errors.Errorf("invalid host id: %d", hostId)
	}

	response, err := h.client.R().
		SetPathParam("host_id", fmt.Sprintf("%d", hostId)).
		Delete("/monitoring/hosts/{host_id}/acknowledgements")

	h.logger.Debugf("Response from delete host acknowledgement: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error on delete host acknowledgement")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil
		}
		return errors.Errorf("delete host acknowledgement failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return nil
}

// ListService retrieves a list of service acknowledgements
func (h *DefaultAckneledgementService) ListService() (acknowledgementResponse *ListResponse[AcknowledgementListServiceResponse], err error) {
	h.logger.Debug("List service acknowledgements")

	acknowledgementResponse = new(ListResponse[AcknowledgementListServiceResponse])
	response, err := h.client.R().
		SetResult(acknowledgementResponse).
		Get("/monitoring/services/acknowledgements")

	h.logger.Debugf("Response from list service acknowledgements: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error on list service acknowledgements")
	}

	if response.IsError() {
		return nil, errors.Errorf("list service acknowledgements failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return acknowledgementResponse, nil
}

// CreateService creates an acknowledgement for a service
func (h *DefaultAckneledgementService) CreateService(id int64, acknowledgement *AcknowledgementCreateServiceRequest) (err error) {

	h.logger.Debugf("Create service acknowledgement on %d: %+v", id, acknowledgement)

	if id <= 0 {
		return errors.New("invalid service ID")
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(acknowledgement); err != nil {
		return errors.Wrap(err, "validation error on create service acknowledgement")
	}

	response, err := h.client.R().
		SetBody(acknowledgement).
		SetPathParam("id", fmt.Sprintf("%d", id)).
		Post("/monitoring/services/{id}/acknowledgements")

	h.logger.Debugf("Response from create service acknowledgement: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error when create service acknowledgement")
	}

	if response.IsError() {
		return errors.Errorf("create service acknowledgement failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return nil
}

// DeleteService deletes an acknowledgement for service by its ID
func (h *DefaultAckneledgementService) DeleteService(serviceId int64) (err error) {
	h.logger.Debugf("Delete service acknowledgement for service Id: %d", serviceId)

	if serviceId <= 0 {
		return errors.Errorf("invalid service id: %d", serviceId)
	}

	response, err := h.client.R().
		SetPathParam("service_id", fmt.Sprintf("%d", serviceId)).
		Delete("/monitoring/services/{service_id}/acknowledgements")

	h.logger.Debugf("Response from delete service acknowledgement: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error on delete service acknowledgement")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil
		}
		return errors.Errorf("delete service acknowledgement failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return nil
}
