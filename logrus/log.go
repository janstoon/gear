package logrus

import (
	"github.com/sirupsen/logrus"

	"gitlab.com/janstun/gear"
)

type levelled struct {
	logger *logrus.Logger
}

var lvlMap = map[gear.LogLevel]logrus.Level{
	gear.FATAL: logrus.FatalLevel,
	gear.ERROR: logrus.ErrorLevel,
	gear.WARN:  logrus.WarnLevel,
	gear.INFO:  logrus.InfoLevel,
	gear.DEBUG: logrus.DebugLevel,
}

func NewLevelledLogger(lvl gear.LogLevel) (gear.LevelledLogger, error) {
	logger := logrus.New()
	logger.Level = lvlMap[lvl]

	return levelled{logger}, nil
}

func (s levelled) Log(level gear.LogLevel, format string, v ...interface{}) {
	switch level {
	case gear.FATAL:
		s.logger.Fatalf(format, v...)

	case gear.ERROR:
		s.logger.Errorf(format, v...)

	case gear.WARN:
		s.logger.Warnf(format, v...)

	case gear.INFO:
		s.logger.Infof(format, v...)

	case gear.DEBUG:
		s.logger.Debugf(format, v...)

	default:
		s.logger.Printf(format, v...)
	}
}
