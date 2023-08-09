package usecase

import (
	"api/model"
	"api/repository"
)

type UserUseCase interface {
	RegisterUser(user model.User) error
	FindUserById(id string) (model.User, error)
	GetAllUser() ([]model.User, error)
}
type userUsecase struct {
	repo repository.UserRepository
}

func (c *userUsecase) RegisterUser(user model.User) error {
	return c.repo.Create(user)
}
func (c *userUsecase) FindUserById(id string) (model.User, error) {
	return c.repo.FindById(id)
}
func (c *userUsecase) GetAllUser() ([]model.User, error) {
	return c.repo.RetrieveAll()
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUsecase{repo: repo}
}
