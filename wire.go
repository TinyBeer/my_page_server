//go:build wireinject
// +build wireinject

package main

import (
	"personal_page/config"
	"personal_page/delivery"
	"personal_page/repository"
	"personal_page/repository/database/mysql"
	"personal_page/repository/driver"
	"personal_page/usecase"

	"github.com/google/wire"
)

func InitializeServer() *delivery.WebDeli {
	wire.Build(
		config.Get,
		driver.GetES,
		driver.GetMYSQL,
		repository.GetUserRepository,
		repository.GetMemoRepository,
		repository.GetMovieRepository,
		mysql.NewTodoRepo,
		usecase.NewUserUc,
		wire.Bind(new(usecase.UserUsecase), new(*usecase.UserUc)),
		usecase.NewMemoUc,
		wire.Bind(new(usecase.MemoUsecase), new(*usecase.MemoUc)),
		usecase.NewMovieUc,
		wire.Bind(new(usecase.MovieUsecase), new(*usecase.MovieUc)),
		usecase.NewTodoUsecase,
		delivery.NewWebDeli,
	)
	return &delivery.WebDeli{}
}
