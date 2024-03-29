package server

import (
	"context"
	"database/sql"
	"os"

	"github.com/sirupsen/logrus"
)

const envKey = "ENV"

// Options store all service configuration options
type Options struct {
	db       *DBOptions
	redis    *RedisOptions
	logger   *LoggerOptions
	firebase *FirebaseOptions

	Context context.Context
}

// NewOptions returns an empty options object.
func NewOptions() *Options {
	return &Options{}
}

// DB set the service database configuration
func (o *Options) DB(dbOptions *DBOptions) {
	GetLogger().Trace("DB options on DB() :")
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

	db *sql.DB
}

const (
	migrationDirKey     = "DATABASE_MIGRATIONS_DIR" //carefull: Will be relative from calling file.
	databaseHostKey     = "DATABASE_HOST"
	databasePasswordKey = "DATABASE_PASSWORD"
	databaseUserKey     = "DATABASE_USER"
	databaseSSLModeKey  = "DATABASE_SSL_MODE"
	mainDatabaseKey     = "SERVICE_DATABASE_NAME"
)

// DefaultDBOptions initilizes and returns a DBOptions struct with default values.
// This values must be environment varialbles
func DefaultDBOptions() *DBOptions {

	if !databaseEnvIsSetting() {
		return nil
	}

	return &DBOptions{
		MigrationDir: os.Getenv(migrationDirKey),
		Host:         os.Getenv(databaseHostKey),
		User:         os.Getenv(databaseUserKey),
		SSLMode:      os.Getenv(databaseSSLModeKey),
		MainDatabase: os.Getenv(mainDatabaseKey),
		Password:     os.Getenv(databasePasswordKey),
	}
}

func databaseEnvIsSetting() bool {
	return envExist(migrationDirKey) &&
		envExist(databaseHostKey) &&
		envExist(databaseUserKey) &&
		envExist(databaseSSLModeKey) &&
		envExist(mainDatabaseKey) &&
		envExist(databasePasswordKey)
}

// WithInjectedDB can be used to database dependency injection
func (dbo *DBOptions) WithInjectedDB(db *sql.DB) *DBOptions {
	dbo.db = db
	return dbo
}

// GetInjectedDB returns the injected database
func (dbo *DBOptions) GetInjectedDB() *sql.DB {
	return dbo.db
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

const (
	firebaseBucketKey         = "FIREBASE_BUCKET"
	firebaseBucketFileNameKey = "FIREBASE_BUCKET_FILE_NAME"
	firebaseConfigPathKey     = "FIREBASE_CONFIG_PATH"
)

// FirebaseOptions stores firebase configurations to authorization/authentication
type FirebaseOptions struct {
	bucket     string
	name       string
	configPath string
}

// NewFirebaseOptions returns an emty FirebaseOptionsObject
func NewFirebaseOptions() *FirebaseOptions {
	return &FirebaseOptions{}
}

// FromFile fills the FirebaseObject with the path to the firebase.json config file.
// You can rename the file to other.json or whatever you want, only be sure that
// the path include the file name.
func (fbo *FirebaseOptions) FromFile(path string) *FirebaseOptions {
	fbo.configPath = path
	return fbo
}

// FromBucket fills the FirebaseObject with the gclout storage bucket that stores
// the firebase.json. If you save the file with other name, you can provide it
// in the name parameter. In any other case, left it to blank string: ""
func (fbo *FirebaseOptions) FromBucket(bucket string, name string) *FirebaseOptions {
	fbo.bucket = bucket
	fbo.name = name
	return fbo
}

// DefaultFirebaseOptions returns a FirebaseOptions fills with the bucket/path
// found on environment variables.
func DefaultFirebaseOptions() *FirebaseOptions {
	switch true {
	case envExist(firebaseBucketKey):
		bucket := os.Getenv(firebaseBucketKey)
		name := os.Getenv(firebaseBucketFileNameKey)
		return NewFirebaseOptions().FromBucket(bucket, name)
	case envExist(firebaseConfigPathKey):
		path := os.Getenv(firebaseConfigPathKey)
		return NewFirebaseOptions().FromFile(path)

	default:
		return NewFirebaseOptions()
	}
}
