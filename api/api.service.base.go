package centreonapi

import (
	"encoding/json"
	"strings"

	"github.com/disaster37/go-centreon-rest/v21/models"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ServiceBaseImpl is implementation of ServiceBase interface
// ServiceBase is used by Service and ServiceTemplate
type serviceBaseImpl struct {
	client *resty.Client
	object string
}

func newServiceBase(client *resty.Client, object string) serviceBaseAPIGeneric {
	return &serviceBaseImpl{
		client: client,
		object: object,
	}
}

// Get permit to retrieve one service. It call with show API
func (s *serviceBaseImpl) Get(host, name string) (service *models.ServiceBaseGet, err error) {
	if host == "" {
		return nil, errors.New("Name must be provided")
	}
	if name == "" {
		return nil, errors.New("Service name must be provided")
	}
	log.Tracef("Host: %s", host)
	log.Tracef("Service Name: %s", name)

	payload := NewPayload("show", s.object, "%s;%s", host, name)
	log.Debugf("Payload: %s", payload)
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 300 {
		return nil, errors.Errorf("Error when get service %s on host %s: %s", name, host, resp.Body())
	}

	result := new(Result)
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		return nil, err
	}
	services := make([]*models.ServiceBaseGet, 0)
	if err = json.Unmarshal(result.Result, &services); err != nil {
		return nil, err
	}
	if len(services) == 0 {
		return nil, nil
	}

	log.Debugf("Get service %s on host %s successfully", name, host)

	return services[0], nil
}

// List permit to retrieve all services. It call with show API
func (s *serviceBaseImpl) List() (services []*models.ServiceBaseGet, err error) {

	payload := NewPayload("show", s.object, "")
	log.Debugf("Payload: %s", payload)
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 300 {
		return nil, errors.Errorf("Error when get all services: %s", resp.Body())
	}

	result := new(Result)
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		return nil, err
	}
	services = make([]*models.ServiceBaseGet, 0, len(result.Result))
	if err = json.Unmarshal(result.Result, &services); err != nil {
		return nil, err
	}

	log.Debugf("Get all services successfully")

	return services, nil
}

// Add permit to create new service. It call with add API
func (s *serviceBaseImpl) Add(host, name, template string) (err error) {
	if host == "" {
		return errors.New("Name must be provided")
	}
	if name == "" {
		return errors.New("Service name must be provided")
	}
	log.Tracef("Host: %s", host)
	log.Tracef("Service Name: %s", name)
	log.Tracef("Template: %s", template)

	payload := NewPayload("add", s.object, "%s;%s;%s", host, name, template)
	log.Debugf("Payload: %s", payload)
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 300 {
		return errors.Errorf("Error when create service %s on host %s: %s", name, host, resp.Body())
	}

	return nil
}

// Delete permit to remove on service. It call with del API
func (s *serviceBaseImpl) Delete(host, name string) (err error) {
	if host == "" {
		return errors.New("Name must be provided")
	}
	if name == "" {
		return errors.New("Service name must be provided")
	}
	log.Tracef("Host: %s", host)
	log.Tracef("Service Name: %s", name)

	payload := NewPayload("del", s.object, "%s;%s", host, name)
	log.Debugf("Payload: %s", payload)
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 300 {
		return errors.Errorf("Error when delete service %s on host %s: %s", name, host, resp.Body())
	}

	log.Debugf("Delete service %s on host %s successfully", name, host)

	return nil
}

// SetParam permit to set property value on service
func (s *serviceBaseImpl) SetParam(host, service, name, value string) (err error) {
	if host == "" {
		return errors.New("Name must be provided")
	}
	if service == "" {
		return errors.New("Service name must be provided")
	}
	if name == "" {
		return errors.New("Property name must be provided")
	}
	log.Tracef("Host: %s", host)
	log.Tracef("Service Name: %s", service)
	log.Tracef("Property Name: %s", name)
	log.Tracef("Value: %s", value)

	payload := NewPayload("setparam", s.object, "%s;%s;%s;%s", host, service, name, value)
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 300 {
		return errors.Errorf("Error when set params %s=%s: %s", name, value, resp.Body())
	}

	log.Debugf("Set param %s on %s/%s successfully", name, host, service)

	return nil
}

// GetParam permit to get property value on service
func (s *serviceBaseImpl) GetParam(host, service string, params []string) (values map[string]string, err error) {
	if host == "" {
		return nil, errors.New("Name must be provided")
	}
	if service == "" {
		return nil, errors.New("Service name must be provided")
	}
	if params == nil || len(params) == 0 {
		return nil, errors.New("Params must be provided")
	}
	log.Tracef("Host: %s", host)
	log.Tracef("Service Name: %s", service)
	log.Tracef("Params Name: %+v", params)

	payload := NewPayload("getparam", s.object, "%s;%s;%s", host, service, strings.Join(params, "|"))
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
	tmpValues := make([]map[string]string, 0, 1)
	if err := json.Unmarshal(result.Result, &tmpValues); err != nil {
		return nil, err
	}

	if len(tmpValues) == 0 {
		return values, nil
	}

	log.Debugf("Get params %s on service  %s/%s successfully", strings.Join(params, "|"), host, service)

	return tmpValues[0], nil
}

// GetMacros permit to get all macros on service
func (s *serviceBaseImpl) GetMacros(host, service string) (macros []*models.Macro, err error) {
	if host == "" {
		return nil, errors.New("Name must be provided")
	}
	if service == "" {
		return nil, errors.New("Service name must be provided")
	}
	log.Tracef("Host: %s", host)
	log.Tracef("Service Name: %s", service)

	payload := NewPayload("getmacro", s.object, "%s;%s", host, service)
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 300 {
		return nil, errors.Errorf("Error when get macros: %s", resp.Body())
	}

	result := new(Result)
	if err := json.Unmarshal(resp.Body(), result); err != nil {
		return nil, err
	}
	macros = make([]*models.Macro, 0, len(result.Result))
	if err := json.Unmarshal(result.Result, &macros); err != nil {
		return nil, err
	}

	log.Debugf("Get all macros on service %s/%s successfully", host, service)

	return macros, nil

}

// SetMacro permit to set macro on service
func (s *serviceBaseImpl) SetMacro(host, service string, macro *models.Macro) (err error) {
	if host == "" {
		return errors.New("Name must be provided")
	}
	if service == "" {
		return errors.New("Service name must be provided")
	}
	if macro == nil || !macro.IsValid() {
		return errors.New("Valid macro must be provided")
	}
	log.Tracef("Host: %s", host)
	log.Tracef("Service Name: %s", service)
	log.Tracef("Macro: %+v", macro)

	payload := NewPayload("setmacro", s.object, "%s;%s;%s;%s;%s;%s", host, service, strings.ToUpper(macro.Name), macro.Value, macro.IsPassword, macro.Description)
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 300 {
		return errors.Errorf("Error when set macro: %s", resp.Body())
	}

	log.Debugf("Set macro %s on service  %s/%s successfully", macro.Name, host, service)

	return nil
}

// DeleteMacro permit to remove macro on service
func (s *serviceBaseImpl) DeleteMacro(host, service, name string) (err error) {
	if host == "" {
		return errors.New("Name must be provided")
	}
	if service == "" {
		return errors.New("Service name must be provided")
	}
	if name == "" {
		return errors.New("Macro name must be provided")
	}
	log.Tracef("Host: %s", host)
	log.Tracef("Service Name: %s", service)
	log.Tracef("Macro: %s", name)

	payload := NewPayload("delmacro", s.object, "%s;%s;%s", host, service, name)
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 300 {
		return errors.Errorf("Error when delete macro: %s", resp.Body())
	}

	log.Debugf("Delete macro %s on service  %s/%s successfully", name, host, service)

	return nil
}

// GetCategories permit to get all categories on service
func (s *serviceBaseImpl) GetCategories(host, service string) (categories []string, err error) {
	if host == "" {
		return nil, errors.New("Name must be provided")
	}
	if service == "" {
		return nil, errors.New("Service name must be provided")
	}
	log.Tracef("Host: %s", host)
	log.Tracef("Service Name: %s", service)

	payload := NewPayload("getcategory", s.object, "%s;%s", host, service)
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 300 {
		return nil, errors.Errorf("Error when get categories: %s", resp.Body())
	}

	result := new(Result)
	if err := json.Unmarshal(resp.Body(), result); err != nil {
		return nil, err
	}
	mapRes := make([]map[string]string, 0, len(result.Result))
	categories = make([]string, 0, len(result.Result))
	if err := json.Unmarshal(result.Result, &mapRes); err != nil {
		return nil, err
	}

	for _, res := range mapRes {
		categories = append(categories, res["name"])
	}

	log.Debugf("Get all categories on service %s/%s successfully", host, service)

	return categories, nil
}

// SetCategories permit to set categories on service
func (s *serviceBaseImpl) SetCategories(host, service string, categories []string) (err error) {
	if host == "" {
		return errors.New("Name must be provided")
	}
	if service == "" {
		return errors.New("Service name must be provided")
	}
	if categories == nil || len(categories) == 0 {
		return errors.New("Categories must be provided")
	}
	log.Tracef("Host: %s", host)
	log.Tracef("Service Name: %s", service)
	log.Tracef("Categories: %+v", categories)

	payload := NewPayload("setcategory", s.object, "%s;%s;%s", host, service, strings.Join(categories, "|"))
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 300 {
		return errors.Errorf("Error when set categories: %s", resp.Body())
	}

	log.Debugf("Set categories on service %s/%s successfully", host, service)

	return nil
}

// Delete categories permit to delete categories on service
func (s *serviceBaseImpl) DeleteCategories(host, service string, categories []string) (err error) {
	if host == "" {
		return errors.New("Name must be provided")
	}
	if service == "" {
		return errors.New("Service name must be provided")
	}
	if categories == nil || len(categories) == 0 {
		return errors.New("Categories must be provided")
	}
	log.Tracef("Host: %s", host)
	log.Tracef("Service Name: %s", service)
	log.Tracef("Categories: %+v", categories)

	payload := NewPayload("delcategory", s.object, "%s;%s;%s", host, service, strings.Join(categories, "|"))
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 300 {
		return errors.Errorf("Error when delete categories: %s", resp.Body())
	}

	log.Debugf("Delete categories on service %s/%s successfully", host, service)

	return nil
}

// GetServiceGroups permit to get all service groups on service
func (s *serviceBaseImpl) GetServiceGroups(host, service string) (serviceGroups []string, err error) {
	if host == "" {
		return nil, errors.New("Name must be provided")
	}
	if service == "" {
		return nil, errors.New("Service name must be provided")
	}
	log.Tracef("Host: %s", host)
	log.Tracef("Service Name: %s", service)

	payload := NewPayload("getservicegroup", s.object, "%s;%s", host, service)
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 300 {
		return nil, errors.Errorf("Error when get service groups: %s", resp.Body())
	}

	result := new(Result)
	if err := json.Unmarshal(resp.Body(), result); err != nil {
		return nil, err
	}
	mapRes := make([]map[string]string, 0, len(result.Result))
	serviceGroups = make([]string, 0, len(result.Result))
	if err := json.Unmarshal(result.Result, &mapRes); err != nil {
		return nil, err
	}

	for _, res := range mapRes {
		serviceGroups = append(serviceGroups, res["name"])
	}

	log.Debugf("Get all service groups on service %s/%s successfully", host, service)

	return serviceGroups, nil
}

// SetServiceGroups permit to set service groups on service
func (s *serviceBaseImpl) SetServiceGroups(host, service string, serviceGroups []string) (err error) {
	if host == "" {
		return errors.New("Name must be provided")
	}
	if service == "" {
		return errors.New("Service name must be provided")
	}
	if serviceGroups == nil || len(serviceGroups) == 0 {
		return errors.New("Service groups must be provided")
	}
	log.Tracef("Host: %s", host)
	log.Tracef("Service Name: %s", service)
	log.Tracef("Service groups: %+v", serviceGroups)

	payload := NewPayload("setservicegroup", s.object, "%s;%s;%s", host, service, strings.Join(serviceGroups, "|"))
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 300 {
		return errors.Errorf("Error when set service groups: %s", resp.Body())
	}

	log.Debugf("Set service groups on service %s/%s successfully", host, service)

	return nil
}

// DeleteServiceGroups permit to delete service groups on service
func (s *serviceBaseImpl) DeleteServiceGroups(host, service string, serviceGroups []string) (err error) {
	if host == "" {
		return errors.New("Name must be provided")
	}
	if service == "" {
		return errors.New("Service name must be provided")
	}
	if serviceGroups == nil || len(serviceGroups) == 0 {
		return errors.New("Service groups must be provided")
	}
	log.Tracef("Host: %s", host)
	log.Tracef("Service Name: %s", service)
	log.Tracef("Service groups: %+v", serviceGroups)

	payload := NewPayload("delservicegroup", s.object, "%s;%s;%s", host, service, strings.Join(serviceGroups, "|"))
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 300 {
		return errors.Errorf("Error when delete service groups: %s", resp.Body())
	}

	log.Debugf("Delete service groups on service %s/%s successfully", host, service)

	return nil
}

// GetTraps permit to get all traps on service
func (s *serviceBaseImpl) GetTraps(host, service string) (traps []string, err error) {
	if host == "" {
		return nil, errors.New("Name must be provided")
	}
	if service == "" {
		return nil, errors.New("Service name must be provided")
	}
	log.Tracef("Host: %s", host)
	log.Tracef("Service Name: %s", service)

	payload := NewPayload("gettrap", s.object, "%s;%s", host, service)
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 300 {
		return nil, errors.Errorf("Error when get traps: %s", resp.Body())
	}

	result := new(Result)
	if err := json.Unmarshal(resp.Body(), result); err != nil {
		return nil, err
	}
	mapRes := make([]map[string]string, 0, len(result.Result))
	traps = make([]string, 0, len(result.Result))
	if err := json.Unmarshal(result.Result, &mapRes); err != nil {
		return nil, err
	}

	for _, res := range mapRes {
		traps = append(traps, res["name"])
	}

	log.Debugf("Get all traps on service %s/%s successfully", host, service)

	return traps, nil
}

// SetTraps permit to set traps on service
func (s *serviceBaseImpl) SetTraps(host, service string, traps []string) (err error) {
	if host == "" {
		return errors.New("Name must be provided")
	}
	if service == "" {
		return errors.New("Service name must be provided")
	}
	if traps == nil || len(traps) == 0 {
		return errors.New("Traps must be provided")
	}
	log.Tracef("Host: %s", host)
	log.Tracef("Service Name: %s", service)
	log.Tracef("Traps: %+v", traps)

	payload := NewPayload("settrap", s.object, "%s;%s;%s", host, service, strings.Join(traps, "|"))
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 300 {
		return errors.Errorf("Error when set traps: %s", resp.Body())
	}

	log.Debugf("Set traps on service %s/%s successfully", host, service)

	return nil
}

// DeleteTraps permit to delete traps on service
func (s *serviceBaseImpl) DeleteTraps(host, service string, traps []string) (err error) {
	if host == "" {
		return errors.New("Name must be provided")
	}
	if service == "" {
		return errors.New("Service name must be provided")
	}
	if traps == nil || len(traps) == 0 {
		return errors.New("Traps must be provided")
	}
	log.Tracef("Host: %s", host)
	log.Tracef("Service Name: %s", service)
	log.Tracef("Traps: %+v", traps)

	payload := NewPayload("deltrap", s.object, "%s;%s;%s", host, service, strings.Join(traps, "|"))
	resp, err := s.client.R().
		SetBody(payload).
		Post("")
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 300 {
		return errors.Errorf("Error when delete traps: %s", resp.Body())
	}

	log.Debugf("Delete traps on service %s/%s successfully", host, service)

	return nil
}
