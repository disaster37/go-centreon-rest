package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	username = "admin"
	password = "f6Uq186LB4kX$"
	url      = "http://localhost:8080/centreon/api/v25.10"
)

type ApiTestSuite struct {
	suite.Suite
	client *resty.Client
	api    API
}

func (s *ApiTestSuite) SetupSuite() {

	// Init logger
	logrus.SetFormatter(new(prefixed.TextFormatter))
	logrus.SetLevel(logrus.DebugLevel)
	logger := logrus.NewEntry(logrus.New())

	restyClient := resty.New().
		SetBaseURL(url).
		SetHeader("Content-Type", "application/json").
		SetDebug(true)

	s.client = restyClient
	s.api = New(restyClient, logger)

	// Wait Centreon API is ready
	isOnline := false
	nbTry := 0
	for isOnline == false {
		_, err := s.api.Authentification().Login(&AuthenticationRequest{
			Security: AuthenticationRequestSecurity{
				Credentials: AuthenticationRequestSecurityCredentials{
					Login:    username,
					Password: password,
				},
			},
		})
		if err == nil {
			isOnline = true
		} else {
			logrus.Error(err.Error())
			time.Sleep(5 * time.Second)
			if nbTry == 10 {
				panic(fmt.Sprintf("We wait 50s that Centreon start: %s", err))
			}
			nbTry++
		}
	}

}

func (s *ApiTestSuite) TearDownSuite() {
	// Delete test hosts
	hostResult, err := s.api.Host().GetByName("test2")
	if err != nil {
		logrus.Error(err.Error())
	} else if hostResult != nil {
		err = s.api.Host().Delete(hostResult.Id)
		if err != nil {
			logrus.Error(err.Error())
		}
	}

	// Logout
	_, err = s.api.Authentification().Logout()
	if err != nil {
		logrus.Error(err.Error())
	}

}

func TestApiTestSuite(t *testing.T) {
	suite.Run(t, new(ApiTestSuite))
}
