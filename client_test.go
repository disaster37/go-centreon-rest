package centreon

import (
	"testing"

	"github.com/disaster37/go-centreon-rest/v21/models"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {

	cfg := &models.Config{
		Address:  "http://localhost",
		Username: "test",
		Password: "test",
	}

	client, err := NewClient(cfg)
	// Auth not work because it need centreon server
	assert.NoError(t, err)
	assert.NotNil(t, client)
}
