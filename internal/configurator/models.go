package configurator

const (
	metalMineType    = "metal_mine"
	crystalMineType  = "crystal_mine"
	gasMineType      = "gas_mine"
	spacePortType    = "space_port"
	laboratoryType   = "laboratory"
	robotFactoryType = "robot_factory"
)

type GameConfig struct {
	Buildings []Building `json:"buildings"`
}

type Building struct {
	Type                string      `json:"type"`
	Level               uint8       `json:"level"`
	UpgradeCost         UpgradeCost `json:"upgrade_cost"`
	ProductionPerSecond uint64      `json:"production_per_second"`
	Bonuses             Bonuses     `json:"bonuses"`
	UpgradeTimeSeconds  uint64      `json:"upgrade_time_seconds"`
}

type Bonuses struct {
	FleetBuildSpeed float64 `json:"fleet_build_speed,omitempty"`
	ResearchSpeed   float64 `json:"research_speed,omitempty"`
	BuildSpeed      float64 `json:"building_speed,omitempty"`
}

type UpgradeCost struct {
	Metal   uint64 `json:"metal"`
	Crystal uint64 `json:"crystal"`
	Gas     uint64 `json:"gas"`
}
