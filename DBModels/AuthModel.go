package DBModels

import (
	"apiGateway/DB"
)

type Auth struct {
	ID       int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// 设置User的表名为`Api`
func (p *Auth) TableName() string {
	return "Auth"
}

func CheckAuth(username, password string) bool {
	var auth Auth
	DB.DBConn().Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}
