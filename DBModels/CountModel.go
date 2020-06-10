package DBModels

import (
	"apiGateway/Common/DB"
	"apiGateway/Utils/ComponentUtil"
)

type Count struct {
	Id      int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	ApiName string `gorm:"column:api_name;type:varchar(50);not null"`
	Time    string `gorm:"column:time;type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
}

// 设置User的表名为`Api`
func (p *Count) TableName() string {
	return "Count"
}

func (p *Count) GetCountByApiName(apiName string) error {
	if err := DB.DBConn().First(&p, "where api_name = ?", apiName).Error; err != nil {
		ComponentUtil.RuntimeLog().Error(err)
		return err
	}
	return nil
}

func (p *Count) SaveCount() bool {
	if err := DB.DBConn().Create(&p).Error; err != nil {
		ComponentUtil.RuntimeLog().Error(err)
		return false
	}
	return true
}

func (p *Count) GetCountListByData(start, end string) ([]Count, error) {
	var countList []Count
	if err := DB.DBConn().Order("time").Find(&countList, "time >= ? and time <= ? ", start, end).Error; err != nil {
		ComponentUtil.RuntimeLog().Error(err)
		return nil, err
	}
	return countList, nil
}
