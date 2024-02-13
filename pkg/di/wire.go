//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/Akshayvij07/thambola-generator/pkg/api"
	"github.com/Akshayvij07/thambola-generator/pkg/api/handler"
	config "github.com/Akshayvij07/thambola-generator/pkg/config"
	db "github.com/Akshayvij07/thambola-generator/pkg/db"
	repository "github.com/Akshayvij07/thambola-generator/pkg/repository"
	usecase "github.com/Akshayvij07/thambola-generator/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewTicketRepository, usecase.NewTicketUseCase, handler.NewUserHandler, http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
