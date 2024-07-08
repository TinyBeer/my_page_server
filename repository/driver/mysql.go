package driver

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GetMYSQL(conf *viper.Viper) *gorm.DB {
	username := conf.GetString("database.username")
	password := conf.GetString("database.password")
	host := conf.GetString("database.host")
	port := conf.GetInt("database.port")
	dbname := conf.GetString("database.dbname")
	timeout := conf.GetInt("database.timeout")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%ds",
		username, password, host, port, dbname, timeout)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		// SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable:       true,
			NameReplacer:        nil,
			NoLowerCase:         false,
			IdentifierMaxLength: 0,
		},
	})
	if err != nil {
		panic(err)
	}

	return db
}
