package log

import (
	"os"

	"github.com/Focinfi/sakura/config"
	"github.com/Sirupsen/logrus"
)

func init() {
	if config.IsProduction() {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetLevel(logrus.WarnLevel)
		logrus.SetOutput(os.Stderr)
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetOutput(os.Stderr)
	}
}

// LogicErrorLogger for internal logic error
var LogicErrorLogger = logrus.New()

// DBErrorLogger for database error logger
var DBErrorLogger = logrus.New()

// LibErrorLogger for lib error logger
var LibErrorLogger = logrus.New()

// ThirdPartyServiceErrorLogger for service error logger
var ThirdPartyServiceErrorLogger = logrus.New()

// LogicError for logic error
func LogicError(funcName string, message interface{}) {
	LogicErrorLogger.
		WithFields(logrus.Fields{"function_name": funcName}).
		Fatal(message)
}

// DBError log databse error
func DBError(sql interface{}, err error, message interface{}) {
	DBErrorLogger.
		WithFields(logrus.Fields{"sql": sql, "error": err}).
		Fatal(message)
}

// LibError for lib error
func LibError(lib string, message interface{}) {
	DBErrorLogger.
		WithFields(logrus.Fields{"lib": lib}).
		Fatal(message)
}

// ThirdPartyServiceError for third-party service error
func ThirdPartyServiceError(thirdPartyService string, err error, message interface{}, params ...string) {
	DBErrorLogger.
		WithFields(logrus.Fields{"third_party_service": thirdPartyService, "error": err, "params": params}).
		Info(message)
}

// Debugf prints string for debuging
func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format+"\n", args...)
}

// Infof prints info string
func Infof(format string, args ...interface{}) {
	logrus.Infof(format+"\n", args...)
}
