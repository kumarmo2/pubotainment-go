package middlewares

import (
	"log"
	"net/http"
	"pubwebservice/constants"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserAuthMiddleWare(c *gin.Context) {
	cookie, err := c.Cookie(constants.USER_AUTH_COOKIE_NAME)

	if err != nil {
		log.Println("err:", err.Error())
		c.JSON(http.StatusUnauthorized, nil)
		c.Abort()
		return
	}
	token, err := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {

		if t.Method.Alg() != jwt.SigningMethodHS256.Name {
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
	c.Keys["companyId"] = value
}
