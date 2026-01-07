package api

func (s *ApiTestSuite) Test_HostSeverity() {

	// Create Host Severity
	hostSeverityToCreate := &HostSeverityUpdateRequest{
		Name:   "test2",
		Alias:  "Test Host Severity Alias",
		Level:  1,
		IconId: 1,
	}
	hostSeverityResponse, err := s.api.HostSeverity().Create(hostSeverityToCreate)
	s.NoError(err)
	s.NotZero(hostSeverityResponse.Id)

	// Update Host Severity
	hostSeverityToCreate.Alias = "Updated Test Host Severity Alias"
	err = s.api.HostSeverity().Update(hostSeverityResponse.Id, hostSeverityToCreate)
	s.NoError(err)

	// Get Host Severity
	getResp, err := s.api.HostSeverity().Get(hostSeverityResponse.Id)
	s.NoError(err)
	s.NotNil(getResp)
	s.Equal("Updated Test Host Severity Alias", getResp.Alias)

	// Get by name
	getByNameResp, err := s.api.HostSeverity().GetByName("test2")
	s.NoError(err)
	s.NotNil(getByNameResp)
	s.Equal("Updated Test Host Severity Alias", getByNameResp.Alias)

	// List all Host Severities and check if the created host severity is present
	listResp, err := s.api.HostSeverity().List(nil)
	s.NoError(err)
	s.NotEmpty(listResp.Result)

	// List Host Severities with filter
	listRespWithFilter, err := s.api.HostSeverity().List(&ListOptions{
		Search: map[string]interface{}{
			"name": "test2",
		},
	})
	s.NoError(err)
	s.Equal(1, len(listRespWithFilter.Result))

	// Delete Host Severity
	err = s.api.HostSeverity().Delete(hostSeverityResponse.Id)
	s.NoError(err)

}
