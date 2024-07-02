package jwt

import (
	gojwt "github.com/golang-jwt/jwt/v5"
)

func CreateJWT(claims gojwt.MapClaims, secret string) (string, error) {
	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string, secret string) (*gojwt.Token, error) {
	token, err := gojwt.Parse(tokenString, func(token *gojwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
