package authentication

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	authBus "pubwebservice/business/authentication"
	cons "pubwebservice/constants"
	authDto "pubwebservice/dtos/authentication"
	"pubwebservice/services/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func getRequest(c *gin.Context) (*authDto.SignInRequest, error) {
	var request authDto.SignInRequest

	if err := c.BindJSON(&request); err != nil {
		return nil, err
	}

	if request.Password == "" {
		return nil, errors.New("Password cannot be empty")
	}

	if request.CompanyId == 0 {
		return nil, errors.New("Invalid companyId")
	}
	return &request, nil
}

func SignInAdmin(c *gin.Context) {
	request, err := getRequest(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isSuccess := authBus.SignInAdmin(request)

	if !isSuccess {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	id, _ := uuid.NewUUID()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"companyId": request.CompanyId,
		"deviceId":  id,
	})
	tokenString, err := token.SignedString(cons.JWT_SECRET_ADMIN)
	if err != nil {
		fmt.Println("error while signing the token, err:", err.Error())
		panic("Something went wrong. Please try again.")
	}
	cookieBuilder := utils.NewCookieBuilder()

	cookieBuilder.SetName("auth").SetValue(tokenString).SetHttpOnly(false).SetSecure(false).SetDomain("localhost")
	cookie, err := cookieBuilder.Build()

	if err != nil {
		log.Println("error while creating cookie, err:", err.Error())
		panic("Something went wrong. Please try again.")
	}
	http.SetCookie(c.Writer, cookie)

}

func SignInUser(c *gin.Context) {
	request, err := getRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isSuccess := authBus.SignInUser(request)
	if !isSuccess {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	id, _ := uuid.NewUUID()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"companyId": request.CompanyId,
		"deviceId":  id,
	})

	tokenString, err := token.SignedString(cons.JWT_SECRET_USER)
	if err != nil {
		fmt.Println("error while signing the token, err:", err.Error())
		panic("Something went wrong. Please try again.")
	}

	cookie, err := utils.NewCookieBuilder().SetName(cons.USER_AUTH_COOKIE_NAME).SetValue(tokenString).SetHttpOnly(true).SetSecure(false).Build()
	if err != nil {
		log.Println("error while creating cookie, err:", err.Error())
		panic("Something went wrong. Please try again.")
	}

	http.SetCookie(c.Writer, cookie)
	c.JSONP(http.StatusOK, nil)
}

func SignOut(c *gin.Context) {
}
