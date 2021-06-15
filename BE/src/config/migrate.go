package config

import (
	"app/src/domain/models"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

// DBMigrate will create & migrate the tables, then make the some relationships if necessary
func DBMigrate() (*gorm.DB, error) {
	log.Println("Migration start")
	conn, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if conn.HasTable(models.SignalEvent{}) {
		conn.DropTable(models.SignalEvent{})
	}

	conn.AutoMigrate(models.SignalEvent{})

	log.Println("Migration has been processed")
	return conn, nil
}

func Seeds() (*gorm.DB, error) {
	conn, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	testSignalEvent1 := &models.SignalEvent{
		Time:        time.Now(),
		ProductCode: "BTC-JPY",
		Side:        "BUY",
		Price:       1000.0,
		Size:        50.0,
	}

	if err := conn.Debug().Create(testSignalEvent1).Error; err != nil {
		return nil, err
	}

	return nil, err
}
