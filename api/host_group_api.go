package api

import (
	"fmt"
	"net/http"
	"strings"

	"emperror.dev/errors"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// HostGroupService defines the interface for managing host groups in Centreon.
type HostGroupService interface {

	// Create creates a new host group.
	Create(hostGroup *HostGroupCreateOrUpdateRequest) (hostGroupResponse *HostGroupResponse, err error)

	//
	Update(id int64, hostGroup *HostGroupCreateOrUpdateRequest) (err error)

	// Get retrieves a host group by its ID.
	Get(id int64) (hostGroupResponse *HostGroupResponse, err error)

	// GetByName retrieves a host group by its Name.
	// It uses the List method to get the host group.
	GetByName(name string) (hostGroupResponse *HostGroupResponse, err error)

	// List retrieves a list of host groups based on the provided options.
	List(options *ListOptions) (hostGroupListResponse *ListResponse[HostGroupResponse], err error)

	// Delete deletes a host group by its ID.
	Delete(id int64) (err error)

	// List retrieves a list of host groups based on the provided options.
	ListFromRealTime(options *ListOptions) (hostGroupListResponse *ListResponse[HostGroupRealTimeResponse], err error)
}

type DefaultHostGroupService struct {
	client *resty.Client
	logger *logrus.Entry
}

// NewHostGroupService creates a new instance of DefaultHostGroupService
func NewHostGroupService(client *resty.Client, logger *logrus.Entry) HostGroupService {
	return &DefaultHostGroupService{
		client: client,
		logger: logger.WithField("service", "hostGroup"),
	}
}

// Create creates a new host group.
func (s *DefaultHostGroupService) Create(hostGroup *HostGroupCreateOrUpdateRequest) (hostGroupResponse *HostGroupResponse, err error) {
	s.logger.Debugf("Create host group: %+v", hostGroup)

	if hostGroup.Hosts == nil {
		hostGroup.Hosts = []int64{}
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(hostGroup); err != nil {
		return nil, errors.Wrap(err, "validation error on create host group")
	}

	hostGroupResponse = new(HostGroupResponse)

	response, err := s.client.R().
		SetBody(hostGroup).
		SetResult(hostGroupResponse).
		Post("/configuration/hosts/groups")

	s.logger.Debugf("Response from create host group: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during create host group request")
	}

	if response.IsError() {
		return nil, errors.Errorf("create host group failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return hostGroupResponse, nil
}

// List retrieves a list of host groups based on the provided options.
func (s *DefaultHostGroupService) List(opts *ListOptions) (hostGroupListResponse *ListResponse[HostGroupResponse], err error) {
	if opts == nil {
		opts = &ListOptions{}
	}

	s.logger.Debugf("Find host groups with options: %+v", opts)

	hostGroupListResponse = new(ListResponse[HostGroupResponse])

	response, err := s.client.R().
		SetQueryParams(opts.GetQueryParams()).
		SetResult(hostGroupListResponse).
		Get("/configuration/hosts/groups")

	s.logger.Debugf("Response from find host groups: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during find host groups request")
	}

	if response.IsError() {
		return nil, errors.Errorf("list host groups failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return hostGroupListResponse, nil
}

// Update updates an existing host group.
func (s *DefaultHostGroupService) Update(id int64, hostGroup *HostGroupCreateOrUpdateRequest) (err error) {
	s.logger.Debugf("Update host group with Id %d: %+v", id, hostGroup)

	if id <= 0 {
		return errors.New("invalid host group id: 0")
	}

	if hostGroup.Hosts == nil {
		hostGroup.Hosts = []int64{}
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(hostGroup); err != nil {
		return errors.Wrap(err, "validation error on update host group")
	}

	response, err := s.client.R().
		SetBody(hostGroup).
		SetPathParam("groupId", fmt.Sprintf("%d", id)).
		Put("/configuration/hosts/groups/{groupId}")

	s.logger.Debugf("Response from update host group: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during update host group request")
	}

	if response.IsError() {
		return errors.Errorf("failed to update host group, status code: %d, response: %s", response.StatusCode(), response.String())
	}

	return nil
}

// Get retrieves a host group by its ID.
func (s *DefaultHostGroupService) Get(id int64) (hostGroupResponse *HostGroupResponse, err error) {
	s.logger.Debugf("Get host group with Id: %d", id)

	if id <= 0 {
		return nil, errors.New("invalid host group id: 0")
	}

	hostGroupResponse = new(HostGroupResponse)

	response, err := s.client.R().
		SetPathParam("groupId", fmt.Sprintf("%d", id)).
		SetResult(hostGroupResponse).
		Get("/configuration/hosts/groups/{groupId}")

	s.logger.Debugf("Response from get host group: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during get host group request")
	}

	if response.IsError() {

		if response.StatusCode() == http.StatusNotFound {
			return nil, nil
		}
		return nil, errors.Errorf("get host group failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return hostGroupResponse, nil
}

// Delete deletes a host group by its ID.
func (s *DefaultHostGroupService) Delete(id int64) (err error) {
	s.logger.Debugf("Delete host group with Id: %d", id)

	if id <= 0 {
		return errors.New("invalid host group id: 0")
	}

	response, err := s.client.R().
		SetPathParam("groupId", fmt.Sprintf("%d", id)).
		Delete("/configuration/hosts/groups/{groupId}")

	s.logger.Debugf("Response from delete host group: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during delete host group request")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil
		}
		return errors.Errorf("delete host group failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return nil
}

// GetByName retrieves a host group by its Name.
// It uses the List method to get the host group.
func (s *DefaultHostGroupService) GetByName(name string) (hostGroupResponse *HostGroupResponse, err error) {
	s.logger.Debugf("Get host group by Name: %s", name)

	if strings.TrimSpace(name) == "" {
		return nil, errors.New("host group name cannot be empty")
	}

	hostGroupListResponse, err := s.List(&ListOptions{
		Search: map[string]interface{}{
			"name": name,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "error during get host group by name: %s", name)
	}
	if hostGroupListResponse.Meta.Total == 1 {
		return &hostGroupListResponse.Result[0], nil
	}

	return nil, nil
}

func (s *DefaultHostGroupService) ListFromRealTime(opts *ListOptions) (hostGroupListResponse *ListResponse[HostGroupRealTimeResponse], err error) {
	if opts == nil {
		opts = &ListOptions{}
	}

	s.logger.Debugf("Find host groups with options: %+v", opts)

	hostGroupListResponse = new(ListResponse[HostGroupRealTimeResponse])

	response, err := s.client.R().
		SetQueryParams(opts.GetQueryParams()).
		SetResult(hostGroupListResponse).
		Get("/monitoring/hostgroups")

	s.logger.Debugf("Response from find host groups: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during find host groups request")
	}

	if response.IsError() {
		return nil, errors.Errorf("list host groups failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return hostGroupListResponse, nil
}
