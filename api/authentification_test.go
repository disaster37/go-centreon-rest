package api

func (s *ApiTestSuite) TestAuthentificationApi() {

	// Test login
	respLogin, err := s.api.Authentification().Login(&AuthenticationRequest{
		Security: AuthenticationRequestSecurity{
			Credentials: AuthenticationRequestSecurityCredentials{
				Login:    username,
				Password: password,
			},
		},
	})
	s.NoError(err)
	s.NotEmpty(respLogin.Security.Token)
	s.Equal("admin_admin", respLogin.Contact.Name)
	s.Equal("admin", respLogin.Contact.Alias)
	s.Equal("admin@no.no", respLogin.Contact.Email)
	s.True(*respLogin.Contact.IsAdmin)
	s.Equal(1, respLogin.Contact.Id)

	// Login must failed with wrong password
	_, err = s.api.Authentification().Login(&AuthenticationRequest{
		Security: AuthenticationRequestSecurity{
			Credentials: AuthenticationRequestSecurityCredentials{
				Login:    username,
				Password: "wrongPassword",
			},
		},
	})
	s.Error(err)

	// Test logout
	respLogout, err := s.api.Authentification().Logout()
	s.NoError(err)
	s.NotEmpty(respLogout.Message)

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

	// Test change password
	_, err = s.api.Authentification().UpdatePassword(
		username,
		&PasswordUpdateRequest{
			OldPassword: password,
			NewPassword: password,
		},
	)
	s.NoError(err)

	// Test change password with wrong old password
	_, err = s.api.Authentification().UpdatePassword(
		username,
		&PasswordUpdateRequest{
			OldPassword: "wrongOldPassword",
			NewPassword: password,
		},
	)
	s.Error(err)

	// Test get providers
	providers, err := s.api.Authentification().GetProviders()
	s.NoError(err)
	s.NotEmpty(providers)
	s.Equal(1, providers[0].Id)
	s.Equal("local", providers[0].Name)
	s.Equal("local", providers[0].Type)
	s.True(*providers[0].IsActive)
	s.True(*providers[0].IsForced)
	s.NotEmpty(providers[0].AuthenticationURI)

	// Test authentification to provider
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
