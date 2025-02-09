package configs

type Config struct {
	Postgres Postgres `json:"postgres"`
	Api      Api      `json:"api"`
}

type Postgres struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Sslmode  string `json:"sslmode"`
}
type Api struct {
	Port int `json:"port"`
}
