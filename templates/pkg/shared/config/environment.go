package config

type Configuration struct {
	DatabaseDriver   string `env:"DATABASE_DRIVER"`
	DatabaseIP       string `env:"DATABASE_IP"`
	DatabasePassword string `env:"DATABASE_PASSWORD"`
	DatabaseScheme   string `env:"DATABASE_SCHEME"`
	DatabaseUser     string `env:"DATABASE_USER"`
	Environment      string `env:"ENVIRONMENT"`
	ServerPort       string `env:"SERVER_PORT"`
	Server           string `env:"SERVER"`
}
