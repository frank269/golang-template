package repositoriesimpl

import (
	"fmt"

	"github.com/mpt-tiendc/golang-template/internal/domain/repositories"
	"gorm.io/gorm"
)

type BaseRepositoryImpl[T any] struct {
	db *gorm.DB
}

func NewBaseRepositoryImpl[T any](db *gorm.DB) *BaseRepositoryImpl[T] {
	return &BaseRepositoryImpl[T]{db: db}
}

func (r *BaseRepositoryImpl[T]) FindAll(offset, limit int, sortBy, sortOrder string) ([]T, error) {
	var entities []T
	result := r.db.Offset(offset).Limit(limit).Order(sortBy + " " + sortOrder).Find(&entities)
	return entities, result.Error
}

func (r *BaseRepositoryImpl[T]) FilterByQueries(queries []repositories.FilterQuery, offset, limit int, sortBy, sortOrder string) ([]T, error) {
	var entities []T
	query := r.db.Offset(offset).Limit(limit).Order(sortBy + " " + sortOrder)
	for _, q := range queries {
		query = query.Where(fmt.Sprintf(" %s %s ?", q.Field, q.Operator), q.Value)
	}
	result := query.Find(&entities)
	return entities, result.Error
}

func (r *BaseRepositoryImpl[T]) FindByID(id string) (*T, error) {
	var entity T
	result := r.db.First(&entity, `id = ?`, id)
	return &entity, result.Error
}

func (r *BaseRepositoryImpl[T]) Create(entity *T) error {
	result := r.db.Create(entity)
	return result.Error
}

func (r *BaseRepositoryImpl[T]) Update(entity *T) error {
	result := r.db.Save(entity)
	return result.Error
}

func (r *BaseRepositoryImpl[T]) Delete(id string) error {
	result := r.db.Delete(new(T), `id = ?`, id)
	return result.Error
}

func (r *BaseRepositoryImpl[T]) Count(queries []repositories.FilterQuery) (int64, error) {
	query := r.db.Model(new(T))
	for _, q := range queries {
		query = query.Where(fmt.Sprintf(" %s %s ?", q.Field, q.Operator), q.Value)
	}
	var count int64
	result := query.Count(&count)
	return count, result.Error
}
