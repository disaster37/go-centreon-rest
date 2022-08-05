package centreonapi

import (
	"encoding/json"
	"strings"

	"github.com/disaster37/go-centreon-rest/v21/models"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	objectServiceGroup = "sg"
)

type ServiceGroupImpl struct {
	client *resty.Client
	object string
}

// NewService permit to get serviceGroup handler that implement ServiceGroupAPI interface
func NewServiceGroup(client *resty.Client) ServiceGroupAPI {
	return &ServiceGroupImpl{
		client: client,
		object: objectServiceGroup,
	}
}

// Get permit to get one service group by it's name
func (s *ServiceGroupImpl) Get(name string) (sg *models.ServiceGroup, err error) {
	if name == "" {
		return nil, errors.New("ServiceGroup name must be provided")
	}
	log.Tracef("ServiceGroup Name: %s", name)

	payload := NewPayload("show", s.object, "%s", name)
	log.Debugf("Payload: %s", payload)
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 300 {
		return nil, errors.Errorf("Error when get serviceGroup %s: %s", name, resp.Body())
	}

	result := new(Result)
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		return nil, err
	}
	serviceGroups := make([]*models.ServiceGroup, 0, 1)
	if err = json.Unmarshal(result.Result, &serviceGroups); err != nil {
		return nil, err
	}
	if len(serviceGroups) == 0 {
		return nil, nil
	}

	log.Debugf("Get serviceGroups %s successfully", name)

	return serviceGroups[0], nil
}

// List permit to get all serviceGroups
func (s *ServiceGroupImpl) List() (serviceGroups []*models.ServiceGroup, err error) {
	payload := NewPayload("show", s.object, "")
	log.Debugf("Payload: %s", payload)
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 300 {
		return nil, errors.Errorf("Error when get all serviceGroups: %s", resp.Body())
	}

	result := new(Result)
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		return nil, err
	}
	serviceGroups = make([]*models.ServiceGroup, 0, len(result.Result))
	if err = json.Unmarshal(result.Result, &serviceGroups); err != nil {
		return nil, err
	}

	log.Debugf("Get all serviceGroups successfully")

	return serviceGroups, nil
}

// Add permit to create to service group
func (s *ServiceGroupImpl) Add(name, description string) (err error) {
	if name == "" {
		return errors.New("ServiceGroup name must be provided")
	}
	if description == "" {
		return errors.New("Description name must be provided")
	}
	log.Tracef("ServiceGroup Name: %s", name)
	log.Tracef("Description: %s", description)

	payload := NewPayload("add", s.object, "%s;%s", name, description)

	log.Debugf("Payload: %s", payload)
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 300 {
		return errors.Errorf("Error when create serviceGroup %s: %s", name, resp.Body())
	}

	return nil
}

// Delete permit to delete the service group
func (s *ServiceGroupImpl) Delete(name string) (err error) {
	if name == "" {
		return errors.New("ServiceGroup name must be provided")
	}
	log.Tracef("ServiceGroup Name: %s", name)

	payload := NewPayload("del", s.object, "%s", name)
	log.Debugf("Payload: %s", payload)
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 300 {
		return errors.Errorf("Error when delete serviceGroup %s: %s", name, resp.Body())
	}

	log.Debugf("Delete serviceGroup %s successfully", name)

	return nil
}

// SetParam permit to set extra properties on service group
func (s *ServiceGroupImpl) SetParam(serviceGroupName, name, value string) (err error) {
	if serviceGroupName == "" {
		return errors.New("ServiceGroup name must be provided")
	}
	if name == "" {
		return errors.New("Property name must be provided")
	}
	log.Tracef("ServiceGroup Name: %s", serviceGroupName)
	log.Tracef("Property Name: %s", name)
	log.Tracef("Value: %s", value)

	payload := NewPayload("setparam", s.object, "%s;%s;%s", serviceGroupName, name, value)
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 300 {
		return errors.Errorf("Error when set params %s=%s: %s", name, value, resp.Body())
	}

	log.Debugf("Set param %s on %s successfully", name, serviceGroupName)

	return nil
}

// GetParam permit to get extra properties from service group
func (s *ServiceGroupImpl) GetParam(serviceGroupName string, params []string) (values map[string]string, err error) {
	if serviceGroupName == "" {
		return nil, errors.New("ServiceGroup name must be provided")
	}
	if params == nil || len(params) == 0 {
		return nil, errors.New("Params must be provided")
	}
	log.Tracef("ServiceGroup Name: %s", serviceGroupName)
	log.Tracef("Params Name: %+v", params)

	payload := NewPayload("getparam", s.object, "%s;%s", serviceGroupName, strings.Join(params, "|"))
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 300 {
		return nil, errors.Errorf("Error when get params: %s", resp.Body())
	}

	result := new(Result)
	if err := json.Unmarshal(resp.Body(), result); err != nil {
		return nil, err
	}

	// It return array
	if len(params) == 1 {
		tmpValues := make([]string, 0, 1)
		if err := json.Unmarshal(result.Result, &tmpValues); err != nil {
			return nil, err
		}
		if len(tmpValues) > 0 {
			values = map[string]string{
				params[0]: tmpValues[0],
			}
		}
	} else {
		// else it return map
		tmpValues := make([]map[string]string, 0, 1)
		if err := json.Unmarshal(result.Result, &tmpValues); err != nil {
			return nil, err
		}

		if len(tmpValues) != 0 {
			values = tmpValues[0]
		}
	}

	log.Debugf("Get params %s on serviceGroup  %s successfully", strings.Join(params, "|"), serviceGroupName)

	return values, nil
}
