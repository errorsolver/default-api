package manager

import "api/repository"

type RepositoryManager interface {
	UserRepo() repository.UserRepository
}
type repositoryManager struct {
	infra InfraManager
}

func (r *repositoryManager) UserRepo() repository.UserRepository {
	return repository.NewUserDbRepository(r.infra.DbConn())
}
func NewRepositoryManager(manager InfraManager) RepositoryManager {
	return &repositoryManager{infra: manager}
}
