package utils

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger
var once sync.Once

// start loggeando
func GetInstance() *logrus.Logger {
	once.Do(func() {
		logger = createLogger()
	})
	return logger
}

func createLogger() *logrus.Logger {
	l := logrus.New()
	return l
}
