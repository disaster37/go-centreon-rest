package centreonapi

import (
	"encoding/json"
	"net/http"

	"github.com/disaster37/go-centreon-rest/v21.10/models"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func (t *APITestSuite) TestServiceBaseGet() {

	serviceBase := newServiceBase(t.client.Client(), objectService)
	var payload *Payload

	// Normale use case
	result := &ResultTest{
		Result: []*models.ServiceBaseGet{
			{
				HostName: "host",
				Name:     "service",
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
	service, err := serviceBase.Get("host", "service")
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), "host", service.HostName)
	assert.Equal(t.T(), "service", service.Name)
	expectedPayload := &Payload{
		Action: "show",
		Object: "service",
		Values: "host;service",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When not found
	result = &ResultTest{
		Result: []*models.ServiceBaseGet{},
	}
	responder, err := httpmock.NewJsonResponder(200, result)
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("POST", testURL, responder)
	service, err = serviceBase.Get("host", "service")
	assert.NoError(t.T(), err)
	assert.Nil(t.T(), service)

	// When error
	responder, err = httpmock.NewJsonResponder(500, result)
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("POST", testURL, responder)
	service, err = serviceBase.Get("host", "service")
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceBaseList() {

	serviceBase := newServiceBase(t.client.Client(), objectService)
	var payload *Payload

	// Normale use case
	result := &ResultTest{
		Result: []*models.ServiceBaseGet{
			{
				HostName: "host",
				Name:     "service",
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
	services, err := serviceBase.List()
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), "host", services[0].HostName)
	assert.Equal(t.T(), "service", services[0].Name)
	expectedPayload := &Payload{
		Action: "show",
		Object: "service",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When not found
	result = &ResultTest{
		Result: []*models.ServiceBaseGet{},
	}
	responder, err := httpmock.NewJsonResponder(200, result)
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("POST", testURL, responder)
	services, err = serviceBase.List()
	assert.NoError(t.T(), err)
	assert.Empty(t.T(), services)

	// When error
	responder, err = httpmock.NewJsonResponder(500, result)
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("POST", testURL, responder)
	services, err = serviceBase.List()
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceBaseAdd() {

	serviceBase := newServiceBase(t.client.Client(), objectService)
	var payload *Payload

	httpmock.RegisterResponder("POST", testURL, func(req *http.Request) (*http.Response, error) {
		payload = &Payload{}
		if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
			panic(err)
		}
		return httpmock.NewStringResponse(200, ""), nil
	})
	err := serviceBase.Add("host", "service", "template")
	assert.NoError(t.T(), err)
	expectedPayload := &Payload{
		Action: "add",
		Object: "service",
		Values: "host;service;template",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	err = serviceBase.Add("host", "service", "template")
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceBaseDelete() {

	serviceBase := newServiceBase(t.client.Client(), objectService)
	var payload *Payload

	httpmock.RegisterResponder("POST", testURL, func(req *http.Request) (*http.Response, error) {
		payload = &Payload{}
		if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
			panic(err)
		}
		return httpmock.NewStringResponse(200, ""), nil
	})
	err := serviceBase.Delete("host", "service")
	assert.NoError(t.T(), err)
	expectedPayload := &Payload{
		Action: "del",
		Object: "service",
		Values: "host;service",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	err = serviceBase.Delete("host", "service")
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceBaseSetParam() {

	serviceBase := newServiceBase(t.client.Client(), objectService)
	var payload *Payload

	httpmock.RegisterResponder("POST", testURL, func(req *http.Request) (*http.Response, error) {
		payload = &Payload{}
		if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
			panic(err)
		}
		return httpmock.NewStringResponse(200, ""), nil
	})
	err := serviceBase.SetParam("host", "service", "key", "value")
	assert.NoError(t.T(), err)
	expectedPayload := &Payload{
		Action: "setparam",
		Object: "service",
		Values: "host;service;key;value",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	err = serviceBase.SetParam("host", "service", "key", "value")
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceBaseGetParam() {

	serviceBase := newServiceBase(t.client.Client(), objectService)
	var payload *Payload

	// Normal use case
	result := &ResultTest{
		Result: []map[string]string{
			{
				"key": "value",
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
	values, err := serviceBase.GetParam("host", "service", []string{"key"})
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), "value", values["key"])
	expectedPayload := &Payload{
		Action: "getparam",
		Object: "service",
		Values: "host;service;key",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	values, err = serviceBase.GetParam("host", "service", []string{"key"})
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceBaseGetMacros() {

	serviceBase := newServiceBase(t.client.Client(), objectService)
	var payload *Payload

	// Normal use case
	result := &ResultTest{
		Result: []*models.Macro{
			{
				Name:  "macro",
				Value: "value",
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
	macros, err := serviceBase.GetMacros("host", "service")
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), "macro", macros[0].Name)
	assert.Equal(t.T(), "value", macros[0].Value)
	expectedPayload := &Payload{
		Action: "getmacro",
		Object: "service",
		Values: "host;service",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	macros, err = serviceBase.GetMacros("host", "service")
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceBaseSetMacro() {

	serviceBase := newServiceBase(t.client.Client(), objectService)
	var payload *Payload

	httpmock.RegisterResponder("POST", testURL, func(req *http.Request) (*http.Response, error) {
		payload = &Payload{}
		if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
			panic(err)
		}
		return httpmock.NewStringResponse(200, ""), nil
	})
	macro := &models.Macro{
		Name:       "macro",
		Value:      "value",
		IsPassword: "0",
	}
	err := serviceBase.SetMacro("host", "service", macro)
	assert.NoError(t.T(), err)
	expectedPayload := &Payload{
		Action: "setmacro",
		Object: "service",
		Values: "host;service;MACRO;value;0;",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	err = serviceBase.SetMacro("host", "service", macro)
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceBaseDeleteMacro() {

	serviceBase := newServiceBase(t.client.Client(), objectService)
	var payload *Payload

	httpmock.RegisterResponder("POST", testURL, func(req *http.Request) (*http.Response, error) {
		payload = &Payload{}
		if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
			panic(err)
		}
		return httpmock.NewStringResponse(200, ""), nil
	})
	err := serviceBase.DeleteMacro("host", "service", "macro")
	assert.NoError(t.T(), err)
	expectedPayload := &Payload{
		Action: "delmacro",
		Object: "service",
		Values: "host;service;macro",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	err = serviceBase.DeleteMacro("host", "service", "macro")
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceBaseGetCategories() {

	serviceBase := newServiceBase(t.client.Client(), objectService)
	var payload *Payload

	// Normal use case
	result := &ResultTest{
		Result: []map[string]string{
			{
				"name": "category",
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
	categories, err := serviceBase.GetCategories("host", "service")
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), "category", categories[0])
	expectedPayload := &Payload{
		Action: "getcategory",
		Object: "service",
		Values: "host;service",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	categories, err = serviceBase.GetCategories("host", "service")
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceBaseSetCategories() {

	serviceBase := newServiceBase(t.client.Client(), objectService)
	var payload *Payload

	httpmock.RegisterResponder("POST", testURL, func(req *http.Request) (*http.Response, error) {
		payload = &Payload{}
		if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
			panic(err)
		}
		return httpmock.NewStringResponse(200, ""), nil
	})
	err := serviceBase.SetCategories("host", "service", []string{"cat1", "cat2"})
	assert.NoError(t.T(), err)
	expectedPayload := &Payload{
		Action: "setcategory",
		Object: "service",
		Values: "host;service;cat1|cat2",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	err = serviceBase.SetCategories("host", "service", []string{"cat1", "cat2"})
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceBaseDeleteCategories() {

	serviceBase := newServiceBase(t.client.Client(), objectService)
	var payload *Payload

	httpmock.RegisterResponder("POST", testURL, func(req *http.Request) (*http.Response, error) {
		payload = &Payload{}
		if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
			panic(err)
		}
		return httpmock.NewStringResponse(200, ""), nil
	})
	err := serviceBase.DeleteCategories("host", "service", []string{"cat1", "cat2"})
	assert.NoError(t.T(), err)
	expectedPayload := &Payload{
		Action: "delcategory",
		Object: "service",
		Values: "host;service;cat1|cat2",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	err = serviceBase.DeleteCategories("host", "service", []string{"cat1", "cat2"})
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceBaseGetServiceGroups() {

	serviceBase := newServiceBase(t.client.Client(), objectService)
	var payload *Payload

	// Normal use case
	result := &ResultTest{
		Result: []map[string]string{
			{
				"name": "sg",
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
	serviceGroups, err := serviceBase.GetServiceGroups("host", "service")
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), "sg", serviceGroups[0])
	expectedPayload := &Payload{
		Action: "getservicegroup",
		Object: "service",
		Values: "host;service",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	serviceGroups, err = serviceBase.GetServiceGroups("host", "service")
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceBaseSetServiceGroups() {

	serviceBase := newServiceBase(t.client.Client(), objectService)
	var payload *Payload

	httpmock.RegisterResponder("POST", testURL, func(req *http.Request) (*http.Response, error) {
		payload = &Payload{}
		if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
			panic(err)
		}
		return httpmock.NewStringResponse(200, ""), nil
	})
	err := serviceBase.SetServiceGroups("host", "service", []string{"sg1", "sg2"})
	assert.NoError(t.T(), err)
	expectedPayload := &Payload{
		Action: "setservicegroup",
		Object: "service",
		Values: "host;service;sg1|sg2",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	err = serviceBase.SetServiceGroups("host", "service", []string{"sg1", "sg2"})
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceBaseDeleteServiceGroups() {

	serviceBase := newServiceBase(t.client.Client(), objectService)
	var payload *Payload

	httpmock.RegisterResponder("POST", testURL, func(req *http.Request) (*http.Response, error) {
		payload = &Payload{}
		if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
			panic(err)
		}
		return httpmock.NewStringResponse(200, ""), nil
	})
	err := serviceBase.DeleteServiceGroups("host", "service", []string{"sg1", "sg2"})
	assert.NoError(t.T(), err)
	expectedPayload := &Payload{
		Action: "delservicegroup",
		Object: "service",
		Values: "host;service;sg1|sg2",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	err = serviceBase.DeleteServiceGroups("host", "service", []string{"sg1", "sg2"})
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceBaseGetTraps() {

	serviceBase := newServiceBase(t.client.Client(), objectService)
	var payload *Payload

	// Normal use case
	result := &ResultTest{
		Result: []map[string]string{
			{
				"name": "trap",
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
	traps, err := serviceBase.GetTraps("host", "service")
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), "trap", traps[0])
	expectedPayload := &Payload{
		Action: "gettrap",
		Object: "service",
		Values: "host;service",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	traps, err = serviceBase.GetTraps("host", "service")
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceBaseSetTraps() {

	serviceBase := newServiceBase(t.client.Client(), objectService)
	var payload *Payload

	httpmock.RegisterResponder("POST", testURL, func(req *http.Request) (*http.Response, error) {
		payload = &Payload{}
		if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
			panic(err)
		}
		return httpmock.NewStringResponse(200, ""), nil
	})
	err := serviceBase.SetTraps("host", "service", []string{"trap1", "trap2"})
	assert.NoError(t.T(), err)
	expectedPayload := &Payload{
		Action: "settrap",
		Object: "service",
		Values: "host;service;trap1|trap2",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	err = serviceBase.SetTraps("host", "service", []string{"trap1", "trap2"})
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceBaseDeleteTraps() {

	serviceBase := newServiceBase(t.client.Client(), objectService)
	var payload *Payload

	httpmock.RegisterResponder("POST", testURL, func(req *http.Request) (*http.Response, error) {
		payload = &Payload{}
		if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
			panic(err)
		}
		return httpmock.NewStringResponse(200, ""), nil
	})
	err := serviceBase.DeleteTraps("host", "service", []string{"trap1", "trap2"})
	assert.NoError(t.T(), err)
	expectedPayload := &Payload{
		Action: "deltrap",
		Object: "service",
		Values: "host;service;trap1|trap2",
	}
	assert.Equal(t.T(), expectedPayload, payload)

	// When error
	httpmock.RegisterResponder("POST", testURL, httpmock.NewStringResponder(500, ""))
	err = serviceBase.DeleteTraps("host", "service", []string{"trap1", "trap2"})
	assert.Error(t.T(), err)
}
