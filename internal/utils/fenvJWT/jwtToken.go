package fenvJWT

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

// 载荷，可添加自己需要的一些信息
type CustomJwtClaims struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	RoleId   int64  `json:"roleId"`
	jwt.StandardClaims
}

const (
	// 这个是需要保密的一段信息
	JwtKey                 = "www.fenv.top"
	Gin_Context_Key string = "claims"
)

var accessSecret = []byte("qyrzr")
var refreshSecret = []byte("ar")

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token is not active yet")
	TokenMalformed   error = errors.New("Malformed token")
	TokenInvalid     error = errors.New("can't handle this token")
)
