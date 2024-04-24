//go:build wireinject
// +build wireinject

package main

import (
	"io"
	"personal_page/config"
	"personal_page/database"
	"personal_page/delivery"
	"personal_page/logger"
	"personal_page/mq"
	"personal_page/repository"
	"personal_page/usecase"

	"github.com/google/wire"
)

func InitializeServer() *delivery.WebDeli {
	wire.Build(
		config.Get,
		mq.GetRabbitMQ,
		logger.NewLoggerOut,
		wire.Bind(new(io.Writer), new(*logger.LoggerOut)),
		database.GetDb,
		repository.NewUserRepo,
		wire.Bind(new(repository.UserRepository), new(*repository.UserRepo)),
		repository.NewMemoRepo,
		wire.Bind(new(repository.MemoRepository), new(*repository.MemoRepo)),
		usecase.NewUserUc,
		wire.Bind(new(usecase.UserUsecase), new(*usecase.UserUc)),
		usecase.NewMemoUc,
		wire.Bind(new(usecase.MemoUsecase), new(*usecase.MemoUc)),
		delivery.NewWebDeli,
	)
	return &delivery.WebDeli{}
}
