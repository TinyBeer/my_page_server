//go:build wireinject
// +build wireinject

package main

import (
	"personal_page/config"
	"personal_page/delivery"
	"personal_page/repository"
	"personal_page/repository/driver"
	"personal_page/usecase"

	"github.com/google/wire"
)

func InitializeServer() *delivery.WebDeli {
	wire.Build(
		config.Get,
		driver.GetES,
		repository.GetUserRepository,
		repository.GetMemoRepository,
		repository.GetMovieRepository,
		usecase.NewUserUc,
		wire.Bind(new(usecase.UserUsecase), new(*usecase.UserUc)),
		usecase.NewMemoUc,
		wire.Bind(new(usecase.MemoUsecase), new(*usecase.MemoUc)),
		usecase.NewMovieUc,
		wire.Bind(new(usecase.MovieUsecase), new(*usecase.MovieUc)),
		delivery.NewWebDeli,
	)
	return &delivery.WebDeli{}
}
