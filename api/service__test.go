package api

import "k8s.io/utils/ptr"

func (s *ApiTestSuite) TestServiceApi() {

	// Create a new service
	serviceCreateRequest := &ServiceCreateRequest{
		Name:                     "test2",
		Comment:                  ptr.To("test"),
		HostId:                   1,
		CheckCommandId:           ptr.To(int64(91)),
		CheckTimeperiodId:        ptr.To(int64(1)),
		NotificationTimeperiodId: ptr.To(int64(1)),
	}
	createServiceResponse, err := s.api.Service().Create(serviceCreateRequest)
	s.NoError(err)
	s.NotNil(createServiceResponse)
	s.Equal("test2", createServiceResponse.Name)

	// Update the service
	serviceUpdateRequest := &ServiceUpdateRequest{
		Comment: ptr.To("updated test"),
	}
	err = s.api.Service().Update(createServiceResponse.Id, serviceUpdateRequest)
	s.NoError(err)

	// Get service by ID
	getServiceResponse, err := s.api.Service().Get(createServiceResponse.Id)
	s.NoError(err)
	s.NotNil(getServiceResponse)
	s.Equal("updated test", *getServiceResponse.Comment)

	// Get service by name
	getByNameResponse, err := s.api.Service().GetByName("test2")
	s.NoError(err)
	s.NotNil(getByNameResponse)
	s.Equal(createServiceResponse.Id, getByNameResponse.Id)

	// Delete the service
	err = s.api.Service().Delete(createServiceResponse.Id)
	s.NoError(err)

}
