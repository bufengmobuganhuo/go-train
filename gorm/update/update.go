package main

import (
	"gorm.io/gorm"
	"mengyu.com/gotrain/gorm/connect"
	"mengyu.com/gotrain/gorm/model"
)

func main() {
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}
	update(db)
}

func update(db *gorm.DB) {
	// 更新单个列
	// UPDATE users SET name='hello', updated_at='now' WHERE age=18;
	db.Model(&model.User{}).Where("age = ?", 18).Update("name", "hello")

	user := model.User{
		ID:   16,
		Name: "xxx",
	}
	// 更新单个列
	// 如果user里的id存在，则会根据ID更新
	// UPDATE `users` SET `name`='hello',`updated_at`='2022-07-12 20:29:52.804' WHERE age = 18 AND `id` = 16
	db.Model(&user).Where("age = ?", 18).Update("name", "hello")

	// 更新多列, 也可以使用map传参
	// UPDATE `users` SET `name`='xxx',`age`=18,`updated_at`='2022-07-12 20:31:31.636' WHERE `id` = 16
	db.Model(&user).Updates(model.User{Name: "xxx", Age: 18})

	// 更新选定的字段
	// 使用 Map 进行 Select
	// UPDATE users SET name='hello' WHERE id=16;
	db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})

	// UPDATE users SET age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=16;
	db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
}

func save(db *gorm.DB) {
	var user model.User
	db.First(&user)
	user.Name = "jinzhu 2"
	user.Age = 100
	// save方法会根据ID更新所有字段，包括为零值的字段
	// UPDATE `users` SET `id`=16,`name`='jinzhu 2',`email`=NULL,`age`=100,`birthday`='2022-08-05 12:00:00',`member_number`=NULL,`activated_at`=NULL,`created_at`='2022-07-07 20:27:10.168',`updated_at`='2022-07-12 20:26:13.7' WHERE `id` = 16
	db.Save(user)
}
