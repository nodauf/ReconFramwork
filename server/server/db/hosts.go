package db

import (
	"errors"

	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
	"gorm.io/gorm"
)

func GetHost(address string) modelsDatabases.Host {
	var host modelsDatabases.Host
	result := db.Where("address = ?", address).Preload("Ports").Preload("Ports.PortComment").First(&host)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return modelsDatabases.Host{}
	}
	return host
}

func GetHostWherePort(address, port string) modelsDatabases.Host {
	var host modelsDatabases.Host
	result := db.Joins("JOIN ports ON ports.host_id = hosts.id ").Where("address = ?", address).Preload("Ports").Preload("Ports.PortComment").First(&host)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return modelsDatabases.Host{}
	}
	return host
}

func AddOrUpdateHost(host *modelsDatabases.Host) modelsDatabases.Host {
	var tmp modelsDatabases.Host
	result := db.Where("address = ? ", host.Address).Preload("Ports").First(&tmp)
	//fmt.Println(result.RowsAffected) // returns found records count
	//fmt.Println(result.Error)        // returns error
	// check error ErrRecordNotFound. If the record does not exist we create it. Otherwise we update it
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db.Create(&host)
	} else {
		//db.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&host)
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&host)
		//fmt.Println(host)
	}
	// Get the full object of the database
	db.Where("address = ? ", host.Address).Preload("Ports").First(host)

	//fmt.Println(host.ID)
	return *host
}

func DeleteHost(host modelsDatabases.Host) bool {
	result := db.Delete(&host)
	if result.RowsAffected > 0 {
		return true
	}
	return false

}
