# go-starter

bootstrap for golang project with logg database and env

# usage/configurations

## Logger:

**Zap**

```go
    logger, err := init.BuildLogger()
    if err != nil {
        log.Fatalf("can't initialize zap logger: %v", err)
    }
    
    //how to add an id to the logger (for example on a request)
    ctx := init.DeriveContextWithRequestId(context.Background())
    logger.Info("Log with an id", config.ZapFieldWithRequestIdFromCtx(ctx))

```

## DB
 For database this project uses postgres and the migrations will be on db/migrations folder
 ```go
    //red env variables
    connectionString, err := starter.BuildDbString()
    if err != nil {
        logger.Error("Fail to get connection string", zap.Error(err))
        os.Exit(1)
    }
    //open DbConnecion
    dbConnection, err := starter.OpenDbConnection(connectionString, logger)
    if err != nil {
        os.Exit(1)
    }
	//to use migrations on //db/migartions
    _ = starter.RunMigrations(dataSource, logger)
    
    //transactionManager initialization 
    transactionManager := starter.NewTransactionManager(dataSource)
    
    //usage example at service or repository level
	var resultToReturn uuid.UUID
	if err := transactionManager.ExecWithTransaction(func(tx *sql.Tx) error {
		//repository method to perform some query to db
		dbResult := uuid.New()
		if err != nil {
			return err
		}
		//attach result to external variable
		resultToReturn = dbResult

		return nil
	}); err != nil {
		//deal with error
	}
```

## Env Variables

|     Variable Name      | Default Value | Required |     Possible values     |
|:----------------------:|:-------------:|:--------:|:-----------------------:|
|        DB_HOST         |      nil      |   true   |            *            |
|        DB_PORT         |      nil      |   true   |            *            |
|        DB_USER         |      nil      |   true   |            *            |
|      DB_PASSWORD       |      nil      |   true   |            *            |
|        DB_NAME         |      nil      |   true   |            *            |
|       HTTP_HOST        |    0.0.0.0    |  false   |            *            |
|       HTTP_PORT        |     8080      |  false   |            *            |
|      CONTEXT_PATH      |       /       |  false   |            *            |
| NET_HTTP_PPROF_ENABLED |     false     |  false   |       true, false       |
|       LOGGER_ENV       |  development  |  false   | development, production |
