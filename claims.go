package main

import jwt "github.com/dgrijalva/jwt-go"

// Claims ...
type Claims struct {
	UserID string `json:"userId"`
	Admin  bool   `json:"admin"`
	jwt.StandardClaims
}
