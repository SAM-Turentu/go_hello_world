package main

import (
	"fmt"
	"math"
)

// 常量定义
const PI = math.Pi

// 全局变量的声明和赋值
var name = "gopher"

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
	fmt.Println()

}
