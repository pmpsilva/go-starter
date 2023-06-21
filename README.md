# go-starter

bootstrap for golang project with logg database and env

# usage/configurations

## Logger:

**Zap**

```
    logger, err := config.BuildLogger()
    if err != nil {
        log.Fatalf("can't initialize zap logger: %v", err)
    }
    
    //how to add an id to the logger (for example on a request)
    ctx := config.DeriveContextWithRequestId(context.Background())
    logger.Info("Log with an id", config.ZapFieldWithRequestIdFromCtx(ctx))

```

## DB
 For database this project uses postgres and the migrations will be on db/migrations folder
 ```
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
        logger.Error("Fail to run migrations", zap.Error(err))
    os.Exit(1)
    }

    _ = config.RunMigrations(dataSource, logger)
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
