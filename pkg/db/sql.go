package db

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ezzycreative1/svc-pokemon/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(config *config.Database) *gorm.DB {
	// gorm mysql orm
	dbConfig := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Schema,
	)

	maxConn := 100
	if config.MaxConn != 0 {
		maxConn = config.MaxConn
	}
	return createDatabaseInstance(dbConfig, maxConn)
}

func createDatabaseInstance(dsn string, maxConn int) *gorm.DB {
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		// Logger:                 logger.Default.LogMode(logger.Silent),
	}

	inst, err := gorm.Open(
		mysql.Open(dsn),
		gormConfig,
	)
	if err != nil {
		log.Panic(err)
	}

	db, err := inst.DB()
	if err != nil {
		log.Panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(maxConn)
	db.SetConnMaxLifetime(time.Minute * 20)

	return inst
}

func OpenMySQLConnection(dbConfig *config.MySQLConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Schema,
	)

	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
	}
	if strings.TrimSpace(strings.ToLower(dbConfig.Environment)) == "prod" {
		gormConfig.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             6 * time.Second, // Slow SQL threshold
				LogLevel:                  logger.Error,    // Log level
				IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,           // Disable color
			},
		)
	}
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		log.Panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Panic(err)
	}
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(dbConfig.ConnMaxLifetime))

	return db
}
