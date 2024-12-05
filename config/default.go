package config

func DefaultConfig() AppConfig {
	return AppConfig{
		Token:  "",
		Prefix: "!",
		Debug:  false,
		ENV:    "development",
	}
}
