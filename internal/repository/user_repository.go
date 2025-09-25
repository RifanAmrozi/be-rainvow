package repository

import (
	"github.com/rifanamrozi/be-rainvow/internal/entity"
)

type UserRepository interface {
	GetAll() ([]*entity.User, error)
	GetByID(id int64) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
}
