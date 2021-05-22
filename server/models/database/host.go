package modelsDatabases

import (
	"strconv"

	"github.com/nodauf/ReconFramwork/server/models"
	"gorm.io/gorm"
)

type Host struct {
	gorm.Model
	//Hostname string   //`gorm:"uniqueindex:idx_hosts"`
	Address string   `gorm:"uniqueindex:idx_hosts"`
	Ports   []Port   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` //`gorm:"many2many:Hosts_Ports;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Domain  []Domain `gorm:"many2many:Hosts_Domains;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	//PortDetail HostsPorts
}

func (host *Host) HasService(serviceCommand map[string]models.CommandService) map[string]string {
	var targets = make(map[string]string)
	for _, port := range host.Ports {

		if _, ok := serviceCommand[port.Service]; ok {
			targets[port.Service] = host.Address + ":" + strconv.Itoa(port.Port)
		}
	}
	return targets
}

func (host *Host) HasSubdomain() bool {
	// A host can't have a subdomain
	return false
}

func (host *Host) GetTarget() string {
	return host.Address
}
