package api



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
	Contact  LoginAuthenticationResponseContact `json:"contact,omitempty"`
	Security TokenResponse                      `json:"security,omitempty"`
}

type LoginAuthenticationResponseContact struct {
	ID      *int   `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Alias   string `json:"alias,omitempty"`
	Email   string `json:"email,omitempty"`
	IsAdmin *bool  `json:"is_admin,omitempty"`
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
