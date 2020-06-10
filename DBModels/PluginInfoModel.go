package DBModels

import (
	"apiGateway/Common/DB"
	"apiGateway/Utils/ComponentUtil"
)

type PluginInfo struct {
	Id           int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	PluginName   string `gorm:"column:plugin_name;type:varchar(50);not null;unique_index"`
	PluginStatus int    `gorm:"column:plugin_status;type:tinyint(1);not null;default:0"`
	Description  string `gorm:"column:description;type:varchar(255)"`
	PluginType   string `gorm:"column:plugin_type;type:varchar(20);not null"`
}

// 设置User的表名为`Api`
func (p *PluginInfo) TableName() string {
	return "PluginInfo"
}

func (p *PluginInfo) GetPluginInfo() error {
	if err := DB.DBConn().First(&p, p).Error; err != nil {
		ComponentUtil.RuntimeLog().Error(err)
		return err
	}
	return nil
}

func (p *PluginInfo) SavePluginInfo() bool {
	// 已存在更新，否则创建
	exist := PluginInfo{}
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

func (p *PluginInfo) DeletePluginInfo() bool {
	if err := DB.DBConn().Where(p).Delete(&p).Error; err != nil {
		ComponentUtil.RuntimeLog().Error(err)
		return false
	}
	return true
}

func (p *PluginInfo) GetPluginInfoList() ([]PluginInfo, error) {
	var pluginInfoList []PluginInfo
	if err := DB.DBConn().Find(&pluginInfoList).Error; err != nil {
		ComponentUtil.RuntimeLog().Error(err)
		return nil, err
	}
	return pluginInfoList, nil
}
