package DBModels

type ProdDB struct {
	Id   int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name string `gorm:"column:name;type:varchar(50);not null;unique_index"`
}

// 设置User的表名为`profiles`

func (p ProdDB) TableName() string {
	if p.Name == "admin" {
		return "ProdDB_Admin"
	} else {
		return "ProdDB"
	}
}
