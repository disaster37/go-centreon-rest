package centreonapi

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
)

const (
	testURL = "http://localhost/centreon/api/index.php?action=action&object=centreon_clapi"
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
		SetHostURL(testURL).
		SetHeader("Content-Type", "application/json").
		SetDebug(false)
	httpmock.ActivateNonDefault(restyClient.GetClient())

	t.client = New(restyClient)
}

func (t *APITestSuite) BeforeTest(suiteName, testName string) {
	httpmock.Reset()
}
