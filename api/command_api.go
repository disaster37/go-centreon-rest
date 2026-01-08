package api

import (
	"emperror.dev/errors"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// CommandService defines the interface for managing commands in Centreon.
type CommandService interface {
	// Create a new command
	Create(command *CommandCreateRequest) (commandResponse *CommandCreateResponse, err error)

	// Get a command by ID
	// It uses Find method to get the command
	Get(id int64) (commandResponse *CommandResponse, err error)

	// GetByName retrieves a command by its Name
	// It uses Find method to get the command
	GetByName(name string) (commandResponse *CommandResponse, err error)

	// Find retrieves commands based on specific criteria
	Find(opts *ListOptions) (commandListResponse *ListResponse[CommandResponse], err error)
}

type DefaultCommandService struct {
	client *resty.Client
	logger *logrus.Entry
}

// NewCommandService creates a new instance of DefaultCommandService
func NewCommandService(client *resty.Client, logger *logrus.Entry) CommandService {
	return &DefaultCommandService{
		client: client,
		logger: logger.WithField("service", "command"),
	}
}

// Get retrieves a command by its ID
func (c *DefaultCommandService) Get(id int64) (commandResponse *CommandResponse, err error) {

	c.logger.Debugf("Get command with Id: %d", id)

	commandListResponse, err := c.Find(&ListOptions{
		Search: map[string]interface{}{
			"id": id,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find command with id %d", id)
	}

	if commandListResponse.Meta.Total == 1 {
		return &commandListResponse.Result[0], nil
	}

	return nil, nil
}

// GetByName retrieves a command by its Name
func (c *DefaultCommandService) GetByName(name string) (commandResponse *CommandResponse, err error) {

	c.logger.Debugf("Get command with Name: %s", name)

	commandListResponse, err := c.Find(&ListOptions{
		Search: map[string]interface{}{
			"name": name,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find command with name %s", name)
	}

	if commandListResponse.Meta.Total == 1 {
		return &commandListResponse.Result[0], nil
	}

	return nil, nil
}

// Find retrieves commands based on specific criteria
func (c *DefaultCommandService) Find(opts *ListOptions) (commandListResponse *ListResponse[CommandResponse], err error) {

	if opts == nil {
		opts = &ListOptions{}
	}

	c.logger.Debugf("Find commands with options: %+v", opts)

	commandListResponse = new(ListResponse[CommandResponse])

	response, err := c.client.R().
		SetResult(commandListResponse).
		SetQueryParams(opts.GetQueryParams()).
		Get("/configuration/commands")

	c.logger.Debugf("Response from list commands: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during list commands request")
	}

	if response.IsError() {
		return nil, errors.Errorf("list commands failed with status code: %d", response.StatusCode())
	}

	return commandListResponse, nil
}
// Create creates a new command
func (c *DefaultCommandService) Create(command *CommandCreateRequest) (commandResponse *CommandCreateResponse, err error) {
	c.logger.Debugf("Create command with data: %+v", command)

	commandResponse = new(CommandCreateResponse)

	response, err := c.client.R().
		SetBody(command).
		SetResult(commandResponse).
		Post("/configuration/commands")

	c.logger.Debugf("Response from create command: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during create command request")
	}

	if response.IsError() {
		return nil, errors.Errorf("create command failed with status code: %d", response.StatusCode())
	}

	return commandResponse, nil
}