package microserver

import (
	"context"
	"os"
)

// Options store all service configuration options
type Options struct {
	db    *DBOptions
	redis *RedisOptions

	Context context.Context
}

// DB set the service database configuration
func (o *Options) DB(dbOptions *DBOptions) {
	o.db = dbOptions
}

// Redis set the service redis database configuration
func (o *Options) Redis(redisOptions *RedisOptions) {
	o.redis = redisOptions
}

// DBOptions stores database configuration.
type DBOptions struct {
	MigrationDir string
	Host         string
	User         string
	SSLMode      string
	MainDatabase string
	Password     string
}

const (
	migrationDirKey     = "./db" //carefull: Will be relative from calling file.
	postgresHostKey     = "POSTGRES_HOST"
	postgresPasswordKey = "POSTGRES_PASSWORD"
	postgresUserKey     = "POSTGRES_USER"
	postgresSSLModeKey  = "POSTGRES_SSL_MODE"
	mainDatabaseKey     = "SERVICE_DATABASE_NAME"
)

// DefaultDBOptions initilizes and returns a DBOptions struct with default values.
// This values must be environment varialbles
func DefaultDBOptions() *DBOptions {
	return &DBOptions{
		MigrationDir: os.Getenv(migrationDirKey),
		Host:         os.Getenv(postgresHostKey),
		User:         os.Getenv(postgresUserKey),
		SSLMode:      os.Getenv(postgresSSLModeKey),
		MainDatabase: os.Getenv(mainDatabaseKey),
		Password:     os.Getenv(postgresPasswordKey),
	}
}

// RedisOptions stores redis database configuration.
type RedisOptions struct {
	Address  string
	Password string
}

const (
	redisAddressKey  = "REDIS_ADDRESS"
	redisPasswordKey = "REDIS_PASSWORD"
)

// DefaultRedisOptions initilizes and returns a DBOptions struct with default values.
// This values must be environment varialbles
func DefaultRedisOptions() *RedisOptions {
	return &RedisOptions{
		Address:  os.Getenv(redisAddressKey),
		Password: os.Getenv(redisPasswordKey),
	}
}
