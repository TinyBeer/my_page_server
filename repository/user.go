package repository

import (
	"errors"
	"log"

	"personal_page/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	admin := &model.User{
		Name: "admin",
		// bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.DefaultCost)
		Password: "$2a$10$1iXYspdrld4iQ.W41m6iaOvbOgoyOncycJQ8pWRCdzyOWg8bsEMnq",
	}
	has := db.Migrator().HasTable(admin)
	if !has {
		err := errors.Join(db.AutoMigrate(admin), db.Create(admin).Error)
		if err != nil {
			log.Fatalln(err)
		}
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
