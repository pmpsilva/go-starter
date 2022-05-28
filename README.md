# go-starter

bootstrap for golang project with logg database and env

# usage/configurations

## Logger:

**Zap**

```
    logger, err := zap.NewProduction()
    if err != nil {
    log.Fatalf("can't initialize zap logger: %v", err)
    }
    defer logger.Sync()
    
    logger.Info("test")
  
    //how to add ran id to a logger for example on a request
    ctx := config.DeriveContextWithRequestId(context.Background())
    logger.Info("Log with an id", zap.Any("request_id", config.AddCtxAndRequestIDIfPresent(ctx)))
```

## DB
 For database this project uses postgres and the migrations will be on db/migrations folder
 ```
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
```

## Env Variables

|     Variable Name      | Default Value | Required |
|:----------------------:|:-------------:|:--------:|
|        DB_HOST         |      nil      |   true   |
|        DB_PORT         |      nil      |   true   |
|        DB_USER         |      nil      |   true   |
|      DB_PASSWORD       |      nil      |   true   |
|        DB_NAME         |      nil      |   true   |
|       HTTP_HOST        |    0.0.0.0    |  false   |
|       HTTP_PORT        |     8080      |  false   |
|      CONTEXT_PATH      |       /       |  false   |
| NET_HTTP_PPROF_ENABLED |     false     |  false   |
