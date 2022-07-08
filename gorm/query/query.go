package main

import (
	"errors"

	"gorm.io/gorm"
	"mengyu.com/gotrain/gorm/connect"
	"mengyu.com/gotrain/gorm/model"
)

func main() {
	db, err := connect.Connect()
	if err != nil {
		panic(err)
	}
	// selectOne(db)
	// selectById(db)
	selectByConditions(db)
}

func selectOne(db *gorm.DB) {
	user := model.User{}
	// 主键升序，找到第一条记录
	// select * from users order by id limit 1
	res := db.First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		// 查询不到记录
		panic(res.Error)
	}
	user.Print()

	// 获取第一条记录
	// select * from users limit 1
	user1 := model.User{}
	res1 := db.Take(&user1)
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		panic(res1.Error)
	}
	user1.Print()

	// 获取最后一条记录
	// select *from users order by id desc limit 1
	user2 := model.User{}
	res2 := db.Last(&user2)
	if errors.Is(res2.Error, gorm.ErrRecordNotFound) {
		panic(res2.Error)
	}
	user2.Print()

	// 如果model里没有主键，上述方法默认以第一个字段排序
	// 无法工作，因为map里没有第一个字段
	result := map[string]interface{}{}
	db.Table("users").First(&result)
}

func selectById(db *gorm.DB) {
	user := model.User{}
	// select *from users where id = 16
	db.First(&user, "16")
	user.Print()

	user1 := model.User{}
	// 如下方法会进行转义，则可以防止SQL注入
	// select *from users where id = '16'
	db.First(&user1, "id = ?", 16)
	user1.Print()

	// 也可以使用model来传参
	var res model.User
	db.Model(model.User{ID: 16}).First(&res)
	res.Print()

	var res2 []model.User
	// select *from users where id in (16, 17 ,10)
	db.Find(&res2, []int{16, 17, 10})
	for _, usr := range res2 {
		usr.Print()
	}
	//res2.Print()
}

func selectByConditions(db *gorm.DB) {
	// 找到第一个match条件的记录
	// select *from users where name = 'jinzhu' order by id limit 1
	var user model.User
	db.Where("name = ?", "jinzhu").First(&user)
	user.Print()

	var users []model.User
	// 获取到所有匹配的记录
	// SELECT * FROM users WHERE name <> 'jinzhu';
	db.Where("name <> ?", "jinzhu").Find(&users)

	var users1 []model.User
	// select *from users where name in ('')
	db.Where("name in ?", []string{"jinzhu", "jinzhu2"}).Find(&users1)
	model.Print(users1)

	var users2 []model.User
	// select *from users where name like '%jin%'
	db.Where("name like ?", "%jin%").Find(&users2)
	model.Print(users2)

	var users3 []model.User
	db.Where("name = ? and age >= ?", "jinzhu", "16").Find(&users3)
	model.Print(users3)

	var user1 model.User
	// 使用struct传参，默认是AND条件，如果struct的字段值是零值，则不会作为查询条件，如果想将零值作为查询条件，可以使用Map传参
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;
	db.Where(&model.User{Name: "jinzhu", Age: 20}).First(&user1)
	user1.Print()

	// 

	var users4 []model.User
	// 使用Map传参，默认是AND的关系
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users4)
	for _, usr := range users4 {
		usr.Print()
	}

	var users5 []model.User
	// 默认使用IN查询id字段
	// SELECT * FROM users WHERE id IN (20, 21, 22);
	db.Where([]int64{16, 17}).Find(&users5)
	for _, usr := range users5 {
		usr.Print()
	}
}
