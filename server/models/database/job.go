package database

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	TaskUUID  string
	Processed bool `gorm:"default:false"`
	HostID    uint `gorm:"default:null"`
	Host      Host `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Parser    string
}