package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/usmhk/goserver/pkg/db"
	"sync"

	"github.com/usmhk/goserver/internal/apiserver/store"
	"gorm.io/gorm"
)

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) Close() error {
	if ds.db == nil {
		return nil
	}

	db, err := ds.db.DB()
	if err != nil {
		return nil
	}
	return db.Close()
}

var (
	mysqlFactory store.Factory
	once         sync.Once
)

// GetMysqlFactoryOr create mysql factory with the give config.
func GetMysqlFactoryOr() (store.Factory, error) {
	var (
		err   error
		dbIns *gorm.DB
	)
	once.Do(func() {
		options := &db.Options{
			Host:                  viper.GetString("db.host"),
			Username:              viper.GetString("db.username"),
			Password:              viper.GetString("db.password"),
			Database:              viper.GetString("db.database"),
			MaxIdleConnections:    viper.GetInt("db.max-idle-connections"),
			MaxOpenConnections:    viper.GetInt("db.max-open-connections"),
			MaxConnectionLifeTime: viper.GetDuration("db.max-connection-life-time"),
			LogLevel:              viper.GetInt("db.log-level"),
			Logger:                nil,
		}
		dbIns, err = db.New(options)

		mysqlFactory = &datastore{dbIns}
	})
	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store factory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}
	return mysqlFactory, nil
}
