package main

import (
	"gateway/config"
	"gateway/internal"
	"gateway/internal/delivery/http"
)

func main() {
	cfg, err := config.LoadAppConfig()
	if err != nil {
		panic(err)
	}

	app := internal.NewApp(cfg)

	err = app.Init()
	if err != nil {
		panic(err)
	}

	server := http.NewHttpServer(cfg)
	server.Init(app)

	err = server.Start()
	if err != nil {
		panic(err)
	}
}
