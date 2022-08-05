package centreonapi

import (
	"encoding/json"
	"net/http"

	"github.com/disaster37/go-centreon-rest/v21/models"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func (t *APITestSuite) TestServiceGroupGet() {

	sgHandler := NewServiceGroup(t.client.Client())
	var payload *Payload

	// Normale use case
	result := &ResultTest{
		Result: []*models.ServiceGroup{
			{
				Name:        "serviceGroup",
				Description: "my SG",
				ID:          "1",
			},
		},
	}
	httpmock.RegisterResponder("POST", testURL, func(req *http.Request) (*http.Response, error) {
		payload = &Payload{}
		if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
			panic(err)
		}
		resp, err := httpmock.NewJsonResponse(200, result)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})
	serviceGroup, err := sgHandler.Get("serviceGroup")
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), "serviceGroup", serviceGroup.Name)
	assert.Equal(t.T(), "my SG", serviceGroup.Description)
	assert.Equal(t.T(), "1", serviceGroup.ID)
	expectedPayload := &Payload{
		Action: "show",
		Object: "sg",
		Values: "serviceGroup",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When not found
	result = &ResultTest{
		Result: []*models.ServiceGroup{},
	}
	responder, err := httpmock.NewJsonResponder(200, result)
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("POST", testURL, responder)
	serviceGroup, err = sgHandler.Get("serviceGroup")
	assert.NoError(t.T(), err)
	assert.Nil(t.T(), serviceGroup)

	// When error
	responder, err = httpmock.NewJsonResponder(500, result)
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("POST", testURL, responder)
	_, err = sgHandler.Get("serviceGroup")
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceGroupList() {

	sgHandler := NewServiceGroup(t.client.Client())
	var payload *Payload

	// Normale use case
	result := &ResultTest{
		Result: []*models.ServiceGroup{
			{
				Name:        "serviceGroup",
				Description: "my SG",
			},
		},
	}

	httpmock.RegisterResponder("POST", testURL, func(req *http.Request) (*http.Response, error) {
		payload = &Payload{}
		if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
			panic(err)
		}
		resp, err := httpmock.NewJsonResponse(200, result)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})
	serviceGroups, err := sgHandler.List()
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), "serviceGroup", serviceGroups[0].Name)
	assert.Equal(t.T(), "my SG", serviceGroups[0].Description)
	expectedPayload := &Payload{
		Action: "show",
		Object: "sg",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When not found
	result = &ResultTest{
		Result: []*models.ServiceGroup{},
	}
	responder, err := httpmock.NewJsonResponder(200, result)
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("POST", testURL, responder)
	serviceGroups, err = sgHandler.List()
	assert.NoError(t.T(), err)
	assert.Empty(t.T(), serviceGroups)

	// When error
	responder, err = httpmock.NewJsonResponder(500, result)
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("POST", testURL, responder)
	_, err = sgHandler.List()
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceGroupAdd() {

	sgHandler := NewServiceGroup(t.client.Client())
	var payload *Payload

	httpmock.RegisterResponder("POST", testURL, func(req *http.Request) (*http.Response, error) {
		payload = &Payload{}
		if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
			panic(err)
		}
		return httpmock.NewStringResponse(200, ""), nil
	})
	err := sgHandler.Add("serviceGroup", "my SG")
	assert.NoError(t.T(), err)
	expectedPayload := &Payload{
		Action: "add",
		Object: "sg",
		Values: "serviceGroup;my SG",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	err = sgHandler.Add("serviceGroup", "my SG")
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceGroupDelete() {

	sgHandler := NewServiceGroup(t.client.Client())
	var payload *Payload

	httpmock.RegisterResponder("POST", testURL, func(req *http.Request) (*http.Response, error) {
		payload = &Payload{}
		if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
			panic(err)
		}
		return httpmock.NewStringResponse(200, ""), nil
	})
	err := sgHandler.Delete("serviceGroup")
	assert.NoError(t.T(), err)
	expectedPayload := &Payload{
		Action: "del",
		Object: "sg",
		Values: "serviceGroup",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	err = sgHandler.Delete("serviceGroup")
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceGroupSetParam() {

	sgHandler := NewServiceGroup(t.client.Client())
	var payload *Payload

	httpmock.RegisterResponder("POST", testURL, func(req *http.Request) (*http.Response, error) {
		payload = &Payload{}
		if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
			panic(err)
		}
		return httpmock.NewStringResponse(200, ""), nil
	})
	err := sgHandler.SetParam("serviceGroup", "key", "value")
	assert.NoError(t.T(), err)
	expectedPayload := &Payload{
		Action: "setparam",
		Object: "sg",
		Values: "serviceGroup;key;value",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	err = sgHandler.SetParam("serviceGroup", "key", "value")
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceGroupGetParam() {

	sgHandler := NewServiceGroup(t.client.Client())
	var payload *Payload

	// Normal use case
	result := &ResultTest{
		Result: []string{"value"},
	}
	httpmock.RegisterResponder("POST", testURL, func(req *http.Request) (*http.Response, error) {
		payload = &Payload{}
		if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
			panic(err)
		}
		resp, err := httpmock.NewJsonResponse(200, result)
		if err != nil {
			panic(err)
		}
		return resp, nil
	})
	values, err := sgHandler.GetParam("serviceGroup", []string{"key"})
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), "value", values["key"])
	expectedPayload := &Payload{
		Action: "getparam",
		Object: "sg",
		Values: "serviceGroup;key",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	_, err = sgHandler.GetParam("serviceGroup", []string{"key"})
	assert.Error(t.T(), err)
}
