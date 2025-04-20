package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func Init(debug bool) {
	Log.SetOutput(os.Stdout)
	Log.SetFormatter(&logrus.JSONFormatter{})

	if debug {
		Log.SetLevel(logrus.DebugLevel)
	} else {
		Log.SetLevel(logrus.InfoLevel)
	}
}
