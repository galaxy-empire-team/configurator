package config

type GameConfig struct {
	PathToConfig string `envconfig:"PATH_TO_CONFIG" default:"./config/config.json"`
}
