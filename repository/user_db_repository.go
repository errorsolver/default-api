package repository

import (
	"api/model"
	"database/sql"
)

type UserRepository interface {
	Create(newUser model.User) error
	RetrieveAll() ([]model.User, error)
	FindById(id string) (model.User, error)
}

type userDbRepository struct {
	db *sql.DB
}

func (c *userDbRepository) Create(newUser model.User) error {
	return nil
}
func (c *userDbRepository) RetrieveAll() ([]model.User, error) {
	return nil, nil
}
func (c *userDbRepository) FindById(id string) (model.User, error) {
	var user model.User
	return user, nil
}

// Like constructor
func NewUserDbRepository(db *sql.DB) UserRepository {
	return &userDbRepository{
		db: db,
	}
}
