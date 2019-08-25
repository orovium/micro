package microserver

import (
	"database/sql"
	"sync"

	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	micro "github.com/micro/go-micro"
	"github.com/sirupsen/logrus"

	proto "github.com/orovium/micro/proto"
)

var onceService sync.Once
var onceDB sync.Once
var onceRedis sync.Once
var instance *Service

// Service models the service and upstream needed capabilities
type Service struct {
	options   *Options
	db        *sql.DB
	dbx       *sqlx.DB
	redisPool *redis.Pool
	service   *micro.Service
	log       *logrus.Logger
}

// Init initializes a service if there are no other service initialized
func Init(options *Options) error {
	if instance != nil {
		return ServiceAlreadyInitializeError()
	}

	var err error

	onceService.Do(func() {

		instance, err = newService(options)

	})

	return err
}

// InitDefault returns a service with default options attached.
// If there are another service already initilized, thows a
// ServiceAlreadyInitilizedError erro
func InitDefault() error {
	options := NewOptions().WithDefaultOptions()

	return Init(options)

}

func newService(options *Options) (*Service, error) {
	service := &Service{
		options: options,
	}

	service.init()

	return service, nil
}

func (s *Service) init() error {
	if s.options == nil {
		log.Warn("Warning: Starting new service with no plugins")
		return nil
	}

	var err error

	s.initLogger()
	err = s.initDB()
	if err != nil && !IsNoDatabaseOptionsError(err) {
		GetLogger().WithError(err).Fatal("Can't connect to provided database")
	}
	return nil
}

func (s *Service) initLogger() {
	setLogger(s.options.logger)
	s.log = GetLogger()
}

// GetService returns the service if initialized.
func GetService() (*Service, error) {
	if instance == nil {
		return nil, ServiceNotYetInitializeError()
	}

	return instance, nil
}

// GetDB returns the main sql database connection.
func (s *Service) GetDB() (*sql.DB, error) {
	if s.db == nil {
		return nil, DatabaseNotYetInitializeError()
	}

	return s.db, nil
}

// StartDefaultService returns an initialized service attached to the Ping handler.
func StartDefaultService(options ...micro.Option) micro.Service {
	GetLogger().Infof("Starting service with default handlers/middleware")
	service := startService(options...)

	proto.RegisterPingHandler(service.Server(), new(Ping))

	return service
}

// StartService returns an initialized service with that is not attached to
// any Handler.
func StartService(options ...micro.Option) micro.Service {
	GetLogger().Info("Starting service with no handlers/middleware attached")
	return startService(options...)
}

func startService(options ...micro.Option) micro.Service {
	service := micro.NewService(options...)

	service.Init()
	return service
}
