package repository

import (
	"api/model"
	"database/sql"
)

type UserRepository interface {
	Create(newUser model.UserCredential) error
	RetrieveAll() ([]model.UserCredential, error)
	FindById(id string) (model.UserCredential, error)
}

type userDbRepository struct {
	db *sql.DB
}

func (c *userDbRepository) Create(newUser model.UserCredential) error {
	return nil
}
func (c *userDbRepository) RetrieveAll() ([]model.UserCredential, error) {
	return nil, nil
}
func (c *userDbRepository) FindById(id string) (model.UserCredential, error) {
	var user model.UserCredential
	return user, nil
}

// Like constructor
func NewUserDbRepository(db *sql.DB) UserRepository {
	return &userDbRepository{
		db: db,
	}
}
