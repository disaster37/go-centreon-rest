package centreonapi

import (
	"github.com/disaster37/go-centreon-rest/v21/models"
	"github.com/go-resty/resty/v2"
)

type API interface {
	Service() ServiceAPI
	ServiceTemplate() ServiceTemplateAPI
	ServiceGroup() ServiceGroupAPI
	Client() *resty.Client
	Auth() (err error)
}

type serviceBaseAPI interface {
	Add(host, name, template string) (err error)
	Delete(host, name string) (err error)
	SetParam(host, service, name, value string) (err error)
	GetParam(host, service string, params []string) (values map[string]string, err error)
	GetMacros(host, service string) (macros []*models.Macro, err error)
	SetMacro(host, service string, macro *models.Macro) (err error)
	DeleteMacro(host, service, name string) (err error)
	GetCategories(host, service string) (categories []string, err error)
	SetCategories(host, service string, categories []string) (err error)
	DeleteCategories(host, service string, categories []string) (err error)
	GetServiceGroups(host, service string) (serviceGroups []string, err error)
	SetServiceGroups(host, service string, serviceGroups []string) (err error)
	DeleteServiceGroups(host, service string, serviceGroups []string) (err error)
	GetTraps(host, service string) (traps []string, err error)
	SetTraps(host, service string, traps []string) (err error)
	DeleteTraps(host, service string, traps []string) (err error)
}

type serviceBaseAPIGeneric interface {
	Get(host, name string) (service *models.ServiceBaseGet, err error)
	List() (services []*models.ServiceBaseGet, err error)
	serviceBaseAPI
}

type ServiceAPI interface {
	Get(host, name string) (service *models.ServiceGet, err error)
	List() (services []*models.ServiceGet, err error)
	SetHost(host, name, newHost string) (err error)
	serviceBaseAPI
}

type ServiceTemplateAPI interface {
	Get(host, name string) (service *models.ServiceTemplateGet, err error)
	List() (services []*models.ServiceTemplateGet, err error)
	serviceBaseAPI
}

type ServiceGroupAPI interface {
	Get(name string) (serviceGroup *models.ServiceGroup, err error)
	List() (serviceGroups []*models.ServiceGroup, err error)
	Add(name, description string) (err error)
	Delete(name string) (err error)
	SetParam(serviceGroupName, name, value string) (err error)
	GetParam(serviceGroupName string, params []string) (values map[string]string, err error)
}
