package microserver

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

const envKey = "ENV"

// Options store all service configuration options
type Options struct {
	db     *DBOptions
	redis  *RedisOptions
	logger *LoggerOptions

	Context context.Context
}

// NewOptions returns an empty options object.
func NewOptions() *Options {
	return &Options{}
}

// DB set the service database configuration
func (o *Options) DB(dbOptions *DBOptions) {
	o.db = dbOptions
}

// Redis set the service redis database configuration
func (o *Options) Redis(redisOptions *RedisOptions) {
	o.redis = redisOptions
}

// Logger set the service logger configuration
func (o *Options) Logger(loggerOptions *LoggerOptions) {
	o.logger = loggerOptions
}

// WithDefaultOptions attach default configuration to the options struct and returns
// a pointer with the default configuration options.
func (o *Options) WithDefaultOptions() *Options {
	o.Logger(DefaultLoggerOptions())
	o.DB(DefaultDBOptions())
	o.Redis(DefaultRedisOptions())
	o.Context = context.Background()

	return o
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
	databaseHostKey     = "DATABASE_HOST"
	databasePasswordKey = "DATABASE_PASSWORD"
	databaseUserKey     = "DATABASE_USER"
	databaseSSLModeKey  = "DATABASE_SSL_MODE"
	mainDatabaseKey     = "SERVICE_DATABASE_NAME"
)

// DefaultDBOptions initilizes and returns a DBOptions struct with default values.
// This values must be environment varialbles
func DefaultDBOptions() *DBOptions {
	return &DBOptions{
		MigrationDir: os.Getenv(migrationDirKey),
		Host:         os.Getenv(databaseHostKey),
		User:         os.Getenv(databaseUserKey),
		SSLMode:      os.Getenv(databaseSSLModeKey),
		MainDatabase: os.Getenv(mainDatabaseKey),
		Password:     os.Getenv(databasePasswordKey),
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

// LoggerOptions stores logrus initialization options
type LoggerOptions struct {
	Env    string
	Level  logrus.Level
	Format logrus.Formatter
}

// DefaultLoggerOptions returns a LoggerOptions filled based on the environment.
func DefaultLoggerOptions() *LoggerOptions {
	env := os.Getenv(envKey)
	level, format := getDefaultLoggerConfByEnv(env)
	return &LoggerOptions{
		Env:    env,
		Level:  level,
		Format: format,
	}
}
