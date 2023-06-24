package starter

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
	"strconv"
)

const (
	dbHost     = "DB_HOST"
	dbPort     = "DB_PORT"
	dbUser     = "DB_USER"
	dbPassword = "DB_PASSWORD"
	dbName     = "DB_NAME"
)

func BuildDbString() (*string, error) {
	host := GetVariableValue(dbHost)
	port, err := strconv.Atoi(GetVariableValue(dbPort))
	if err != nil {
		return nil, errors.New("fail to get env variable")
	}
	user := GetVariableValue(dbUser)
	password := GetVariableValue(dbPassword)
	name := GetVariableValue(dbName)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, name)

	return &psqlInfo, nil
}

func OpenDbConnection(connectionString *string, logger *zap.Logger) (*sql.DB, error) {
	dataSource, err := sql.Open("postgres", *connectionString)
	if err != nil {
		logger.Error("Unable to connect to database", zap.Error(err))
		return nil, err
	}
	defer func(dataSource *sql.DB) {
		err := dataSource.Close()
		if err != nil {
			return
		}
	}(dataSource)
	if err != nil {
		logger.Error("Fail to close connection to  database", zap.Error(err))
		return nil, err
	}
	return dataSource, nil
}

func RunMigrations(db *sql.DB, log *zap.Logger) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Error(err.Error())
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		GetVariableValue(dbName), driver)
	if err != nil {
		log.Warn(err.Error())
		return err
	}

	err = m.Up()

	if err != nil {
		log.Info(err.Error())
		return err
	}
	return nil
}
