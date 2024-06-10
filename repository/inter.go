package repository

import (
	"personal_page/model"
	"personal_page/repository/mysql"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(name string, password string) error
	GetUserByName(name string) (*model.User, error)
}

type MemoRepository interface {
	ListMemo() ([]*model.Memo, error)
	CreateMemo(content string) error
	DeleteMemoById(id string) error
	CompleteWithId(id string) error
}

func GetUserRepository(v *viper.Viper, sql *gorm.DB, mongo *mongo.Database) UserRepository {
	return mysql.NewUserRepo(sql)
}

func GetMemoRepository(v *viper.Viper, sql *gorm.DB, mongo *mongo.Database) MemoRepository {
	return mysql.NewMemoRepo(sql)
}

var (
	_ UserRepository = &mysql.UserRepo{}
	_ MemoRepository = &mysql.MemoRepo{}
)
