package main

import (
	"fmt"
	"strconv"
)

func Second() {
	fmt.Println("************************Second Learn Golang************************")
	fmt.Println("Second Learn Golang")

	//region 浮点数简写和科学计数法
	//num := .234 // python不能这么写
	//fmt.Printf("num = %f\n", num)
	//num1 := 1. // python不能这么写
	//fmt.Printf("num1 = %f\n", num1)
	//num2 := 12033443e-5 // 可以使用科学计算法，python中不可以
	//fmt.Printf("num2 = %F\n", num2)
	//
	//const Avogadro = 6.02214129e23 // 阿伏伽德罗常数
	//const Planck = 6.62606957e-34  // 普朗克常数
	//fmt.Println(Avogadro)          // 6.02214129e23
	//fmt.Println(Planck)            // 6.62606957e-34
	//endregion

	//region if 中 一个值在数组中的判断
	arr := []int{1, 2, 3, 4}
	//if 1 in arr{} // 在python等语言中可以使用 in，go中不能使用
	for _, v := range arr {
		if v == 1 {

		}
	}
	_boo, _ := strconv.ParseBool("0")
	if _boo {
		fmt.Println(_boo)
	}

	//endregion

	//region 整型范围
	//fmt.Printf("int max:%d, min:%d\n", math.MaxInt, math.MinInt)       // int = int64
	//fmt.Printf("int8 max:%d, min:%d\n", math.MaxInt8, math.MinInt8)    // -128, 127
	//fmt.Printf("int16 max:%d, min:%d\n", math.MaxInt16, math.MinInt16) // -32768, 32767
	//fmt.Printf("int32 max:%d, min:%d\n", math.MaxInt32, math.MinInt32) // -2147483648, 2147483647
	//fmt.Printf("int64 max:%d, min:%d\n", math.MaxInt64, math.MinInt64) // -9223372036854775808, 9223372036854775807

	//endregion

	//region 指针
	//var model = flag.String("params", "value", "help text")
	//flag.Parse()
	//fmt.Println(*model) // 打印配置参数（program arguments）  test_params_value_flagString

	//endregion

	//region 类型的别名
	//type newInt = int // newInt 即 int
	//var a newInt = 10
	//fmt.Println(a)

	//endregion

	//region 切片
	//arr1 := []int {1,2,3,4}
	//fmt.Println(arr1[len(arr1) - 1])  // error fmt.Println(arr1[len(arr1)])
	//
	//var a = []int{1,2,3}
	//a = append([]int{0}, a...) // 在开头添加1个元素
	//fmt.Println(a) // [0,1,2,3]
	//a = append([]int{-3,-2,-1}, a...) // 在开头添加1个切片
	//fmt.Println(a) // [-3 -2 -1 0 1 2 3]
	// todo 切片链式操作
	//b := make([]int, 10)
	//fmt.Println(b)  // [0,0,0,0,0,0,0,0,0,0]
	//b = append(b[:2], append([]int{10}, b[2:]...)...) // 在第2个位置插入10
	//fmt.Println(b)  // [0,0,10,0,0,  0,0,0,0,0]
	//b = append(b[:3], append([]int{1,2,3}, b[3:]...)...) // 在第3个位置插入切片
	//fmt.Println(b)  // [0,0,10,1,2,3,0,0,0,  0,0,0,0,0]
	//
	//b = make([]int, 10)
	//fmt.Println(b)  // [0,0,0,0,0,  0,0,0,0,0]
	//b = append(b[:4], append([]int{5}, b[7:]...)...) // 在第3个位置插入切片
	//fmt.Println(b) // [0 0 0 0 5 0 0 0]
	//endregion

	//region

	//endregion

	//region

	//endregion

	fmt.Println("************************Second END************************")
}

// 可以写一个in方法
func in(p int, arr []int) bool {
	result := false
	for _, v := range arr {
		if v == p {
			result = true
		}
	}
	return result
}

//func in(s string, arr []int) bool {
//
//}
