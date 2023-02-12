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

	//strings := []string{"google", "runoob"}
	//fmt.Printf("strings=%s\n", strings)
	//for i, val := range strings {
	//	fmt.Printf("i=%d, val=%s\n", i, val)
	//}
	//data_a := read_data()
	//fmt.Printf("data_a=%d\n", data_a)
	//
	//data_b := data()
	//fmt.Printf("data_b=%d\n", data_b)
	//fmt.Printf("data_b()=%d\n", data_b())

	//region 调用其他文件方法，注意配置问题，不然无法运行调试
	//a := PrintHello()
	//fmt.Println(a)
	//fmt.Println("***************************")
	//hello()
	//endregion

	//region 矩形结构体和方法
	//r := Rectangle{}
	//fmt.Printf("矩形的长：%f\n矩形的宽：%f\n矩形的面积：%f\n", r.length, r.width, r.getArea())
	//fmt.Println("******************************")
	//r.length = 20
	//r.width = 10
	//fmt.Printf("矩形的长：%f\n矩形的宽：%f\n矩形的面积：%f\n", r.length, r.width, r.getArea())
	//fmt.Println("******************************")
	//r.setLength(10)
	//r.setWidth(5)
	//fmt.Printf("矩形的长：%f\n矩形的宽：%f\n矩形的面积：%f\n", r.length, r.width, r.getArea())
	//endregion

	//region 圆的结构体
	//fmt.Println("**************函数方法**************")
	//circle := Circle{}
	//circle.radius = 10
	//area := circle.getArea()
	//fmt.Printf("圆的面积：%f\n", area)
	//fmt.Printf("圆的周长：%f\n", circle.getLen())
	//endregion

	//region 一维数组
	//fmt.Println("**************一维数组**************")
	//var data_list []float64
	//fmt.Println(data_list)
	//fmt.Println(len(data_list))
	////向数组中添加元素
	//data_list = append(data_list, 32.4)
	//fmt.Println(data_list)
	//fmt.Println(len(data_list))
	//
	////一维数组遍历
	//strings := []string{"google", "runoob"}
	//fmt.Printf("strings=%s\n", strings)
	//for i, val := range strings {
	//	fmt.Printf("i=%d, val=%s\n", i, val)
	//}
	//
	//删除一维数组中的元素

	//修改数组元素
	//arr_1 := [] int {1,2,3,4}
	//arr_2 := change_array(arr_1)
	//fmt.Println(arr_1)
	//fmt.Println(arr_2)
	////数组和切片问题，数组不可修改，切片可以修改，当作为函数参数传递时：数组是作为值传递，切片是作为类似于 指针传递（切片包含对底层数组内容的引用）
	////如果使用数组作为值传递时，调用方法使用了指针，可以修改数组中的值
	//var nums=[3]int{1,2,3}
	//change(nums)   //nums并未被改变
	//fmt.Println(nums[0])
	//fmt.Println(nums)
	//
	//fmt.Println("*****************一维数组：方法****************")
	//
	//d := []int{1, 2, 3, 4, 5}
	//fmt.Println(len(d))
	//sum_f := GetAverage(d)
	//fmt.Printf("数组的平均值为：%f\n", sum_f)

	//endregion

	//region 二维数组
	//fmt.Println("*****************二维数组：遍历1****************")
	//two_body := [2][3]int{
	//	{2, 3, 4},
	//	{4, 5, 3},
	//}
	//
	////二维数组遍历
	//for i, val := range two_body {
	//	for j, val_1 := range val {
	//		fmt.Printf("two_body[%d][%d] = %d\n", i, j, val_1)
	//
	//	}
	//}
	//fmt.Println("*****************二维数组：遍历2****************")
	//two_body_1 := [...][3]string{
	//	{"1", "2", "3"},
	//	{"4", "5", "6"},
	//}
	//
	//for i, v1 := range two_body_1 {
	//	for j := range v1 {
	//		fmt.Printf("two_body_1[%d][%d] = %s\n", i, j, two_body_1[i][j])
	//
	//	}
	//}
	//
	//var arr = [3][4]int{
	//	{0, 1, 2, 3},
	//	{4, 5, 6, 7},
	//	{8, 9, 10, 11},
	//}
	//fmt.Println(arr[1][1])

	//向二维数组添加元素

	//删除二维数组中的元素

	//endregion

	//region 浮点数精度丢失问题
	//a := 1.69
	//b := 1.7
	//c := a * b      // 结果应该是2.873
	//fmt.Println(c)  // 输出的是2.8729999999999998

	//a := 1690           // 表示1.69
	//b := 1700           // 表示1.70
	//c := a * b          // 结果应该是2873000表示 2.873
	//fmt.Println(c)      // 内部编码
	//fmt.Println(float64(c) / 1000000) // 显示
	//endregion

	//region 指针
	//a := 5
	//var p *int
	//p = &a
	//fmt.Println(p)
	//
	//fmt.Println("使用指针交换两个变量的值")
	//b := 10
	//swap(&a, &b)
	//fmt.Printf("a=%d, b=%d\n", a, b)

	//endregion

	student := Students{}
	student.name = "SAM"
	fmt.Println(student.name)
	student.rename("Lemon")
	fmt.Println(student.name)

	student.name = "Turentu"
	fmt.Println(student.name)

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

//region 矩形结构体和矩形的方法

// 定义一个矩形的结构体
type Rectangle struct {
	length float64
	width  float64
}

// 获取矩形的长
func (r Rectangle) getLength() float64 {
	return r.length
}

// 设置矩形的长
func (r *Rectangle) setLength(length float64) {
	r.length = length
}

// //获取矩形的宽
func (r Rectangle) getWidth() float64 {
	return r.width
}

// 设置矩形的宽
func (r *Rectangle) setWidth(width float64) {
	r.width = width
}

// 矩形面积
func (r Rectangle) getArea() float64 {
	return r.length * r.width
}

//endregion

// region 定义一个结构体
type Circle struct {
	radius float64
}

// 类似于 类方法;
// 圆面积方法
func (c Circle) getArea() float64 {
	return c.radius * c.radius * PI
}

// 圆周长
func (c Circle) getLen() float64 {
	return 2 * PI * c.radius
}

//endregion

func GetAverage(params []int) float64 {
	var sum int
	for _, v := range params {
		sum += v
	}
	fmt.Printf("数组的和为：%d\n", sum)
	fmt.Printf("数组的长度为：%d\n", len(params))
	return float64(sum / len(params))

}

func change_array(p []int) []int {
	p[1] = 23
	return p
}

func change(nums [3]int) {
	nums[0] = 100
}

func swap(a *int, b *int) (int, int) {
	*a, *b = *b, *a // 简洁的写法
	//var temp int
	//temp = *a
	//*a = *b
	//*b = temp
	return *a, *b
}

type Students struct {
	name  string
	sex   int
	age   int
	phone string
}

func (s *Students) rename(new_name string) {
	s.name = new_name
}
