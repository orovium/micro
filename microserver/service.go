package microserver

import (
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
	db        *sqlx.DB
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

	s.initLogger()
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
