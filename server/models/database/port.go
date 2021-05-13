package database

import "gorm.io/gorm"

type Port struct {
	gorm.Model
	Port    int `gorm:"uniqueindex:idx_port"`
	Service string
	Version string
	HostID  uint `gorm:"uniqueindex:idx_port"`
	//Comment string `gorm:"primaryKey"`
	//Hosts       []Host        `gorm:"many2many:Hosts_Ports;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` //foreign_key:Port"`
	PortComment []PortComment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //`gorm:"many2many:Port_PortComment"`
}

/*type PortDetail struct {
	gorm.Model
	//Port    int    `gorm:"primaryKey"`
	Service string `gorm:"index:idx_portdetail"`
	Tool    string `gorm:"index:idx_portdetail"`
	Version string
	Comment string
	HostID  int  `gorm:"type:bigint unsigned;index:idx_portdetail"`
	Host    Host `gorm:"foreignKey:HostID"`
	//Port    Port `gorm:"many2many:Hosts_Ports"`
}*/

type PortComment struct {
	gorm.Model
	Tool    string `gorm:"uniqueindex:idx_portcomment"`
	Comment string `gorm:"uniqueindex:idx_portcomment"`
	PortID  uint   `gorm:"uniqueindex:idx_portcomment"` //`gorm:"type:bigint unsigned;"`
	//Port    Port `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
