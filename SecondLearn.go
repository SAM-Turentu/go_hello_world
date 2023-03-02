package main

import (
	"bytes"
	"fmt"
	"sync"
)

func Second() {
	fmt.Println("*********************Second Learn Golang*********************")
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
	//arr := []int{1, 2, 3, 4}
	////if 1 in arr{} // 在python等语言中可以使用 in，go中不能使用
	//for _, v := range arr {
	//	if v == 1 {
	//
	//	}
	//}
	//_boo, _ := strconv.ParseBool("0")
	//if _boo {
	//	fmt.Println(_boo)
	//}

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
	//a := make([]int, 5, 10)
	//fmt.Printf("a: %v, len(a): %d, cap(a): %d\n", a, len(a), cap(a))
	//a[0] = 1 // 向a切片中添加5个元素
	//a[1] = 2 // 向a切片中添加5个元素
	//a[2] = 3 // 向a切片中添加5个元素
	//a[3] = 4 // 向a切片中添加5个元素
	//a[4] = 5 // 向a切片中添加5个元素
	//fmt.Printf("a: %v\n", a)
	////b := []int{}
	//b := make([]int, 3, 10)
	//
	//fmt.Printf("b: %v, len(b): %d, cap(b): %d\n", b, len(b), cap(b))
	//copy(b, a) // 将a复制给b，len(b) = 3 所以只复制a[:3]
	//fmt.Printf("b: %v, len(b): %d, cap(b): %d\n", b, len(b), cap(b))
	//
	//c := make([]int, 10, 10)
	//fmt.Printf("c: %v, len(c): %d, cap(c): %d\n", c, len(c), cap(c))
	//copy(c, a) // 将c复制给a，len(a) = 5 所以只复制c[:5]
	//fmt.Printf("c: %v, len(c): %d, cap(c): %d\n", c, len(c), cap(c))
	//fmt.Printf("a: %v, len(a): %d, cap(a): %d\n", a, len(a), cap(a))

	//endregion

	//region 双链表
	//l := list.New() // 初始化一个新链表
	//fmt.Printf("双向链表的长度：%d\n", l.Len())
	//l.PushFront(10) // 向链表中添加元素
	//l.PushFront(11)
	//l.PushFront("推一个元素")
	//l.PushFront(Students{name: "", age: 18, sex: 1, phone: "1829200"})
	//fmt.Printf("查看链表最后一个值：%v\n", l.Front().Value) // 查看链表最后一个值
	//fmt.Printf("查看链表第一个值：%v\n", l.Back().Value)   //查看链表第一个值
	//
	///*
	//	遍历双向链表的所有元素
	//	先进先出，第一个被打印的是最后一个添加到链表中的元素
	//*/
	//fmt.Println("************************开始遍历链表数据************************")
	//for e := l.Front(); e != nil; e = e.Next() {
	//	fmt.Println(e.Value)
	//}

	//endregion

	//region range 返回的是引用
	//arr := []int{10,20, 30,40}
	//var a *int
	//for i, v := range arr {
	//	fmt.Printf("value: %d, value-attr: %X elem-addr %X\n", v, &v, &arr[i])
	//	a = &v
	//}
	//fmt.Printf("&a: %v, *a: %v\n", a, *a)
	//endregion

	//region sync.Map 并发环境中使用map
	//sync_map()
	//endregion

	//region nil 赋值【不建议使用】
	//fmt.Println(nil)
	//fmt.Printf("nil 类型 %T\n", nil)  // nil 没有类型
	//nil := 1  // nil 不是关键词或保留字，可以当做变量使用【不建议使用】
	//fmt.Println(nil)
	//endregion

	//region switch
	//var a = "daddy" // 或 num
	//switch a {
	//case "mum", "daddy":
	//	fmt.Println("family")
	//}
	//endregion

	//region goto
	//go_to()
	//goto goto_func
	//goto_func:
	//	fmt.Println("可以跳转到此处")
	//OuterLoop:
	//	for i := 0; i < 4; i++ {
	//	fmt.Println(i)
	//
	//	if i == 1 {
	//		break OuterLoop
	//	}
	//}
	//endregion

	//region 匿名函数
	//func(d int) {
	//	fmt.Printf("d: %d\n", d)
	//}(20)
	//
	//d := func(d int) {
	//	fmt.Printf("d: %d\n", d)
	//}
	//d(101)
	//arr := []int{1, 2, 3, 4, 5}
	//visit(arr, func(v int) { // 匿名函数作为回调函数
	//	fmt.Printf("%d * %d = %d\n", v, v, v*v)
	//})
	//endregion

	//region 函数当接口
	////结构体的接口
	//var invoker Invoker
	//s := new(Studentss)
	//invoker = s
	//invoker.Call("hello")
	//
	////函数体实现接口
	//invoker = FuncCaller(func(v interface{}) {
	//	fmt.Println("from function", v)
	//})
	//invoker.Call("hello")
	//endregion

	//region 闭包
	//accumulator := Accumulate(1)
	//fmt.Println(accumulator())
	//fmt.Println(accumulator())
	//fmt.Printf("打印累加器的地址：%p\n", &accumulator)
	//
	//accumulator2 := Accumulate(10)
	//fmt.Println(accumulator2())
	//fmt.Printf("打印累加器的地址：%p\n", &accumulator2) // 两个累加器地址不同
	//
	//generator := playerGen("high noon")
	//
	//name, hp := generator()
	//fmt.Println(name, hp)

	//endregion

	//region 可变参数类型
	//myfunc(1, 2, 3, 4)
	//myfunc(7, 8)
	//
	////将输入连接成字符串
	//fmt.Println(joinStrings("pig", "and", "cat"))
	//fmt.Println(joinStrings("hammer", "mom", "and", "hawk"))

	//endregion

	//region

	//endregion

	fmt.Println("*************************Second END*************************")

}

// region 可以写一个in方法
/*func in(p int, arr []int) bool {
	result := false
	for _, v := range arr {
		if v == p {
			result = true
		}
	}
	return result
}*/

//func in(s string, arr []int) bool {
//
//}

//endregion

// region 定义链表  https://zhuanlan.zhihu.com/p/404754444   ==  import container/list
// Element 定义链表中的节点结构体
//type Element struct {
//	next, prev *Element
//	list       *List // 元素所在的链表
//	Value      interface{}
//}
//
//// List 表示双向链表
//type List struct {
//	/*
//		根节点，链表的起点，不存值
//		直接访问 root.prev root.next 分别是链表的前一个几点 和 后一个节点
//	*/
//	root Element
//	len  int // 链表的长度，不包括根节点
//}
//
//// Init 链表初始化（清空链表）
//func (l *List) Init() *List {
//	l.root.next = &l.root
//	l.root.prev = &l.root
//	l.len = 0
//	return l
//}
//
//// New 创建一个新链表
//func New() *List {
//	return new(List).Init()
//}
//
//// Len 查看链表长度
//func (l *List) Len() int {
//	return l.len
//}
//
//// Front 返回链表首节点
//func (l *List) Front() *Element {
//	if l.len == 0 {
//		return nil
//	}
//	return l.root.next
//}
//
//// Back 返回链表的末节点
//func (l *List) Back() *Element {
//	if l.len == 0 {
//		return nil
//	}
//	return l.root.prev
//}
//
//// insert 在链表特定元素后插入新元素
//func (l *List) insert(e, at *Element) *Element {
//	e.prev = at      // 将e的前节点指向at
//	e.next = at.next // 将原有的at后一个节点 给 e的后一个节点
//	e.prev.next = e  // e的前一个节点的后一个节点 是 e本身
//	e.next.prev = e  // e的后一个节点的前一个节点 是 e本身
//	e.list = l       // l 元素所在的链表
//	l.len++
//	return e
//}
//
//// 直接向链表插入一个值
//func (l *List) insertValue(v interface{}, at *Element) *Element {
//	return l.insert(&Element{Value: v}, at)
//}
//
//// 删除链表指定节点
//func (l *List) remove(e *Element) *Element {
//	e.prev.next = e.next
//	e.next.prev = e.prev
//	e.next = nil // 赋空，避免内存泄露
//	e.prev = nil
//	e.list = nil
//	l.len--
//	return e
//}
//
//// 将一个节点e移动到另一个节点后
//func (l *List) move(e, at *Element) *Element {
//	if e == at {
//		return e
//	}
//	//先删除一个元素
//	e.prev.next = e.next
//	e.next.prev = e.prev
//
//	//将元素插入在一个元素后
//	e.prev = at
//	e.next = at.next
//	e.prev.next = e
//	e.next.prev = e
//	//链表长度不变
//	return e
//}
//
//// 删除节点并返回元素值
//func (l *List) Remove(e *Element) interface{} {
//	if e.list == l {
//		//判断节点是否在唉该链表中
//		// 如果链表是空的话，而且e是nil的话会panic
//		l.remove(e)
//	}
//	return e.Value
//}
//
//// 插入一个值到链表头
//func (l *List) PushFront(v interface{}) *Element {
//	return l.insertValue(v, &l.root)
//}
//
//// 在链尾添加一个值
//func (l *List) PushBack(v interface{}) *Element {
//	return l.insertValue(v, l.root.prev)
//}
//
//// 插入一个值到某个节点前面
//func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
//	if mark.list != l {
//		return nil
//	}
//	return l.insertValue(v, mark.prev)
//}
//
//// 插入一个元素在某个节点后
//func (l *List) InsertBack(v interface{}, mark *Element) *Element {
//	if mark.list != l {
//		return nil
//	}
//	return l.insertValue(v, mark)
//}
//
//// 将一个节点移动到链表首
//func (l *List) MoveToFront(e *Element) {
//	if e.list != l || l.root.next != e {
//		return
//	}
//	l.move(e, &l.root)
//}
//
//// 将一个元素移动到链表尾部
//func (l *List) MoveToBack(e *Element) {
//	if e.list != l || l.root.prev != e {
//		return
//	}
//	l.move(e, l.root.prev)
//}
//
//// 将一个节点移动到另一个节点前
//func (l *List) MoveBefore(e, mark *Element) {
//	if e.list != l || e == mark || mark.list != l {
//		return
//	}
//	l.move(e, mark.prev)
//}
//
//// 将一个节点移动到另一个节点后
//func (l *List) MoveAfter(e, mark *Element) {
//	if e.list != l || e == mark || mark.list != l {
//		return
//	}
//	l.move(e, mark)
//}
//
//////在链表后面添加一个链表
////func (l *List) PushBackList(other *List) {
////	for i, e := other.Len(), other.Front(); i >0; i, e = i-1, e.Next(){
////		l.insertValue(e.Value, l.root.prev)
////	}
////}
////
//////将一个链表逆序添加到链表到前面
////func (l *List) PushFrontList(other *List) {
////
////}

//endregion

// region sync.Map 并发环境中使用map
// 需要并发读写时，要加锁，
func sync_map() {
	var s sync.Map
	//将键值保存到sync.Map
	s.Store("name", "sam")
	s.Store("london", []int{10, 20, 30})
	s.Store("egypt", 200)

	fmt.Println(s.Load("name"))   // 返回 sam true
	fmt.Println(s.Load("green"))  // 没有键，返回 nil false
	fmt.Println(s.Load("london")) // 返回 [10 20 30] true

	s.Delete("egypt") // 删除键值对

	//遍历所有sync.Map中的键值对
	s.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})

}

//endregion

// region goto
func go_to() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			if i%2 != 0 {
				fmt.Printf("i = %d, j = %d\n", i, j)
				goto break_print
			}
		}
	}
break_print:
	fmt.Println("跳出循环")
}

//endregion

// region 匿名函数
// 匿名函数作为回调函数
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

//endregion

// region 函数体实现接口 和 结构体实现接口 对比
// 结构体的接口
type Invoker interface {
	Call(interface{})
}

type Studentss struct {
}

func (s *Studentss) Call(p interface{}) {
	fmt.Println("from struct", p)
}

// 函数的接口
type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}) {
	f(p)
}

//endregion

//region

// Accumulate 累加器
func Accumulate(value int) func() int {
	//闭包
	return func() int {
		value++
		return value
	} // 返回一个int类型的func()
}

// 生成器（工厂模式）
func playerGen(name string) func() (string, int) {
	hp := 150
	//返回创建的闭包
	return func() (string, int) {
		//将变量应用到闭包中
		return name, hp
	}
}

//endregion

//region 可变参数

// myfunc 可变参数类型
func myfunc(args ...int) {
	for _, v := range args {
		fmt.Printf("%d", v)
	}
	fmt.Println()
}

// joinStrings 任意类型的可变参数
func joinStrings(slist ...string) string {
	// 定义一个字节缓冲， 快速地连接字符串
	var b bytes.Buffer
	// 遍历可变参数列表slist， 类型为[]string
	for _, s := range slist {
		// 将遍历出的字符串连续写入字节数组
		b.WriteString(s)
	}
	// 将连接好的字节数组转换为字符串输出
	return b.String()
}

//endregion

//region

//endregion
