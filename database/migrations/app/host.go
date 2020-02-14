package app

import "github.com/jinzhu/gorm"

// Host data model
type Host struct {
	gorm.Model
	Host string `sql:"type:varchar(255);"`
	Time string `sql:"type:varchar(255);"`
}
