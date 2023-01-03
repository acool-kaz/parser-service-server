package main

import (
	"log"

	"github.com/acool-kaz/parser-service-server/internal/app"
	"github.com/acool-kaz/parser-service-server/internal/config"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.InitConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	app, err := app.InitApp(cfg)
	if err != nil {
		log.Fatal(err)
	}

	app.RunApp()
}
