package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pmpsilva/go-starter/config"
	"go.uber.org/zap"
	"log"
	"os"
)

//todo to be removed on future versions.
func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	//how to add ran id to a logger for example on a request
	ctx := config.DeriveContextWithRequestId(context.Background())
	logger.Info("Log with an id", zap.Any("request_id", config.AddCtxAndRequestIDIfPresent(ctx)))

	connectionString, err := config.BuildDbString()
	if err != nil {
		logger.Error("Fail to get connection string", zap.Error(err))
		return
	}
	conn, err := pgx.Connect(context.Background(), *connectionString)
	if err != nil {
		logger.Error("Unable to connect to database", zap.Error(err))
	}
	defer conn.Close(context.Background())

	_ = config.RunMigrations(*connectionString, logger)
	if err != nil {
		logger.Error("Fail to run migrations", zap.Error(err))
		os.Exit(1)
	}

}
