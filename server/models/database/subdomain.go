package database

import "gorm.io/gorm"

type Domain struct {
	gorm.Model
	Domain string `gorm:"uniqueindex:idx_domain"`
	Host   []Host `gorm:"many2many:Hosts_Domains;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
