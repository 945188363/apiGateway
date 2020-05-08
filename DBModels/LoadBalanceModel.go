package DBModels

import (
	"apiGateway/DB"
	"log"
)

type LoadBalance struct {
	Id           int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name         int    `gorm:"column:name;type:varchar(50);not null;unique_index"`
	RegistryName string `gorm:"column:registry_name;type:varchar(50);not null;"`
	Strategy     string `gorm:"column:strategy;type:varchar(50);not null;default:'random'"`
	ServiceName  string `gorm:"column:service_name;type:text"`
}

// 设置User的表名为`Api`
func (p *LoadBalance) TableName() string {
	return "LoadBalance"
}

func (p *LoadBalance) GetLoadBalance() error {
	if err := DB.DBConn().First(&p, p).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (p *LoadBalance) GetLoadBalanceByServiceName(serviceName string) error {
	if err := DB.DBConn().First(&p, "log_period like ?", serviceName).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (p *LoadBalance) SaveLoadBalance() bool {
	// 已存在更新，否则创建
	exist := LogInfo{}
	DB.DBConn().First(&exist, p)
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

func (p *LoadBalance) DeleteLoadBalance() bool {
	if err := DB.DBConn().Where(p).Delete(&p).Error; err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (p *LoadBalance) GetLoadBalanceList() ([]LoadBalance, error) {
	var LoadBalanceList []LoadBalance
	if err := DB.DBConn().Find(&LoadBalanceList).Error; err != nil {
		log.Fatal(err)
		return nil, err
	}
	return LoadBalanceList, nil
}