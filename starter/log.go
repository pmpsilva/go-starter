package starter

import (
	"go.uber.org/zap"
	"strings"
)

// BuildLogger build a zap logger by default is development or production if is set on LOGGER_ENV variable
func BuildLogger() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			return
		}
	}(logger)
	if err != nil {
		return nil, err
	}

	if strings.EqualFold(ReadEnvVarOrDefault(loggerEnv, defaultLoggerEnv), "production") {
		logger, err = zap.NewProduction()
		defer func(logger *zap.Logger) {
			err := logger.Sync()
			if err != nil {
				return
			}
		}(logger)
		if err != nil {
			return nil, err
		}

	}

	return logger, nil

}
