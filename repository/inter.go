package repository

import "personal_page/model"

type UserRepository interface {
	CreateUser(name string, password string) error
	GetUserByName(name string) (*model.User, error)
	UpdateUserPasswordByName(name string, password string) error
	DeleteUserByName(name string) error
}

type MemoRepository interface {
	ListMemo() ([]*model.Memo, error)
	CreateMemo(content string) error
	GetMemoById(id uint) (*model.Memo, error)
	DeleteMemoById(id uint) error
	CompleteWithId(id uint) error
}
