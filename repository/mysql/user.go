package mysql

import (
	"errors"
	"log"
	"strconv"

	"personal_page/model"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type User struct {
	ID        int                   `gorm:"size:10"`
	Name      string                `gorm:"size:10;not null"`
	Password  string                `gorm:"size:60;not null"`
	CreatedAt int64                 `gorm:"autoCreateTime"`
	UpdatedAt int64                 `gorm:"autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name"`
}

func (u *User) toModelUser() *model.User {
	return &model.User{
		ID:       strconv.Itoa(u.ID),
		Name:     u.Name,
		Password: u.Password,
	}
}

type UserRepo struct {
	db *gorm.DB
}

func (u *UserRepo) CreateUser(name string, password string) error {
	return u.db.Create(&User{
		Name:     name,
		Password: password,
	}).Error
}

func (u *UserRepo) GetUserByName(name string) (*model.User, error) {
	user := &User{}
	err := u.db.Take(user, "name = ?", name).Error
	return user.toModelUser(), err
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	admin := &User{
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
