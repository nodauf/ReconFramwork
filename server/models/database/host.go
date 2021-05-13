package database

import "gorm.io/gorm"

type Host struct {
	gorm.Model
	Hostname string `gorm:"uniqueindex:idx_hosts"`
	Address  string `gorm:"uniqueindex:idx_hosts"`
	Ports    []Port //`gorm:"many2many:Hosts_Ports;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	//PortDetail HostsPorts
}
