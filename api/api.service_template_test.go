package centreonapi

import (
	"github.com/disaster37/go-centreon-rest/v21/models"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func (t *APITestSuite) TestServiceTemplateGet() {

	// Normale use case
	result := &ResultTest{
		Result: []*models.ServiceTemplateGet{
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
	service, err := t.client.ServiceTemplate().Get("host", "service")
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), "host", service.HostName)
	assert.Equal(t.T(), "service", service.Name)

	// When not found
	result = &ResultTest{
		Result: []*models.ServiceTemplateGet{},
	}
	responder, err = httpmock.NewJsonResponder(200, result)
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("POST", testURL, responder)
	service, err = t.client.ServiceTemplate().Get("host", "service")
	assert.NoError(t.T(), err)
	assert.Nil(t.T(), service)

	// When error
	responder, err = httpmock.NewJsonResponder(500, result)
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("POST", testURL, responder)
	_, err = t.client.ServiceTemplate().Get("host", "service")
	assert.Error(t.T(), err)
}

func (t *APITestSuite) TestServiceTemplateList() {

	// Normale use case
	result := &ResultTest{
		Result: []*models.ServiceTemplateGet{
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
	services, err := t.client.ServiceTemplate().List()
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), "host", services[0].HostName)
	assert.Equal(t.T(), "service", services[0].Name)

	// When not found
	result = &ResultTest{
		Result: []*models.ServiceTemplateGet{},
	}
	responder, err = httpmock.NewJsonResponder(200, result)
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("POST", testURL, responder)
	services, err = t.client.ServiceTemplate().List()
	assert.NoError(t.T(), err)
	assert.Empty(t.T(), services)

	// When error
	responder, err = httpmock.NewJsonResponder(500, result)
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("POST", testURL, responder)
	_, err = t.client.ServiceTemplate().List()
	assert.Error(t.T(), err)
}
