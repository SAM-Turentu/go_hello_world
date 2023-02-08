package main

import (
	"fmt"
	"math"
)

// 常量定义
const PI = math.Pi

// 全局变量的声明和赋值
var name = "sam"

// 定义多个变量, 默认值：0
var identifier, age int

// 定义字符串变量，默认值：空字符串
var user string

// 定义布尔变量，默认值：false
var _boo bool

// user_age := 1  // 只能存在于函数内，全局变量无法使用

// 其他类型默认值：nil

// 一般类型声明
type newType int

// 结构的声明
type gopher struct{}

// 接口的声明
type golang interface{}

func main() {
	fmt.Println("hello wold")
	fmt.Println("golang")
	fmt.Println("good good study, day day up")
	fmt.Printf("math pi = %f", PI)
	fmt.Println(name)
	fmt.Println("/************************/")
	fmt.Println(identifier)
	fmt.Println(age)
	fmt.Printf("user:%s", user)
	fmt.Println(_boo)

	user_age := 18 // 只能存在于函数内，全局变量无法使用

	fmt.Println(user_age)

}
