//go:build wireinject

package main

import (
	"context"
	"github.com/aragorn-yang/go-camp-03/internal/repository"
	"github.com/aragorn-yang/go-camp-03/internal/repository/ent"
	"github.com/aragorn-yang/go-camp-03/internal/service"
	"github.com/aragorn-yang/go-camp-03/internal/usecase"
	"github.com/google/wire"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

func InitUserService() (*service.UserService, error) {
	wire.Build(UserSet, NewDB, InitConfig)
	return new(service.UserService), nil
}

var UserSet = wire.NewSet(
	service.NewUserService,
	repository.NewRepository,
	usecase.NewUser,
)

func NewDB(v *viper.Viper) (*ent.Client, error) {
	client, err := ent.Open(
		v.Sub("db").GetString("type"),
		v.Sub("db").GetString("dsn"),
	)
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}

func InitConfig() (*viper.Viper, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	return viper.GetViper(), nil
}
