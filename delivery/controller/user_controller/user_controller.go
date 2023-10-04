package usercontroller

import (
	"api/model"
	"api/usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	authUseCase usecase.AuthUseCase
}

func (cc *UserController) userAuth(c *gin.Context) {
	var user model.UserCredential
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "can't bind struct",
			"error":   err,
		})
		return
	}

	token, err := cc.authUseCase.UserAuth(user)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Auth fail",
		})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}

func (cc *UserController) getUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user",
	})
}
