package database

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"workout_app_backend/services/workoutAppServices/internal/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
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
	DatabaseType    string        `json:"database_type"`
}

// AWSConfig holds AWS-related configuration
type AWSConfig struct {
	Region          string
	SecretName      string
	SecretCacheTime time.Duration
}

// GetSecretsManagerClient creates a new Secrets Manager client
func GetSecretsManagerClient() (*secretsmanager.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		return nil, err
	}
	return secretsmanager.NewFromConfig(cfg), nil
}

func GetLocalDBConfig() (*DBConfig, error) {

	var dbConfig DBConfig
	dbConfig.Username = os.Getenv("DB_USER_LOCAL")
	dbConfig.Host = "localhost"
	dbConfig.Port = os.Getenv("DB_PORT")
	dbConfig.DBName = os.Getenv("DB_NAME")
	dbConfig.MaxOpenConns = 25
	dbConfig.MaxIdleConns = 5
	dbConfig.ConnMaxLifetime = 5 * time.Minute
	dbConfig.DatabaseType = "local-postgres"
	return &dbConfig, nil
}

// getDBConfigFromAWS retrieves database configuration from AWS Secrets Manager
func getDBConfigFromAWS() (*DBConfig, error) {
	client, err := GetSecretsManagerClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create Secrets Manager client: %v", err)
	}

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(os.Getenv("AWS_SECRET_NAME")),
	}

	result, err := client.GetSecretValue(context.Background(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to get secret value: %v", err)
	}

	var dbConfig DBConfig
	if err := json.Unmarshal([]byte(*result.SecretString), &dbConfig); err != nil {
		return nil, fmt.Errorf("failed to unmarshal secret data: %v", err)
	}

	dbConfig.Host = os.Getenv("DB_HOST")
	dbConfig.Port = os.Getenv("DB_PORT")
	dbConfig.DBName = "postgres"
	dbConfig.MaxOpenConns = 25
	dbConfig.MaxIdleConns = 5
	dbConfig.ConnMaxLifetime = 5 * time.Minute
	dbConfig.DatabaseType = "aws-postgres"
	return &dbConfig, nil
}

// GetDBConfig returns database configuration
func GetDBConfig() (*DBConfig, error) {
	if utils.IsTestEnv() {
		return GetLocalDBConfig()
	}
	return getDBConfigFromAWS()
}

// GetConnectionString returns a PostgreSQL connection string
func (c *DBConfig) GetConnectionString() string {

	if utils.IsTestEnv() {
		return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", c.Host, c.Port, c.Username, c.DBName)
	}

	return c.GetAWSDBConfigConnectionString()
}

func (c *DBConfig) GetAWSDBConfigConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.Username, c.Password, c.DBName)
}
