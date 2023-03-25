package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
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

	//region

	//endregion

	fmt.Println("*************************Advance END*************************")
}

var wg sync.WaitGroup

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

//endregion

//region

//endregion
