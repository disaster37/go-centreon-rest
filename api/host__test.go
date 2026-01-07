package api

import "k8s.io/utils/ptr"

func (s *ApiTestSuite) TestHost() {

	// Create Host
	hostToCreate := &HostUpdateRequest{
		MonitoringServerId: 1,
		Name:               "test2",
		Alias:              "Test Host Alias",
		Address:            "127.0.0.1",
		IsActivated:        ptr.To(true),
	}

	createResp, err := s.api.Host().Create(hostToCreate)
	s.NoError(err)
	s.NotZero(createResp.Id)

	// Update Host
	hostToCreate.Alias = "Updated Test Host Alias"
	err = s.api.Host().Update(createResp.Id, hostToCreate)
	s.NoError(err)

	// Get Host
	getResp, err := s.api.Host().Get(createResp.Id)
	s.NoError(err)
	s.Equal("Updated Test Host Alias", getResp.Alias)

	// List all hosts and check if the created host is present
	listResp, err := s.api.Host().ListFromRealTime(nil)
	s.NoError(err)
	s.NotEmpty(listResp.Result)

	// List hosts with filter
	listRespWithFilter, err := s.api.Host().ListFromRealTime(&HostListOptions{
		ListOptions: ListOptions{
			Search: map[string]interface{}{
				"host.name": map[string]string{
					"$eq": "test",
				},
			},
		},
	})
	s.NoError(err)
	s.Equal(1, len(listRespWithFilter.Result))

	// Find all hosts and check if the created host is present
	findResp, err := s.api.Host().Find(nil)
	s.NoError(err)
	s.NotEmpty(findResp.Result)

	// Find with filter
	findRespWithFilter, err := s.api.Host().Find(&ListOptions{
		Search: map[string]interface{}{
			"name": map[string]string{
				"$eq": "test2",
			},
		},
	})
	s.NoError(err)
	s.Equal(1, len(findRespWithFilter.Result))

	// Delete Host
	err = s.api.Host().Delete(createResp.Id)
	s.NoError(err)

}
