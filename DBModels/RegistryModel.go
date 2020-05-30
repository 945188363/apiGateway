package DBModels

import (
	"apiGateway/Common/DB"
	"log"
)

type Registry struct {
	Id           int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name         string `gorm:"column:name;type:varchar(50);not null;unique_index"`
	RegistryType string `gorm:"column:registry_type;type:varchar(50);not null;"`
	Addr         string `gorm:"column:addr;type:varchar(50);not null"`
}

// 设置User的表名为`Api`
func (p *Registry) TableName() string {
	return "Registry"
}

func (p *Registry) GetRegistry() error {
	if err := DB.DBConn().First(&p, p).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (p *Registry) SaveRegistry() bool {
	// 已存在更新，否则创建
	exist := Registry{}
	DB.DBConn().First(&exist, "name = ?", p.Name)
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

func (p *Registry) DeleteRegistry() bool {
	if err := DB.DBConn().Where(p).Delete(&p).Error; err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (p *Registry) GetRegistryList() ([]Registry, error) {
	var RegistryList []Registry
	if err := DB.DBConn().Find(&RegistryList).Error; err != nil {
		log.Fatal(err)
		return nil, err
	}
	return RegistryList, nil
}
