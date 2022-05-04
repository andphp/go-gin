package jwtService

import "github.com/dgrijalva/jwt-go"

// Custom Session structure
type CustomSession struct {
	Session
	BufferTime int64
	jwt.StandardClaims
}

type Session struct {
	ID          int64
	AccountName string
	NickName    string
	RoleID      int
}

