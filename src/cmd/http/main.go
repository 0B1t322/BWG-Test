package main

import (
	"github.com/0B1t322/BWG-Test/internal/app"
	"github.com/0B1t322/BWG-Test/internal/config"
)

// @title BWG API
// @version 1.0
// @BasePath /api
func main() {
	config.InitGlobalConfig()
	app := app.NewApp(config.GlobalConfig)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
