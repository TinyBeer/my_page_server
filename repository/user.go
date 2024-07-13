package repository

import (
	"personal_page/model"
	mongor "personal_page/repository/database/mongo"
	"personal_page/repository/database/mysql"
	"personal_page/repository/driver"

	"github.com/spf13/viper"
)

type UserRepository interface {
	CreateUser(name, password string) error
	GetUserByName(name string) (*model.User, error)
}

var _ UserRepository = &mysql.UserRepo{}

func GetUserRepository(v *viper.Viper) UserRepository {
	switch v.GetString("database.driver") {
	case "mysql":
		return mysql.NewUserRepo(driver.GetMYSQL(v))
	case "mongo":
		return mongor.NewUserRepo(driver.GetMongoDB(v))
	}

	panic("GetUserRepository: no avaliable database driver")
}
