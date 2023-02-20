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

	////切片中添加一个切片
	//单链式操作
	//b := make([]int, 10)
	//fmt.Println(b)                                         // [0,0,0,0,0,0,0,0,0,0]
	//b = append(b[:2], append([]int{10}, b[2:]...)...)      // 在第2个位置插入10
	//fmt.Println(b)                                         // [0,0,10,0,0,  0,0,0,0,0]
	//b = append(b[:3], append([]int{1, 2, 3}, b[3:]...)...) // 在第3个位置插入切片
	//fmt.Println(b)                                         // [0,0,10,1,2,3,0,0,0,  0,0,0,0,0]
	//
	//b = make([]int, 10)
	//fmt.Println(b) // [0,0,0,0,0,  0,0,0,0,0]
	//a := append([]int{5}, b[7:]...)
	//fmt.Printf("a: %v\n", a)
	//b = append(b[:4], append([]int{5}, b[7:]...)...) // 在第3个位置插入切片
	//fmt.Println(b)                                   // [0 0 0 0 5 0 0 0]
	//
	////使用append也可以删除索引元素
	//b = append(b[:4], b[5:]...) // 删除 b[4] = 5 的元素
	//fmt.Printf("删除b[4]后，b:%v\n", b)

	//endregion

	//region 切片copy
	// 按照其中较小的那个数组切片的元素个数进行复制
	a := make([]int, 5, 10)
	fmt.Printf("a: %v, len(a): %d, cap(a): %d\n", a, len(a), cap(a))
	a[0] = 1 // 向a切片中添加5个元素
	a[1] = 2 // 向a切片中添加5个元素
	a[2] = 3 // 向a切片中添加5个元素
	a[3] = 4 // 向a切片中添加5个元素
	a[4] = 5 // 向a切片中添加5个元素
	fmt.Printf("a: %v\n", a)
	//b := []int{}
	b := make([]int, 3, 10)

	fmt.Printf("b: %v, len(b): %d, cap(b): %d\n", b, len(b), cap(b))
	copy(b, a) // 将a复制给b，len(b) = 3 所以只复制a[:3]
	fmt.Printf("b: %v, len(b): %d, cap(b): %d\n", b, len(b), cap(b))

	c := make([]int, 10, 10)
	fmt.Printf("c: %v, len(c): %d, cap(c): %d\n", c, len(c), cap(c))
	copy(c, a) // 将c复制给a，len(a) = 5 所以只复制c[:5]
	fmt.Printf("c: %v, len(c): %d, cap(c): %d\n", c, len(c), cap(c))
	fmt.Printf("a: %v, len(a): %d, cap(a): %d\n", a, len(a), cap(a))

	//endregion

	//region 双链表
	new_list := New() // 新建一个链表
	//new_list.insertValue()

	//endregion

	//region

	//endregion

	//region

	//endregion

	fmt.Println("************************Second END************************")
}

// region 可以写一个in方法
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

//endregion

// region 定义链表  https://zhuanlan.zhihu.com/p/404754444
// Element 定义链表中的节点结构体
type Element struct {
	next, prev *Element
	list       *List // 元素所在的链表
	Value      interface{}
}

// List 表示双向链表
type List struct {
	/*
		根节点，链表的起点，不存值
		直接访问 root.prev root.next 分别是链表的前一个几点 和 后一个节点
	*/
	root Element
	len  int // 链表的长度，不包括根节点
}

// Init 链表初始化（清空链表）
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// New 创建一个新链表
func New() *List {
	return new(List).Init()
}

// Len 查看链表长度
func (l *List) Len() int {
	return l.len
}

// Front 返回链表首节点
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back 返回链表的末节点
func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// insert 在链表特定元素后插入新元素
func (l *List) insert(e, at *Element) *Element {
	e.prev = at      // 将e的前节点指向at
	e.next = at.next // 将原有的at后一个节点 给 e的后一个节点
	e.prev.next = e  // e的前一个节点的后一个节点 是 e本身
	e.next.prev = e  // e的后一个节点的前一个节点 是 e本身
	e.list = l       // l 元素所在的链表
	l.len++
	return e
}

// 直接向链表插入一个值
func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// 删除链表指定节点
func (l *List) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // 赋空，避免内存泄露
	e.prev = nil
	e.list = nil
	l.len--
	return e
}

// 将一个节点e移动到另一个节点后
func (l *List) move(e, at *Element) *Element {
	if e == at {
		return e
	}
	//先删除一个元素
	e.prev.next = e.next
	e.next.prev = e.prev

	//将元素插入在一个元素后
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	//链表长度不变
	return e
}

//endregion
