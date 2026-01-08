package api

import "k8s.io/utils/ptr"

func (s *ApiTestSuite) TestHostApi() {

	// Create Host
	hostToCreate := &HostCreateRequest{
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
	hostToUpdate := &HostUpdateRequest{
		Alias: ptr.To("Updated Test Host Alias"),
	}
	err = s.api.Host().Update(createResp.Id, hostToUpdate)
	s.NoError(err)

	// Get Host
	getResp, err := s.api.Host().Get(createResp.Id)
	s.NoError(err)
	s.Equal("Updated Test Host Alias", getResp.Alias)

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

	// Get host from real time by ID
	getFromRealTimeResp, err := s.api.Host().GetFromRealTime(1)
	s.NoError(err)
	s.NotNil(getFromRealTimeResp)
	s.Equal("test", getFromRealTimeResp.Name)

	// Get host from real time by Name
	getByNameFromRealTimeResp, err := s.api.Host().GetByNameFromRealTime("test")
	s.NoError(err)
	s.NotNil(getByNameFromRealTimeResp)
	s.Equal(int64(1), getByNameFromRealTimeResp.Id)

	// Count status from real time
	countStatusResp, err := s.api.Host().CountHostsByStatusFromRealTime()
	s.NoError(err)
	s.NotNil(countStatusResp)
	s.NotZero(countStatusResp.Total)

	// Delete Host
	err = s.api.Host().Delete(createResp.Id)
	s.NoError(err)

}
