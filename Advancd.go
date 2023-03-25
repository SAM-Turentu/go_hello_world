package main

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo"
	"math/rand"
	"net/http"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func AdvanceLearn() {
	fmt.Println("*********************Advance Learn Golang********************")

	//region web
	//Start()
	//endregion

	//region 导入package
	//handles.SAM()
	//endregion

	//region 并发锁
	//no_sync()
	//have_lock()
	//have_lock_v2()

	//ch := make(chan struct{}, 2)
	//for i := 0; i < 5; i++ {
	//	go ReadLock(i, ch)
	//}
	//for i := 0; i < 5; i++ {
	//	go WriteLock(i, ch)
	//}
	//for i := 0; i < 10; i++ {
	//	<-ch
	//}

	//endregion

	//region 时间和日期
	//usetime()
	//endregion

	//region 生产者与消费者 并发

	//c := make(chan int)  // 创建一个channel
	//
	//go consumer(c) // 并发执行消费者
	//
	//// 生产者
	//for i := 1; i < 10; i++ {  // 从1开始
	//	c <- i  // 将数据通过channel 传给consumer
	//}
	//c <- 0 // 通知并发读consumer结束死循环
	//<-c  // 等待consumer结束

	//endregion

	//region 单向通道
	//ch := make(chan int)
	//var chSendOnly chan<- int = ch // 只写channel
	//var chRecyOnly <-chan int = ch // 只读channel
	//// 只写只读都指向相同都地址
	//fmt.Println(chSendOnly)
	//fmt.Println(chRecyOnly)
	//
	//onlyRead := make(<-chan int)  // 只读channel
	//onlyWrite := make(chan<- int) // 只写channel
	//fmt.Println(onlyRead)
	//fmt.Println(onlyWrite)
	//
	//close(onlyWrite) // 关闭channel
	//endregion

	//region 玩个球 无缓冲channel
	//ch := make(chan int)
	//wg.Add(2)
	//go player("球手1", ch)
	//go player("球手2", ch)
	//
	//ch <- 1
	//wg.Wait() // 等待游戏结束

	//endregion

	//region 接力跑 无缓冲channel
	//baton := make(chan int)
	//wg.Add(1)
	//go run(baton)
	//baton <- 1
	//wg.Wait()
	//endregion

	//region 超时channel  select case
	//delayChan()
	//endregion

	//region 设置cpu核心数
	//fmt.Println(runtime.NumCPU()) // 处理器核数
	//runtime.GOMAXPROCS(8)  // 设置需要用到的CPU核心数
	//endregion

	//region flag 代码执行参数
	//name := flag.String("name", "SAM", "帮助信息")
	//age := flag.Int("age", 18, "年龄")
	//sex := flag.Bool("sex", true, "性别")
	//d := flag.Duration("d", 0, "时间间隔")
	////var score []int
	////scores := flag.Var(newScoreValues([]int{}, &score), "scores", "分数数组") //  flag.Var 自定义一个参数字段类型
	//
	//flag.Parse() // 需要在定义好flag参数后使用
	//fmt.Println(*name, *age, *sex, *d)
	//endregion

	//region 需要使用命令 go mod 下载依赖包 echo
	//EchoWeb()
	//endregion

	//region 并发 goroutine

	//go running()
	//
	////接受命令行输入，不做任何事
	//var input string
	//fmt.Scanln(&input) // 控制台输入字符立即停止go

	//var lock sync.Mutex
	//for i := 0; i < 10; i++ {
	//	go Count(&lock)
	//}
	//// 共享 Counter 变量，每次操作都需要加锁 解锁
	//for {
	//	lock.Lock()
	//	c := Counter
	//	lock.Unlock()
	//	runtime.Gosched()
	//	if c >= 10 {
	//		break
	//	}
	//}

	//锁住共享资源
	//wg.Add(2)
	//go incCounter(1)
	//go incCounter(2)
	//wg.Wait() // 等待 goroutine 结束
	//fmt.Println(counter)

	//endregion

	//region
	//wg.Add(2)
	//go incCount1()
	//go incCount1()
	//wg.Wait()
	//fmt.Println(counter) // 结果可能是2,3,4... // counter 没有同步保护，两个 goroutine 运行互相读写，结果被覆盖

	//wg.Add(2)
	//go doWork("A")
	//go doWork("B")
	//
	//time.Sleep(time.Second)
	//fmt.Println("Shutdown Now")
	//atomic.StoreInt64(&shutDown, 1)
	//wg.Wait()

	//互斥锁
	//wg.Add(2)
	//go incCounter2(1)
	//go incCounter2(2)
	//wg.Wait()
	//fmt.Println(counter)

	//endregion

	//region 等待组 sync.WaitGroup
	//fmt.Println(runtime.NumCPU())  // 16核处理器
	//waitgroup()
	//endregion

	//region 并发中的死锁
	//deadlock()
	//endregion

	//region 活锁，两个并发都在改变状态
	//livelock()
	//endregion

	//region 饥饿
	//hungry()
	//endregion

	//region 并发相关
	//orderGoroutine()
	//orderGoroutine2()
	//for i := 0; i < 10; i++ {
	//	orderGoroutine3()
	//	fmt.Println()
	//	fmt.Println()
	//}

	//orderGoroutine4()
	//endregion

	//region 反射三大定律
	//var a int
	//var b *int
	//type c struct{}
	//var d chan string
	//typeOfA := reflect.TypeOf(a)
	//typeOfB := reflect.TypeOf(b)
	//typeOfC := reflect.TypeOf(c{})
	//typeOfD := reflect.TypeOf(d)
	//// Name() 类型名称的字符串
	//// Kind() 类型归属的种类
	//fmt.Println(typeOfA.Name(), typeOfA.Kind()) // 'int' int
	//fmt.Println(typeOfB.Name(), typeOfB.Kind()) // '' ptr
	//fmt.Println(typeOfC.Name(), typeOfC.Kind()) // 'c' struct
	//fmt.Println(typeOfD.Name(), typeOfD.Kind()) // '' chan
	//
	//type Enum int
	//const (
	//	Zero Enum = 0
	//)
	//// 获取Zero常量的反射类型对象
	//typeOfZ := reflect.TypeOf(Zero)
	//// 显示反射类型对象的名称和种类
	//fmt.Println(typeOfZ.Name(), typeOfZ.Kind())  // Enum int

	//type cat struct {
	//	CatName string
	//	CatAge  int
	//	CatType int `json:"cat_type" id:"100" desc:"field desc"`
	//}
	//ins := &cat{}
	//// 获取结构体实例的 反射类型对象
	//typeOfCat := reflect.TypeOf(ins)
	//// 显示反射类型对象的名称和种类
	//fmt.Printf("name: '%v' kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind()) // '', ptr
	//// 取类型的元素
	//typeOfCat = typeOfCat.Elem() // 不可以使用一个非指正类型获取它的指针类型
	//// 显示反射类型对象的名称和种类
	//fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind()) // cat, struct
	//
	//// 获取结构体成员名称和种类
	//smallcat := cat{CatName: "mimi", CatAge: 2, CatType: 1}
	//typeOfS := reflect.TypeOf(smallcat)
	//for i := 0; i < typeOfS.NumField(); i++ {
	//	fieldType := typeOfS.Field(i)
	//	fmt.Printf("name: '%v', tag: '%v'\n", fieldType.Name, fieldType.Tag)
	//}
	//if catType, ok := typeOfCat.FieldByName("CatType"); ok{
	//	// 从tag中取出需要的tag
	//	fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	//}

	//反射第一定律：反射可以将 "接口类型变量" 转换为 "反射类型对象"
	//var x float64 = 3.4
	//fmt.Println("value:", reflect.ValueOf(x)) // valueOf 入参是一个接口类型

	//反射第二定律：反射可以将 "反射类型对象" 转换为 "接口类型变量"

	//反射第三定律：如果要修改 "反射类型对象" 其值必须是 "可写的"
	//var xx float64 =3.4
	//v := reflect.ValueOf(xx)
	//// v时不可写的，可写性 时放射类型变量的一个属性，但不是所有反射类型变量都有这个属性
	////v.SetFloat(7.1)  // panic: reflect: reflect.Value.SetFloat using unaddressable value
	//fmt.Println("settabillity of v: ", v.CanSet())  // 输出false
	//
	//// 指针p指向的数据
	//fmt.Println("使用指针可以修改 不可写 类型")
	//p := reflect.ValueOf(&xx)
	//v = p.Elem()  // Elem 对指针进行 "解引用"
	//fmt.Println("settabillity of v: ", v.CanSet())
	//v.SetFloat(7.1)
	//fmt.Println(v.Interface())
	//fmt.Println(xx)

	//结构体
	//reflectStruct()
	//endregion

	//region 反射之 IsNil IsValid
	//var a int
	//valueOfA := reflect.ValueOf(a)
	//fmt.Println(valueOfA.IsNil())  // 值类型不是通道（channel）、函数、接口、map、指针或 切片时发生 panic
	//fmt.Println(valueOfA.IsValid())

	//var b *int
	//
	//valueOfB := reflect.ValueOf(b)
	//fmt.Println(valueOfB.IsNil())  // 值类型不是通道（channel）、函数、接口、map、指针或 切片时发生 panic
	//fmt.Println(valueOfB.IsValid())

	//type T struct {
	//
	//}
	//s := struct {}{}
	//valueOfT := reflect.ValueOf(&T{})
	//fmt.Println("T.IsNil", valueOfT.IsNil()) // false
	//fmt.Println("T.IsValid", valueOfT.IsValid()) // true
	//
	//fmt.Println("s method isValid", reflect.ValueOf(s).MethodByName("").IsValid()) // false
	//endregion

	//region 反射 通过类型信息创建实例
	//var x int
	//typeOfX := reflect.TypeOf(x)
	//xx := reflect.New(typeOfX)
	//fmt.Println(xx, xx.Type(), xx.Kind())

	//endregion

	//region 通过反射 调用函数（进一步可以实现依赖注入）
	//funcValue := reflect.ValueOf(add)
	//// 构造入参
	//params := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(2)}
	////调用函数
	//ret := funcValue.Call(params)
	//fmt.Println("通过反射调用函数add 结果为：", ret[0].Int())

	//endregion

	//region 通过反射实现 依赖注入

	//inj := inject.New()
	//// 实参注入
	//inj.Map("SAM")
	//inj.Map(23)
	//inj.Map("1829200") // 当有两个以上相同类型的变量时，inject.Map 会覆盖前一个参数值，需要使用MapTo
	//inj.Map(1)
	////inj.MapTo("公司名", (*S1)(nil))
	////inj.MapTo("等级", (*S2)(nil))
	//inj.Invoke(FormatString) // 函数反转调用
	//
	//// 可以实现对struct 的注入
	//type SpecialString interface{}
	//type injStruct struct {
	//	Name   string        `inject`
	//	Gender SpecialString `inject`
	//	Age    int           `inject`
	//	Uid    int           `inject`
	//	Nick   []byte
	//}
	//
	//s := injStruct{}
	//inj = inject.New()
	//inj.Map("张三")
	//inj.MapTo("男", (*SpecialString)(nil))
	//inj.Map(18)
	//inj2 := inject.New()
	//inj2.MapTo(228392, (*int)(nil)) // todo 怎么设置uid？？
	//inj.SetParent(inj2)
	//
	//inj.Apply(&s)
	//
	//fmt.Println(s) // {张三 男 18 18 []}

	//inj.Apply(&struct)

	//endregion

	//region

	//endregion

	fmt.Println("*************************Advance END*************************")
}

// region web 服务
func Start() {
	http.HandleFunc("/", index)
	http.ListenAndServe("0.0.0.0:8081", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Wold!")
	fmt.Println("控制台：hello world！")
}

//endregion

//region 并发中的锁

// no_sync 没有并发锁时，a++ 是无序的
func no_sync() {
	a := 0
	for i := 0; i < 1000; i++ {
		go func(idx int) {
			a++
			fmt.Println(a)
		}(i)
	}
	time.Sleep(time.Second)
}

// 互斥锁会引起阻塞，多个协程抢夺资源
func have_lock() {
	a := 0
	var lock sync.Mutex // 互斥锁
	for i := 0; i < 1000; i++ {
		go func(idx int) {
			lock.Lock()
			defer lock.Unlock() // 函数释放资源时，解锁
			a++
			fmt.Printf("goroutine %d, a = %d\n", idx, a)
		}(i)
	}
	time.Sleep(time.Second)
}

func have_lock_v2() {
	ch := make(chan struct{}, 2) // 声明两个通道
	var l sync.Mutex
	go func() {
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine1: 锁定时间大概5秒")
		time.Sleep(time.Second * 5)
		fmt.Println("goroutine1: 解锁了，开始抢夺")
		ch <- struct{}{}
	}()
	go func() {
		fmt.Println("goroutine2: 等待解锁")
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine2: 我也解锁了")
		ch <- struct{}{}
	}()
	for i := 0; i < 2; i++ {
		<-ch
	}
}

//endregion

// region 读写锁，读不互斥（多个读同时），写需要互斥，读必须等写操作完成才能继续
var count int
var rw sync.RWMutex //读写锁

// ReadLock 读
func ReadLock(n int, ch chan struct{}) {
	rw.RLock()
	fmt.Printf("goroutine %d 读入读操作... \n", n)
	v := count
	fmt.Printf("gorountine %d 读取结束，值为：%d\n", n, v)
	rw.RUnlock()
	ch <- struct{}{}
}

// WriteLock 写
func WriteLock(n int, ch chan struct{}) {
	rw.Lock()
	fmt.Printf("gorountine %d 写入写操作...\n", n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("gorountine %d 写入结束，新值为：%d\n", n, v)
	rw.Unlock()
	ch <- struct{}{}
}

//endregion

// region
func usetime() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Local())
	fmt.Println(now.UTC())
	fmt.Println(now.Year())
	fmt.Println(now.UnixNano())
	//haomiao := now.UnixNano() / 1e6

	timestamp1 := now.Unix() //时间戳
	timestamp2 := now.UnixNano() / 1e9
	fmt.Println(timestamp1) // 时间戳 1678609912
	fmt.Println(timestamp2) // 时间戳 1678609912

	//将时间戳转为时间
	current_time := time.Unix(timestamp1, 0)
	fmt.Println(current_time)

}

//endregion

//region 生产者与消费者

// Consumer 消费者
func consumer(c chan int) {
	// 开始死循环
	for {
		data := <-c    // 将channel数据给data
		if data == 0 { // channel 数据为0 时，跳出死循环
			break
		}
		fmt.Println(data) // 打印消费读数据
	}
	c <- 0 // 通知生产者已经结束循环
}

//endregion

// region 玩个球
func player(name string, ch chan int) {
	defer wg.Done()
	for {
		ball, ok := <-ch
		if !ok {
			fmt.Printf("Player %s Won.\n", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 { // 随机数，判断是否丢球，关闭channel
			fmt.Printf("Player %s Missed.\n", name)
			close(ch)
			return
		}
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		ch <- ball
	}
}

//endregion

// region 跑个步
func run(baton chan int) {
	//defer wg.Done()
	var newRunner int
	runner := <-baton // 等待接力棒
	fmt.Printf("Runner %d Running With Baton\n", runner)

	// 创建下一个跑者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go run(baton)
	}

	time.Sleep(100 * time.Millisecond) // 正在跑步中

	// 最后一个人了，比赛结束
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)

	baton <- newRunner
}

//endregion

// region 超时channel
func delayChan() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case num := <-ch: // 每个case语句，必须是一个IO操作
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quit <- true
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	//_, ok := <- quit
	//if !ok{
	//	fmt.Println("结束2")
	//}
	<-quit
	fmt.Println("结束")
}

//endregion

//region

//type scoreValues []int

//func newScoreValues(vals []int, p *[]int) *scoreValues {
//	*p = vals
//	return (*scoreValues)(p)
//}
//
//
//func (s *scoreValues) Score() string {
//
//	scores, _ := strconv.Atoi(strings.Split("defalut is me", ","))
//	*s = scoreValues{scores}
//	return "It's none of my business"
//}
//
//// DefineFlagVar 定义一个flag.Var 参数类型
//func DefineFlagVar() {
//
//}

//endregion

// region echo web服务
func EchoWeb() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

//endregion

// region goroutine
func running() {
	var times int
	for { // 无限循环
		times++
		fmt.Println("tick", times)
		time.Sleep(time.Second)
	}
}

var Counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	defer lock.Unlock()
	Counter++
	fmt.Printf("Counter value: %d\n", Counter)
}

//endregion

// region 竞争状态
var (
	counter  int64
	shutDown int64
	wg       sync.WaitGroup
	mutex    sync.Mutex
)

func incCount1() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := counter
		runtime.Gosched()
		value++
		counter = value
	}
}

// 使用原子函数 加锁同步访问整型变量和指针
func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1) // 安全的队counter加1；强制同一时刻只能有一个goroutine 运行
		runtime.Gosched()
	}
}

func doWork(name string) {
	defer wg.Done()
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Microsecond)

		if atomic.LoadInt64(&shutDown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}

// 互斥锁
func incCounter2(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		// 同一时刻只允许一个 goroutine 进入这个临界区
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock() // 释放锁，允许其他正在等待的 goroutine 进入临界区
	}
}

//endregion

//region  等待组 sync.WaitGroup

func waitgroup() {
	//wg.Add(-1) == wg.Done() 两者相等
	var urls = []string{
		"http://www.github.com/",
		"https://www.baidu.com/",
		"https://www.bing.com/",
	}
	for _, url := range urls {
		wg.Add(1) // 每一个任务开始时, 将等待组增加1
		go func(url string) {
			defer wg.Done() // 使用defer, 表示函数完成时将等待组值减1
			response, err := http.Get(url)
			body := response.Body
			header := response.Header
			fmt.Println(body, "\n", header)
			fmt.Println(url, response, err)
			fmt.Println()
		}(url)
	}
	wg.Wait() // 等待所有的任务完成
	fmt.Println("结束")
}

//endregion

//region 试图运行死锁

type value struct {
	memAccess sync.Mutex
	value     int
}

func deadlock() {
	runtime.GOMAXPROCS(3)
	sum := func(v1, v2 *value) {
		defer wg.Done()
		v1.memAccess.Lock()
		fmt.Println("v1 上锁")
		time.Sleep(2 * time.Second)
		fmt.Println("v1 休眠2秒结束")
		v2.memAccess.Lock() // fatal error: all goroutines are asleep - deadlock!
		fmt.Println("v2 再次上锁")
		fmt.Printf("sum = %d\n", v1.value+v2.value)
		v1.memAccess.Unlock()
		v2.memAccess.Unlock()
	}

	product := func(v1, v2 *value) {
		defer wg.Done()
		v2.memAccess.Lock()
		fmt.Println("v2 上锁")
		time.Sleep(2 * time.Second)
		fmt.Println("v2 休眠2秒结束")
		v1.memAccess.Lock() // fatal error: all goroutines are asleep - deadlock!
		fmt.Println("v1 再次上锁")
		fmt.Printf("product = %d\n", v1.value*v2.value)
		v1.memAccess.Unlock()
		v2.memAccess.Unlock()
	}
	var v1, v2 value
	v1.value = 2
	v2.value = 3
	wg.Add(2)
	go sum(&v1, &v2)
	go product(&v1, &v2)
	wg.Wait()
}

//endregion

//region 活锁

func livelock() {
	runtime.GOMAXPROCS(3) // 设置使用3核处理器
	cv := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Second) { // 通过tick 控制两个人的步调
			cv.Broadcast()
		}
	}()

	takeStep := func() {
		cv.L.Lock()
		cv.Wait()
		cv.L.Unlock()
	}

	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, "%+v", dirName)
		atomic.AddInt32(dir, 1)
		takeStep()                      // 走上一步
		if atomic.LoadInt32(dir) == 1 { // 走成功就返回
			fmt.Fprintf(out, ". Success!")
			return true
		}
		takeStep() // 没有成功，再走回来
		atomic.AddInt32(dir, -1)
		return false
	}

	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool {
		return tryDir(" 向左走", &left, out)
	}
	tryRight := func(out *bytes.Buffer) bool {
		return tryDir(" 向右走", &right, out)
	}

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer walking.Done()
		defer func() { fmt.Println(out.String()) }()
		fmt.Fprintf(&out, "%v is trying to scoot:", name)

		for i := 0; i < 5; i++ {
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v is tried!", name)
	}
	wg.Add(2)
	go walk(&wg, "男人")
	go walk(&wg, "女人")
	wg.Wait()

}

//endregion

//region 饥饿

func hungry() {
	runtime.GOMAXPROCS(3)
	const runtime = 1 * time.Second
	var shareLock sync.Mutex

	greedyWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			shareLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			shareLock.Unlock()
			count++
		}
		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	}

	politeWorder := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			shareLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			shareLock.Unlock()

			shareLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			shareLock.Unlock()

			shareLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			shareLock.Unlock()

			count++
		}
		fmt.Printf("polite worker was able to execute %v work loops\n", count)
	}

	wg.Add(2)
	go greedyWorker()
	go politeWorder()
	wg.Wait()

}

//endregion

//region 并发相关

func orderGoroutine() {
	var mu sync.Mutex // 互斥锁
	mu.Lock()
	go func() {
		fmt.Println("打印 互斥锁")
		mu.Unlock()
	}()
	mu.Lock()
}

func orderGoroutine2() {
	ch := make(chan int)
	go func() {
		fmt.Println("打印 ch 无缓存")
		<-ch
	}()
	ch <- 1
}

func orderGoroutine3() {
	ch := make(chan int, 1)
	go func() {
		fmt.Println("打印 ch 有缓存")
		ch <- 1
		fmt.Println(1)
	}()
	fmt.Println(2)
	<-ch
	fmt.Println(3)
}

// 等待一组事件
func orderGoroutine4() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("打印")
			wg.Done()
		}()
	}
	wg.Wait()
}

//endregion

// region 反射 结构体 可写
func reflectStruct() {
	type T struct {
		A int
		B string
		//c string
	}
	//t:=T{123, "test", "小写字母c"} // 结构体中只有可导出的字段是“可设置”的
	t := T{123, "test"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(22)
	fmt.Println("t is now: ", t)
}

//endregion

// region 通过反射调用函数
func add(a, b int) int {
	return a + b
}

//endregion

//region

type S1 interface{}
type S2 interface{}

func FormatString(name string, age int, phone string, gender int) {
	//func FormatString(name string, company S1, level S2, age int)  {
	genderMap := make(map[int]string)
	genderMap[1] = "man"
	genderMap[0] = "woman"
	genderMap[9] = "Unkonwn"
	//fmt.Println(genderMap[gender])
	// 当有两个以上相同类型的变量时，inject.Map 会覆盖前一个参数值，需要使用MapTo
	fmt.Printf("My name is %s, %d years old, gender: %s, phone: %s\n", name, age, genderMap[gender], phone)
	//fmt.Printf("My name is %s, %d years old\n", name, age)
	//fmt.Printf("company %s, leve %s\n", company, level)
}

//endregion

//region

//endregion

//region

//endregion
