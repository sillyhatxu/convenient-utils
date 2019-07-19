package webjwt

import (
	"encoding/base64"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

const (
	HS256 = "HS256"
	HS384 = "HS384"
	HS512 = "HS512"
)

func createTokenString(secretKey, signingKey string, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod(signingKey), claims)
	tokenstring, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

func CreateTokenStringHS256(secretKey string, claims jwt.Claims) (string, error) {
	return createTokenString(secretKey, HS256, claims)
}

func CreateTokenStringHS384(secretKey string, claims jwt.Claims) (string, error) {
	return createTokenString(secretKey, HS384, claims)
}

func CreateTokenStringHS512(secretKey string, claims jwt.Claims) (string, error) {
	return createTokenString(secretKey, HS512, claims)
}

func ParseToken(tokenSrc, secretKey string, claimFunc func(subject string) error) error {
	standardClaims := jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenSrc, &standardClaims, func(token *jwt.Token) (interface{}, error) {
		dec, err := base64.URLEncoding.DecodeString(secretKey)
		if err != nil {
			return nil, err
		}
		return dec, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("Token check failed.")
	}
	err = claimFunc(standardClaims.Subject)
	if err != nil {
		return err
	}
	return nil
}
