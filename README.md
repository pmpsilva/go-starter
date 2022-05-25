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
```

## DB

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
	if err != nil {
		logger.Error("Fail to run migrations", zap.Error(err))
		os.Exit(1)
	}
```

## Env Variables

| Variable Name | Default Value | Required |
|:-------------:|:-------------:|:--------:|
|    DB_HOST    |      nil      |   true   |
|    DB_PORT    |      nil      |   true   |
|    DB_USER    |      nil      |   true   |
|  DB_PASSWORD  |      nil      |   true   |
|    DB_NAME    |      nil      |   true   |
