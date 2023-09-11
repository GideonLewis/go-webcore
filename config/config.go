package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/megaqstar/web-core/common"
	log "github.com/sirupsen/logrus"
)

type MySQL struct {
	Host            string `envconfig:"MYSQL_HOST"`
	Port            string `envconfig:"MYSQL_PORT"`
	User            string `envconfig:"MYSQL_USER"`
	Password        string `envconfig:"MYSQL_PASSWORD"`
	DBName          string `envconfig:"MYSQL_NAME"`
	ConnMaxIdleTime int    `envconfig:"MYSQL_CONN_MAX_IDLE_TIME"`
	ConnMaxLifeTime int    `envconfig:"MYSQL_CONN_MAX_LIFE_TIME"`
	DBMaxIdleConns  int    `envconfig:"MYSQL_MAX_IDLE_CONNS"`
	DBMaxOpenConns  int    `envconfig:"MYSQL_MAX_OPEN_CONNS"`
}

type Firebase struct {
	ProjectID string `json:"project_id" `
}

type Json struct {
	Firebase Firebase
}

type Env struct {
	MySQL MySQL
}

type Config struct {
	Json Json
	Env  Env
}

var appConfig *Config

func GetConfig() (*Config, error) {
	if appConfig == nil {
		err := loadENVConfig()
		if err != nil {
			return nil, err
		}

		err = loadJSONConfig()
		if err != nil {
			return nil, err
		}
	}
	return appConfig, nil
}

func loadENVConfig() error {
	envCfg := &Env{}
	err := godotenv.Load()
	if err != nil {
		log.Errorf(fmt.Sprintf(common.ERR_ENV_LOADING, err))
		return err
	}

	err = envconfig.Process("", envCfg)
	if err != nil {
		log.Errorf(fmt.Sprintf(common.ERR_ENV_PROCESS, err))
		return err
	}

	if appConfig == nil {
		appConfig = &Config{
			Env: *envCfg,
		}
		return nil
	}

	appConfig.Env = *envCfg
	return nil
}

func loadJSONConfig() error {
	jsonCfg := &Json{}
	configFile, err := os.Open(common.FIREBASE_SERVICE_ACCOUNT)
	if err != nil {
		log.Errorf(fmt.Sprintf(common.ERR_JSON_LOADING, err))
		return err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&jsonCfg)

	if appConfig == nil {
		appConfig = &Config{
			Json: *jsonCfg,
		}
		return nil
	}

	appConfig.Json = *jsonCfg
	return nil
}
