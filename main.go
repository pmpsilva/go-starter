package main

import (
	"context"
	"database/sql"
	"github.com/pmpsilva/go-starter/config"
	"go.uber.org/zap"
	"log"
	"os"
)

// todo to be removed on future versions.
func main() {

	logger, err := config.BuildLogger()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	//how to add an id to the logger (for example on a request)
	ctx := config.DeriveContextWithRequestId(context.Background())
	logger.Info("Log with an id", config.ZapFieldWithRequestIdFromCtx(ctx))

	connectionString, err := config.BuildDbString()
	if err != nil {
		logger.Error("Fail to get connection string", zap.Error(err))
		os.Exit(1)
	}

	dataSource, err := sql.Open("postgres", *connectionString)
	if err != nil {
		logger.Error("Unable to connect to database", zap.Error(err))
		os.Exit(1)
	}
	defer func(dataSource *sql.DB) {
		err := dataSource.Close()
		if err != nil {
			return
		}
	}(dataSource)
	if err != nil {
		logger.Error("Fail to close connection to  database", zap.Error(err))
		os.Exit(1)
	}
	_ = config.RunMigrations(dataSource, logger)

}
