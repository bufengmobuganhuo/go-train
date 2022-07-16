package main

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"mengyu.com/gotrain/gorm/connect"
	"mengyu.com/gotrain/gorm/model"
)

func main() {
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}
	delete(db)
}

func delete(db *gorm.DB) {
	user := model.User{
		ID: 16,
	}
	// 根据主键删除
	// DELETE FROM users WHERE id = 16
	db.Delete(&user)
	// DELETE FROM users WHERE id = 16
	db.Delete(&model.User{}, 17)
	// DELETE FROM users WHERE id in (16, 17)
	db.Delete(&model.User{}, []int{16, 17})

	// 带条件的删除
	// DELETE FROM users WHERE name LIKE % AND id = 16
	db.Where("name LIKE ?", "%").Delete(&user)

	// 如果在没有任何条件的情况下执行批量删除，GORM 不会执行该操作，并返回 ErrMissingWhereClause 错误
	res := db.Delete(&model.User{})
	if errors.Is(res.Error, gorm.ErrMissingWhereClause) {
		fmt.Printf("cant delete all:%v", res.Error)
	}

	// 软删除：如果模型中包含了一个类型为 gorm.DeleteAt字段，它不会真正删除数据，而是将delete_at字段设置成当前时间，并且无法通过普通方法找到这个记录
	// 如果要找到这个记录，可以使用：db.Unscoped().Where("age = 20").Find(&users)
	type email struct {
		ID        int
		emial     string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeleteAt  gorm.DeletedAt
	}
	emi := email{
		emial: "123@qq.com",
	}
	db.AutoMigrate(&email{})
	db.Create(&emi)
	// UPDATE `emails` SET `delete_at`='2022-07-13 20:12:01.309' WHERE `emails`.`id` = 1 AND `emails`.`delete_at` IS NULL
	db.Delete(&emi)
	// 普通查找无法找到记录
	res1 := db.First(&emi)
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("cant find record:%v", res1.Error)
	}
	var emi2 = email{}
	// 可以找到记录
	db.Unscoped().First(&emi2)
	fmt.Printf("found record: %d", emi2.ID)
	// 彻底删除：DELETE FROM `emails` WHERE `emails`.`id` = 1
	db.Unscoped().Delete(&email{ID: 1})

	// 也可以使用1/0作为软删除的flag
	type address struct {
		ID      int
		address string
		IsDel   soft_delete.DeletedAt `gorm:"softDelete:flag"`
	}
	addr := address{
		address: "xxxxxx",
	}
	db.AutoMigrate(&address{})
	db.Create(&addr)
	// UPDATE `addresses` SET `is_del`=1 WHERE `addresses`.`id` = 1 AND `addresses`.`is_del` = 0
	db.Delete(&addr)
	// SELECT * FROM `addresses` WHERE `addresses`.`is_del` = 0 AND `addresses`.`id` = 1 ORDER BY `addresses`.`id` LIMIT 1
	res3 := db.First(&addr)
	if errors.Is(res3.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("cant find record:%v", res3.Error)
	}
}
