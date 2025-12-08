package api

import (
	"fmt"
	"strings"

	"emperror.dev/errors"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type AuthenticationInterface interface {
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

// AuthentifcationImpl implements AuthenticationInterface
type AuthentifcationImpl struct {
	client *resty.Client
	logger *logrus.Entry
}

// NewAuthentication creates a new instance of AuthentifcationImpl
func NewAuthentication(client *resty.Client) AuthenticationInterface {
	return &AuthentifcationImpl{
		client: client,
		logger: logrus.WithField("module", "authentication"),
	}
}

// Login method to authenticate a user and obtain a token. It inject token in client header upon success.
func (h *AuthentifcationImpl) Login(request *AuthenticationRequest) (*LoginAuthenticationResponse, error) {
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
func (h *AuthentifcationImpl) Logout() (*LogoutAuthenticationResponse, error) {

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

func (h *AuthentifcationImpl) UpdatePassword(username string, request *PasswordUpdateRequest) (*PasswordUpdateResponse, error) {
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

func (h *AuthentifcationImpl) GetProviders() ([]ProviderConfiguration, error) {
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

func (h *AuthentifcationImpl) AuthentificationToProvider(providerName string, request *AuthenticationRequestSecurityCredentials) (*AuthentificationToProviderResponse, error) {

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

// AuthenticationRequest represents the login request structure
type AuthenticationRequest struct {
	Security AuthenticationRequestSecurity `json:"security" validate:"required"`
}

type AuthenticationRequestSecurity struct {
	Credentials AuthenticationRequestSecurityCredentials `json:"credentials" validate:"required"`
}

type AuthenticationRequestSecurityCredentials struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginAuthenticationResponse represents the response after successful login
type LoginAuthenticationResponse struct {
	Contact struct {
		ID      *int   `json:"id,omitempty"`
		Name    string `json:"name,omitempty"`
		Alias   string `json:"alias,omitempty"`
		Email   string `json:"email,omitempty"`
		IsAdmin *bool  `json:"is_admin,omitempty"`
	} `json:"contact,omitempty"`
	Security TokenResponse `json:"security,omitempty"`
}

// LogoutAuthenticationResponse represents the response after logout
type LogoutAuthenticationResponse struct {
	Message string `json:"message"`
}

// PasswordUpdateRequest represents the request structure for password update
type PasswordUpdateRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type PasswordUpdateResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// ProviderConfiguration represents authentication provider configuration
type ProviderConfiguration struct {
	ID                *int   `json:"id,omitempty"`
	Type              string `json:"type" validate:"required"`
	Name              string `json:"name" validate:"required"`
	AuthenticationURI string `json:"authentication_uri,omitempty"`
	IsActive          *bool  `json:"is_active,omitempty"`
	IsForced          *bool  `json:"is_forced,omitempty"`
}

// SAMLProviderConfiguration represents SAML provider configuration
type SAMLProviderConfiguration struct {
	ID                     *int   `json:"id,omitempty"`
	Type                   string `json:"type" validate:"required"`
	Name                   string `json:"name" validate:"required"`
	IsActive               *bool  `json:"is_active,omitempty"`
	IsForced               *bool  `json:"is_forced,omitempty"`
	RequestedAuthnContext  string `json:"requested_authn_context,omitempty"`
	EntityIDURL            string `json:"entity_id_url,omitempty"`
	SingleSignOnServiceURL string `json:"single_sign_on_service_url,omitempty"`
	SingleLogoutServiceURL string `json:"single_logout_service_url,omitempty"`
	X509Certificate        string `json:"x509_certificate,omitempty"`
}

// LDAPProviderConfiguration represents LDAP provider configuration
type LDAPProviderConfiguration struct {
	ID                 *int     `json:"id,omitempty"`
	Type               string   `json:"type" validate:"required"`
	Name               string   `json:"name" validate:"required"`
	IsActive           *bool    `json:"is_active,omitempty"`
	IsForced           *bool    `json:"is_forced,omitempty"`
	ConnectionSecurity string   `json:"connection_security,omitempty"`
	LDAPServers        []string `json:"ldap_servers,omitempty"`
	BaseDN             string   `json:"base_dn,omitempty"`
	LoginAttribute     string   `json:"login_attribute,omitempty"`
	BindDN             string   `json:"bind_dn,omitempty"`
	BindPassword       string   `json:"bind_password,omitempty"`
}

// TokenRequest represents a token generation request
type TokenRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// TokenResponse represents a token generation response
type TokenResponse struct {
	Token string `json:"token"`
}

// AuthentificationToProviderResponse is the reponse when call authentification on given provider
type AuthentificationToProviderResponse struct {
	RedirectUri string `json:"redirect_uri,omitempty"`
}

// Validation functions for AuthenticationRequest
func (ar *AuthenticationRequest) ValidateForLogin() error {

	if ar.Security.Credentials.Login == "" || len(strings.TrimSpace(ar.Security.Credentials.Login)) == 0 {
		return fmt.Errorf("login is required for authentication")
	}
	if ar.Security.Credentials.Password == "" || len(strings.TrimSpace(ar.Security.Credentials.Password)) == 0 {
		return fmt.Errorf("password is required for authentication")
	}
	return nil
}

// Validation functions for PasswordUpdateRequest
func (pur *PasswordUpdateRequest) ValidateForUpdate() error {
	if pur.OldPassword == "" || len(strings.TrimSpace(pur.OldPassword)) == 0 {
		return fmt.Errorf("old_password is required for password update")
	}
	if pur.NewPassword == "" || len(strings.TrimSpace(pur.NewPassword)) == 0 {
		return fmt.Errorf("new_password is required for password update")
	}
	if pur.OldPassword == pur.NewPassword {
		return fmt.Errorf("new_password must be different from old_password")
	}
	if len(pur.NewPassword) < 8 {
		return fmt.Errorf("new_password must be at least 8 characters long")
	}
	return nil
}
