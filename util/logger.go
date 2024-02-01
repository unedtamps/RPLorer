package util

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	Log.SetOutput(os.Stdout)
	Log.SetReportCaller(true)
}
