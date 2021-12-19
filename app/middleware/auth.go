package middleware

import (
	"backend/controllers"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	Id uint `json:"id"`
	jwt.StandardClaims
}

type JwtCustomClaimsTch struct {
	Idt uint `json:"idt"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT string
}

type ConfigsJWT struct {
	SecretJWTch string
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return controllers.ErrorResponse(c, http.StatusForbidden, e.Error(), e)
		}),
	}
}

func (jwtConfs *ConfigsJWT) Init1() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaimsTch{},
		SigningKey: []byte(jwtConfs.SecretJWTch),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return controllers.ErrorResponse(c, http.StatusForbidden, e.Error(), e)
		}),
	}
}

func (jwtConf *ConfigJWT) GenerateJWT(claims JwtCustomClaims) string {
	jsonData := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := jsonData.SignedString([]byte(jwtConf.SecretJWT))
	// if err != nil {
	// 	return err.Error()
	// }
	return token
}

func (jwtConfs *ConfigsJWT) GenerateJWTtch(claimss JwtCustomClaimsTch) string {
	jsonData1 := jwt.NewWithClaims(jwt.SigningMethodHS256, claimss)
	token1, _ := jsonData1.SignedString([]byte(jwtConfs.SecretJWTch))
	// if err != nil {
	// 	return err.Error()
	// }
	return token1
}

func GetIdFromJWT(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	id := claims.Id
	return int(id)
}

func GetIdFromJWTtch(c echo.Context) int {
	tch := c.Get("user").(*jwt.Token)
	claims := tch.Claims.(*JwtCustomClaimsTch)
	id := claims.Idt
	return int(id)
}
