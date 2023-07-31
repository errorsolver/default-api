package middleware

import (
	"fmt"
	"strings"

	"api/utils/authenticator"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}
type AuthTokenMiddleware interface {
	RequireToken() gin.HandlerFunc
}

type authTokenMiddleware struct {
	acctToken authenticator.AccessToken
}

func (a *authTokenMiddleware) RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}
		if err := c.ShouldBindHeader(&h); err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
		fmt.Println(tokenString)
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		token, err := a.acctToken.VerifyAccessToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		fmt.Println("token: ", token)

		if token != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Unauthorized",
			})
			return
		}
	}
}

func NewTokenValidator(acctToken authenticator.AccessToken) AuthTokenMiddleware {
	return &authTokenMiddleware{
		acctToken: acctToken,
	}
}
