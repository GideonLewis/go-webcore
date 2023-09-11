package mysql

import (
	"fmt"
	"time"

	"github.com/megaqstar/web-core/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

var db *gorm.DB

func connect(cfg config.MySQL) error {
	connectionString := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	dbConn, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Error),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Error(fmt.Sprintf("Failed to open database connection: %s", err))
		return err
	}

	db = dbConn.Debug()
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(time.Minute * time.Duration(cfg.ConnMaxIdleTime))
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(cfg.ConnMaxLifeTime))
	sqlDB.SetMaxIdleConns(cfg.DBMaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.DBMaxOpenConns)

	err = sqlDB.Ping()
	if err != nil {
		log.Error(fmt.Sprintf("MySQL database is no longer active : %s", err))
		return err
	}
	db = dbConn
	return nil
}

func GetClient(cfg config.MySQL) (*gorm.DB, error) {
	var err error
	if db == nil {
		if err = connect(cfg); err != nil {
			return nil, err
		}
		return db, nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func Disconnect() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Error(fmt.Sprintf("MySQL database is no longer active : %s", err))
		}
		sqlDB.Close()
	}
}
