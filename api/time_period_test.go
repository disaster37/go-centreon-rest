package api

func (s *ApiTestSuite) TestTimePeriodApi() {

	// Create Time Period
	timePeriodToCreate := &TimePeriodCreateOrUpdateRequest{
		Name:  "test2",
		Alias: "Test Time Period Alias",
		Days: []TimePeriodDays{
			{
				Day:       TimeDayMonday,
				TimeRange: "00:00-24:00",
			},
		},
	}

	createResp, err := s.api.TimePeriod().Create(timePeriodToCreate)
	s.NoError(err)
	s.NotNil(createResp)
	s.NotZero(createResp.Id)

	// Update Time Period
	timePeriodToCreate.Alias = "Updated Test Time Period Alias"
	err = s.api.TimePeriod().Update(createResp.Id, timePeriodToCreate)
	s.NoError(err)

	// Get Time Period by id
	getResp, err := s.api.TimePeriod().Get(createResp.Id)
	s.NoError(err)
	s.NotNil(getResp)

	// Get Time Period by name
	getByNameResp, err := s.api.TimePeriod().GetByName("test2")
	s.NoError(err)
	s.NotNil(getByNameResp)
	s.Equal(createResp.Id, getByNameResp.Id)

	// List all Time Periods and check if the created time period is present
	listResp, err := s.api.TimePeriod().List(nil)
	s.NoError(err)
	s.NotEmpty(listResp.Result)

	// List Time Periods with filter
	listRespWithFilter, err := s.api.TimePeriod().List(&ListOptions{
		Search: map[string]interface{}{
			"name": "test2",
		},
	})
	s.NoError(err)
	s.Equal(1, len(listRespWithFilter.Result))

	//Delete Time Period
	err = s.api.TimePeriod().Delete(createResp.Id)
	s.NoError(err)

}
