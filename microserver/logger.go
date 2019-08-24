package microserver

import (
	"os"
	"strings"

	stackdriver "github.com/TV4/logrus-stackdriver-formatter"
	"github.com/sirupsen/logrus"
)

const envKey = "ENV"

var log = logrus.New()

func init() {
	env := strings.ToLower(os.Getenv(envKey))

	switch env {
	case "local":
		localLogging()
	case "dev":
		nonProdServerLogging()
	case "pre":
		nonProdServerLogging()
	case "prod":
		nonProdServerLogging()

	default:
		nonProdServerLogging()
		log.Warn("Can't get environment. Setting to non production server")
	}

	log.Infof("Start to loggin at %v with %v level", env, log.GetLevel())
}

// GetLogger returns a logger ready to log to stackdriver
func GetLogger() *logrus.Logger {
	return log
}

func localLogging() {
	log.SetLevel(logrus.TraceLevel)
}

func nonProdServerLogging() {
	setLoggerForServer()
	log.Level = logrus.DebugLevel
}

func prodServerLogging() {
	setLoggerForServer()
	log.SetLevel(logrus.InfoLevel)
}

func setLoggerForServer() {
	log.Formatter = stackdriver.NewFormatter(
		stackdriver.WithService("your-service"),
		stackdriver.WithVersion("v0.1.0"),
	)
}
