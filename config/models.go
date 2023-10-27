package config

import "time"

type AppConfig struct {
	Postgres Postgres `mapstructure:"postgres"`
	OPA      OPA      `mapstructure:"opa"`
	JWT      JWT      `mapstructure:"jwt"`
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

type OPA struct {
	Enabled bool          `mapstructure:"enabled"`
	URL     string        `mapstructure:"url"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type JWT struct {
	Secret             string        `mapstructure:"secret"`
	ExpirationDuration time.Duration `mapstructure:"expiration_duration"`
}
