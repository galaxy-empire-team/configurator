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

	err = c.upsertBuildings(ctx, gameConfig)
	if err != nil {
		return fmt.Errorf("upsertBuildings(): %w", err)
	}

	c.logger.Info("Buildings upserted successfully")

	err = c.upsertFleet(ctx, gameConfig)
	if err != nil {
		return fmt.Errorf("upsertFleet(): %w", err)
	}

	c.logger.Info("Fleet upserted successfully")

	err = c.upsertMissions(ctx, gameConfig)
	if err != nil {
		return fmt.Errorf("upsertMissions(): %w", err)
	}

	c.logger.Info("Missions upserted successfully")

	err = c.upsertNotifications(ctx, gameConfig)
	if err != nil {
		return fmt.Errorf("upsertNotifications(): %w", err)
	}

	c.logger.Info("Notifications upserted successfully")

	c.logger.Info("Configurator finished")

	return nil
}

func (c *Configurator) upsertBuildings(ctx context.Context, cfg GameConfig) (err error) {
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

func (c *Configurator) upsertFleet(ctx context.Context, cfg GameConfig) (err error) {
	const upsertFleetQuery = `
		INSERT INTO session_beta.fleet (
			ship_type,
			attack, 
			defense,
			speed,
			cargo_capacity,
			metal_cost,
			crystal_cost,
			gas_cost, 
			build_time_s
		) VALUES (
		 	$1,        --- ship_type
			$2,        --- attack
			$3,        --- defense
			$4,        --- speed
			$5,        --- cargo_capacity
			$6,        --- metal_cost
			$7,        --- crystal_cost
			$8,        --- gas_cost
			$9         --- build_time_s
		)
		ON CONFLICT (ship_type) DO UPDATE SET
			attack = EXCLUDED.attack,
			defense = EXCLUDED.defense,
			speed = EXCLUDED.speed,
			cargo_capacity = EXCLUDED.cargo_capacity,
			metal_cost = EXCLUDED.metal_cost,
			crystal_cost = EXCLUDED.crystal_cost,
			gas_cost = EXCLUDED.gas_cost,
			build_time_s = EXCLUDED.build_time_s;	
	`

	batch := &pgx.Batch{}

	for _, fleetUnit := range cfg.Fleet {
		batch.Queue(upsertFleetQuery,
			fleetUnit.Type,
			fleetUnit.Attack,
			fleetUnit.Defense,
			fleetUnit.Speed,
			fleetUnit.CargoCapacity,
			fleetUnit.BuildCost.Metal,
			fleetUnit.BuildCost.Crystal,
			fleetUnit.BuildCost.Gas,
			fleetUnit.BuildTimeSeconds,
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

func (c *Configurator) upsertMissions(ctx context.Context, cfg GameConfig) (err error) {
	const upsertMissionTypesQuery = `
		INSERT INTO session_beta.missions (
			mission_type
		) VALUES (
		 	$1        --- mission_type
		)
		ON CONFLICT (mission_type) DO NOTHING;
	`

	batch := &pgx.Batch{}

	for _, missionType := range cfg.MissionTypes {
		batch.Queue(upsertMissionTypesQuery, missionType.Type)
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

func (c *Configurator) upsertNotifications(ctx context.Context, cfg GameConfig) (err error) {
	const upsertNotificationTypesQuery = `
		INSERT INTO session_beta.notifications (
			notification_type
		) VALUES (
		 	$1        --- notification_type
		)
		ON CONFLICT (notification_type) DO NOTHING;
	`

	batch := &pgx.Batch{}

	for _, notificationType := range cfg.NotificationTypes {
		batch.Queue(upsertNotificationTypesQuery, notificationType.Type)
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
