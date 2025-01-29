package helpers

import "github.com/sirupsen/logrus"

type LoggerInterface interface {
	Error(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	WithField(key string, value interface{}) LoggerInterface
}

type LogrusWrapper struct {
	entry *logrus.Entry
}

func (l *LogrusWrapper) Error(args ...interface{}) {
	l.entry.Error(args...)
}

func (l *LogrusWrapper) Info(args ...interface{}) {
	l.entry.Info(args...)
}

func (l *LogrusWrapper) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

func (l *LogrusWrapper) WithField(key string, value interface{}) LoggerInterface {
	return &LogrusWrapper{entry: l.entry.WithField(key, value)}
}

var Logger LoggerInterface

func SetupLogger() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	Logger = &LogrusWrapper{entry: logrus.NewEntry(logger)}
	Logger.Info("Logger initialized")
}
