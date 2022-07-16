package main

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"mengyu.com/gotrain/gorm/connect"
)

type User struct {
	gorm.Model
	CreditCardID uint
	CreditCard   CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
}

func crud(db *gorm.DB) {
	user := User{
		CreditCard: CreditCard{
			Number: "12344567",
		},
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&CreditCard{})
	db.Create(&user)

	var users []User
	db.Preload("CreditCard").Find(&users)
	db.Joins("CreditCard").Find(&users)
	fmt.Println()
}

func preload(db *gorm.DB) {
	var users1 []User
	// 一个model中可能包含多个其他的model，此时可以使用如下方法：
	db.Preload(clause.Associations).Find(&users1)
	// 但是对于嵌套的model，无法使用上述方法查询出来，可以使用如下方法
	// db.Preload("Orders.OrderItems.Product").Preload(clause.Associations).Find(&users1)

	// preload时指定条件
	// SELECT * FROM `credit_cards` WHERE `credit_cards`.`id` = 1 AND number = '12345' AND `credit_cards`.`deleted_at` IS NULL
	// SELECT * FROM `users` WHERE id = '1' AND `users`.`deleted_at` IS NULL
	db.Preload("CreditCard", "number = ?", "12345").Where("id = ?", "1").Find(&users1)
}

func main() {
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}
	// crud(db)
	preload(db)
}
