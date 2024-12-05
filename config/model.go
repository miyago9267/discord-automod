package config

type AppConfig struct {
	Token  string `mapstructure:"token" env:"TOKEN" default:""`
	Prefix string `mapstructure:"prefix" env:"PREFIX" default:"!"`
	ENV    string `mapstructure:"env" env:"ENV" default:"development"`
	Debug  bool   `mapstructure:"debug" env:"DEBUG" default:"false"`
}
