package api

import "os"

func (s *ApiTestSuite) TestMediaApi() {

	// Create Media
	imageContend, err := os.ReadFile("fixtures/test.png")
	if err != nil {
		s.Fail(err.Error())
	}
	mediaToCreate := &MediaCreateRequest{
		Name:      "test2.png",
		Directory: "test",
		Data:      imageContend,
	}
	createResp, err := s.api.Media().Create(mediaToCreate)
	s.NoError(err)
	s.NotNil(createResp)
	s.NotZero(createResp.Result[0].Id)

	// Update Media
	updateResp, err := s.api.Media().Update(createResp.Result[0].Id, mediaToCreate.Name, mediaToCreate.Data)
	s.NoError(err)
	s.NotNil(updateResp)

	// Get Media by id
	getResp, err := s.api.Media().Get(createResp.Result[0].Id)
	s.NoError(err)
	s.NotNil(getResp)
	s.Equal("test2.png", getResp.Filename)

	// Get Media by name
	getByNameResp, err := s.api.Media().GetByName("test2.png")
	s.NoError(err)
	s.NotNil(getByNameResp)
	s.Equal(createResp.Result[0].Id, getByNameResp.Id)

	// Find all Media and check if the created media is present
	listResp, err := s.api.Media().Find(nil)
	s.NoError(err)
	s.NotEmpty(listResp.Result)

	// Find Media with filter
	listRespWithFilter, err := s.api.Media().Find(&ListOptions{
		Search: map[string]interface{}{
			"name": "test2.png",
		},
	})
	s.NoError(err)
	s.Equal(1, len(listRespWithFilter.Result))

	//Delete Media
	err = s.api.Media().Delete(createResp.Result[0].Id)
	s.NoError(err)

}
