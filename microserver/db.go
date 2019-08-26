package microserver

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/eapache/go-resiliency/retrier"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Here we initializes the database
)

const pluginDatabaseDriver = "postgres"

// GetDB returns the main sql database connection.
func GetDB() (*sql.DB, error) {
	service, err := GetService()
	if err != nil {
		GetLogger().
			WithError(err).
			Warn("Can't retrieve database. Does you call microserver.Init()??")
		return nil, err
	}

	db, err := service.GetDB()
	if err != nil {
		GetLogger().
			WithError(err).
			Warn("Can't retrieve database. Does you add a database options??")
	}

	return db, err
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
		s.db, s.dbx, err = initializeDBFromOptions(s.options.db)
	})

	return err
}

func mustInitilizeDB(options *Options) bool {
	return options.db != nil
}

func initializeDBFromOptions(options *DBOptions) (db *sql.DB, dbx *sqlx.DB, err error) {
	db, err = getDBFromOptions(options)
	dbx = sqlx.NewDb(db, pluginDatabaseDriver)
	assertDBExists(dbx, options)
	return
}

func getDBFromOptions(options *DBOptions) (db *sql.DB, err error) {
	if injectedDB := options.GetInjectedDB(); injectedDB != nil {
		db = injectedDB
		return
	}

	db, err = connectToDBServer(options)
	return
}

func assertDBExists(dbx *sqlx.DB, options *DBOptions) error {
	if dbExists(dbx, options) {
		return nil
	}
	createDB(dbx, options)

	return nil
}

func dbExists(dbx *sqlx.DB, options *DBOptions) bool {
	query := "SELECT true from pg_database WHERE datname=$1"
	var exists bool
	err := dbx.Get(&exists, query, options.MainDatabase)
	return err == nil && exists
}

func createDB(dbx *sqlx.DB, options *DBOptions) {
	query := fmt.Sprintf("CREATE DATABASE %v;", options.MainDatabase)
	dbx.MustExec(query)
}

func connectToDBServer(options *DBOptions) (db *sql.DB, err error) {
	connectionParams := getServerConnectionString(options)
	ret := retrier.New(retrier.ExponentialBackoff(5, 1*time.Second), retrier.DefaultClassifier{})
	ret.Run(func() error {
		db, err = sql.Open(pluginDatabaseDriver, connectionParams)
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
