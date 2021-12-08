package models

import "time"

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
