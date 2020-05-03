package DBModels

import (
	"apiGateway/DB"
	"log"
)

type MonitorInfo struct {
	Id            int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	MonitorType   string `gorm:"column:monitor_type;type:varchar(50);not null;unique_index"`
	MonitorStatus int    `gorm:"column:monitor_status;type:tinyint(1);not null;default:0"`
	MonitorConfig string `gorm:"column:monitor_config;type:varchar(255)"`
}

// 设置User的表名为`Api`
func (p *MonitorInfo) TableName() string {
	return "MonitorInfo"
}

func (p *MonitorInfo) GetMonitorInfo() error {
	if err := DB.DBConn().Find(&p, "monitor_type = ?", p.MonitorType).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (p *MonitorInfo) SaveMonitorInfo() bool {
	// 已存在更新，否则创建
	exist := MonitorInfo{}
	DB.DBConn().Find(&exist, "monitor_type=?", p.MonitorType)
	if exist.Id != 0 {
		if err := DB.DBConn().Model(&exist).Updates(&p).Error; err != nil {
			log.Fatal(err)
			return false
		}
		return true
	}
	if err := DB.DBConn().Create(&p).Error; err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (p *MonitorInfo) DeleteMonitorInfo() bool {
	if err := DB.DBConn().Where("monitor_type = ?", p.MonitorType).Delete(&p).Error; err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (p *MonitorInfo) GetMonitorInfoList() ([]MonitorInfo, error) {
	var monitorInfoList []MonitorInfo
	if err := DB.DBConn().Find(&monitorInfoList).Error; err != nil {
		log.Fatal(err)
		return nil, err
	}
	return monitorInfoList, nil
}
