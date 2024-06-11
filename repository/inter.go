package repository

import (
	"personal_page/model"
	mongor "personal_page/repository/mongo"
	"personal_page/repository/mysql"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(name, password string) error
	GetUserByName(name string) (*model.User, error)
}

type MemoRepository interface {
	ListMemo() ([]*model.Memo, error)
	CreateMemo(content string) error
	DeleteMemoById(id string) error
	CompleteWithId(id string) error
}

func GetUserRepository(sql *gorm.DB, mongo *mongo.Database) UserRepository {
	if sql != nil {
		return mysql.NewUserRepo(sql)
	}

	if mongo != nil {
		return mongor.NewUserRepo(mongo)
	}

	panic("GetUserRepository: no avaliable database")
}

func GetMemoRepository(sql *gorm.DB, mongo *mongo.Database) MemoRepository {
	if sql != nil {
		return mysql.NewMemoRepo(sql)
	}

	if mongo != nil {
		return mongor.NewMemoRepo(mongo)
	}
	panic("GetMemoRepository: no avaliable database")
}

var (
	_ UserRepository = &mysql.UserRepo{}
	_ MemoRepository = &mysql.MemoRepo{}
)

var (
	_ UserRepository = &mongor.UserRepo{}
	_ MemoRepository = &mongor.MemoRepo{}
)
