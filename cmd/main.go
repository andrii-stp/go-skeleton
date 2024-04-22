package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/andrii-stp/go-skeleton/config"

	"github.com/andrii-stp/go-libs/service"
)

func loadFromConfig(path string) (*config.Config, error) {
	return nil, nil
}

func main() {
	//create logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// read config
	cfg, err := loadFromConfig("")
	if err != nil {
		logger.Error("failed to load config", slog.String("description", err.Error()))
		os.Exit(1)
	}

	// db migrate
	exit, err := migrateDB(logger, cfg.DB)
	if err != nil {
		logger.Error("failed to migrate db", slog.String("description", err.Error()))
		os.Exit(1)
	}

	if exit {
		logger.Info("migration finished successfully")
		os.Exit(0)
	}

	services := NewService(logger, cfg)

	err := service.Run(services)

}

func migrateDB(logger *slog.Logger, cfg *config.Database) (bool, error) {
	var up, down bool

	flag.BoolVar(&up, "db-migrate-up", false, "migrate db up")
	flag.BoolVar(&down, "db-migrate-down", false, "migrate db down")

	flag.Parse()

	if up {
		logger.Info("starting migration up")
		// use golang-migrate to migrate up

		logger.Info("finished migration up")
		return true, err
	}

	if down {
		logger.Info("starting migration down")
		// use golang-migrate to migrate up

		logger.Info("finished migration down")
		return true, err
	}

	return false, nil
}
