package database

import (
	"strconv"

	"github.com/nodauf/ReconFramwork/server/models"
	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
	Domain string `gorm:"uniqueindex:idx_domain"`
	Host   []Host `gorm:"many2many:Hosts_Domains;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (domain *Domain) HasService(serviceCommand map[string]models.CommandService) map[string]string {
	var targets = make(map[string]string)
	for _, host := range domain.Host {
		for _, port := range host.Ports {

			if _, ok := serviceCommand[port.Service]; ok {
				targets[port.Service] = host.Address + ":" + strconv.Itoa(port.Port)
			}
		}
	}
	return targets
}
