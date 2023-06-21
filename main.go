package main

import (
	"context"
	"database/sql"
	"github.com/pmpsilva/go-starter/init"
	"go.uber.org/zap"
	"log"
	"os"
)

// todo to be removed on future versions.
func main() {

	logger, err := init.BuildLogger()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	//how to add an id to the logger (for example on a request)
	ctx := init.DeriveContextWithRequestId(context.Background())
	logger.Info("Log with an id", init.ZapFieldWithRequestIdFromCtx(ctx))

	connectionString, err := init.BuildDbString()
	if err != nil {
		logger.Error("Fail to get connection string", zap.Error(err))
		os.Exit(1)
	}

	//open DbConnecion
	dbConnection, err := init.OpenDbConnection(connectionString, logger)
	if err != nil {
		os.Exit(1)
	}

	//to use migrations at //db/migrations
	_ = init.RunMigrations(dbConnection, logger)

	//transaction manager usage
	//initialization
	transactionManager := init.NewTransactionManager(dbConnection)
	//at service or repository level
	if err := transactionManager.ExecWithTransaction(func(tx *sql.Tx) error {
		//repository method to perform some query to db
		//dbResult := uuid.New()
		if err != nil {
			return err
		}
		//attach result to external variable
		//resultToReturn = dbResult

		return nil
	}); err != nil {
		//deal with error
	}

}
