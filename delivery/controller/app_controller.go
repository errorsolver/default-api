package controller

import (
	"api/delivery/middleware"
	"api/usecase"
	"api/delivery/controller/user_controller"

	"github.com/gin-gonic/gin"
)

type AppController struct {
	rg          *gin.RouterGroup
	authUseCase usecase.AuthUseCase
	userUseCase usecase.UserUseCase
}

func NewAppController(rGroup *gin.RouterGroup, authUseCase usecase.AuthUseCase, tokenMdw middleware.AuthTokenMiddleware) *AppController {
	controller := AppController{
		rg:          rGroup,
		authUseCase: authUseCase,
	}

	controller.rg.POST("/auth", usercontroller.UserController)
	protectedGroup := controller.rg.Group("/protected", tokenMdw.RequireToken())
	protectedGroup.GET("/user", controller.getUser)
	return &controller
}
