package centreon

import (
	"crypto/tls"
	"net/http"
	"time"

	centreonapi "github.com/disaster37/go-centreon-rest/v21/api"
	"github.com/disaster37/go-centreon-rest/v21/models"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// Client contain the REST client and the API specification
type Client struct {
	API    centreonapi.API
	config *models.Config
}

// NewDefaultClient init client with empty config
func NewDefaultClient() (*Client, error) {
	return NewClient(&models.Config{})
}

// NewClient init client with custom config
func NewClient(cfg *models.Config) (*Client, error) {
	if cfg.Address == "" {
		cfg.Address = "http://localhost/api/v2.0"
	}

	restyClient := resty.New().
		SetBaseURL(cfg.Address).
		SetHeader("Content-Type", "application/json").
		SetTimeout(cfg.Timeout).
		SetDebug(cfg.Debug).
		SetCookieJar(nil).
		SetRetryCount(1).
		SetRetryWaitTime(time.Second * 1).
		SetQueryParams(map[string]string{
			"action": "action",
			"object": "centreon_clapi",
		})

	if cfg.Logger != nil {
		restyClient.SetLogger(cfg.Logger)
	}

	for _, path := range cfg.CAs {
		restyClient.SetRootCertificate(path)
	}

	if cfg.DisableVerifySSL {
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	client := &Client{
		API:    centreonapi.New(restyClient, cfg),
		config: cfg,
	}

	// handle refresh token when get Unauthorized
	restyClient.AddRetryCondition(func(r *resty.Response, e error) bool {
		if r.StatusCode() == http.StatusUnauthorized || r.StatusCode() == http.StatusForbidden {
			if err := client.API.Auth(); err != nil {
				logrus.Errorf("Error when refresh token: %s", err.Error())
				time.Sleep(1 * time.Second)
				return true
			}
		}
		return false
	})

	return client, nil

}
