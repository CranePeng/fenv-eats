package fenvJWT

import (
	"context"
	"errors"
	"fenv-eats/cmd"
	"fenv-eats/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *CustomJwtClaims, error) {
	Claims := &CustomJwtClaims{}

	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return token, nil, TokenMalformed
		} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
			// Token is expired
			return token, nil, TokenExpired
		} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
			return token, nil, TokenNotValidYet
		} else {
			return token, nil, TokenInvalid
		}
	}
	if Claims, ok := token.Claims.(*CustomJwtClaims); ok && token.Valid {
		return token, Claims, err
	}
	return token, nil, TokenInvalid
}

// 生成token
func CreateToken(user *models.User) (string, error) {
	// 过期时间
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &CustomJwtClaims{
		UserId:   user.Id,
		UserName: user.Email,
		RoleId:   1,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "fenv.top",        // 签名颁发者
			Subject:   "fenv user token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JwtKey))
	cmd.Logger.Debug(context.TODO(), "当前用户：%v，生成token:%v\n", user.Email, tokenString)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return tokenString, nil
}

// 续期
func RefreshToken(refresh string) (bool, error) {
	// 以后上redis需要改变同步存在redis

	return false, nil
}
