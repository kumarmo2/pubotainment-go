package middlewares

import (
	"errors"
	"log"
	"net/http"
	"pubwebservice/constants"

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

	value, ok := claims["companyId"]
	if !ok {
		log.Println("no companyId found")
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}
	c.Keys = make(map[string]any)
	c.Keys["companyId"] = value

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
	value, ok := claims["companyId"]
	if !ok {
		log.Println("no companyId found")
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}
	c.Keys = make(map[string]any)
	c.Keys["companyId"] = value
}
