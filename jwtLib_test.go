package main

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	jwtLib := GetJwtLib("test")
	claims := &Claims{
		UserID: "1",
		Admin:  true,
	}
	tokenStr, err := jwtLib.Generate(claims)
	assert.Nil(t, err)
	assert.NotNil(t, tokenStr)
}

func TestGenerateJWTEmptyClaims(t *testing.T) {
	jwtLib := GetJwtLib("test")
	tokenStr, err := jwtLib.Generate(nil)
	assert.Equal(t, errors.New("claims cant be empty"), err)
	assert.Equal(t, "", tokenStr)
}

func TestValidateJWT(t *testing.T) {
	jwtLib := GetJwtLib("test")
	claims := &Claims{
		UserID: "1",
		Admin:  true,
	}
	tokenStr, err := jwtLib.Generate(claims)
	assert.Nil(t, err)
	assert.NotNil(t, tokenStr)
	valid, err := jwtLib.Validate(tokenStr)
	assert.Nil(t, err)
	assert.True(t, valid)
}

func TestValidateJWTInvalidToken(t *testing.T) {
	jwtLib := GetJwtLib("test")
	claims := &Claims{
		UserID: "1",
		Admin:  true,
	}
	tokenStr, err := jwtLib.Generate(claims)
	subStr := string(tokenStr[0:2])
	tokenStr = strings.Replace(tokenStr, subStr, base64.StdEncoding.EncodeToString([]byte("bad")), 1)
	assert.Nil(t, err)
	assert.NotNil(t, tokenStr)
	valid, err := jwtLib.Validate(tokenStr)
	assert.NotNil(t, err)
	assert.False(t, valid)
}

func TestValidateJWTExpiredToken(t *testing.T) {
	jwtLib := GetJwtLib("test")
	claims := &Claims{
		UserID: "1",
		Admin:  true,
	}
	claims.StandardClaims.ExpiresAt = time.Now().Add(-5 * time.Minute).Unix()
	tokenStr, err := jwtLib.Generate(claims)
	assert.Nil(t, err)
	assert.NotNil(t, tokenStr)
	valid, err := jwtLib.Validate(tokenStr)
	assert.NotNil(t, err)
	assert.False(t, valid)
}

func TestGetClaims(t *testing.T) {
	jwtLib := GetJwtLib("test")
	claims := &Claims{
		UserID: "1",
		Admin:  true,
	}
	claims.StandardClaims.ExpiresAt = time.Now().Add(5 * time.Minute).Unix()
	tokenStr, err := jwtLib.Generate(claims)
	assert.Nil(t, err)
	assert.NotNil(t, tokenStr)
	responseClaim, err := jwtLib.GetClaims(tokenStr)
	assert.Nil(t, err)
	assert.Equal(t, claims, responseClaim)
}

func TestGetClaimsExpired(t *testing.T) {
	jwtLib := GetJwtLib("test")
	claims := &Claims{
		UserID: "1",
		Admin:  true,
	}
	claims.StandardClaims.ExpiresAt = time.Now().Add(-5 * time.Minute).Unix()
	tokenStr, err := jwtLib.Generate(claims)
	assert.Nil(t, err)
	assert.NotNil(t, tokenStr)
	responseClaim, err := jwtLib.GetClaims(tokenStr)
	assert.NotNil(t, err)
	assert.Nil(t, responseClaim)
}
