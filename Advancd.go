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
	usetime()
	//endregion

	//region

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

//region

//endregion
