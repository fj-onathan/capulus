package models

import (
	. "capulus/database"

	"github.com/jinzhu/gorm"
)

type Host struct {
	ID   int    `json:"id"`
	Host string `json:"host"`
	Time string `json:"time"`
}

// GetHosts gets a list of hosts
func GetHosts() ([]Host, error) {
	var (
		hosts []Host
		err   error
	)

	err = Db.Order("id desc").Find(&hosts).Error

	// build
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return hosts, nil
}

// AddHost Add a new host
func AddHost(host, time string) error {
	hostNew := Host{
		Host: host,
		Time: time,
	}
	if err := Db.Create(&hostNew).Error; err != nil {
		return err
	}

	return nil
}

// DeleteHost delete a host
func DeleteHost(id int) error {
	if err := Db.Where("id = ?", id).Delete(&Host{}).Error; err != nil {
		return err
	}
	return nil
}
