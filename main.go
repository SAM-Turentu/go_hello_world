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
	//fmt.Println("hello wold")
	//fmt.Println("golang")
	//fmt.Println("good good study, day day up")
	//fmt.Printf("math pi = %f", PI)
	//fmt.Println(name)
	//fmt.Println("/************************/")
	//fmt.Println(identifier)
	//fmt.Println(age)
	//fmt.Printf("user:%s", user)
	//fmt.Println(_boo)
	//
	////region 初始化声明
	//// 变量必须先声明
	////只能存在于函数内，全局变量无法使用；
	////不使用变量也会报错；
	//user_age := 18
	//fmt.Println(&user_age) // 添加 &  获取内存中的地址 0xc00000e0d8  0xc000122058  0xc00000e0d8  0xc0000a6058
	////endregion
	//
	////a, b, c := 5, 7, "abc"  // 同时初始化声明
	////fmt.Println(a,b,c)
	//
	////iota 初始值0，每增加一个iota类型 自动加 1
	//const (
	//	e = iota
	//	f = iota
	//)
	//fmt.Println(e)
	//fmt.Println(f)
	//
	//const (
	//	h = iota // 0
	//	k        // 1
	//	l        // 2
	//	j = "as" // as
	//	m = iota // 4
	//)
	//fmt.Println(h)
	//fmt.Println(k)
	//fmt.Println(l)
	//fmt.Println(j)
	//fmt.Println(m)
	//type ByteSize float64
	//const (
	//	_           = iota
	//	kb ByteSize = 1 << (iota * 10)
	//	mb
	//	gb
	//)
	//fmt.Println(kb, mb, gb)
	//fmt.Println("*********************第一种遍历*********************")
	//sum := 2
	//for ; sum < 10; sum++ {
	//	fmt.Printf("sum=%d\n", sum)
	//}
	//
	//fmt.Println("*********************第二种遍历*********************")
	//sum = 2
	//for sum < 10 {
	//	fmt.Printf("sum=%d\n", sum)
	//	sum++
	//}

	strings := []string{"google", "runoob"}
	fmt.Printf("strings=%s\n", strings)
	for i, val := range strings {
		fmt.Printf("i=%d, val=%s\n", i, val)
	}
	data_a := read_data()
	fmt.Printf("data_a=%d\n", data_a)

	data_b := data()
	fmt.Printf("data_b=%d\n", data_b)
	fmt.Printf("data_b()=%d\n", data_b())

}

// 函数放在调用函数的前后没有关系
func read_data() int {
	return 10
}

func data() func() int {
	i := 1
	// 这是个闭包函数
	return func() int {
		i++
		return i
	}
}
