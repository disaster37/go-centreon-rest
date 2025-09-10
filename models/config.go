package models

import (
	"time"

	"github.com/go-resty/resty/v2"
)

// Config contain the value to access on Kibana API
type Config struct {
	Address          string        `json:"address"`
	Username         string        `json:"username"`
	Password         string        `json:"password"`
	Token            string        `json:"token"`
	DisableVerifySSL bool          `json:"disableVerifySSL"`
	CAs              []string      `json:"cas"`
	Timeout          time.Duration `json:"timeout"`
	Debug            bool          `json:"debug"`
	Logger           resty.Logger  `json:"-"`
}
