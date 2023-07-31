package controller

import (
	"api/delivery/middleware"
	"api/model"
	"api/usecase"

	"github.com/gin-gonic/gin"
)

type AppController struct {
	rg          *gin.RouterGroup
	authUseCase usecase.AuthUseCase
}

func (cc *AppController) userAuth(c *gin.Context) {
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

func (cc *AppController) getCustomer(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user",
	})
}

func NewAppController(rGroup *gin.RouterGroup, authUseCase usecase.AuthUseCase, tokenMdw middleware.AuthTokenMiddleware) *AppController {
	controller := AppController{
		rg:          rGroup,
		authUseCase: authUseCase,
	}

	controller.rg.POST("/auth", controller.userAuth)
	protectedGroup := controller.rg.Group("/protected", tokenMdw.RequireToken())
	protectedGroup.GET("/user", controller.getCustomer)
	return &controller
}
