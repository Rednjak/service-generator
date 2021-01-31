package main

import (
	"github.com/Rednjak/zlat/pkg/app"
)

func main() {
	server := app.Run()
	server.StartServer()
}
