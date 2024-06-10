//go:build wireinject
// +build wireinject

package main

import (
	"personal_page/config"
	"personal_page/database"
	"personal_page/delivery"
	"personal_page/repository"
	"personal_page/usecase"

	"github.com/google/wire"
)

func InitializeServer() *delivery.WebDeli {
	wire.Build(
		config.Get,
		database.GetDb,
		database.GetMongoDB,
		repository.GetUserRepository,
		repository.GetMemoRepository,
		usecase.NewUserUc,
		wire.Bind(new(usecase.UserUsecase), new(*usecase.UserUc)),
		usecase.NewMemoUc,
		wire.Bind(new(usecase.MemoUsecase), new(*usecase.MemoUc)),
		delivery.NewWebDeli,
	)
	return &delivery.WebDeli{}
}
