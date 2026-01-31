package configurator

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func (c *Configurator) Run(ctx context.Context) error {
	c.logger.Info("Configurator started")

	data, err := os.ReadFile(c.cfg.PathToConfig)
	if err != nil {
		return fmt.Errorf("os.ReadFile(): %w", err)
	}

	var gameConfig GameConfig
	err = json.Unmarshal(data, &gameConfig)
	if err != nil {
		return fmt.Errorf("json.Unmarshal(): %w", err)
	}

	err = c.upsertConfig(ctx, gameConfig)
	if err != nil {
		return fmt.Errorf("upsertConfig(): %w", err)
	}

	c.logger.Info("Configurator finished")

	return nil
}

func (c *Configurator) upsertConfig(ctx context.Context, cfg GameConfig) (err error) {
	const upsertQuery = `
		INSERT INTO session_beta.buildings (
			building_type,
			level, 
			metal_cost,
			crystal_cost,
			gas_cost, 
			production_s,
			bonuses,
			upgrade_time_s
		) VALUES (
		 	$1,        --- building_type
			$2,        --- level
			$3,        --- metal_cost
			$4,        --- crystal_cost
			$5,        --- gas_cost
			$6,        --- production_s
			$7::jsonb, --- bonuses
			$8         --- upgrade_time_s
		)
		ON CONFLICT (building_type, level) DO UPDATE SET
			metal_cost = EXCLUDED.metal_cost,
			crystal_cost = EXCLUDED.crystal_cost,
			gas_cost = EXCLUDED.gas_cost,
			production_s = EXCLUDED.production_s,
			bonuses = EXCLUDED.bonuses,
			upgrade_time_s = EXCLUDED.upgrade_time_s;	
	`

	batch := &pgx.Batch{}

	for _, building := range cfg.Buildings {
		bonusesJSON, err := json.Marshal(building.Bonuses)
		if err != nil {
			return fmt.Errorf("json.Marshal(): %w", err)
		}

		batch.Queue(
			upsertQuery,
			building.Type,
			building.Level,
			building.UpgradeCost.Metal,
			building.UpgradeCost.Crystal,
			building.UpgradeCost.Gas,
			building.ProductionPerSecond,
			bonusesJSON,
			building.UpgradeTimeSeconds,
		)
	}

	br := c.db.SendBatch(ctx, batch)
	defer br.Close()

	for i := range batch.Len() {
		_, err := br.Exec()
		if err != nil {
			return fmt.Errorf("br.Exec(idx: %d): %w", i, err)
		}
	}

	return nil
}
