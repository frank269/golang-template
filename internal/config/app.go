package config

type AppConfig struct {
	Host           string `env:"HOST" envDefault:"0.0.0.0"`
	Port           int    `env:"PORT" envDefault:"8080"`
	LogLevel       string `env:"LOG_LEVEL" envDefault:"error"`
	DatabaseConfig *DatabaseConfig
}
