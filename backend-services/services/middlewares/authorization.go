package middlewares

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"pubwebservice/constants"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AdminAuthMiddleWare(c *gin.Context) {
	cookie, err := c.Cookie(constants.ADMIN_AUTH_COOKIE_NAME)
	if err != nil {
		log.Println("err: ", err.Error())
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}

	token, err := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {

		if t.Method.Alg() != jwt.SigningMethodHS256.Name {
			return nil, errors.New("token signing algorithm didn't match")
		}
		return constants.JWT_SECRET_ADMIN, nil
	})

	if err != nil {
		log.Println("err:", err.Error())
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		log.Println("err:", err.Error())
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}
	setCommonContextKeys(c, claims)
}

func setCommonContextKeys(c *gin.Context, claims jwt.MapClaims) {
	c.Keys = make(map[string]any)

	// set companyId
	value, ok := claims["companyId"]
	if !ok {
		log.Println("no companyId found")
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}
	companyId, _ := strconv.ParseInt(fmt.Sprintf("%v", value), 10, 64)
	c.Keys["companyId"] = companyId

	// set deviceId
	value, ok = claims["deviceId"]
	if !ok {
		log.Println("no deviceId found")
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}
	deviceId := fmt.Sprintf("%v", value)
	c.Keys["deviceId"] = deviceId
}

func UserAuthMiddleWare(c *gin.Context) {
	log.Printf("cookies: %+v\n", c.Request.Cookies())
	cookie, err := c.Cookie(constants.USER_AUTH_COOKIE_NAME)

	if err != nil {
		log.Println("err:", err.Error())
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}
	token, err := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {

		if t.Method.Alg() != jwt.SigningMethodHS256.Name {
			return nil, errors.New("token signing algorithm didn't match")
		}
		return constants.JWT_SECRET_USER, nil

	})
	if err != nil {
		log.Println("err:", err.Error())
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}
	setCommonContextKeys(c, claims)
}
