package config

import (
	"fmt"
	"time"
)

// DBConfig holds database configuration
type DBConfig struct {
	Host            string        `json:"host"`
	Port            string        `json:"port"`
	Username        string        `json:"username"`
	Password        string        `json:"password"`
	DBName          string        `json:"dbname"`
	MaxOpenConns    int           `json:"max_open_conns"`
	MaxIdleConns    int           `json:"max_idle_conns"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime"`
}

// GetDBConfig returns database configuration
func GetDBConfig() (*DBConfig, error) {
	// Try to get credentials from AWS Secrets Manager first
	config, err := getDBConfigFromAWS()
	if err == nil {
		return config, nil
	}
	//!There is no fallback, if aws secrets manager is not working, the program will not run
	return nil, fmt.Errorf("failed to get database configuration")
}

// GetConnectionString returns a PostgreSQL connection string
func (c *DBConfig) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.Username, c.Password, c.DBName)
}
