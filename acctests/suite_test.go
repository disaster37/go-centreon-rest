package acctests

import (
	"os"
	"testing"

	"github.com/disaster37/go-centreon-rest/v21"
	"github.com/disaster37/go-centreon-rest/v21/models"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type AccTestSuite struct {
	suite.Suite
	client *centreon.Client
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(AccTestSuite))
}

func (t *AccTestSuite) SetupSuite() {
	logrus.SetLevel(logrus.DebugLevel)

	// Init Centreon client
	url, found := os.LookupEnv("CENTREON_URL")
	if !found {
		url = "http://localhost/centreon/api/index.php"
	}
	username, found := os.LookupEnv("CENTREON_USERNAME")
	if !found {
		username = "admin"
	}
	password, found := os.LookupEnv("CENTREON_PASSWORD")
	if !found {
		password = "admin"
	}
	disableSSLCheck := os.Getenv("CENTREON_DISABLE_SSL_CHECK")
	cfg := &models.Config{
		Address:  url,
		Username: username,
		Password: password,
		Debug:    true,
	}
	if disableSSLCheck == "true" {
		cfg.DisableVerifySSL = true
	} else {
		cfg.DisableVerifySSL = false
	}
	centreonClient, err := centreon.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	t.client = centreonClient

	logrus.Infof("Connect to centreon: %s", url)

	err = t.client.API.Auth()
	if err != nil {
		panic(err)
	}

}

func (t *AccTestSuite) BeforeTest(suiteName, testName string) {
}

func (t *AccTestSuite) AfterTest(suiteName, testName string) {
}
