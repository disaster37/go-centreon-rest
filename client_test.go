package centreon

import (
	"io"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {

	cfg := Config{
		Address:  "http://localhost",
		Username: "test",
		Password: "test",
	}

	client, err := NewClient(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, client)
}

func TestAuth(t *testing.T) {
	cfg := Config{
		Address:  "http://localhost",
		Username: "user",
		Password: "password",
	}

	client, _ := NewClient(cfg)

	httpmock.ActivateNonDefault(client.API.Client().GetClient())
	var data string
	httpmock.RegisterResponder("POST", "http://localhost?action=authenticate&object=", func(req *http.Request) (*http.Response, error) {
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
	err := client.Auth()
	assert.NoError(t, err)
	assert.Equal(t, "password=password&username=user", data)

}
