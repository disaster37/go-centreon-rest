package api

func (s *ApiTestSuite) TestCommandApi() {

	// Create Command
	commandToCreate := &CommandCreateOrUpdateRequest{
		Name:        "test2",
		Type:        CommandTypeCheck,
		CommandLine: "/usr/bin/test_command",
	}

	createResp, err := s.api.Command().Create(commandToCreate)
	s.NoError(err)
	s.NotNil(createResp)
	s.NotZero(createResp.Id)

	// Update Command
	commandToCreate.CommandLine = "/usr/bin/updated_test_command"
	err = s.api.Command().Update(createResp.Id, commandToCreate)
	s.NoError(err)

	// Get Command
	getResp, err := s.api.Command().Get(createResp.Id)
	s.NoError(err)
	s.NotNil(getResp)
	s.Equal("test2", getResp.Name)

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

	// Delete Command
	err = s.api.Command().Delete(createResp.Id)
	s.NoError(err)

}
