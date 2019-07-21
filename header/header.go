package header

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sillyhatxu/convenient-utils/webjwt"
	"golang.org/x/text/language"
)

const UseJWTKey = "USE_JWT"

func GetDeviceId(context *gin.Context) string {
	values := context.Request.Header["Device-Id"]
	if len(values) > 0 {
		return values[0]
	}
	return ""
}

func GetAcceptLanguage(context *gin.Context) language.Tag {
	local := language.English
	acceptLanguage := context.Request.Header["Accept-Language"]
	if acceptLanguage != nil && len(acceptLanguage) > 0 {
		tag, err := language.Parse(acceptLanguage[0])
		if err == nil {
			local = tag
		}
	}
	return local
}

func GetCurrentUserId(context *gin.Context, secretKey string) (string, error) {
	value, ok := context.Get(UseJWTKey)
	if ok && value.(bool) {
		authorization := context.Request.Header["Authorization"]
		if len(authorization) > 0 {
			return authorization[0], nil
		}
		return "", errors.New("header data error.")
	}
	cookie, err := context.Request.Cookie("X-ADV-TOKEN")
	if err != nil {
		return "", err
	}
	err = webjwt.ParseToken(cookie.Value, secretKey, func(subject string) error {
		return nil
	})
	if err != nil {
		return "", err
	}
	return "", nil
}
