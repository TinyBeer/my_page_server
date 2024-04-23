package repository

import "personal_page/model"

type UserRepository interface {
	CreateUser(name string, password string) error
	GetUserByName(name string) (*model.User, error)
	UpdateUserPasswordByName(name string, password string) error
	DeleteUserByName(name string) error
}
