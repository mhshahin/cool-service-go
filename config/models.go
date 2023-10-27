package config

type AppConfig struct {
	Postgres Postgres
	OPA      OPA
}

type Postgres struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	SSLMode  string
}

type OPA struct {
	PoliciesFile string
}
