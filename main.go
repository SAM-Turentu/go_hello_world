package main

import (
	"encoding/json"
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
	fmt.Println("**************************开始**************************")
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
	//endregion∂

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

	//region 结构体
	//student := Students{}
	//student.name = "SAM"
	//fmt.Println(student.name)
	//student.rename("Lemon")
	//fmt.Println(student.name)
	//
	//student.name = "Turentu"
	//fmt.Println(student.name)

	//结构体成员首字母大写的字段，才可以转换为json格式
	//class := ClassStruct{Class: "one", grade: "tow", head_teacher: "SAM"}
	//fmt.Printf("class %s, grade %s\n", class.Class, class.grade)
	////json.Marshal 返回两个值
	//if result, err := json.Marshal(&class); err == nil { //前声明语句; 后逻辑判断语句
	//	fmt.Println(string(result))
	//}
	////结构体成员首字母大写的字段，才可以转换为json格式
	//user := UserInfo{UserName: "sam", Phone: "1829200", PassWord: "264264"}
	//if result, err := json.Marshal(&user); err == nil { //前声明语句; 后逻辑判断语句
	//	fmt.Println(string(result))
	//}
	//结构体成员首字母大写的字段，才可以转换为json格式
	//user2 := UserInfoV2{UserName: "sam", Phone: "182", PassWord: "264264"}
	//if result, err := json.Marshal(&user2); err == nil { //前声明语句; 后逻辑判断语句
	//	fmt.Println(string(result))
	//}
	//json.Unmarshal()  JSON反序列化

	//var p *UserInfo
	//p = &user
	//fmt.Println(p)          // &{sam 1829200 264264}
	//fmt.Println(p.UserName) //  sam
	//fmt.Println(*p)         //  {sam 1829200 264264}
	//fmt.Println(&user)      //  &{sam 1829200 264264}
	//fmt.Println(user)       //  {sam 1829200 264264}
	//fmt.Println(user.Phone) //  1829200

	//endregion

	//region 切片
	//var slice0 []int
	//slice1 := []int{3, 2, 3}
	//slice2 := make([]int, 10)
	//var slice3 = make([]int, 10)
	//slice4 := make([]int, 10, 20)
	//
	//fmt.Println(slice0)
	//fmt.Println(slice1)
	//fmt.Println(slice2)
	//fmt.Println(slice3)
	//
	//if slice0 == nil {
	//	fmt.Println("slice0 是空切片")
	//}
	//if slice1 == nil {
	//	fmt.Println("slice1 是空切片")
	//}
	//if slice2 == nil {
	//	fmt.Println("slice2 是空切片")
	//}
	//if slice4 == nil {
	//	fmt.Println("slice4 是空切片")
	//}
	//
	//slice1 = append(slice1, 1)
	//slice1 = append(slice1, 2)
	//slice1 = append(slice1)
	//fmt.Printf("slice1 = %v\n", slice1) // slice1 = [3 2 3 1 2]
	//
	//s := slice1[1:3]
	//fmt.Println(s) //[2 3]
	//
	//s = slice1[1:]
	//fmt.Println(s) //[2 3 1 2]
	//
	//s = slice1[:]
	//fmt.Println(s) //[3 2 3 1 2]
	//
	//s = slice1[:3] //没有负数slice1[-1]  slice1[:-1]
	//fmt.Println(s) //[3 2 3]
	//
	////cap 计算切片容量的方法，测量切片最长可以达到多少
	////切片的容量(capacity)问题  https://zhuanlan.zhihu.com/p/413972333
	//fmt.Printf("len=%d cap=%d slice=%v\n", len(s), cap(s), s)
	//s = append(s, 5)
	//s = append(s, 6)
	//s = append(s, 7)
	//fmt.Printf("len=%d cap=%d slice=%v\n", len(s), cap(s), s)
	//s = append(s, 8)
	//fmt.Printf("len=%d cap=%d slice=%v\n", len(s), cap(s), s)

	////两层切片（两层不定长的数组）
	//var slice0 []int
	//var slice1 [][]int
	//slice2 := make([][]int, 5)
	//slice3 := [][]int{{1, 2}, {3, 4}, {5, 6}}
	//
	////i := 6
	////slice4 := [i]int{1, 2, 3, 4}  // 无法通过参数设置数组长度
	//slice4 := [6]int{1, 2, 3, 4}
	//
	//fmt.Println(slice0)
	//fmt.Println(slice1)
	//fmt.Println(slice2)
	//fmt.Println(slice3)
	//fmt.Printf("len=%d cap=%d slice=%v\n", len(slice3), cap(slice3), slice3)
	//fmt.Println(slice4)
	//fmt.Printf("len=%d cap=%d slice=%v\n", len(slice4), cap(slice4), slice4)
	//slice4[4] = 5 //定长数组不能使用append
	//fmt.Println(slice4)
	//fmt.Printf("len=%d cap=%d slice=%v\n", len(slice4), cap(slice4), slice4)
	//
	//
	////slice1 = append(slice1, {2,3,4})  // error
	////slice2 = append(slice2, {1,2})  // error
	////slice3[3] = {7,8}

	//endregion

	//region map 序列化与反序列化
	//serialization_of_map()
	//deserialization_of_map()
	//endregion

	//region 范围Range
	//map1 := make(map[int]float32)
	//map1[1] = 1.0
	//map1[2] = 2.0
	//map1[3] = 3.0
	//map1[4] = 4.0
	//fmt.Println(map1)
	//
	//// 读取 key 和 value
	//for key, value := range map1 {
	//	fmt.Printf("key is: %d - value is: %f\n", key, value)
	//}
	//
	//// 读取 key
	//for key := range map1 {
	//	fmt.Printf("key is: %d\n", key)
	//}
	//
	//// 读取 value
	//for _, value := range map1 {
	//	fmt.Printf("value is: %f\n", value)
	//}
	//
	//s := "中国人"
	//ss := "chinese"
	//fmt.Println(s)
	//for i := 0; i < len(ss); i++ { // 中文不能使用该方法遍历汉字
	//	fmt.Printf("%x ", ss[i])
	//}
	//for _, val := range s {
	//	fmt.Printf("%x %c,", val, val)
	//}

	//endregion

	//region map 集合
	//map1 := make(map[string]string, 10)
	//fmt.Printf("map1 = %v\n", map1)
	//map2 := make(map[string]int, 10)
	//fmt.Printf("map2 = %v\n", map2)
	//
	//map1["name"] = "sam"
	//map1["age"] = "18"
	//fmt.Printf("map1[name] = %s\n", map1["name"])
	//fmt.Printf("map1[sex] = %s\n", map1["sex"]) //不存在的键值 输出零值 int 为0， string 为 ""， bool 为 false ...
	//
	//map3 := map[int]string{1: "hello", 2: "world"} // 键值可以是int型
	//fmt.Printf("map3[1] = %s\n", map3[1])
	//fmt.Printf("map3 = %v\n", map3)
	//map4 := map[string]string{
	//	"UserName": "SAM",
	//	"Phone":    "1829200",
	//}
	//fmt.Printf("map4 = %v\n", map4) //map无序,每次都不相同
	//
	//map5 := map[string]interface{}{}
	//map5["UserName"] = "Turentu"
	//map5["Password"] = "264264"
	//map5["Details"] = map[string]string{"Province": "Shanghai", "City": "Shanghai", "Country": "China"}
	//map5["Age"] = []int{1, 2, 3}
	//map5["Address"] = "Shanghai, China"
	//map5["Address"] = "Shanghai, China"
	//fmt.Printf("全部输出map5 = %v\n", map5) //map无序,每次都不相同
	//for k, v := range map5 {
	//	fmt.Printf("map5[%s] = %v,  type(%s) = %v\n", k, v, k, reflect.TypeOf(v))
	//}
	//delete(map5, "Age")                 // 删除键值对
	//fmt.Printf("全部输出map5 = %v\n", map5) //map无序,每次都不相同
	//map5["Age"] = 18
	//fmt.Printf("map5[Age] = %d\n", map5["Age"])

	//endregion

	fmt.Println("**************************结束**************************")
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

// 使用指针修改结构体成员的值
func (s *Students) rename(new_name string) {
	s.name = new_name
}

// ClassStruct 结构体成员首字母大写的字段，才可以转换为json格式
type ClassStruct struct {
	Class        string // 年级；首字母小写为私有结构体成员，json包无法使用
	grade        string // 班级；且小写开头的字段在其他包内不能被调用
	head_teacher string // 班主任
}

// UserInfo 结构体成员首字母大写的字段，才可以转换为json格式
type UserInfo struct {
	UserName string
	Phone    string
	PassWord string
}

type UserInfoV2 struct {
	UserName string `json:"user_name"` // `json:"重命名"`
	Phone    string `json:"phone"`
	PassWord string `json:"-"` // `json:"-"`  标记忽略该字段
}

// map 序列化
func serialization_of_map() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["Name"] = "Turentu"
	a["Age"] = 18
	a["skill"] = "python,java,.net,golang..."
	if data, err := json.Marshal(a); err != nil {
		fmt.Printf("Map序列化失败！err = %v\n", err)
	} else {
		fmt.Printf("Map序列化后：%v\n", string(data))
	}
}

// 反序列化
func deserialization_of_map() {

}
