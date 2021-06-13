package db

import (
	"errors"
	"net"

	modelsDatabases "github.com/nodauf/ReconFramwork/server/server/models/database"
	"gorm.io/gorm"
)

func GetDomain(domainStr string) modelsDatabases.Domain {
	var domain modelsDatabases.Domain
	// When we do a lot of select with go routing we need to use transaction to lock the table
	tx := db.Begin()

	result := tx.Where("domain = ?", domainStr).Preload("Host").Preload("Host.Ports").First(&domain)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return modelsDatabases.Domain{}
	}
	return domain
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
		db.Create(&domain)
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

func DeleteDomain(domain *modelsDatabases.Domain) bool {
	result := db.Unscoped().Delete(&domain)
	if result.RowsAffected > 0 {
		return true
	}
	return false

}

func GetDomainSubdomain(domain *modelsDatabases.Domain) []string {
	var subdomain []string
	var domains []modelsDatabases.Domain
	db.Where("domain LIKE ?", "%"+domain.Domain).Find(&domains)
	for _, domainDB := range domains {
		subdomain = append(subdomain, domainDB.Domain)
	}
	return subdomain
}
