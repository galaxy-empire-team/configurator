package configurator

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/galaxy-empire-team/configurator/internal/config"
)

type Configurator struct {
	cfg    config.GameConfig
	db     *pgxpool.Pool
	logger *zap.Logger
}

func New(config config.GameConfig, connPool *pgxpool.Pool, logger *zap.Logger) *Configurator {
	return &Configurator{
		cfg:    config,
		db:     connPool,
		logger: logger,
	}
}
