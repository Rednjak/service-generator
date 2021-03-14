package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/repo/package/pkg/app"
	"github.com/repo/package/pkg/repository"
	"github.com/repo/package/pkg/service"
	"github.com/repo/package/pkg/shared"
	"github.com/tkanos/gonfig"
)

func main() {
	cfg := initConfig()
	storage := initStorage(cfg)
	service := service.SetupService(storage)
	server := app.NewServer()
	server.Service = service

	server.StartServer()
}

func initStorage(cfg *shared.Configuration) repository.Storage {
	storage := repository.NewStorage()
	defer storage.CloseConnection()

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseIP, cfg.DatabaseScheme)
	err := storage.InitMySQLConn(cfg.DatabaseDriver, connection)
	if err != nil {
		log.Fatal("issue connecting to the database", err)
	}

	return storage
}

func initConfig() *shared.Configuration {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	cfg := &shared.Configuration{}
	gonfig.GetConf("", cfg)

	return cfg
}
