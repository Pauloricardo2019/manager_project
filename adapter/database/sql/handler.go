package sql

import (
	"gerenciador/adapter/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var gormDb *gorm.DB
var mutexDB sync.Mutex

func GetGormDB() (*gorm.DB, error) {
	mutexDB.Lock()
	defer mutexDB.Unlock()

	if gormDb != nil {
		return gormDb, nil
	}

	newDb, err := gorm.Open(postgres.Open(config.GetConfig().DbConnString), &gorm.Config{FullSaveAssociations: true})
	if err != nil {
		return nil, err
	}

	gormDb = newDb

	return gormDb, nil
}
