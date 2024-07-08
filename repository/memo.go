package repository

import (
	"personal_page/model"
	mongor "personal_page/repository/database/mongo"
	"personal_page/repository/database/mysql"
	"personal_page/repository/driver"

	"github.com/spf13/viper"
)

var (
	_ UserRepository = &mongor.UserRepo{}
	_ MemoRepository = &mongor.MemoRepo{}
)

type MemoRepository interface {
	ListMemo() ([]*model.Memo, error)
	CreateMemo(content string) error
	DeleteMemoById(id string) error
	CompleteWithId(id string) error
}

func GetMemoRepository(v *viper.Viper) MemoRepository {
	switch v.GetString("database.driver") {
	case "mysql":
		return mysql.NewMemoRepo(driver.GetMYSQL(v))
	case "mongo":
		return mongor.NewMemoRepo(driver.GetMongoDB(v))
	}

	panic("GetMemoRepository: no avaliable database")
}
