package gohttp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHttpGet(t *testing.T) {
	httpClient := New("http://localhost:18883/order-internal-api", "/orders/O28446394101553111448058095558")
	response, err := httpClient.HttpGet()
	assert.Nil(t, err)
	httpResponse, err := GetHttpResponseEntity(response)
	assert.Nil(t, err)
	assert.NotNil(t, httpResponse)
}
