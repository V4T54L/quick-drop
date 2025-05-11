package config

type Config struct {
	DBURI     string
	ServerURL string
	PORT      string
}

func GetConfig() *Config {
	return &Config{
		DBURI:     "postgres://user:password@localhost:5432/db?sslmode=disable",
		ServerURL: "",
		PORT:      ":8000",
	}
}
