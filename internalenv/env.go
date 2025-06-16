package internalenv

import (
	"log"
	"os"
	"strconv"
)

type Env struct {
	AppEnv            string `json:"app_env"`
	JwtExpTimeMinutes int    `json:"JWT_EXP_TIME_MINUTES"`
	JwtSecret         string `json:"JWT_SECRET"`
	ServerAddress     string `json:"address"`
	ServerPort        string `json:"port"`
	TimeOut           int    `json:"timeout"`
	DBName            string `json:"db_name"`
	DBHost            string `json:"db_host"`
	DBPort            int    `json:"db_port"`
	DBUser            string `json:"db_user"`
	DBPassword        string `json:"db_password"`
}

type EnvInteface interface {
	CombineServerAddress() string
}

func (env *Env) CombineServerAddress() string {
	return env.ServerAddress + ":" + env.ServerPort
}

const (
	envName           = "APP_ENV"
	jwtExpTimeMinutes = "SERVER_CONFIG_JWT_EXP_TIME_MINUTES"
	jwtSecret         = "SERVER_CONFIG_JWT_SECRET"
	serverAddress     = "SERVER_CONFIG_ADDRESS"
	serverPort        = "SERVER_CONFIG_PORT"
	timeOut           = "SERVER_CONFIG_TIMEOUT"
	dbname            = "SERVER_CONFIG_DB_NAME"
	dbhost            = "SERVER_CONFIG_DB_HOST"
	dbport            = "SERVER_CONFIG_DB_PORT"
	dbuser            = "SERVER_CONFIG_DB_USER"
	dbpassword        = "SERVER_CONFIG_DB_PASSWORD"
)

func LoadEnv() *Env {
	env, err := parseEnv()
	if err != nil {
		log.Default().Fatalf("Failed to parse env")
		log.Default().Fatal(err)
		return nil
	}
	log.Default().Printf("The App is running in %s env", env.AppEnv)
	return env
}

func parseEnv() (*Env, error) {
	appEnvName := os.Getenv(envName)
	appJwtExpTimeMinutes, err := strconv.Atoi(os.Getenv(jwtExpTimeMinutes))
	if err != nil {
		return nil, err
	}
	appJwtSecret := os.Getenv(jwtSecret)
	appServerAddress := os.Getenv(serverAddress)
	appServerPort := os.Getenv(serverPort)
	appTimeOut, err := strconv.Atoi(os.Getenv(timeOut))
	if err != nil {
		return nil, err
	}
	appDbName := os.Getenv(dbname)
	appDbHost := os.Getenv(dbhost)
	appDbPort, err := strconv.Atoi(os.Getenv(dbport))
	if err != nil {
		return nil, err
	}
	appDbUser := os.Getenv(dbuser)
	appDbPassword := os.Getenv(dbpassword)
	return &Env{AppEnv: appEnvName,
		JwtExpTimeMinutes: appJwtExpTimeMinutes,
		JwtSecret:         appJwtSecret,
		ServerAddress:     appServerAddress,
		ServerPort:        appServerPort,
		TimeOut:           appTimeOut,
		DBName:            appDbName,
		DBHost:            appDbHost,
		DBPort:            appDbPort,
		DBUser:            appDbUser,
		DBPassword:        appDbPassword}, nil
}
