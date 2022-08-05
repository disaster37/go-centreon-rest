package acctests

import (
	"github.com/disaster37/go-centreon-rest/v21/models"
	"github.com/stretchr/testify/assert"
)

func (t *AccTestSuite) TestServiceGroup() {
	var (
		err       error
		s         *models.ServiceGroup
		expectedS *models.ServiceGroup
	)

	// Create complete serviceGroup
	err = t.client.API.ServiceGroup().Add("test-acc", "Test custom SG")
	assert.NoError(t.T(), err)
	err = t.client.API.ServiceGroup().SetParam("test-acc", "comment", "created from acc test")
	assert.NoError(t.T(), err)

	// Get service
	expectedS = &models.ServiceGroup{
		Name: "test-acc",
		Description: "Test custom SG",
	}
	s, err = t.client.API.ServiceGroup().Get("test-acc")
	assert.NoError(t.T(), err)
	assert.NotNil(t.T(), s)
	assert.NotEmpty(t.T(), s.ID)
	s.ID = ""
	assert.Equal(t.T(), expectedS, s)

	params, err := t.client.API.ServiceGroup().GetParam("test-acc", []string{"activate", "comment"})
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), "created from acc test", params["comment"])
	assert.Equal(t.T(), "1", params["activate"])


	// Get not existing service
	s, err = t.client.API.ServiceGroup().Get("fake")
	assert.NoError(t.T(), err)
	assert.Nil(t.T(), s)

	// Update service
	// @TODO
}
