package microserver

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/eapache/go-resiliency/retrier"
	_ "github.com/lib/pq" // Here we initializes the database
)

// GetDB returns the main sql database connection.
func GetDB() *sql.DB {
	service, err := GetService()
	if err != nil {
		GetLogger().
			WithError(err).
			Warn("Can't retrieve database. Does you call microserver.Init()??")
	}

	db, err := service.GetDB()
	if err != nil {
		GetLogger().
			WithError(err).
			Warn("Can't retrieve database. Does you add a database options??")
	}

	return db
}

func (s *Service) initDB() error {
	if !mustInitilizeDB(s.options) {
		return NoDatabaseOptionsError()
	}

	if s.db != nil {
		return DatabaseAlreadyInitializeError()
	}

	var err error
	onceDB.Do(func() {
		s.db, err = getDBFromOptions(s.options.db)
	})

	return err
}

func mustInitilizeDB(options *Options) bool {
	return options.db != nil
}

func getDBFromOptions(options *DBOptions) (db *sql.DB, err error) {
	if injectedDB := options.GetInjectedDB(); injectedDB != nil {
		db = injectedDB
		return
	}

	db, err = connectToDBServer(options)
	return
}

func connectToDBServer(options *DBOptions) (db *sql.DB, err error) {
	connectionParams := getServerConnectionString(options)
	ret := retrier.New(retrier.ExponentialBackoff(5, 1*time.Second), retrier.DefaultClassifier{})
	ret.Run(func() error {
		db, err = sql.Open("postgres", connectionParams)
		if err != nil {
			return err
		}
		err = db.Ping()
		return err
	})
	if err == nil {
		GetLogger().Infof("database connection stablished")
	}
	return
}

func getServerConnectionString(options *DBOptions) string {
	return fmt.Sprintf(
		"user=%v password=%v sslmode=%v host=%v",
		options.User,
		options.Password,
		options.SSLMode,
		options.Host,
	)
}
