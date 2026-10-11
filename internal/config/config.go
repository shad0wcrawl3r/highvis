// Package config provides the configuration for the application.
package config

type Listen struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Config struct {
	Listen Listen `mapstructure:"listen"`
}

func InitConfig() Config {
	return Config{
		Listen: Listen{
			Host: "0.0.0.0",
			Port: 8080,
		},
	}
}
