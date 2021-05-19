package db

import (
	"errors"
	"fmt"

	"github.com/nodauf/ReconFramwork/server/models/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Init() {
	dsn := "gorm:gorm@tcp(127.0.0.1:3306)/ReconFramwork?charset=utf8&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//conn.Logger = conn.Logger.LogMode(logger.Info)
	conn.Logger = conn.Logger.LogMode(logger.Silent)
	// Auto Migrate
	conn.AutoMigrate(&database.Host{}, &database.Port{}, &database.PortComment{}, &database.Job{}, &database.Domain{}) //, &database.HostsPorts{})
	// Set table options
	//db.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&database.Host{})
	/*conn.Debug().Migrator().CreateConstraint(&database.Host{}, "Ports")
	conn.Debug().Migrator().CreateConstraint(&database.Port{}, "Hosts")
	conn.Debug().Migrator().CreateConstraint(&database.Port{}, "PortComment")
	conn.Debug().Migrator().CreateConstraint(&database.PortComment{}, "Host")
	conn.Debug().Migrator().CreateConstraint(&database.PortComment{}, "Port")*/
	db = conn
}

func GetHost(address string) database.Host {
	var host database.Host
	result := db.Where("address = ? or hostname = ?", address, address).Preload("Ports").Preload("Ports.PortComment").First(&host)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return database.Host{}
	}
	return host
}

func GetDomain(domainStr string) database.Domain {
	var domain database.Domain
	result := db.Where("domain = ?", domainStr).Preload("Host").Preload("Host.Ports").First(&domain)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return database.Domain{}
	}
	return domain
}

func GetHostWherePort(address, port string) database.Host {
	var host database.Host
	result := db.Joins("JOIN ports ON ports.host_id = hosts.id ").Where("address = ?", address).Preload("Ports").Preload("Ports.PortComment").First(&host)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return database.Host{}
	}
	return host
}

func AddOrUpdateHost(host *database.Host) uint {
	db.Session(&gorm.Session{FullSaveAssociations: true}).Where("address = ? ", host.Address).Debug().FirstOrCreate(host)

	//fmt.Println(result.RowsAffected) // returns found records count
	//fmt.Println(result.Error)        // returns error

	// check error ErrRecordNotFound
	/*if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		//db.Debug().Create(&host)
		fmt.Println("create")
		db.Debug().Create(&host)
	} else {
		fmt.Println("update")
		//db.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&host)
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&host)
		//fmt.Println(host)
	}
	//fmt.Println(host.ID)*/
	return host.ID
}

func AddOrUpdateDomain(domain *database.Domain) uint {
	db.Session(&gorm.Session{FullSaveAssociations: true}).Where("domain = ? ", domain.Domain).Debug().FirstOrCreate(domain)

	//fmt.Println(result.RowsAffected) // returns found records count
	//fmt.Println(result.Error)        // returns error

	// check error ErrRecordNotFound
	/*if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		//db.Debug().Create(&host)
		fmt.Println("create")
		db.Debug().Create(&host)
	} else {
		fmt.Println("update")
		//db.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&host)
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&host)
		//fmt.Println(host)
	}
	//fmt.Println(host.ID)*/
	return domain.ID
}

/*func HostHasService(target string, serviceCommand map[string]models.CommandService) map[string]string {
	var targets = make(map[string]string)
	var host database.Host
	db.Where("address = ?", target).Preload("Ports").First(&host)
	for _, port := range host.Ports {

		if _, ok := serviceCommand[port.Service]; ok {
			targets[port.Service] = target + ":" + strconv.Itoa(port.Port)
		}
	}
	return targets
}*/

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

func AddJob(target, parser, taskUUID string) (database.Job, error) {
	var err error
	host := GetHost(target)
	domain := GetDomain(target)
	var job database.Job
	if host.ID != 0 {
		job.Host = host
		job.TaskUUID = taskUUID
		job.Processed = false
		job.Parser = parser
		db.Debug().Create(&job)
		db.Preload("Host").Debug().First(&job)
	} else if domain.ID != 0 {
		job.Domain = domain
		job.TaskUUID = taskUUID
		job.Processed = false
		job.Parser = parser
		db.Debug().Create(&job)
		db.Preload("Host").Debug().First(&job)

	} else {
		err = errors.New("Cannot attach the job to an host or domain")
	}
	return job, err
}

func AddDomain(domain database.Domain) {
	result := db.Where("domain = ? ", domain.Domain).Debug().First(&domain)
	//fmt.Println(result.RowsAffected) // returns found records count
	//fmt.Println(result.Error)        // returns error

	// check error ErrRecordNotFound
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		//db.Debug().Create(&host)
		fmt.Println("create domain")
		db.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Create(&domain)
	} /*else {
		fmt.Println("update")
		//db.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&host)
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&host)
		//fmt.Println(host)
	}*/
}

func RemoveJob(job *database.Job) {
	db.Model(&database.Job{}).Where("id = ?", job.ID).Update("processed", true)
}

func GetNonProcessedTasks() []database.Job {
	var jobs []database.Job
	db.Where("processed = ?", false).Preload("Host").Find(&jobs)
	return jobs
}
