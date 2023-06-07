package app

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/hexiaopi/blog-service/internal/config"
	"github.com/hexiaopi/blog-service/internal/util"
)

type Claims struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(config.AppEngine.JWT.Secret)
}

func GenerateToken(id int, username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(config.AppEngine.JWT.Expire)
	claims := Claims{
		UserId:   id,
		UserName: username,
		PassWord: util.EncodeMD5(password),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.AppEngine.JWT.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
