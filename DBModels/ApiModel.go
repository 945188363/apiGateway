package DBModels

import (
	"apiGateway/DB"
	"log"
)

type Api struct {
	Id               int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	ApiName          string `gorm:"column:api_name;type:varchar(50);not null;unique_index"`
	ApiUrl           string `gorm:"column:api_url;type:varchar(255);not null"`
	ProtocolType     string `gorm:"column:protocol_type;type:varchar(50);not null;default:'http'"`
	BackendUrl       string `gorm:"column:backend_url;type:varchar(255);not null"`
	ApiMethod        string `gorm:"column:api_method;type:varchar(50);not null;default:'post'"`
	ApiTimeout       int    `gorm:"column:api_timeout;type:int(11);DEFAULT:3000"`
	ApiRetry         int    `gorm:"column:api_retry;type:int(11);DEFAULT:3"`
	ApiReturnType    string `gorm:"column:api_return_type;type:varchar(50);default:'RAW'"`
	ApiReturnContent string `gorm:"column:api_return_content;type:text"`
	ApiGroup         string `gorm:"column:api_group;type:varchar(50);not null;"`
}

// 设置User的表名为`Api`
func (p *Api) TableName() string {
	return "Api"
}

func (p *Api) GetApiByGroup(apiGroupName string) ([]Api, error) {
	var apiList []Api
	if err := DB.DBConn().Find(&apiList, "api_group = ?", apiGroupName).Error; err != nil {
		log.Fatal(err)
		return nil, err
	}
	return apiList, nil
}

func (p *Api) GetApi() error {
	if err := DB.DBConn().First(&p, p).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (p *Api) SaveApi() bool {
	// 已存在更新，否则创建
	exist := Api{}
	DB.DBConn().First(&exist, "api_name = ?", p.ApiName)
	if exist.Id != 0 {
		updateApi := Api{
			ApiName:          p.ApiName,
			ApiUrl:           p.ApiUrl,
			BackendUrl:       p.BackendUrl,
			ApiMethod:        p.ApiMethod,
			ApiTimeout:       p.ApiTimeout,
			ApiRetry:         p.ApiRetry,
			ApiReturnType:    p.ApiReturnType,
			ApiReturnContent: p.ApiReturnContent,
			ApiGroup:         p.ApiGroup,
		}
		if err := DB.DBConn().Model(&exist).Updates(&updateApi).Error; err != nil {
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

func (p *Api) DeleteApi() bool {
	if err := DB.DBConn().Where(p).Delete(&p).Error; err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (p *Api) GetApiList() ([]Api, error) {
	var apiList []Api
	if err := DB.DBConn().Find(&apiList).Error; err != nil {
		log.Fatal(err)
		return nil, err
	}
	return apiList, nil
}
