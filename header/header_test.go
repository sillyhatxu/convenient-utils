package header

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"testing"
)

func TestGetAcceptLanguage(t *testing.T) {
	fmt.Println(language.Chinese)
	fmt.Println(language.English)
	tag, err := language.Parse("zh")
	assert.Nil(t, err)
	assert.EqualValues(t, tag, language.Chinese)
}

func TestGetCurrentUserId(t *testing.T) {
	context := gin.Context{}
	//context.Set(UseJWTKey, true)
	userId, err := GetCurrentUserId(&context, "")
	assert.Nil(t, err)
	assert.EqualValues(t, userId, "")
}
