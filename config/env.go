package config

import (
	"os"
	"strconv"
	"strings"
)

const (
	loggerEnv        = "LOGGER_ENV"
	defaultLoggerEnv = "development"
)

func GetVariableValue(variableName string) string {
	variable := os.Getenv(variableName)
	return variable
}

func ReadEnvVarOrDefault(varName string, defaultValue string) (varVal string) {
	varVal = strings.TrimSpace(os.Getenv(varName))

	if varVal == "" {
		varVal = defaultValue
	}

	return
}

func ReadEnvBoolVarOrDefault(varName string, defaultValue bool) bool {
	varVal, err := strconv.ParseBool(os.Getenv(varName))

	if err != nil {
		varVal = defaultValue
	}

	return varVal
}
