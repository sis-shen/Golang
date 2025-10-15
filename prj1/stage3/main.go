package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func task() {
	time.Sleep(1 * time.Second)
	fmt.Println("I am task")
}

var g_sum = 0
var g_mutex = sync.Mutex{}

func add(num int) {
	g_sum += num
}

func con_add(num int) {
	g_mutex.Lock()
	defer g_mutex.Unlock()
	g_sum += num
}

//func main() {
//	//ch := make(chan string)
//	//go func() {
//	//	task()
//	//	ch <- "task done"
//	//}()
//	//fmt.Println("I am main")
//	//msg := <-ch
//	//println(msg)
//	//time.Sleep(5 * time.Second)
//	cond := sync.NewCond(&sync.Mutex{})
//	loopRound := 10000
//	wg := sync.WaitGroup{}
//	wg.Add(loopRound)
//	for i := 0; i < loopRound; i++ {
//		go func() {
//			defer wg.Done()
//			cond.L.Lock()
//			cond.Wait()
//			con_add(1000)
//			cond.L.Unlock()
//		}()
//	}
//	wg.Wait()
//	fmt.Println(g_sum)
//}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, stop := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		watchDog(ctx, "【看门狗】")
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("发送真的指令")
	stop()
	wg.Wait()
}

func watchDog(ctx context.Context, name string) {
	// 开启 for select 循环
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "停止指令已收到，马上停止")
			return
		default:
			fmt.Println(name, "正在监控")
		}
		time.Sleep(time.Second)
	}
}
