package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"sync"
)

func Second() {
	//start := time.Now()  // 获取当前时间
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

	//region defer 可以用来关闭文件，释放资源
	//_deferr()
	//v := readValue("key_name")
	//v2 := readValue2("key_name_2")
	//fmt.Println(v)
	//fmt.Println(v2)
	//endregion

	//region error
	//var err = errors.New("this is error")  // 相当于 python的 raise "发生了错误！"
	//s := err.Error() //  this is error
	//fmt.Println(err) //  this is error
	//fmt.Println(s)
	//endregion

	//region panic 宕机
	//panic("好疼！")  // 类似与其他语言的 Exception
	//MustComplie("正则发生了错误")
	//endregion

	//region 宕机后继续运行

	// 允许一段手动触发的错误
	//ProtectRun(func() {
	//	fmt.Println("手动宕机前")
	//	// 使用panic传递上下文
	//	panic(&paincContext{
	//		"手动触发panic",
	//	})
	//})

	// 故意造成空指针访问错误
	//ProtectRun(func() {
	//	fmt.Println("赋值宕机前")
	//	var a *int
	//	*a = 1
	//	fmt.Println("赋值宕机后")
	//})

	//fmt.Println("运行后")

	//endregion

	//region 代码运行时间

	//duration := time.Since(start)  // 将开始时间传入，返回与此刻时间的差值，即程序运行时间
	//fmt.Println("程序运行时间：", duration)

	//s := time.Now().UnixNano() // 纳秒
	//s_haomiao := s / 1e6       // 毫秒
	//fmt.Println(s)
	//fmt.Println(s_haomiao)
	//for i := 0; i < 100000; i++ {
	//	sum := 0
	//	sum++
	//}
	//e := time.Now().UnixNano()
	//e_haomiao := e / 1e6
	//fmt.Println(e_haomiao)
	//d := e - s
	//fmt.Println(d)                          // 使用纳秒 有数值
	//fmt.Printf("%d\n", e_haomiao-s_haomiao) // 毫秒差值太小

	//endregion

	//region 结构体
	//cat := Cat{"花花", 1, "母"} // 匿名结构体字段
	//fmt.Println(cat)
	//
	//car := Cars{
	//	Wheel: Wheel{Size: 18},
	//
	//	// 初始化引擎；嵌入式匿名结构体 需要重新声明
	//	Engine: struct {
	//		Power int
	//		Type  string
	//	}{Power: 220, Type: "BMW"},
	//}
	//fmt.Println(car)

	//endregion

	//region SetFinalizer 终止器
	//entry() // 终止器只有在对象被 GC 时，才会被执行。其他情况下，都不会被执行，即使程序正常结束或者发生错误。
	//for i := 0; i < 5; i++ {
	//	time.Sleep(time.Second)
	//	runtime.GC()
	//}

	//endregion

	//region io 读
	// data := []byte("Hello World！ 你好世界！")
	// read(data)

	// readByte(data)

	// data = []byte("Hello World, 你好世界！")
	// readBytes(data)

	// data = []byte("Hello World！\r\n 你好世界！")
	// readLine(data)

	// data = []byte("你好世界！\r\n Hello World！")
	// readRune(data)

	// data = []byte("Hello World, 你好世界！")
	// readSlice(data)

	// readString(data)

	// // data = []byte("Go语言入门教程")
	// read_buffered(data)

	//endregion

	//region io 写
	// data := []byte("Hello World, 你好世界！")
	// avaliable(data)

	// wirte_buffered(data)

	// WriteByte(data)

	// writeRune()

	// str := string(data)

	// writeString(str)

	//endregion

	//region 接口

	// file := new(files) // 初始化files结构体
	// var writer Writer  // 声明一个接口
	// writer = file
	// writer.WriteData("data...")
	// val, err := writer.(interface{}) // 判断接口类型
	// fmt.Println(val, err)

	//endregion

	//region 排序
	// names := MyStringList{
	// 	"3. Triple Kill",
	// 	"5. Penta Kill",
	// 	"2. Double Kill",
	// 	"4. Quadra Kill",
	// 	"1. First Blood",
	// }
	// sort.Sort(names)          // 排序
	// for _, v := range names { // 遍历排序后打印
	// 	fmt.Printf("%s\n", v)
	// }

	// ages := IntSlice{
	// 	10,
	// 	3,
	// 	13,
	// 	5,
	// 	3,
	// 	1,
	// 	0,
	// }
	// sort.Sort(ages)
	// for _, v := range ages { // 遍历排序后打印
	// 	fmt.Printf("%d\n", v)
	// }

	// score := MyFloat{
	// 	23.5,
	// 	56.5,
	// 	478.5,
	// 	45.151213,
	// 	10.0,
	// 	23.5,
	// }
	// sort.Sort(score)
	// for _, v := range score { // 遍历排序后打印
	// 	fmt.Printf("%f\n", v)
	// }

	// hero := Heros{
	// 	&Hero{"吕布", Tank},
	// 	&Hero{"李白", Assassin},
	// 	&Hero{"妲己", Mage},
	// 	&Hero{"貂蝉", Assassin},
	// 	&Hero{"关羽", Tank},
	// 	&Hero{"dema", Assassin},
	// 	&Hero{"Axi", Assassin},
	// 	&Hero{"诸葛亮", Mage},
	// }
	// sort.Sort(hero)          // 先排英文，后排中文
	// for _, v := range hero { // 遍历排序后打印
	// 	fmt.Printf("%+v\n", v)
	// }

	//endregion

	//region 断言 switch
	// var a int = 1
	// var i interface{} = a
	// fmt.Println(a == i) // 不同类型可以和 interface 比较
	// // 不能比较空接口中的动态值，比如 切片， map

	// assert()

	// var name string = "SAM"
	// var age int = 18
	// var sex bool = true
	// FuncSwitch(name) // 使用switch断言
	// FuncSwitch(age)
	// FuncSwitch(sex)
	// FuncSwitch(Students{name: "ni"})

	// print(new(Alipay))
	// print(new(Cash))

	//endregion

	//region error
	// result, err := Sqrt(-13)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }

	//endregion

	fmt.Println("*************************Second END*************************")

}

//region 可以写一个in方法
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

//region 定义链表  https://zhuanlan.zhihu.com/p/404754444   ==  import container/list
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
//
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
//
// 匿名函数作为回调函数
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

//endregion

// region 函数体实现接口 和 结构体实现接口 对比
//
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

//region 闭包

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

// region defer
func _deferr() {
	fmt.Println("start...")
	a := 1
	defer fmt.Println(a + 1) // 延迟执行语句
	defer fmt.Println(a + 2) //类似栈，后进先出
	defer fmt.Println(a + 3) //在函数退出时，释放资源
	a++
	fmt.Printf("a + 1: %d\n", a)
	fmt.Printf("a: %d\n", a)
	fmt.Println("end")
}

var (
	valueByKey      = make(map[string]int)
	valueByKeyGuard sync.Mutex // 保证使用映射时的并发安全的互斥锁
)

func readValue(key string) int {
	valueByKeyGuard.Lock() // 对共享资源加锁
	v := valueByKey[key]
	valueByKeyGuard.Unlock() // 对共享资源解锁
	return v
}

// readValue2 使用延迟并发解锁 由上面简化后
func readValue2(key string) int {
	valueByKeyGuard.Lock()         // 对共享资源加锁
	defer valueByKeyGuard.Unlock() // 函数执行完后，释放资源
	return valueByKey[key]
}

func fileSize(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		return 0
	}
	//defer f.Close()

	info, err := f.Stat()
	if err != nil {
		f.Close()
		return 0
	}

	size := info.Size()
	f.Close()
	return size

}

//endregion

// region go中的错误机制
func Dial(network, address string) (net.Conn, error) {
	var d net.Dialer
	return d.Dial(network, address)
}

//type Writer interface {
//	Writer(p []byte) (n int, err error)
//}
//
//type Closer interface {
//	Closer() error
//}

//type errorString struct {
//	s string
//}
//
//// New 自定义一个错误
//func New(text string) error {
//	return &errorString{text}
//}
//
//func (e errorString) Error() string {
//	return e.s
//}

//endregion

// region 正则表达式发生错误时，主动宕机（抛错）
func MustComplie(s string) *regexp.Regexp {
	regexp, err := regexp.Compile(s)
	if err != nil {
		panic(`regexp: Compile(` + strconv.Quote(s) + `): ` + err.Error())
		defer fmt.Println("宕机后要做的事") // 方法资源释放时执行
	}
	return regexp
}

//endregion

//region 宕机继续运行

// paincContext 崩溃时需要传递的上下文信息
type PaincContext struct {
	function string // 所在函数
}

// ProtectRun 保护方式允许一个函数
func ProtectRun(entry func()) {
	defer func() {
		// 发生宕机时，获取painc传递的上下文并打印
		err := recover()
		switch err.(type) { // 断言 判断类型
		case runtime.Error: //运行时错误
			fmt.Println("runtime error: ", err)
		default: // 非运行时错误
			fmt.Println("error: ", err)
		}
	}()

	entry()

}

//endregion

//region 结构体

type Command struct {
	Name    string // 指令名称
	Var     *int   // 指令绑定的变量
	Comment string // 指令的注释
}

func NewCommand(name string, varref *int, command string) *Command {
	return &Command{
		Name:    name,
		Var:     varref,
		Comment: command,
	}
}

type Cat struct {
	Name   string
	int    //匿名字段
	string //匿名字段
}

type Wheel struct { //轮子
	Size int // 尺寸
}

type Cars struct {
	Wheel
	// 嵌入的结构体不回被外部引用；初始化这个嵌入式结构体时，需要重新声明才能赋值
	Engine struct { // 引擎
		Power int    // 功率
		Type  string // 类型
	}
}

//endregion

// region 垃圾回收GC
type Road int

func findRoad(r *Road) {
	log.Println("road: ", *r)
}

func entry() {
	var rd Road = Road(999)
	r := &rd

	runtime.SetFinalizer(r, findRoad)

}

//endregion

//region IO 读方法（Read,ReadByte, ReadBytes,ReadLine, ReadRune, ReadSlice, ReadString,UnreadByte, UnreadRune, Buffered,Peek）

func read(data []byte) {
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [50]byte // buf 用来存放读取数据的字节切片
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:]), n, err)
}

func readByte(data []byte) {
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	c, err := r.ReadByte()
	fmt.Println(string(c), err)
}

func readBytes(data []byte) {
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','            // 在读到第一个"标识符"前出错，返回已读取的数据和错误，
	line, err := r.ReadBytes(delim) // 当返回的数据不以“delim”结尾时，返回的 err 才不为空值   EOF
	fmt.Println(string(line), err)
}

func readLine(data []byte) {
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	line, prefix, err := r.ReadLine() // 一行数据太多，超出缓冲区，只会返回前面的数据，isPrefix会返回ture，剩余的数据会在之后调用中返回
	fmt.Println(string(line), prefix, err)
}

func readRune(data []byte) {
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	ch, size, err := r.ReadRune() // 读取一个utf-8字符，并返回字节数
	fmt.Println(string(ch), size, err)
}

func readSlice(data []byte) {
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ',' // 分隔符
	line, err := r.ReadSlice(delim)
	fmt.Println("start", string(line), err)
	line, err = r.ReadSlice(delim)
	fmt.Println(string(line), err)
	line, err = r.ReadSlice(delim)
	fmt.Println(string(line), err)
	fmt.Println("end.")
}

func readString(data []byte) {
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','
	line, err := r.ReadString(delim)
	fmt.Println(string(line), err)
}

func read_buffered(data []byte) {
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [14]byte
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err) // Go语言入门 14 <nil>
	rn := r.Buffered()
	fmt.Println(rn) // 6

	n, err = r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err) // 教程 6 <nil>
	rn = r.Buffered()
	fmt.Println(rn) // 0

}

//endregion

//region IO write 写方法（Avaliable, Buffered, Flush, Wirte, WriteByte, WriteRune, WriteString）

// avaliable 返回缓冲区未使用的字节数
func avaliable(data []byte) {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	fmt.Println("写入前未使用的缓冲区为：", w.Available())
	w.Write(data)
	fmt.Printf("写入%q后，未使用的缓冲区为：%d\n", string(data), w.Available())
	fmt.Println()
}

// buffered 返回已写入当前缓冲区的字节数
func wirte_buffered(data []byte) {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)

	fmt.Println("写入前未使用的缓冲区为：", w.Available())

	w.Write(data) // 写入缓冲区，返回已写入的字节数，写入长度和原始长度不同，返回error

	fmt.Printf("写入%q后，未使用的缓冲区为：%d\n", string(data), w.Buffered())

	w.Flush() // 把缓冲区中的主句写入底层io.Writer，并返回错误信息，写入成功返回error=nil

	fmt.Println("执行Flush方法后，写入的字节数为：", w.Buffered())
	fmt.Println()
}

func WriteByte(data []byte) {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	var c byte = 'G'
	err := w.WriteByte(c) // 写入一个字节，成功返回nil
	w.Flush()
	fmt.Println(string(wr.Bytes()), err)
	fmt.Println()
}

func writeRune() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)

	var r rune = 'G'
	size, err := w.WriteRune(r) // 以utf-8编码写入一个unicode字符
	w.Flush()
	fmt.Println(string(wr.Bytes()), size, err)
	fmt.Println()
}

func writeString(data string) {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	n, err := w.WriteString(data) // 写入一个字符串，返回字节数和错误信息
	fmt.Println(string(wr.Bytes()), n, err)
	fmt.Println()
}

//endregion

//region 接口

// Students 在main.go中定义了一个结构体

type Writer interface { // 接口类型名一般以er结尾
	WriteData(data interface{}) error
	// Write(data interface{}) // 所有方法都必须实现才能正确编译
}

type files struct {
}

// 实现Writer接口的WriteData方法
func (d *files) WriteData(data interface{}) error {
	fmt.Println("WriteData:", data)
	return nil
}

//endregion

// region sort

// 字符串排序
type MyStringList []string

func (a MyStringList) Len() int           { return len(a) }           // 获取元素数量
func (a MyStringList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] } // 交换元素
func (a MyStringList) Less(i, j int) bool { return a[i] < a[j] }      // 比较元素

// int 切片排序
type IntSlice []int

func (a IntSlice) Len() int           { return len(a) }
func (a IntSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IntSlice) Less(i, j int) bool { return a[i] < a[j] }

type MyFloat []float64 // type MyFloat []float32

func (a MyFloat) Len() int           { return len(a) }
func (a MyFloat) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a MyFloat) Less(i, j int) bool { return a[i] < a[j] }

type HeroKind int

const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

type Hero struct {
	Name string   //英雄名字
	Kind HeroKind // 英雄种类
}

type Heros []*Hero

func (a Heros) Len() int      { return len(a) }
func (a Heros) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Heros) Less(i, j int) bool {

	if a[i].Kind != a[j].Kind { // 如果英雄的分类不一致时，优先对分类进行排序
		return a[i].Kind < a[j].Kind
	}
	return a[i].Name < a[j].Name // 默认按英雄名字字符升序排序
}

// sort只能使用 stirng, int, float64

//todo 是不是可以将 sort类似于封装

//endregion

// region 断言 switch

func assert() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	// c := w.(*bytes.Buffer)  // 发生panic  断言失败
	fmt.Println(f)
}

// FuncSwitch switch是增强版断言
func FuncSwitch(s interface{}) {
	switch s.(type) {
	case string:
		fmt.Println(s, " type is string")
	case int:
		fmt.Println(s, " type is int")
	case bool:
		fmt.Println(s, " type is bool")
	case interface{}:
		fmt.Println(s, " type is interface{}")
	}

}

// 电子支付
type Alipay struct {
}

// 电子支付可以支持刷脸
func (a *Alipay) CanUseFaceID() {

}

// 现金支付
type Cash struct {
}

// 现金支付，可能被偷
func (a *Cash) Stolen() {

}

type CantainCanUseFaceID interface {
	CanUseFaceID()
}

type CantainStolen interface {
	Stolen()
}

func print(payMethod interface{}) {
	switch payMethod.(type) {
	case CantainCanUseFaceID:
		fmt.Printf("%T can use faceid\n", payMethod)
	case CantainStolen:
		fmt.Printf("%T may be stolen\n", payMethod)
	}
}

//endregion

//region 自定义错误类型

type dualError struct {
	Num     float64
	problem string
}

func (e dualError) Error() string {
	return fmt.Sprintf("Wrong!!!, because \"%f\" is a negative number", e.Num)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, dualError{Num: f}
	}
	return math.Sqrt(f), nil
}

//endregion
