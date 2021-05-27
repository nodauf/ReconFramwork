package modelsDatabases

import (
	"strconv"

	modelsConfig "github.com/nodauf/ReconFramwork/server/models/config"
	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
	Domain        string        `gorm:"unique"`
	SubdomainOfID *uint         //`gorm:"uniqueindex:idx_domain"`
	SubdomainOf   *Domain       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:id;foreignkey:SubdomainOfID"`
	Subdomain     []Domain      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:id;foreignkey:SubdomainOfID"`
	Host          []Host        `gorm:"many2many:Hosts_Domains;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PortComment   []PortComment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (domain *Domain) HasService(serviceCommand map[string]modelsConfig.CommandService) map[string]string {
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

func (domain *Domain) HasPort(port int) int {
	for _, host := range domain.Host {
		if index := host.HasPort(port); index != -1 {
			return index
		}
	}
	return -1
}

// We assume the domain is the same on each IP
func (domain *Domain) AddPortComment(port int, portComment PortComment) ([]Host, error) {
	var listHost []Host
	// Add the domain ID as it will not be added by gorm
	portComment.DomainID = domain.ID
	for _, host := range domain.Host {
		hostUpdated, err := host.AddPortComment(port, portComment)
		if err != nil {
			return []Host{}, err
		}
		listHost = append(listHost, hostUpdated...)
	}
	return listHost, nil
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
