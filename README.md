[![build](https://github.com/disaster37/go-centreon-rest/actions/workflows/workflow.yaml/badge.svg)](https://github.com/disaster37/go-centreon-rest/actions/workflows/workflow.yaml)
[![GoDoc](https://godoc.org/github.com/disaster37/go-centreon-rest?status.svg)](http://godoc.org/github.com/disaster37/go-centreon-rest)
[![codecov](https://codecov.io/gh/disaster37/go-centreon-rest/branch/21.10.x/graph/badge.svg)](https://codecov.io/gh/disaster37/go-centreon-rest/branch/21.10.x)

# go-centreon-rest
Golang Rest client for Centreon
The GO client is actually use [Rest API v1](https://docs.centreon.com/current/fr/api/rest-api-v1.html) because of the [Rest API v2](https://docs.centreon.com/current/fr/api/rest-api-v2.html) not yet support the action to create, update and delete objects.

The following API call is implemented:
- Service / Service template:
  - show
  - add
  - del
  - setparam
  - getparam
  - getmacro
  - setmacro
  - delmacro
  - gettrap
  - settrap
  - deltrap
  - getcategory
  - setcategory
  - delcategory
  - getservicegroup
  - setservicegroup
  - delservicegroup

## Sample

```go

```