[![build](https://github.com/disaster37/go-centreon-rest/actions/workflows/workflow.yaml/badge.svg)](https://github.com/disaster37/go-centreon-rest/actions/workflows/workflow.yaml)
[![GoDoc](https://godoc.org/github.com/disaster37/go-centreon-rest?status.svg)](http://godoc.org/github.com/disaster37/go-centreon-rest)
[![Go Report Card](https://goreportcard.com/badge/github.com/disaster37/go-centreon-rest)](https://goreportcard.com/report/github.com/disaster37/go-centreon-rest)
[![codecov](https://codecov.io/gh/disaster37/go-centreon-rest/branch/21.10.x/graph/badge.svg)](https://codecov.io/gh/disaster37/go-centreon-rest/branch/21.10.x)

# go-centreon-rest
Golang Rest client for Centreon
The GO client is actually use [Rest API v1](https://docs.centreon.com/current/fr/api/rest-api-v1.html) because of the [Rest API v2](https://docs.centreon.com/current/fr/api/rest-api-v2.html) not yet support the action to create, update and delete objects.

This project provide a native API call and a handler  to manage Centreon object with best user friendly way.

## Native API call

**Authentification**:
 - [X] Login
 - [X] Logout
 - [X] Update password
 - [X] Get providers
 - [X] Login on specific provider


**Host**:
- [X] Create host
- [X] Update host
- [X] Delete host
- [] Get host: The endpoint not exist in Centreon Rest API v2. We use Find instead but not all fields are retrieve.
- [X] Find host
- [X] Get host from real time API
- [X] List hosts from real time API
- [X] Count hosts status from real time API

**Host template**:
- [X] Create host template
- [X] Update host template
- [X] Delete host template
- [X] Get host template: The endpoint not exist in Centreon Rest API v2. We use Find instead.
- [X] Find host template

**Host category**:
- [X] Create host category
- [X] Update host category
- [X] Delete host category
- [X] Get host category
- [X] List host category. List with filter not working in Centreon Rest API v2.
- [X] List host category from real time API

**Host severity**:
- [X] Create host severity
- [X] Update host severity
- [X] Delete host severity
- [X] Get host severity
- [X] List host severity

**Host group**:
- [X] Create host group
- [X] Update host group
- [X] Delete host group
- [X] Get host group
- [X] List host group
- [] Delete multiple hosts groups
- [] Enable/disable multiple host groups
- [] List all host groups by host id
- [X] List host group from real time API