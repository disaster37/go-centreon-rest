package api

func (s *ApiTestSuite) Test_HostTemplate() {

	// Create Host Template
	hostTemplateToCreate := &HostTemplateUpdateRequest{
		Name:  "test2",
		Alias: "Test Host Template Alias",
	}

	createResp, err := s.api.HostTemplate().Create(hostTemplateToCreate)
	s.NoError(err)
	s.NotZero(createResp.Id)

	// Update Host Template
	hostTemplateToCreate.Alias = "Updated Test Host Template Alias"
	err = s.api.HostTemplate().Update(createResp.Id, hostTemplateToCreate)
	s.NoError(err)

	// Get Host Template
	getResp, err := s.api.HostTemplate().Get(createResp.Id)
	s.NoError(err)
	s.Equal("Updated Test Host Template Alias", getResp.Alias)

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
