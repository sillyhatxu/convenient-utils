package alioss

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	Provider    = "Provider"
	ImageBucket = "ImageBucket"
	Endpoint    = "Endpoint"
	AccessKey   = "AccessKey"
	SecretKey   = "SecretKey"
)

func init() {
	InitialOssClient(Provider, ImageBucket, AccessKey, SecretKey, Endpoint)
}

func TestGetImageURL(t *testing.T) {
	objectName := "photos/bl/2808bd726a018935bf794c1ba4ece2d3.png"
	image, err := Client.GetSizeImageURL(objectName, 720)
	assert.Nil(t, err)
	fmt.Println(image.URL)
	assert.NotNil(t, image)
}
