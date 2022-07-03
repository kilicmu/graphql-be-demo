package database

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	username string
	password string
	protocol string
	host     string
	port     int
	db       string
}

var cachedConfig *DatabaseConfig

func getConfigByEnv() *DatabaseConfig {
	env := os.Getenv("API_ENV")
	if cachedConfig != nil {
		return cachedConfig
	}
	switch env {
	case "test":
		cachedConfig = &DatabaseConfig{
			username: "root",
			password: "123456",
			protocol: "tcp",
			host:     "127.0.0.1",
			port:     4006,
			db:       "geekbang_cheap_test",
		}
		return cachedConfig
	// production api config
	default:
		cachedConfig = &DatabaseConfig{
			username: "root",
			password: "123456",
			protocol: "tcp",
			host:     "127.0.0.1",
			port:     4006,
			db:       "geekbang_cheap_test",
		}
		return cachedConfig
	}
}

func formatConfigToDSN(config *DatabaseConfig) string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.username, config.password, config.protocol, config.host, config.port, config.db)
}
