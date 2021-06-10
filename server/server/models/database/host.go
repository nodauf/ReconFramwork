package modelsDatabases

import (
	"errors"
	"strconv"

	modelsConfig "github.com/nodauf/ReconFramwork/server/server/models/config"
	"github.com/nodauf/ReconFramwork/utils"
	"gorm.io/gorm"
)

type Host struct {
	gorm.Model
	//Hostname string   //`gorm:"uniqueindex:idx_hosts"`
	Address string   `gorm:"uniqueindex:idx_hosts; NOT NULL"`
	Ports   []Port   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` //`gorm:"many2many:Hosts_Ports;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Domain  []Domain `gorm:"many2many:Hosts_Domains;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	//PortDetail HostsPorts
}

func (host *Host) HasService(serviceCommand map[string]modelsConfig.CommandService) map[string]string {
	var targets = make(map[string]string)
	for _, port := range host.Ports {

		if _, ok := serviceCommand[port.Service]; ok {
			targets[port.Service] = host.Address + ":" + strconv.Itoa(port.Port)
		}
	}
	return targets
}

func (host *Host) HasPort(port int) int {
	var portList []int
	for _, portHost := range host.Ports {
		portList = append(portList, portHost.Port)
	}
	if index, ok := utils.IntInSlice(port, portList); ok {
		return index
	}
	return -1
}

func (host *Host) AddPortComment(port int, portComment PortComment) ([]Host, error) {
	if index := host.HasPort(port); index != -1 {
		host.Ports[index].PortComment = append(host.Ports[index].PortComment, portComment)

	} else {
		return []Host{}, errors.New("The target " + host.Address + " has not the port " + strconv.Itoa(port))
	}
	return []Host{*host}, nil
}

func (host *Host) GetSubdomain() []string {
	// A host can't have a subdomain
	return []string{}
}

func (host *Host) GetDomain() []string {
	var domains []string
	if len(host.Domain) > 0 {
		for _, domain := range host.Domain {
			domains = append(domains, domain.Domain)
		}
	}
	return domains
}

func (host *Host) GetTarget() string {
	return host.Address
}
