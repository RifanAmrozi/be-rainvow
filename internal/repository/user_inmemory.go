package repository

import "github.com/rifanamrozi/be-rainvow/internal/entity"

type inMemoryUserRepository struct {
	users  []*entity.User
	nextID int64
}

func NewInMemoryUserRepository() UserRepository {
	return &inMemoryUserRepository{
		users:  []*entity.User{},
		nextID: 1,
	}
}

func (r *inMemoryUserRepository) GetAll() ([]*entity.User, error) {
	return r.users, nil
}

func (r *inMemoryUserRepository) GetByID(id int64) (*entity.User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, nil
}

func (r *inMemoryUserRepository) Create(user *entity.User) (*entity.User, error) {
	user.ID = r.nextID
	r.nextID++
	r.users = append(r.users, user)
	return user, nil
}
