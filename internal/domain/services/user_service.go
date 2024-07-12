package services

import (
	"github.com/mpt-tiendc/golang-template/internal/domain/entities"
	"github.com/mpt-tiendc/golang-template/internal/domain/repositories"
)

type UserService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) GetUserById(id string) (*entities.User, error) {
	return s.repository.FindByID(id)
}

func (s *UserService) Create(user *entities.User) error {
	return s.repository.Create(user)
}

func (s *UserService) Delete(id string) error {
	return s.repository.Delete(id)
}

func (s *UserService) Update(user *entities.User) error {
	return s.repository.Update(user)
}

func (s *UserService) FindAll(offset, limit int, sortBy string, sortOrder string) ([]entities.User, error) {
	return s.repository.FindAll(offset, limit, sortBy, sortOrder)
}

func (s *UserService) Filter(queries []repositories.FilterQuery, offset, limit int, sortBy string, sortOrder string) ([]entities.User, error) {
	return s.repository.FilterByQueries(queries, offset, limit, sortBy, sortOrder)
}

func (s *UserService) Count(queries []repositories.FilterQuery) (int64, error) {
	return s.repository.Count(queries)
}
