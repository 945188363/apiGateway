package DBModels

import (
	"apiGateway/Common/DB"
	"apiGateway/Utils/ComponentUtil"
)

type ApiGroup struct {
	Id           int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	ApiGroupName string `gorm:"column:api_group_name;type:varchar(50);not null;unique_index"`
	Description  string `gorm:"column:description;type:varchar(255)"`
}

// 设置User的表名为`Api`
func (p *ApiGroup) TableName() string {
	return "ApiGroup"
}

func (p *ApiGroup) GetApiGroup() error {
	if err := DB.DBConn().First(&p, p).Error; err != nil {
		ComponentUtil.RuntimeLog().Error(err)
		return err
	}
	return nil
}

func (p *ApiGroup) GetApiGroupListByGroupName() ([]ApiGroup, error) {
	var apiGroupList []ApiGroup

	if err := DB.DBConn().Find(&apiGroupList, "api_group_name like ? ", "%"+p.ApiGroupName+"%").Error; err != nil {
		ComponentUtil.RuntimeLog().Error(err)
		return nil, err
	}
	return apiGroupList, nil
}

func (p *ApiGroup) SaveApiGroup() bool {
	// 已存在更新，否则创建
	exist := ApiGroup{}
	DB.DBConn().First(&exist, "api_group_name = ?", p.ApiGroupName)
	if exist.Id != 0 {
		updateApi := ApiGroup{
			ApiGroupName: p.ApiGroupName,
			Description:  p.Description,
		}
		if err := DB.DBConn().Model(&exist).Updates(&updateApi).Error; err != nil {
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

func (p *ApiGroup) DeleteApiGroup() bool {
	if err := DB.DBConn().Where(p).Delete(&p).Error; err != nil {
		ComponentUtil.RuntimeLog().Error(err)
		return false
	}
	return true
}

func (p *ApiGroup) GetApiGroupList() ([]ApiGroup, error) {
	var apiGroupList []ApiGroup
	if err := DB.DBConn().Find(&apiGroupList).Error; err != nil {
		ComponentUtil.RuntimeLog().Error(err)
		return nil, err
	}
	return apiGroupList, nil
}
