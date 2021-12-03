.PHONY: all fmt test build-mock build
all: help

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w ./

test: build-mock fmt
	go test ./... -v -count 1 -parallel 1 -race -coverprofile=coverage.txt -covermode=atomic $(TESTARGS) -timeout 120s

build-mock:
	go install github.com/golang/mock/mockgen@v1.6.0
	mockgen --build_flags=--mod=mod -destination=mocks/mock_api.go -package=mocks github.com/disaster37/go-centreon-rest/v21/api API
	mockgen --build_flags=--mod=mod -destination=mocks/mock_service.go -package=mocks github.com/disaster37/go-centreon-rest/v21/api ServiceAPI
	mockgen --build_flags=--mod=mod -destination=mocks/mock_service_template.go -package=mocks github.com/disaster37/go-centreon-rest/v21/api ServiceTemplateAPI

build: fmt
	go build .