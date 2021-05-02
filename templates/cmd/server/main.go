package main

import (
	"context"
	"fmt"
	"log"
	"module_placeholder/pkg/ports"
	"module_placeholder/pkg/service"
	"module_placeholder/pkg/shared/config"
	"module_placeholder/pkg/shared/server"
	"strings"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/tkanos/gonfig"
)

func main() {
	cfg := initConfig()
	context := context.Background()
	application := service.NewApplication(context, cfg)

	serverType := strings.ToLower(cfg.Server)
	switch serverType {
	case "http":
		server.RunHTTPServerOnAddr(cfg.ServerPort, func(r chi.Router) chi.Router {
			return ports.InitializeRoutes(
				ports.NewHttpServer(application),
				r,
			)
		})

	default:
		panic(fmt.Sprintf("server type: %s is not supported", serverType))
	}
}

func initConfig() *config.Configuration {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file. error:", err)
	}

	cfg := &config.Configuration{}
	gonfig.GetConf("", cfg)

	return cfg
}
