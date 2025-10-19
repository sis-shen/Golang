package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func normal_task(idx int) {
	fmt.Printf("normal_task(%d)\n", idx)
}

func select_pattern(ctx context.Context) {
	cnt := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("done\n")
			return
		default:
			cnt++
			cnt %= 100
			normal_task(cnt)
			time.Sleep(1 * time.Second)
		}
	}
}

func select_range_pattern(ctx context.Context, arr []int) {
	for _, v := range arr {
		select {
		case <-ctx.Done():
			fmt.Printf("done\n")
			return
		default:
			normal_task(v)
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Printf("select_range_pattern finished\n")
}

func time_pattern(ctx context.Context, ch chan int) {
	select {
	case <-ctx.Done():
		fmt.Printf("canceled\n")
		return
	case a := <-ch:
		fmt.Printf("recv msg: %d\n", a)
		return
	case <-time.After(2 * time.Second):
		fmt.Printf("timeout 2s\n")
		return
	}
}

func perchase(ctx context.Context, inCh chan int, outCh chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("perchase, done\n")
			return
		default:

		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		select_pattern(ctx)
	}()
	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()

	fmt.Println("-------stage2------------")
	wg = sync.WaitGroup{}
	wg.Add(1)
	ctx, cancel = context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		select_range_pattern(ctx, []int{1, 2, 3, 4, 5})
	}()
	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()

	fmt.Println("-------stage3------------")
	wg = sync.WaitGroup{}
	wg.Add(1)
	ctx, cancel = context.WithCancel(context.Background())
	ch := make(chan int)
	go func() {
		fmt.Println("发起了一个请求:")
		time.Sleep(5 * time.Second)
		ch <- 666
	}()

	go func() {
		defer wg.Done()
		time_pattern(ctx, ch)
	}()
	wg.Wait()
}
