// Open Source: MIT License
// Author: Leon Ding <ding@ibyte.me>
// Date: 2022/2/28 - 7:07 下午 - UTC/GMT+08:00

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/auula/bottle"
)

func init() {
	//bottle.Open(bottle.DefaultOption)
	//
	//option := bottle.Option{
	//	Directory:       "./data",
	//	Enable:          true,
	//	Secret:          bottle.Secret,
	//	DataFileMaxSize: 1048576,
	//}

	bottle.Load("./config.yaml")

	bottle.SetIndexSize(1000)
}

type UserInfo struct {
	Name  string
	Age   uint8
	Skill []string
}

func main() {

	// PUT Data
	bottle.Put([]byte("foo"), []byte("66.6"))

	// 如果转成string那么就是字符串
	fmt.Println(bottle.Get([]byte("foo")).String())

	// 如果不存在默认值就是0
	fmt.Println(bottle.Get([]byte("foo")).Int())

	// 如果不成功就是false
	fmt.Println(bottle.Get([]byte("foo")).Bool())

	// 如果不成功就是0.0
	fmt.Println(bottle.Get([]byte("foo")).Float())

	user := UserInfo{
		Name:  "Leon Ding",
		Age:   22,
		Skill: []string{"Java", "Go", "Rust"},
	}

	// 通过Bson保存数据对象,并且设置超时时间为5秒
	bottle.Put([]byte("user"), bottle.Bson(&user), bottle.TTL(5))

	var u UserInfo

	// 通过Unwrap解析出结构体
	bottle.Get([]byte("user")).Unwrap(&u)

	data := bottle.Get([]byte("user"))

	if data.IsError() {
		fmt.Println(data.Err)
	} else {
		fmt.Println(data.Value)
	}

	// 打印取值
	fmt.Println(u)

	if err := bottle.Close(); err != nil {
		fmt.Println(err)
	}
}
