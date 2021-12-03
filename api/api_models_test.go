package centreonapi

import (
	"github.com/stretchr/testify/assert"
)

func (t *APITestSuite) TestNewPayload() {

	expected := &Payload{
		Action: "show",
		Object: "service",
		Values: "key;value",
	}

	payload := NewPayload("show", "service", "%s;%s", "key", "value")

	assert.Equal(t.T(), expected, payload)

}
