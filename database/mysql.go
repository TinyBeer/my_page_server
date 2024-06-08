package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GetDb(conf *viper.Viper) *gorm.DB {
	username := conf.GetString("mysql.username")
	password := conf.GetString("mysql.password")
	host := conf.GetString("mysql.host")
	port := conf.GetInt("mysql.port")
	dbname := conf.GetString("mysql.dbname")
	timeout := conf.GetInt("mysql.timeout")

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
			TablePrefix:         dbname + "-",
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
