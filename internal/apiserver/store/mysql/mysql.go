package mysql

import (
	"fmt"
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

func GetMysqlFactoryOr() (store.Factory, error) {
	var (
		err   error
		dbIns *gorm.DB
	)
	once.Do(func() {
		//TODO
	})
	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store factory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}
	return mysqlFactory, nil
}
