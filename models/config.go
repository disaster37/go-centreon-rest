package models

import (
	"time"

	"github.com/go-resty/resty/v2"
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
	Logger           resty.Logger
}
