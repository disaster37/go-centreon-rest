package centreonapi

import (
	"github.com/disaster37/go-centreon-rest/v21/models"
	"github.com/go-resty/resty/v2"
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
