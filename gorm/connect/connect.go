package connect

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (db *gorm.DB, err error) {
	dsn := "root:@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
