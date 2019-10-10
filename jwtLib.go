package main

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtLib ...
type JwtLib struct {
	MySigningKey string
}

// GetJwtLib ...
func GetJwtLib(mySigningKey string) *JwtLib {
	return &JwtLib{
		MySigningKey: mySigningKey,
	}
}

// Generate ...
func (jwtLib *JwtLib) Generate(claims *Claims) (string, error) {

	if claims == nil {
		return "", errors.New("claims cant be empty")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtLib.MySigningKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

// Validate ...
func (jwtLib *JwtLib) Validate(tokenStr string) (bool, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtLib.MySigningKey), nil
	})

	if err != nil {
		return false, err
	}
	if token.Valid {
		return true, nil
	}
	return false, nil
}

// GetClaims ...
func (jwtLib *JwtLib) GetClaims(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtLib.MySigningKey), nil
	})

	if err != nil {
		return nil, err
	}
	return claims, nil
}
