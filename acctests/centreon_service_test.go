package acctests

import (
	"github.com/disaster37/go-centreon-rest/v21/models"
	"github.com/stretchr/testify/assert"
)

func (t *AccTestSuite) TestService() {
	var (
		err       error
		s         *models.ServiceGet
		expectedS *models.ServiceGet
	)

	// Create complete service
	macro := &models.Macro{
		Name:       "MAC1",
		Value:      "test",
		IsPassword: "0",
	}
	err = t.client.API.Service().Add("localhost", "test-acc", "template-test")
	assert.NoError(t.T(), err)
	err = t.client.API.Service().SetParam("localhost", "test-acc", "check_command", "ping")
	assert.NoError(t.T(), err)
	err = t.client.API.Service().SetServiceGroups("localhost", "test-acc", []string{"sg1"})
	assert.NoError(t.T(), err)
	err = t.client.API.Service().SetMacro("localhost", "test-acc", macro)
	assert.NoError(t.T(), err)
	err = t.client.API.Service().SetCategories("localhost", "test-acc", []string{"Ping"})
	assert.NoError(t.T(), err)
	err = t.client.API.Service().SetTraps("localhost", "test-acc", []string{"ccmCLIRunningConfigChanged"})
	assert.NoError(t.T(), err)

	// Get service
	expectedS = &models.ServiceGet{
		ServiceBaseGet: &models.ServiceBaseGet{
			HostName:            "localhost",
			Name:                "test-acc",
			CheckCommand:        "ping",
			Activated:           "1",
			PassiveCheckEnabled: "2",
			ActiveCheckEnabled:  "2",
		},
	}
	s, err = t.client.API.Service().Get("localhost", "test-acc")
	assert.NoError(t.T(), err)
	assert.NotNil(t.T(), s)
	assert.NotEmpty(t.T(), s.ID)
	assert.NotEmpty(t.T(), s.HostId)
	s.ID = ""
	s.HostId = ""
	assert.Equal(t.T(), expectedS, s)

	params, err := t.client.API.Service().GetParam("localhost", "test-acc", []string{"template", "comment"})
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), "template-test", params["template"])

	macros, err := t.client.API.Service().GetMacros("localhost", "test-acc")
	assert.NoError(t.T(), err)
	expectedMacro := &models.Macro{
		Name:   "MAC1",
		Value:  "test",
		Source: "direct",
	}
	assert.Equal(t.T(), expectedMacro, macros[0])

	sgs, err := t.client.API.Service().GetServiceGroups("localhost", "test-acc")
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), []string{"sg1"}, sgs)

	cats, err := t.client.API.Service().GetCategories("localhost", "test-acc")
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), []string{"Ping"}, cats)

	traps, err := t.client.API.Service().GetTraps("localhost", "test-acc")
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), []string{"ccmCLIRunningConfigChanged"}, traps)

	// Get not existing service
	s, err = t.client.API.Service().Get("localhost", "fake")
	assert.NoError(t.T(), err)
	assert.Nil(t.T(), s)

	// Update service
	// @TODO
}
