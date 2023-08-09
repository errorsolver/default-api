package manager

import (
	"api/usecase"
	// "api/utils/authenticator"
)

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
}
type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManager.UserRepo())
}

//	func (u *useCaseManager) UserAuthUseCase() usecase.AuthUseCase {
//		return usecase.NewAuthUseCase(authenticator.NewAccessToken(u.))
//	}
func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
