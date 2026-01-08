package api

import "k8s.io/utils/ptr"

func (s *ApiTestSuite) TestHostTemplateApi() {

	// Create Host Template
	hostTemplateToCreate := &HostTemplateCreateRequest{
		Name:  "test2",
		Alias: "Test Host Template Alias",
	}

	createResp, err := s.api.HostTemplate().Create(hostTemplateToCreate)
	s.NoError(err)
	s.NotZero(createResp.Id)

	// Update Host Template
	hostTemplateToUpdate := &HostTemplateUpdateRequest{
		Alias: ptr.To("Updated Test Host Template Alias"),
	}
	err = s.api.HostTemplate().Update(createResp.Id, hostTemplateToUpdate)
	s.NoError(err)

	// Get Host Template by id
	getResp, err := s.api.HostTemplate().Get(createResp.Id)
	s.NoError(err)
	s.NotNil(getResp)
	s.Equal("Updated Test Host Template Alias", getResp.Alias)

	// Get Host Template by name
	getByNameResp, err := s.api.HostTemplate().GetByName("test2")
	s.NoError(err)
	s.NotNil(getByNameResp)
	s.Equal(createResp.Id, getByNameResp.Id)

	// List all Host Templates and check if the created host template is present
	listResp, err := s.api.HostTemplate().Find(nil)
	s.NoError(err)
	s.NotEmpty(listResp.Result)

	// List Host Templates with filter
	listRespWithFilter, err := s.api.HostTemplate().Find(&ListOptions{
		Search: map[string]interface{}{
			"name": "test2",
		},
	})
	s.NoError(err)
	s.Equal(1, len(listRespWithFilter.Result))

	//Delete Host Template
	err = s.api.HostTemplate().Delete(createResp.Id)
	s.NoError(err)
}
