package config

import "time"

type AppConfig struct {
	Postgres Postgres `mapstructure:"postgres"`
	OPA      OPA      `mapstructure:"opa"`
	JWT      JWT      `mapstructure:"jwt"`
}

type Postgres struct {
	Host               string `mapstructure:"host"`
	Port               int    `mapstructure:"port"`
	Username           string `mapstructure:"username"`
	Password           string `mapstructure:"password"`
	Database           string `mapstructure:"database"`
	SSLMode            string `mapstructure:"ssl_mode"`
	MaxOpenConnections int    `mapstructure:"max_open_connections"`
	MaxIdleConnections int    `mapstructure:"max_idle_connections"`
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
