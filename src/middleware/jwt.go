package middleware

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/unedtamps/go-backend/config"
	"github.com/unedtamps/go-backend/util"
)

type Credentials struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

func GetCredentials(c *gin.Context) Credentials {
	return c.Value("cred").(Credentials)
}

func CreateJwtToken(payload interface{}) (string, error) {
	var t *jwt.Token
	claims := jwt.MapClaims{
		"iss":     "rplorer",
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
		"payload": payload,
	}
	key := []byte(config.Env.JWTSecret)
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(key)
}

func VerifiyJwtToken(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString != "" {
		tokenString = strings.Split(tokenString, "Bearer ")[1]
	} else {
		// get from params
		tokenString = c.Query("Authorization")
		if tokenString == "" {
			util.UnauthorizedError(c, errors.New("Token Not Provided"))
			c.Abort()
			return
		}
		tokenString = strings.Split(tokenString, "Bearer ")[1]
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Env.JWTSecret), nil
	})
	if err != nil {
		util.UnauthorizedError(c, err)
		c.Abort()
		return
	}
	if !token.Valid {
		util.ForbiddenError(c, err)
		c.Abort()
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	payload := claims["payload"].(map[string]interface{})
	data := Credentials{
		Id:    payload["id"].(string),
		Email: payload["email"].(string),
		Name:  payload["name"].(string),
		Role:  payload["role"].(string),
	}
	c.Set("cred", data)
	c.Next()
}

func IsAdmin(c *gin.Context) {
	cred := c.Value("cred").(Credentials)
	if cred.Role != "ADMIN" {
		util.ForbiddenError(c, errors.New("You are not authorized to perform this action"))
		c.Abort()
		return
	}
	c.Next()
}
