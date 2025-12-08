package api

func (s *ApiTestSuite) TestLogin() {
	resp, err := s.api.Authentification().Login(&AuthenticationRequest{
		Security: AuthenticationRequestSecurity{
			Credentials: AuthenticationRequestSecurityCredentials{
				Login:    username,
				Password: password,
			},
		},
	})
	s.NoError(err)
	s.NotEmpty(resp.Security.Token)
	s.NotEmpty(resp.Contact.Name)
}

func (s *ApiTestSuite) TestLogout() {
	resp, err := s.api.Authentification().Logout()
	s.NoError(err)
	s.NotEmpty(resp.Message)

	_, err = s.api.Authentification().Login(&AuthenticationRequest{
		Security: AuthenticationRequestSecurity{
			Credentials: AuthenticationRequestSecurityCredentials{
				Login:    username,
				Password: password,
			},
		},
	})
	if err != nil {
		s.FailNow(err.Error())
	}

}

func (s *ApiTestSuite) TestUpdatePassword() {
	_, err := s.api.Authentification().UpdatePassword(
		username,
		&PasswordUpdateRequest{
			OldPassword: password,
			NewPassword: password,
		},
	)

	s.NoError(err)

}

func (s *ApiTestSuite) TestGetProviders() {
	providers, err := s.api.Authentification().GetProviders()
	s.NoError(err)
	s.NotEmpty(providers)
	s.Equal("local", providers[0].Name)
}

func (s *ApiTestSuite) TestAuthentificationToProvider() {
	resp, err := s.api.Authentification().AuthentificationToProvider(
		"local",
		&AuthenticationRequestSecurityCredentials{
			Login:    username,
			Password: password,
		},
	)

	s.NoError(err)
	s.NotEmpty(resp.RedirectUri)
}
