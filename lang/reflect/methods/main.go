package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)
	for i := 0; i < t.NumMethod(); i++ {
		// v.Method(i)：返回一个reflect.Value，代表一个已绑定接收者的方法
		methType := v.Method(i).Type()
		// t.Medhod(i)：返回一个reflect.Method实例，描述了方法的名称和类型
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name, strings.TrimPrefix(methType.String(), "func"))
	}
}

func main()  {
	print(time.Hour)
}
