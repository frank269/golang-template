package repositoriesimpl

import (
	"github.com/mpt-tiendc/golang-template/internal/domain/entities"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	*BaseRepositoryImpl[entities.User]
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		BaseRepositoryImpl: NewBaseRepositoryImpl[entities.User](db),
	}
}
