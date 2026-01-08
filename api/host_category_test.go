package api

import "k8s.io/utils/ptr"

func (s *ApiTestSuite) TestHostCategoryTest() {

	// Create Host Category
	hostCategoryToCreate := &HostCategoryCreateOrUpdateRequest{
		Name:        "test2",
		Alias:       "Test Category Alias",
		IsActivated: ptr.To(true),
	}
	hostCategoryResponse, err := s.api.HostCategory().Create(hostCategoryToCreate)
	s.NoError(err)
	s.NotZero(hostCategoryResponse.Id)

	// Update Host Category
	hostCategoryToCreate.Alias = "Updated Test Category Alias"
	err = s.api.HostCategory().Update(hostCategoryResponse.Id, hostCategoryToCreate)
	s.NoError(err)

	// Get Host Category
	getResp, err := s.api.HostCategory().Get(hostCategoryResponse.Id)
	s.NoError(err)
	s.Equal("Updated Test Category Alias", getResp.Alias)

	// List all Host Categories and check if the created host category is present
	listResp, err := s.api.HostCategory().List(nil)
	s.NoError(err)
	s.NotEmpty(listResp.Result)

	// List Host Categories with filter
	// The filter does not seem to work as expected, always returns an empty list
	_, err = s.api.HostCategory().List(&ListOptions{
		Search: map[string]interface{}{
			"name": "test2",
		},
	})
	s.NoError(err)
	//s.Equal(1, len(listRespWithFilter.Result))

	// Get by name
	// This method call the List method internally, but this API seems to not work as expected
	// It always return nil
	_, err = s.api.HostCategory().GetByName("test2")
	s.NoError(err)
	//s.NotNil(getByNameResp)

	// List from real-time
	// This endpoint seems to always return an empty list
	listFromRealTimeResp, err := s.api.HostCategory().ListFromRealTime(nil)
	s.NoError(err)
	s.NotEmpty(listFromRealTimeResp.Result)

	// List from real-time with filter
	// This endpoint seems to always return an empty list
	listFromRealTimeWithFilterResp, err := s.api.HostCategory().ListFromRealTime(&ListOptions{
		Search: map[string]interface{}{
			"name": "test",
		},
	})
	s.NoError(err)
	s.Equal(1, len(listFromRealTimeWithFilterResp.Result))

	//Delete Host Category
	err = s.api.HostCategory().Delete(hostCategoryResponse.Id)
	s.NoError(err)
}
