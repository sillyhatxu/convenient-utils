package webjwt

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Id           int64  `json:"id"`
	UserName     string `json:"user_name"`
	MobileNumber string `json:"mobile_number"`
	UserType     int    `json:"user_type"`
	IsDelete     bool   `json:"is_delete"`
	jwt.StandardClaims
}

const secretKey = "COOKIE_SECRET_KEY"

func TestCreateTokenStringHS256(t *testing.T) {
	userToken := *&User{Id: 1001, UserName: "Cookie", MobileNumber: "99999999", UserType: 1, IsDelete: false}
	token, err := CreateTokenStringHS256(secretKey, userToken)
	assert.Nil(t, err)
	assert.EqualValues(t, token, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAwMSwidXNlcl9uYW1lIjoiQ29va2llIiwibW9iaWxlX251bWJlciI6Ijk5OTk5OTk5IiwidXNlcl90eXBlIjoxLCJpc19kZWxldGUiOmZhbHNlfQ.bw2NQKDA5zuNFXLtUB0kdatB1AfvKv1Qn4ap5U29-K8")
}

func TestCreateTokenStringHS384(t *testing.T) {
	userToken := *&User{Id: 1001, UserName: "Cookie", MobileNumber: "99999999", UserType: 1, IsDelete: false}
	token, err := CreateTokenStringHS384(secretKey, userToken)
	assert.Nil(t, err)
	assert.EqualValues(t, token, "eyJhbGciOiJIUzM4NCIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAwMSwidXNlcl9uYW1lIjoiQ29va2llIiwibW9iaWxlX251bWJlciI6Ijk5OTk5OTk5IiwidXNlcl90eXBlIjoxLCJpc19kZWxldGUiOmZhbHNlfQ.V-wFyd2gY56KepHwVeiog2CyBJT3xclSgmhboAhxo9jVoTZ8iu_AVn3LXC5gf4Qf")
}

func TestCreateTokenStringHS512(t *testing.T) {
	userToken := *&User{Id: 1001, UserName: "Cookie", MobileNumber: "99999999", UserType: 1, IsDelete: false}
	token, err := CreateTokenStringHS512(secretKey, userToken)
	assert.Nil(t, err)
	assert.EqualValues(t, token, "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAwMSwidXNlcl9uYW1lIjoiQ29va2llIiwibW9iaWxlX251bWJlciI6Ijk5OTk5OTk5IiwidXNlcl90eXBlIjoxLCJpc19kZWxldGUiOmZhbHNlfQ.T_QJq0W2cu-kCJ92UCm1RmwSF2x5J6XEAUKqEQbN7tugCsTr7al5x5xLT_mJhYhOYQyaG2d12WCh0C_CjMg8lw")
}

func TestParseToken(t *testing.T) {
	tokenSrc := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAwMSwidXNlcl9uYW1lIjoiQ29va2llIiwibW9iaWxlX251bWJlciI6Ijk5OTk5OTk5IiwidXNlcl90eXBlIjoxLCJpc19kZWxldGUiOmZhbHNlfQ.T_QJq0W2cu-kCJ92UCm1RmwSF2x5J6XEAUKqEQbN7tugCsTr7al5x5xLT_mJhYhOYQyaG2d12WCh0C_CjMg8lw"
	var user User
	err := ParseToken(tokenSrc, secretKey, func(subject string) error {
		if err := json.Unmarshal([]byte(subject), &user); err != nil {
			return err
		}
		return nil
	})
	assert.Nil(t, err)
}
