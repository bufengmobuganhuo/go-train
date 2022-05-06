package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID string `json:"id"`
	// 省略空值，“，”不可有空格
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price"`
}

type Order struct {
	// 首字母大写的才会被json看到
	// 在后面添加一个tag，可以设置json的键名
	ID         string `json:"id"`
	Items      []*OrderItem
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

// 转成json字符串
func marshal() string {
	o := Order{
		ID: "1234",
		Items: []*OrderItem{
			{
				ID:    "1",
				Name:  "name1",
				Price: 3,
			},
			{
				ID:    "2",
				Name:  "name2",
				Price: 3,
			},
		},
		Quantity:   3,
		TotalPrice: 30,
	}
	//
	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func unmarshal(s string) {
	// 结果会赋值给他，也可以使用map来接收 map := make(map[string]interface{})
	var o Order
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
}

func main() {
	s := marshal()

	// 方法一：使用map来接收
	m := make(map[string]interface{})
	json.Unmarshal([]byte(s), &m)
	// 获取map内的值
	fmt.Printf("%+v\n", m["Items"].([]interface{})[1].(map[string]interface{})["name"])

	// 方法二：使用一个匿名struct接收json
	sct := struct {
		ID         string `json:"id"`
		Items      []*OrderItem
		Quantity   int     `json:"quantity"`
		TotalPrice float64 `json:"total_price"`
	}{}
	json.Unmarshal([]byte(s), &sct)
	fmt.Printf("%+v\n", sct)
	unmarshal(s)
}
