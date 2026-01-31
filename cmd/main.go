package main

import (
	"fmt"
	"log"

	"github.com/galaxy-empire-team/configurator/internal/app"
	"github.com/galaxy-empire-team/configurator/internal/config"
	"github.com/galaxy-empire-team/configurator/internal/configurator"
	"github.com/galaxy-empire-team/configurator/internal/db"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("config.New(): %w", err)
	}

	ctx, app, err := app.New(cfg.App)
	if err != nil {
		return fmt.Errorf("app.New(): %w", err)
	}

	db, err := db.New(ctx, cfg.PgConn)
	if err != nil {
		return fmt.Errorf("db.New(): %w", err)
	}
	defer db.Close()

	configurator := configurator.New(cfg.GameConfig, db, app.ComponentLogger("configurator"))
	if err := configurator.Run(ctx); err != nil {
		return fmt.Errorf("configurator.Run(): %w", err)
	}
	return nil
}
