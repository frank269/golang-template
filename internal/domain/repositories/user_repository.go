package repositories

import "github.com/mpt-tiendc/golang-template/internal/domain/entities"

type UserRepository interface {
	BaseRepository[entities.User]
}
