package centreonapi

import (
	"github.com/disaster37/go-centreon-rest/v21/models"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	objectService = "service"
)

type ServiceImpl struct {
	*serviceBaseImpl
}

// NewService permit to get service hendler that implement ServiceAPI interface
func NewService(client *resty.Client) ServiceAPI {
	return &ServiceImpl{
		serviceBaseImpl: newServiceBase(client, objectService).(*serviceBaseImpl),
	}
}

// Get permit to get one service
func (s *ServiceImpl) Get(host, name string) (service *models.ServiceGet, err error) {
	data, err := s.serviceBaseImpl.Get(host, name)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil
	}

	service = &models.ServiceGet{
		ServiceBaseGet: data,
	}

	return service, nil
}

// List permit to get all services
func (s *ServiceImpl) List() (services []*models.ServiceGet, err error) {
	datas, err := s.serviceBaseImpl.List()
	if err != nil {
		return nil, err
	}

	services = make([]*models.ServiceGet, 0, len(datas))
	for _, data := range datas {
		service := &models.ServiceGet{
			ServiceBaseGet: data,
		}
		services = append(services, service)
	}

	return services, nil

}

func (s *ServiceImpl) SetHost(host, name, newHost string) (err error) {
	if host == "" {
		return errors.New("Host must be provided")
	}
	if name == "" {
		return errors.New("Service name must be provided")
	}
	if newHost == "" {
		return errors.New("New host must be provided")
	}
	log.Tracef("Host: %s", host)
	log.Tracef("Service Name: %s", name)
	log.Tracef("New host: %s", newHost)

	payload := NewPayload("sethost", s.object, "%s;%s;%s", host, name, newHost)
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 300 {
		return errors.Errorf("Error when set new host %s/%s: %s", newHost, name, resp.Body())
	}

	log.Debugf("Set new host %s on %s successfully", newHost, name)

	return nil
}
