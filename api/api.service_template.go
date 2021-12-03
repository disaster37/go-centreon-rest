package centreonapi

import (
	"github.com/disaster37/go-centreon-rest/v21.10/models"
	"github.com/go-resty/resty/v2"
)

const (
	objectServiceTemplate = "stpl"
)

type ServiceTemplateImpl struct {
	*serviceBaseImpl
}

// NewServiceTemplate permit to get service template handler that implement ServiceTemplateAPI interface
func NewServiceTemplate(client *resty.Client) ServiceTemplateAPI {
	return &ServiceTemplateImpl{
		serviceBaseImpl: newServiceBase(client, objectServiceTemplate).(*serviceBaseImpl),
	}
}

// Get permit to get one service template
func (s *ServiceTemplateImpl) Get(host, name string) (service *models.ServiceTemplateGet, err error) {
	data, err := s.serviceBaseImpl.Get(host, name)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil
	}

	service = &models.ServiceTemplateGet{
		ServiceBaseGet: data,
	}

	return service, nil
}

// List permit to get all service templates
func (s *ServiceTemplateImpl) List() (services []*models.ServiceTemplateGet, err error) {
	datas, err := s.serviceBaseImpl.List()
	if err != nil {
		return nil, err
	}

	services = make([]*models.ServiceTemplateGet, 0, len(datas))
	for _, data := range datas {
		service := &models.ServiceTemplateGet{
			ServiceBaseGet: data,
		}
		services = append(services, service)
	}

	return services, nil

}
