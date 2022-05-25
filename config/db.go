package config

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

func RunMigrations(connectionString string, log *zap.Logger) error {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Error(err.Error())
		return err
	}
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
	err = m.Steps(2)

	if err != nil {
		log.Warn(err.Error())
		return err
	}
	return nil
}
