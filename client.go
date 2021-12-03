package centreon

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"time"

	"github.com/disaster37/go-centreon-rest/v21.10/api"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Config contain the value to access on Kibana API
type Config struct {
	Address          string
	Username         string
	Password         string
	DisableVerifySSL bool
	CAs              []string
	Timeout          time.Duration
	Debug            bool
}

// Client contain the REST client and the API specification
type Client struct {
	API    centreonapi.API
	config Config
}

// NewDefaultClient init client with empty config
func NewDefaultClient() (*Client, error) {
	return NewClient(Config{})
}

// NewClient init client with custom config
func NewClient(cfg Config) (*Client, error) {
	if cfg.Address == "" {
		cfg.Address = "http://localhost/api/v2.0"
	}

	restyClient := resty.New().
		SetHostURL(cfg.Address).
		SetHeader("Content-Type", "application/json").
		SetTimeout(cfg.Timeout).
		SetDebug(cfg.Debug).
		SetCookieJar(nil).
		SetRetryCount(2).
		SetQueryParams(map[string]string{
			"action": "action",
			"object": "centreon_clapi",
		})

	for _, path := range cfg.CAs {
		restyClient.SetRootCertificate(path)
	}

	if cfg.DisableVerifySSL == true {
		restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	client := &Client{
		API:    centreonapi.New(restyClient),
		config: cfg,
	}

	// handle refresh token when get Unauthorized
	restyClient.AddRetryCondition(func(r *resty.Response, e error) bool {
		if r.StatusCode() == http.StatusUnauthorized {
			if err := client.getToken(); err != nil {
				log.Errorf("Error when refresh token: %s", err.Error())
				return false
			}
			return true
		}
		return false
	})

	return client, nil

}

func (c *Client) Auth() (err error) {
	return c.getToken()
}

func (c *Client) getToken() (err error) {
	// Get Token
	resp, err := c.API.Client().R().
		SetFormData(map[string]string{
			"username": c.config.Username,
			"password": c.config.Password,
		}).
		SetQueryParams(map[string]string{
			"action": "authenticate",
			"object": "",
		}).
		Post("")
	if err != nil {
		return err
	}
	if resp.StatusCode() >= 300 {
		return errors.Errorf("Error when signin: %s", resp.Body())
	}
	result := map[string]string{}
	if err = json.Unmarshal(resp.Body(), &result); err != nil {
		return err
	}
	if result["authToken"] == "" {
		return errors.New("We get an empty token...")
	}
	c.API.Client().SetHeader("centreon-auth-token", result["authToken"])

	return nil
}
