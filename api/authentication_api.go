package api

import (
	"fmt"
	"strings"

	"emperror.dev/errors"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// Authentication defines methods for authentication operations
type AuthenticationService interface {
	// Login method to authenticate a user and obtain a token
	Login(request *AuthenticationRequest) (*LoginAuthenticationResponse, error)

	// Logout method to invalidate the current user session
	Logout() (*LogoutAuthenticationResponse, error)

	// UpdatePassword method to change the password for a given user
	UpdatePassword(username string, request *PasswordUpdateRequest) (*PasswordUpdateResponse, error)

	// GetProviders method to retrieve all authentication providers
	GetProviders() ([]ProviderConfiguration, error)

	// AuthentificationToProvider method to authenticate a user via a specific provider
	AuthentificationToProvider(providerName string, request *AuthenticationRequestSecurityCredentials) (*AuthentificationToProviderResponse, error)
}

// DefaultAuthentificationService implements Authentication
type DefaultAuthentificationService struct {
	client *resty.Client
	logger *logrus.Entry
}

// NewAuthenticationService creates a new instance of DefaultAuthentificationService
func NewAuthenticationService(client *resty.Client, logger *logrus.Entry) AuthenticationService {
	return &DefaultAuthentificationService{
		client: client,
		logger: logger.WithField("service", "authentication"),
	}
}

// Login method to authenticate a user and obtain a token. It inject token in client header upon success.
func (h *DefaultAuthentificationService) Login(request *AuthenticationRequest) (*LoginAuthenticationResponse, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(request); err != nil {
		return nil, errors.Wrap(err, "validation error on Login")
	}

	loginAuthenticationResponse := new(LoginAuthenticationResponse)

	response, err := h.client.R().
		SetBody(request).
		SetResult(loginAuthenticationResponse).
		Post("/login")

	h.logger.Debugf("Response from Login: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during Login request")
	}

	if response.IsError() {
		return nil, errors.Errorf("login failed with status code: %d", response.StatusCode())
	}

	h.client.SetHeader("X-AUTH-TOKEN", loginAuthenticationResponse.Security.Token)

	h.logger.Infof("User %s logged in successfully", loginAuthenticationResponse.Contact.Name)

	return loginAuthenticationResponse, nil
}

// Logout method to invalidate the current user session and remove the token from client header.
func (h *DefaultAuthentificationService) Logout() (*LogoutAuthenticationResponse, error) {

	logoutAuthenticationResponse := new(LogoutAuthenticationResponse)

	response, err := h.client.R().
		SetResult(logoutAuthenticationResponse).
		Get("/logout")

	h.logger.Debugf("Response from Logout: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during Logout request")
	}

	if response.IsError() {
		return nil, errors.Errorf("logout failed with status code: %d", response.StatusCode())
	}

	h.logger.Info("User logged out successfully")

	h.client.SetHeader("X-AUTH-TOKEN", "")

	return logoutAuthenticationResponse, nil

}

// UpdatePassword method to change the password for a given user
func (h *DefaultAuthentificationService) UpdatePassword(username string, request *PasswordUpdateRequest) (*PasswordUpdateResponse, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(request); err != nil {
		return nil, errors.Wrap(err, "validation error on UpdatePassword")
	}

	passwordUpdateResponse := new(PasswordUpdateResponse)

	response, err := h.client.R().
		SetBody(request).
		SetResult(passwordUpdateResponse).
		Put(fmt.Sprintf("/authentication/users/%s/password", username))

	h.logger.Debugf("Response from UpdatePassword: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during UpdatePassword request")
	}

	if response.IsError() {
		return nil, errors.Errorf("password update failed with status code: %d", response.StatusCode())
	}

	h.logger.Infof("Password for user %s updated successfully", username)

	return passwordUpdateResponse, nil
}

// GetProviders method to retrieve all authentication providers
func (h *DefaultAuthentificationService) GetProviders() ([]ProviderConfiguration, error) {
	var providers []ProviderConfiguration

	response, err := h.client.R().
		SetResult(&providers).
		Get("/authentication/providers/configurations")

	h.logger.Debugf("Response from GetProviders: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during GetProviders request")
	}

	if response.IsError() {
		return nil, errors.Errorf("get providers failed with status code: %d", response.StatusCode())
	}

	h.logger.Infof("Retrieved %d authentication providers", len(providers))

	return providers, nil
}

// AuthentificationToProvider method to authenticate a user via a specific provider
func (h *DefaultAuthentificationService) AuthentificationToProvider(providerName string, request *AuthenticationRequestSecurityCredentials) (*AuthentificationToProviderResponse, error) {

	if providerName == "" || len(strings.TrimSpace(providerName)) == 0 {
		return nil, errors.Errorf("providerName is required for AuthentificationToProvider")
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(request); err != nil {
		return nil, errors.Wrap(err, "validation error on AuthentificationToProvider")
	}

	authentificationToProviderResponse := new(AuthentificationToProviderResponse)

	response, err := h.client.R().
		SetBody(request).
		SetResult(authentificationToProviderResponse).
		Post(fmt.Sprintf("/authentication/providers/configurations/%s", providerName))

	h.logger.Debugf("Response from AuthentificationToProvider: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during AuthentificationToProvider request")
	}

	if response.IsError() {
		return nil, errors.Errorf("authentication to provider failed with status code: %d", response.StatusCode())
	}

	h.logger.Infof("User %s authenticated successfully via provider %s", request.Login, providerName)

	return authentificationToProviderResponse, nil
}
