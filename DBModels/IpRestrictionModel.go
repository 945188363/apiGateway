package DBModels

import (
	"apiGateway/Common/DB"
	"apiGateway/Utils/ComponentUtil"
)

type IpRestriction struct {
	Id           int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name         string `gorm:"column:name;type:varchar(50);not null;unique_index"`
	Global       int    `gorm:"column:global;type:tinyint(1);not null;default:0"`
	Status       int    `gorm:"column:status;type:tinyint(1);not null;default:1"`
	IpWhiteList  string `gorm:"column:ip_white_list;type:text"`
	IpBlackList  string `gorm:"column:ip_black_list;type:text"`
	ApiList      string `gorm:"column:api_list;type:text"`
	ApiGroupList string `gorm:"column:api_group_list;type:text"`
}

// 设置User的表名为`Api`
func (p *IpRestriction) TableName() string {
	return "IpRestriction"
}

func (p *IpRestriction) GetIpRestrictionByApi(api string) error {
	if err := DB.DBConn().First(&p, "api_list like ? and global = ?", api, 0).Error; err != nil {
		ComponentUtil.RuntimeLog().Error(err)
		return err
	}
	return nil
}

func (p *IpRestriction) GetIpRestrictionByApiGroup(apiGroup string) error {
	if err := DB.DBConn().First(&p, "api_group LIKE ? and  global = ?", apiGroup, 0).Error; err != nil {
		ComponentUtil.RuntimeLog().Error(err)
		return err
	}
	return nil
}

func (p *IpRestriction) GetGlobalIpRestriction() error {
	if err := DB.DBConn().First(&p, "global = ?", 1).Error; err != nil {
		ComponentUtil.RuntimeLog().Error(err)
		return err
	}
	return nil
}

func (p *IpRestriction) SaveIpRestriction() bool {
	// 已存在更新，否则创建
	exist := IpRestriction{}
	DB.DBConn().First(&exist, p)
	if exist.Id != 0 {
		if err := DB.DBConn().Model(&exist).Updates(&p).Error; err != nil {
			ComponentUtil.RuntimeLog().Error(err)
			return false
		}
		return true
	}
	if err := DB.DBConn().Create(&p).Error; err != nil {
		ComponentUtil.RuntimeLog().Error(err)
		return false
	}
	return true
}

func (p *IpRestriction) DeleteIpRestriction() bool {
	if err := DB.DBConn().Where(p).Delete(&p).Error; err != nil {
		ComponentUtil.RuntimeLog().Error(err)
		return false
	}
	return true
}

func (p *IpRestriction) GetIpRestrictionList() ([]IpRestriction, error) {
	var ipRestrictionList []IpRestriction
	if err := DB.DBConn().Find(&ipRestrictionList).Error; err != nil {
		ComponentUtil.RuntimeLog().Error(err)
		return nil, err
	}
	return ipRestrictionList, nil
}
