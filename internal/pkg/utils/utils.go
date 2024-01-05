package utils

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

const pubKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvkedJ4DdkdOSmKr42/rIZLUOSDc97w0UKlzWdEApXcilL1eFucEr+R4cZujeWsvOdWxyaA3nSGrJfqC3oiqZ1VK27wMgKRN9Q949q+iasLrmQCKAgZYRbp/TkevJdLqpZhtBdsfhjlbQXXsDgj2Tz+DWE/UEhZtl5BgUuSHvqn1YYuJCJ097lB6UK8i0hZWvejTxteTVc2mdopbXLdzoteCvgOhlhZn93Jbn8kMYlt9lidno18wJ8JN2w/9JXGFUW/xd4i/O4mStCNkWKJVqrl4OFhyHusS82xA7/V/vL7i4X1FUTfWGmX2xqoj3K4Cm29W8VpK8hsLQ+IO5rAmXCQIDAQAB"

func VerifyJWTToken(target string) (*jwt.Token, error) {
	token, err := jwt.Parse(target, func(token *jwt.Token) (interface{}, error) {
		return []byte(pubKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
