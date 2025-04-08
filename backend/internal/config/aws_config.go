package config

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// AWSConfig holds AWS-related configuration
type AWSConfig struct {
	Region          string
	SecretName      string
	SecretCacheTime time.Duration
}

// GetAWSConfig returns AWS configuration
func GetAWSConfig() *AWSConfig {
	return &AWSConfig{
		Region:          os.Getenv("AWS_REGION"),
		SecretName:      os.Getenv("AWS_SECRET_NAME"),
		SecretCacheTime: 5 * time.Minute,
	}
}

// GetAWSClientConfig creates a new AWS client configuration
func GetAWSClientConfig() (aws.Config, error) {
	awsConfig := GetAWSConfig()
	return config.LoadDefaultConfig(context.Background(),
		config.WithRegion(awsConfig.Region),
	)
}

// GetSecretsManagerClient creates a new Secrets Manager client
func GetSecretsManagerClient() (*secretsmanager.Client, error) {
	cfg, err := GetAWSClientConfig()
	if err != nil {
		return nil, err
	}
	return secretsmanager.NewFromConfig(cfg), nil
}

// getDBConfigFromAWS retrieves database configuration from AWS Secrets Manager
func getDBConfigFromAWS() (*DBConfig, error) {
	client, err := GetSecretsManagerClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create Secrets Manager client: %v", err)
	}

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(GetAWSConfig().SecretName),
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

	return &dbConfig, nil
}
