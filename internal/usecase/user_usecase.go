package usecase

import (
	"github.com/rifanamrozi/be-rainvow/internal/entity"
	"github.com/rifanamrozi/be-rainvow/internal/repository"
)

type UserUsecase interface {
	ListUsers() ([]*entity.User, error)
	GetUser(id int64) (*entity.User, error)
	CreateUser(name, email string) (*entity.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: repo}
}

func (uc *userUsecase) ListUsers() ([]*entity.User, error) {
	return uc.userRepo.GetAll()
}

func (uc *userUsecase) GetUser(id int64) (*entity.User, error) {
	return uc.userRepo.GetByID(id)
}

func (uc *userUsecase) CreateUser(name, email string) (*entity.User, error) {
	user := &entity.User{
		Name:  name,
		Email: email,
	}
	return uc.userRepo.Create(user)
}
