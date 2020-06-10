package LogModel

import (
	"apiGateway/Common/DB"
	"log"
)

type LogInfo struct {
	Id              int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	LogType         int    `gorm:"column:log_type;type:int;not null"`
	LogName         string `gorm:"column:log_name;type:varchar(50);not null;unique_index"`
	LogRecordStatus int    `gorm:"column:log_record_status;not null;type:tinyint(1);default:1"`
	LogAddress      string `gorm:"column:log_address;type:varchar(128);not null"`
	LogPeriod       string `gorm:"column:log_period;type:varchar(10);not null;default:'day'"`
	LogSavedTime    int    `gorm:"column:log_save_time;type:int;not null;default:7"`
	LogRecordField  string `gorm:"column:log_record_field;type:text"`
}

// 设置User的表名为`Api`
func (p *LogInfo) TableName() string {
	return "LogInfo"
}

func (p *LogInfo) GetLogInfo() error {
	if err := DB.DBConn().First(&p, p).Error; err != nil {
		// ComponentUtil.RuntimeLog().Error(err)
		log.Fatal(err)
		return err
	}
	return nil
}

func (p *LogInfo) GetLogInfoByType() error {
	if err := DB.DBConn().Find(&p, "log_Type = ?", p.LogType).Error; err != nil {
		// ComponentUtil.RuntimeLog().Error(err)
		log.Fatal(err)
		return err
	}
	return nil
}

func (p *LogInfo) SaveLogInfo() bool {
	// 已存在更新，否则创建
	exist := LogInfo{}
	DB.DBConn().First(&exist, "log_Type = ?", p.LogType)
	if exist.Id != 0 {
		if err := DB.DBConn().Model(&exist).Updates(&p).Error; err != nil {
			// ComponentUtil.RuntimeLog().Error(err)
			log.Fatal(err)
			return false
		}
		return true
	}
	if err := DB.DBConn().Create(&p).Error; err != nil {
		// ComponentUtil.RuntimeLog().Error(err)
		log.Fatal(err)
		return false
	}
	return true
}

func (p *LogInfo) DeleteLogInfo() bool {
	if err := DB.DBConn().Where(p).Delete(&p).Error; err != nil {
		// ComponentUtil.RuntimeLog().Error(err)
		log.Fatal(err)
		return false
	}
	return true
}

func (p *LogInfo) GetLogInfoList() ([]LogInfo, error) {
	var LogInfoList []LogInfo
	if err := DB.DBConn().Find(&LogInfoList).Error; err != nil {
		// ComponentUtil.RuntimeLog().Error(err)
		log.Fatal(err)
		return nil, err
	}
	return LogInfoList, nil
}
