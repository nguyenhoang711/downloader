// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wiring

import (
	"github.com/google/wire"
	"github.com/nguyenhoang711/downloader/internal/app"
	"github.com/nguyenhoang711/downloader/internal/configs"
	"github.com/nguyenhoang711/downloader/internal/dataaccess"
	"github.com/nguyenhoang711/downloader/internal/dataaccess/cache"
	"github.com/nguyenhoang711/downloader/internal/dataaccess/database"
	"github.com/nguyenhoang711/downloader/internal/handler"
	"github.com/nguyenhoang711/downloader/internal/handler/grpc"
	"github.com/nguyenhoang711/downloader/internal/logic"
	"github.com/nguyenhoang711/downloader/internal/utils"
)

// Injectors from wire.go:

func InitializeStandaloneServer(configFilePath configs.ConfigFilePath) (grpc.Server, func(), error) {
	config, err := configs.NewConfig(configFilePath)
	if err != nil {
		return nil, nil, err
	}
	configsDatabase := config.Database
	db, cleanup, err := database.InitializeDB(configsDatabase)
	if err != nil {
		return nil, nil, err
	}
	goquDatabase := database.InitializeGoquDB(db)
	log := config.Log
	logger, cleanup2, err := utils.InitializeLogger(log)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	accountDataAccessor := database.NewAccountDataAccessor(goquDatabase, logger)
	accountPasswordDataAccessor := database.NewAccountPasswordDataAccesor(goquDatabase, logger)
	auth := config.Auth
	hash := logic.NewHash(auth)
	tokenPublicKeyDataAccessor := database.NewTokenPublicKeyDataAccessor(goquDatabase, logger)
	configsCache := config.Cache
	client, err := cache.NewClient(configsCache, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tokenPublicKey := cache.NewTokenPublicKey(client, logger)
	token, err := logic.NewToken(accountDataAccessor, tokenPublicKeyDataAccessor, auth, logger, tokenPublicKey)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	takenAccountName := cache.NewTakenAccountName(client, logger)
	account := logic.NewAccount(goquDatabase, accountDataAccessor, accountPasswordDataAccessor, hash, token, takenAccountName, logger)
	goLoadServiceServer := grpc.NewHandler(account)
	server := grpc.NewServer(goLoadServiceServer)
	return server, func() {
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

var wireSet = wire.NewSet(configs.WireSet, dataaccess.WireSet, handler.WireSet, logic.WireSet, utils.WireSet, app.WireSet)
