package config

import (
	"os"
	"strconv"
)

type GRPCConfig struct {
	Port int
}

type PostgresConfig struct {
	Port     int
	Host     string
	DB       string
	Username string
	Password string
}

type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

type KafkaConfig struct {
	Server string
	Topic  string
}

type Config struct {
	Psql  PostgresConfig
	Grpc  GRPCConfig
	Redis RedisConfig
	Kafka KafkaConfig
}

// New returns a new Config struct
func NewConfig() *Config {
	return &Config{
		Grpc: GRPCConfig{
			Port: getEnvInt("APP_GRPC_PORT", 8001),
		},
		Psql: PostgresConfig{
			Host:     getEnv("APP_POSTGRES_HOST", "127.0.0.1"),
			Port:     getEnvInt("APP_POSTGRES_PORT", 5432),
			DB:       getEnv("APP_POSTGRES_DB", "example"),
			Username: getEnv("APP_POSTGRES_USERNAME", "postgres"),
			Password: getEnv("APP_POSTGRES_PASSWORD", "example"),
		},
		Redis: RedisConfig{
			Address:  getEnv("APP_REDIS_ADDRESS", "localhost:6379"),
			Password: getEnv("APP_REDIS_PASSWORD", ""),
			DB:       getEnvInt("APP_REDIS_DB", 0),
		},
		Kafka: KafkaConfig{
			Server: getEnv("APP_KAFKA_SERVER", "localhost:9092"),
			Topic:  getEnv("APP_KAFKA_TOPIC", "user_creation"),
		},
	}
}

// getEnv returns the value of the environment by name if is exists,
// otherwise it returns default value.
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// getEnvInt is similar getEnv but returns int
func getEnvInt(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		i, err := strconv.Atoi(value)
		if err == nil {
			return i
		}
	}
	return defaultVal
}
