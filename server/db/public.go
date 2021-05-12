package db

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/nodauf/ReconFramwork/server/models/database"
	"github.com/nodauf/ReconFramwork/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB

func Init() {
	dsn := "gorm:gorm@tcp(127.0.0.1:3306)/ReconFramwork?charset=utf8&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	//db.LogMode(true)
	// Auto Migrate
	conn.AutoMigrate(&database.Host{}, &database.Port{}, &database.PortComment{}) //, &database.HostsPorts{})
	// Set table options
	//db.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&database.Host{})
	conn.Debug().Migrator().CreateConstraint(&database.Host{}, "Ports")
	conn.Debug().Migrator().CreateConstraint(&database.Port{}, "Hosts")
	conn.Debug().Migrator().CreateConstraint(&database.Port{}, "PortComment")
	conn.Debug().Migrator().CreateConstraint(&database.PortComment{}, "Host")
	//conn.Debug().Migrator().CreateConstraint(&database.PortComment{}, "Port")
	db = conn
}

func GetHost(address string) database.Host {
	var host database.Host
	result := db.Where("address = ?", address).First(&host)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return database.Host{}
	}
	return host
}

func AddOrUpdateHost(host database.Host) uint {
	result := db.First(&host)
	//fmt.Println(result.RowsAffected) // returns found records count
	//fmt.Println(result.Error)        // returns error

	// check error ErrRecordNotFound
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db.Create(&host)
		fmt.Println("Create")
	} else {
		fmt.Println("before update")
		db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Debug().Save(&host)
		fmt.Println("Update")
		//fmt.Println(host)
	}
	//fmt.Println(host.ID)
	return host.ID
}

func HostHasService(target, serviceString string) []string {
	var targets []string
	var host database.Host
	services := utils.ParseList(serviceString)

	db.Where("address = ?", target).First(&host)
	for _, port := range host.Ports {

		if utils.StringInSlice(port.Service, services) {
			targets = append(targets, target+":"+strconv.Itoa(port.Port))
		}
	}
	return targets
}

/*func UpdateHost(host database.Host) bool {
	result := db.First(&host)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.ERROR.Println("Host " + host.Address + "not found in the database")
		return false
	}

	result = db.Save(&host)
	if result.RowsAffected > 0 {
		return true
	}
	return false

}*/

func DeleteHost(host database.Host) bool {
	result := db.Delete(&host)
	if result.RowsAffected > 0 {
		return true
	}
	return false

}
