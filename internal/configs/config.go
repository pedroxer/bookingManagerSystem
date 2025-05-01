package configs

type Config struct {
	Postgres        Postgres           `json:"postgres"`
	BookingService  ServiceConnectInfo `json:"booking_service"`
	ResourceService ServiceConnectInfo `json:"resource_service"`
	AuthService     ServiceConnectInfo `json:"auth_service"`
	Api             Api                `json:"api"`
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
	Port         int `json:"port"`
	ReadTimeout  int `json:"read_timeout"`
	WriteTimeout int `json:"write_timeout"`
	IdleTimeout  int `json:"idle_timeout"`
}

type ServiceConnectInfo struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
