package api

func (s *ApiTestSuite) TestCommandApi() {

	// Create Command
	commandToCreate := &CommandCreateRequest{
		Name:        "test2",
		Type:        CommandTypeCheck,
		CommandLine: "/usr/bin/test_command",
	}

	createResp, err := s.api.Command().Create(commandToCreate)
	s.NoError(err)
	s.NotNil(createResp)
	s.NotZero(createResp.Id)

	// Get Command
	getResp, err := s.api.Command().Get(createResp.Id)
	s.NoError(err)
	s.NotNil(getResp)

	// Get by name
	getByNameResp, err := s.api.Command().GetByName("test2")
	s.NoError(err)
	s.NotNil(getByNameResp)
	s.Equal(createResp.Id, getByNameResp.Id)

	// Find all Commands and check if the created command is present
	listResp, err := s.api.Command().Find(nil)
	s.NoError(err)
	s.NotEmpty(listResp.Result)

	// List Commands with filter
	listRespWithFilter, err := s.api.Command().Find(&ListOptions{
		Search: map[string]interface{}{
			"name": "test2",
		},
	})
	s.NoError(err)
	s.Equal(1, len(listRespWithFilter.Result))

}
