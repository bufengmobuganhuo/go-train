package main

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"mengyu.com/gotrain/gorm/connect"
	"mengyu.com/gotrain/gorm/model"
)

func main() {
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}
	// 插入单条记录
	create(db)
	// 批量插入
	createInBatches(db)
}

func create(db *gorm.DB) {
	loc, _ := time.LoadLocation("Local")
	birthday, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-08-05 12:00:00", loc)
	user := User{Name: "Jinzhu", Age: 18, Birthday: &birthday}
	// 自动创建表格
	db.AutoMigrate(user)
	result := db.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Printf("affect rows: %d, id: %d\n", result.RowsAffected, user.ID)

	user1 := model.User{Name: "Jinzhu", Age: 18, Birthday: &birthday}
	// 此时user.birthday字段中不会被插入到数据库中
	result = db.Select("Name", "Age", "CreateAt").Create(&user1)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Printf("affect rows: %d, id: %d\n", result.RowsAffected, user1.ID)

	user2 := model.User{Name: "Jinzhu", Age: 18, Birthday: &birthday}
	// 与select相反，插入除了Name, Age, CreatedAt字段的其他字段
	result = db.Omit("Name", "Age", "CreatedAt").Create(&user2)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Printf("affect rows: %d, id: %d\n", result.RowsAffected, user2.ID)
}

func createInBatches(db *gorm.DB) {
	var users = []model.User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	// db.CreateInBatches(users, 100)：指定每次插入100条数据
	// gorm会生成一条单独SQL来插入所有数据，并回填主键的值
	db.Create(&users)

	for _, user := range users {
		fmt.Printf("id: %d\n", user.ID)
	}
}
