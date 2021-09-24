package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"playlist-saver/exceptions"
	"time"
)

type jwtCustomClaims struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT string
	ExpiredIn int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
	}
}

//generate new token
func (jwtConf *ConfigJWT) GenerateToken(id int, name string, role string) string {
	claims := &jwtCustomClaims{
		id,
		name,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(int64(jwtConf.ExpiredIn))).Unix(),
		},
	}
	log.Println("JWT CONFIG LOOKa AT DIS", jwtConf)
	//membuat token dari claims yang isinya data" tersebut
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtConf.SecretJWT))
	log.Println("INI TOKEN", t)
	exceptions.PanicIfError(err)

	return t
}