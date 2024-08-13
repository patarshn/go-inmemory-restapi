package users

import (
	"errors"
)

type IUserRepository interface {
	Create(req UserModel) (UserModel, error)
	Read(id int64) (UserModel, error)
	ReadAll() ([]UserModel, error)
	Update(req UserModel) (UserModel, error)
	Delete(id int64) error
}

type UserRepository struct {
	users  map[int64]UserModel
	nextID int64
}

func NewUserRepository() IUserRepository {
	return &UserRepository{
		users:  make(map[int64]UserModel),
		nextID: 1,
	}
}

func (r *UserRepository) Create(req UserModel) (user UserModel, err error) {
	nextID := r.nextID
	req.ID = &nextID
	r.users[*req.ID] = req
	user = req
	r.nextID += 1

	return user, err
}

func (r *UserRepository) Read(id int64) (user UserModel, err error) {
	user, ok := r.users[id]
	if !ok {
		return user, errors.New("ID Not Found")
	}

	return user, err
}

func (r *UserRepository) ReadAll() (users []UserModel, err error) {
	for _, v := range r.users {
		users = append(users, v)
	}

	return users, err
}
func (r *UserRepository) Update(req UserModel) (user UserModel, err error) {
	if req.ID == nil {
		return user, errors.New("ID Not Found")
	}

	if _, exist := r.users[*req.ID]; !exist {
		return user, errors.New("ID Not found")
	}

	user = req
	r.users[*user.ID] = user

	return user, err
}
func (r *UserRepository) Delete(id int64) (err error) {
	delete(r.users, id)

	return err
}
