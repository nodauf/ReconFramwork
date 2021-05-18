package database

import "gorm.io/gorm"

type Host struct {
	gorm.Model
	Hostname string   //`gorm:"uniqueindex:idx_hosts"`
	Address  string   `gorm:"uniqueindex:idx_hosts"`
	Ports    []Port   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` //`gorm:"many2many:Hosts_Ports;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Domain   []Domain `gorm:"many2many:Hosts_Domains;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	//PortDetail HostsPorts
}
