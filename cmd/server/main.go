package main

import (
	"github.com/radahn42/onetime-note/internal/app"
	"github.com/radahn42/onetime-note/internal/config"
)

func main() {
	cfg := config.MustLoad()
	a := app.New(cfg)

	a.Run()
}
