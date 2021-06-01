package db

import (
	"errors"
	"net"
	"strconv"

	"github.com/nodauf/ReconFramwork/server/server/models"
	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
	"github.com/nodauf/ReconFramwork/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	dsn := "gorm:gorm@tcp(127.0.0.1:3306)/ReconFramwork?charset=utf8&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//conn.Logger = conn.Logger.LogMode(logger.Info)
	conn.Logger = conn.Logger.LogMode(logger.Silent)
	// Auto Migrate
	conn.AutoMigrate(&modelsDatabases.Host{}, &modelsDatabases.Port{}, &modelsDatabases.PortComment{}, &modelsDatabases.Job{}, &modelsDatabases.Domain{}) //, &database.HostsPorts{})
	// The foreign key is not needed by gorm and we don't need it as NULL != NULL, it will break unique index
	conn.Migrator().DropConstraint(&modelsDatabases.PortComment{}, "fk_domains_port_comment")
	// Set table options
	//db.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&database.Host{})
	/*conn.Debug().Migrator().CreateConstraint(&database.Host{}, "Ports")
	conn.Debug().Migrator().CreateConstraint(&database.Port{}, "Hosts")
	conn.Debug().Migrator().CreateConstraint(&database.Port{}, "PortComment")
	conn.Debug().Migrator().CreateConstraint(&database.PortComment{}, "Host")
	conn.Debug().Migrator().CreateConstraint(&database.PortComment{}, "Port")*/
	db = conn
}

func GetHost(address string) modelsDatabases.Host {
	var host modelsDatabases.Host
	result := db.Where("address = ?", address).Preload("Ports").Preload("Ports.PortComment").First(&host)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return modelsDatabases.Host{}
	}
	return host
}

func GetDomain(domainStr string) modelsDatabases.Domain {
	var domain modelsDatabases.Domain
	result := db.Where("domain = ?", domainStr).Preload("Host").Preload("Host.Ports").Preload("Subdomain").First(&domain)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return modelsDatabases.Domain{}
	}
	return domain
}

func GetTarget(target string) models.Target {
	var targetObject models.Target
	host := GetHost(target)
	domain := GetDomain(target)
	// If there is nothing in the datbase for this target
	if host.Address != "" || domain.Domain != "" {
		if host.Address != "" {
			targetObject = &host
		} else {
			targetObject = &domain
		}
	}
	return targetObject
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

func AddOrUpdateDomain(domain *modelsDatabases.Domain) modelsDatabases.Domain {
	var tmp modelsDatabases.Domain
	var hosts []modelsDatabases.Host
	result := db.Where("domain = ? ", domain.Domain).Preload("Host").Preload("Host.Ports").First(&tmp)

	// If the domain has no host we see if there is the IP and add the IP to the host
	if len(tmp.Host) == 0 {
		resolveDomain, err := net.LookupHost(domain.Domain)
		if err == nil {
			result := db.Debug().Where("address IN ?", resolveDomain).Find(&hosts)
			if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
				domain.Host = hosts
			}
		}
	}
	//fmt.Println(result.RowsAffected) // returns found records count
	//fmt.Println(result.Error)        // returns error

	// check error ErrRecordNotFound
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db.Debug().Create(&domain)
	} else {
		//db.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&host)
		db.Session(&gorm.Session{FullSaveAssociations: true}).Preload("Host").Save(&domain)
		//fmt.Println(host)
	}
	// Get the full object of the database
	db.Where("domain = ? ", domain.Domain).Preload("Host").Preload("Host.Ports").First(domain)

	//fmt.Println(host.ID)
	return *domain
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

func DeleteHost(host modelsDatabases.Host) bool {
	result := db.Delete(&host)
	if result.RowsAffected > 0 {
		return true
	}
	return false

}

func AddJob(target, parser, taskUUID, machineryTask, MachineryTaskArgs string) (modelsDatabases.Job, error) {
	var err error
	host := GetHost(target)
	domain := GetDomain(target)
	var job modelsDatabases.Job
	job.TaskUUID = taskUUID
	job.Processed = false
	job.Parser = parser
	job.MachineryTask = machineryTask
	job.MachineryTaskArgs = MachineryTaskArgs
	if host.ID != 0 {
		job.Host = host
		db.Create(&job)
		db.Preload("Host").First(&job)
	} else if domain.ID != 0 {
		job.Domain = domain
		db.Create(&job)
		db.Preload("Domain").First(&job)

	} else {
		err = errors.New("Cannot attach the job to an host or domain")
	}
	return job, err
}

func AddDomain(domain modelsDatabases.Domain) {
	result := db.Where("domain = ? ", domain.Domain).First(&domain)
	//fmt.Println(result.RowsAffected) // returns found records count
	//fmt.Println(result.Error)        // returns error

	// check error ErrRecordNotFound
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		//db.Debug().Create(&host)
		db.Session(&gorm.Session{FullSaveAssociations: true}).Create(&domain)
	} /*else {
		fmt.Println("update")
		//db.Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&host)
		db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&host)
		//fmt.Println(host)
	}*/
}

func RemoveJob(job *modelsDatabases.Job) {
	db.Model(&modelsDatabases.Job{}).Where("id = ?", job.ID).Update("processed", true)
}

func GetNonProcessedTasks() []modelsDatabases.Job {
	var jobs []modelsDatabases.Job
	db.Where("processed = ?", false).Preload("Host").Find(&jobs)
	return jobs
}

func AddOrUpdateTarget(target models.Target) models.Target {
	var targetToReturn models.Target
	if utils.IsIP(target.GetTarget()) {
		host := AddOrUpdateHost(target.(*modelsDatabases.Host))
		targetToReturn = &host
	} else {
		domain := AddOrUpdateDomain(target.(*modelsDatabases.Domain))
		targetToReturn = &domain
	}
	return targetToReturn
}

func AddPortComment(targetObject models.Target, port int, portComment modelsDatabases.PortComment) error {

	if index := targetObject.HasPort(port); index != -1 {
		targetList, err := targetObject.AddPortComment(port, portComment)
		if err != nil {
			return err
		}
		for _, target := range targetList {
			AddOrUpdateHost(&target)
		}
	} else {
		return errors.New("The target " + targetObject.GetTarget() + " has not the port " + strconv.Itoa(port))
	}
	return nil
}
