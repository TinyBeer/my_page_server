package repository

import (
	"personal_page/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) CreateUser(name string, password string) error {
	return u.db.Create(&model.User{
		Name:     name,
		Password: password,
	}).Error
}

func (u *UserRepo) GetUserByName(name string) (*model.User, error) {
	user := &model.User{}
	err := u.db.Take(user, "name = ?", name).Error
	return user, err
}

func (u *UserRepo) UpdateUserPasswordByName(name string, password string) error {
	return u.db.Model(&model.User{}).Where("name = ?", name).Update("password", password).Error
}

func (u *UserRepo) DeleteUserByName(name string) error {
	return u.db.Delete(&model.User{}, "name = ?", name).Error
}

var _ UserRepository = &UserRepo{}
