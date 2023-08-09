package delivery

import (
	"api/config"
	"api/delivery/controller"
	"api/delivery/manager"
	"api/delivery/middleware"
	"api/usecase"
	"api/utils/authenticator"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine         *gin.Engine
	host           string
	authUseCase    usecase.AuthUseCase
	useCaseManager manager.UseCaseManager
	tokenService   authenticator.AccessToken
}

// Input engine and usecase to controller
func (s *Server) initController() {
	publicRoute := s.engine.Group("/enigma")
	tokenMdw := middleware.NewTokenValidator(s.tokenService)
	controller.NewAppController(publicRoute, s.authUseCase, tokenMdw)
}

// last step to run the server
func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

// Get config and inisiated engine to store it to Server struct
func NewServer() *Server {
	c := config.NewConfig()
	r := gin.Default()

	infra := manager.NewInfraManager(c)
	repo := manager.NewRepositoryManager(infra)
	usecaseManager := manager.NewUseCaseManager(repo)

	tokenService := authenticator.NewAccessToken(c.TokenConfig)
	authUseCase := usecase.NewAuthUseCase(tokenService)
	// fmt.Println(c)
	if c.ApiHost == "" || c.ApiPort == "" {
		panic("Host or Port not defined")
	}
	host := fmt.Sprintf("%s:%s", c.ApiHost, c.ApiPort)
	// host := fmt.Sprintf("%s:%s", "localhost", "8888")
	return &Server{
		engine:         r,
		host:           host,
		authUseCase:    authUseCase,
		useCaseManager: usecaseManager,
		tokenService:   tokenService,
	}
}
