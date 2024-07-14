//go:build wireinject
// +build wireinject

package main

import (
	"personal_page/config"
	"personal_page/delivery"
	"personal_page/repository/database/mysql"
	"personal_page/repository/driver"
	"personal_page/usecase"

	"github.com/google/wire"
)

func InitializeServer() *delivery.WebDeli {
	wire.Build(
		config.Get,
		driver.GetMYSQL,
		mysql.NewUserRepository,
		usecase.NewTokenUsecase,
		mysql.NewTodoRepository,
		usecase.NewTodoUsecase,
		delivery.NewWebDeli,
	)
	return &delivery.WebDeli{}
}
