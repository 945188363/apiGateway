package DB

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@tcp(localhost:3306)/fileserver?charset=utf8")
	if err != nil {
		log.Fatal(err)

	}
	db.DB().SetMaxOpenConns(1000)
	db.DB().SetMaxIdleConns(10)
	// 全局禁用表名复数
	db.SingularTable(true)
}

func DBConn() *gorm.DB {
	return db
}
