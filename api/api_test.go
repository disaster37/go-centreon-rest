package centreonapi

import (
	"io"
	"net/http"
	"testing"

	"github.com/disaster37/go-centreon-rest/v21/models"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	testURL = "http://localhost/centreon/api/index.php?action=action&object=centreon_clapi"
	URL     = "http://localhost/centreon/api/index.php"
)

type APITestSuite struct {
	suite.Suite
	client API
}

func TestAPISuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}

func (t *APITestSuite) SetupTest() {
	restyClient := resty.New().
		SetHostURL(URL).
		SetHeader("Content-Type", "application/json").
		SetDebug(false).
		SetQueryParams(map[string]string{
			"action": "action",
			"object": "centreon_clapi",
		})
	cfg := &models.Config{
		Address:  URL,
		Username: "user",
		Password: "password",
	}
	httpmock.ActivateNonDefault(restyClient.GetClient())

	t.client = New(restyClient, cfg)
}

func (t *APITestSuite) BeforeTest(suiteName, testName string) {
	httpmock.Reset()
}

func (t *APITestSuite) TestAuth() {

	var data string
	httpmock.RegisterResponder("POST", "http://localhost/centreon/api/index.php?action=authenticate&object=", func(req *http.Request) (*http.Response, error) {
		b, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		data = string(b)
		resp, err := httpmock.NewJsonResponse(200, map[string]string{
			"authToken": "token",
		})
		if err != nil {
			panic(err)
		}
		return resp, nil
	})
	err := t.client.Auth()
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), "password=password&username=user", data)

}
