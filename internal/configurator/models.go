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
	Buildings         []Building         `json:"buildings"`
	Fleet             []FleetUnit        `json:"fleet"`
	MissionTypes      []MissionType      `json:"missions"`
	NotificationTypes []NotificationType `json:"notifications"`
	Researches        []Research         `json:"researches"`
}

type Building struct {
	Type                string  `json:"type"`
	Level               uint8   `json:"level"`
	UpgradeCost         Cost    `json:"upgrade_cost"`
	ProductionPerSecond uint64  `json:"production_per_second"`
	Bonuses             Bonuses `json:"bonuses"`
	UpgradeTimeSeconds  uint64  `json:"upgrade_time_seconds"`
}

type Research struct {
	Type                string  `json:"type"`
	Level               uint8   `json:"level"`
	Bonuses             Bonuses `json:"bonuses"`
	ResearchCose        Cost    `json:"research_cost"`
	ResearchTimeSeconds uint64  `json:"research_time_seconds"`
}

type Bonuses struct {
	// Buildings
	FleetBuildSpeed float64 `json:"fleet_build_speed,omitempty"`
	ResearchSpeed   float64 `json:"research_speed,omitempty"`
	BuildSpeed      float64 `json:"building_speed,omitempty"`

	// Researches
	Colonize             uint8 `json:"availiable_colonize_count,omitempty"`

	ResourceGain             float64 `json:"resource_gain,omitempty"`
	FleetCostReduse          float64 `json:"fleet_cost_reduce,omitempty"`
	FleetConstructTimeReduce float64 `json:"fleet_construct_time_reduce,omitempty"`
	DefencePower             float64 `json:"defense_power,omitempty"`

	AttackPower         float64 `json:"attack_power,omitempty"`
	ArmorStrength       float64 `json:"armor_strength,omitempty"`
	AttackDeffencePower float64 `json:"attack_deffence_power,omitempty"`
	SuccessSpyChance    float64 `json:"success_spy_chance,omitempty"`
}

type Cost struct {
	Metal   uint64 `json:"metal"`
	Crystal uint64 `json:"crystal"`
	Gas     uint64 `json:"gas"`
}

type FleetUnit struct {
	Type             string       `json:"type"`
	Speed            int          `json:"speed"`
	Attack           int          `json:"attack"`
	Defense          int          `json:"defense"`
	CargoCapacity    int          `json:"cargo_capacity"`
	BuildCost        ResourceCost `json:"build_cost"`
	BuildTimeSeconds int          `json:"build_time_seconds"`
}

type ResourceCost struct {
	Metal   int `json:"metal"`
	Crystal int `json:"crystal"`
	Gas     int `json:"gas"`
}

type MissionType struct {
	Type string `json:"type"`
}

type NotificationType struct {
	Type string `json:"type"`
}
