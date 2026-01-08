package api

import "k8s.io/utils/ptr"

func (s *ApiTestSuite) TestHostGroupApi() {

	// Create Host Group
	hostGroupToCreate := &HostGroupCreateOrUpdateRequest{
		Name:  "test2",
		Alias: ptr.To("Test Host Group Alias"),
	}

	createResp, err := s.api.HostGroup().Create(hostGroupToCreate)
	s.NoError(err)
	s.NotNil(createResp)
	s.NotZero(createResp.Id)

	// Update Host Group
	hostGroupToUpdate := &HostGroupCreateOrUpdateRequest{
		Name:  "test2",
		Alias: ptr.To("Updated Test Host Group Alias"),
	}
	err = s.api.HostGroup().Update(createResp.Id, hostGroupToUpdate)
	s.NoError(err)

	// Get Host Group by id
	getResp, err := s.api.HostGroup().Get(createResp.Id)
	s.NoError(err)
	s.NotNil(getResp)
	s.Equal("Updated Test Host Group Alias", *getResp.Alias)

	// Get Host Group by name
	getByNameResp, err := s.api.HostGroup().GetByName("test2")
	s.NoError(err)
	s.NotNil(getByNameResp)
	s.Equal(createResp.Id, getByNameResp.Id)

	// List all Host Groups and check if the created host group is present
	listResp, err := s.api.HostGroup().List(nil)
	s.NoError(err)
	s.NotEmpty(listResp.Result)

	// List Host Groups with filter
	listRespWithFilter, err := s.api.HostGroup().List(&ListOptions{
		Search: map[string]interface{}{
			"name": "test2",
		},
	})
	s.NoError(err)
	s.Equal(1, len(listRespWithFilter.Result))

	// List from real-time
	_, err = s.api.HostGroup().ListFromRealTime(nil)
	s.NoError(err)

	// List from real-time with filter
	_, err = s.api.HostGroup().ListFromRealTime(&ListOptions{
		Search: map[string]interface{}{
			"name": "test2",
		},
	})
	s.NoError(err)

	//Delete Host Group
	err = s.api.HostGroup().Delete(createResp.Id)
	s.NoError(err)

}
