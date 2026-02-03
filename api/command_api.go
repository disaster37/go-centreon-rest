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

// CommandService defines the interface for managing commands in Centreon.
type CommandService interface {
	// Create a new command
	Create(command *CommandCreateOrUpdateRequest) (commandResponse *CommandResponse, err error)

	// Update an existing command by ID
	// Need PR https://github.com/centreon/centreon/pull/9359
	Update(id int64, command *CommandCreateOrUpdateRequest) (err error)

	// Delete a command by ID
	// Need PR https://github.com/centreon/centreon/pull/9359
	Delete(id int64) (err error)

	// Get a command by ID
	// Need PR https://github.com/centreon/centreon/pull/9359
	Get(id int64) (commandResponse *CommandResponse, err error)

	// GetByName retrieves a command by its Name
	// It uses Find method to get the command
	GetByName(name string) (commandResponse *CommandFindResponse, err error)

	// Find retrieves commands based on specific criteria
	Find(opts *ListOptions) (commandListResponse *ListResponse[CommandFindResponse], err error)
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

	if id <= 0 {
		return nil, errors.Errorf("invalid command id: %d", id)
	}

	commandResponse = new(CommandResponse)

	response, err := c.client.R().
		SetResult(commandResponse).
		SetPathParam("id", fmt.Sprintf("%d", id)).
		Get("/configuration/commands/{id}")

	c.logger.Debugf("Response from get command: %s", response.String())

	if err != nil {
		return nil, errors.Wrapf(err, "failed to get command with id %d", id)
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil, nil
		}
		return nil, errors.Errorf("get command failed with status code: %d", response.StatusCode())
	}

	return commandResponse, nil
}

// GetByName retrieves a command by its Name
func (c *DefaultCommandService) GetByName(name string) (commandResponse *CommandFindResponse, err error) {

	c.logger.Debugf("Get command with Name: %s", name)

	if strings.TrimSpace(name) == "" {
		return nil, errors.New("command name cannot be empty")
	}

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
func (c *DefaultCommandService) Find(opts *ListOptions) (commandListResponse *ListResponse[CommandFindResponse], err error) {

	if opts == nil {
		opts = &ListOptions{}
	}

	c.logger.Debugf("Find commands with options: %+v", opts)

	commandListResponse = new(ListResponse[CommandFindResponse])

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
func (c *DefaultCommandService) Create(command *CommandCreateOrUpdateRequest) (commandResponse *CommandResponse, err error) {
	c.logger.Debugf("Create command with data: %+v", command)

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(command); err != nil {
		return nil, errors.Wrap(err, "validation error on create command")
	}

	commandResponse = new(CommandResponse)

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

// Update update and existing command
func (c *DefaultCommandService) Update(id int64, command *CommandCreateOrUpdateRequest) (err error) {
	c.logger.Debugf("Update command with id %d and data: %+v", id, command)

	if id <= 0 {
		return errors.Errorf("invalid command id: %d", id)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(command); err != nil {
		return errors.Wrap(err, "validation error on update host")
	}

	response, err := c.client.R().
		SetBody(command).
		SetPathParam("id", fmt.Sprintf("%d", id)).
		Put("/configuration/commands/{id}")

	c.logger.Debugf("Response from update command: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during update command request")
	}

	if response.IsError() {
		return errors.Errorf("update command failed with status code: %d and message: %s", response.StatusCode(), response.String())
	}

	return nil
}

// Delete delete and existing command
func (c *DefaultCommandService) Delete(id int64) (err error) {
	c.logger.Debugf("Delete command with id: %d", id)

	if id <= 0 {
		return errors.Errorf("invalid command id: %d", id)
	}

	response, err := c.client.R().
		SetPathParam("id", fmt.Sprintf("%d", id)).
		Delete("/configuration/commands/{id}")

	c.logger.Debugf("Response from delete command: %s", response.String())

	if err != nil {
		return errors.Wrap(err, "error during delete command request")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil
		}
		return errors.Errorf("delete command failed with status code: %d and message: %s", response.StatusCode(), response.String())
	}

	return nil
}
