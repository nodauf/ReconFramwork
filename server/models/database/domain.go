package modelsDatabases

import (
	"strconv"

	"github.com/nodauf/ReconFramwork/server/models"
	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
	Domain        string   `gorm:"uniqueindex:idx_domain"`
	SubdomainOfID *uint    `gorm:"uniqueindex:idx_domain"`
	SubdomainOf   *Domain  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:id;foreignkey:SubdomainOfID"`
	Subdomain     []Domain `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:id;foreignkey:SubdomainOfID"`
	Host          []Host   `gorm:"many2many:Hosts_Domains;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (domain *Domain) HasService(serviceCommand map[string]models.CommandService) map[string]string {
	var targets = make(map[string]string)
	for _, host := range domain.Host {
		for _, port := range host.Ports {

			if _, ok := serviceCommand[port.Service]; ok {
				targets[port.Service] = domain.Domain + ":" + strconv.Itoa(port.Port)
			}
		}
	}
	return targets
}

func (domain *Domain) GetSubdomain() []string {
	var subdomains []string
	for _, subdomain := range domain.Subdomain {
		subdomains = append(subdomains, subdomain.Domain)
	}
	return subdomains
}

func (domain *Domain) GetTarget() string {
	return domain.Domain
}
