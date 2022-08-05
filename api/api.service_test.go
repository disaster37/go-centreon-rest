package centreonapi

import (
	"encoding/json"
	"net/http"

	"github.com/disaster37/go-centreon-rest/v21/models"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func (t *APITestSuite) TestServiceGet() {

	// Normale use case
	result := &ResultTest{
		Result: []*models.ServiceGet{
			{
				ServiceBaseGet: &models.ServiceBaseGet{
					HostName: "host",
					Name:     "service",
				},
			},
		},
	}
	responder, err := httpmock.NewJsonResponder(200, result)
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("POST", testURL, responder)
	service, err := t.client.Service().Get("host", "service")
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), "host", service.HostName)
	assert.Equal(t.T(), "service", service.Name)

	// When not found
	result = &ResultTest{
		Result: []*models.ServiceGet{},
	}
	responder, err = httpmock.NewJsonResponder(200, result)
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("POST", testURL, responder)
	service, err = t.client.Service().Get("host", "service")
	assert.NoError(t.T(), err)
	assert.Nil(t.T(), service)

	// When error
	responder, err = httpmock.NewJsonResponder(500, result)
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("POST", testURL, responder)
	_, err = t.client.Service().Get("host", "service")
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceList() {

	// Normale use case
	result := &ResultTest{
		Result: []*models.ServiceGet{
			{
				ServiceBaseGet: &models.ServiceBaseGet{
					HostName: "host",
					Name:     "service",
				},
			},
		},
	}
	responder, err := httpmock.NewJsonResponder(200, result)
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("POST", testURL, responder)
	services, err := t.client.Service().List()
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), "host", services[0].HostName)
	assert.Equal(t.T(), "service", services[0].Name)

	// When not found
	result = &ResultTest{
		Result: []*models.ServiceGet{},
	}
	responder, err = httpmock.NewJsonResponder(200, result)
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("POST", testURL, responder)
	services, err = t.client.Service().List()
	assert.NoError(t.T(), err)
	assert.Empty(t.T(), services)

	// When error
	responder, err = httpmock.NewJsonResponder(500, result)
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("POST", testURL, responder)
	_, err = t.client.Service().List()
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceSetHost() {

	var payload *Payload

	httpmock.RegisterResponder("POST", testURL, func(req *http.Request) (*http.Response, error) {
		payload = &Payload{}
		if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
			panic(err)
		}
		return httpmock.NewStringResponse(200, ""), nil
	})
	err := t.client.Service().SetHost("host1", "service", "host2")
	assert.NoError(t.T(), err)
	expectedPayload := &Payload{
		Action: "sethost",
		Object: "service",
		Values: "host1;service;host2",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	err = t.client.Service().SetHost("host1", "service", "host2")
	assert.Error(t.T(), err)
}
