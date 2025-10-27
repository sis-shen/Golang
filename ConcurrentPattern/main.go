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

func perchase(ctx context.Context, inCh <-chan int, outCh chan int) {
	defer fmt.Printf("perchase, done\n")
	defer close(outCh)
	for {
		select {
		case <-ctx.Done():
			return
		case num, ok := <-inCh:
			if !ok {
				return
			}
			fmt.Printf("inCh: 采购了%d\n", num)
			select {
			case <-ctx.Done():
				return
			case outCh <- num:
			}
		}
	}

}

func machine(ctx context.Context, inCh <-chan int, outCh chan int) {
	defer fmt.Printf("machine, done\n")
	defer close(outCh)

	for {
		select {
		case <-ctx.Done():
			return
		case num := <-inCh:
			fmt.Printf("machine:收到加工任务%d\n", num)
			for i := 0; i < num; i++ {
				fmt.Printf("machine:正在加工一个产品，编号%d\n", i)
				select {
				case <-ctx.Done():
					return
				case outCh <- 1:
				}
				time.Sleep(1 * time.Second)
			}
			fmt.Printf("inCh: 生产了%d\n", num)
		}
	}
}

func sell(ctx context.Context, inCh <-chan int) {
	defer fmt.Printf("sell, done\n")
	for {
		select {
		case <-ctx.Done():
			return
		case num := <-inCh:
			fmt.Printf("sell:获取到销售任务%d\n", num)
			for i := 0; i < num; i++ {
				fmt.Printf("sell:卖出去了一个产品，编号%d\n", i)
				time.Sleep(1 * time.Second)
			}
			fmt.Printf("本批次任务完成，共计产品%d\n", num)
		}

	}
}

func test_select_pattern() {
	fmt.Println("----------- test select_pattern --------------\n")
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
}

func test_select_range_pattern() {
	fmt.Println("-------test select_range_pattern --------------\n")
	wg := sync.WaitGroup{}
	wg = sync.WaitGroup{}
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		select_range_pattern(ctx, []int{1, 2, 3, 4, 5})
	}()
	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()
}

func test_time_pattern() {
	fmt.Println("----------- test time_pattern --------------\n")
	wg := sync.WaitGroup{}
	wg = sync.WaitGroup{}
	wg.Add(1)
	ctx, _ := context.WithCancel(context.Background())
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

func merge(ins ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}

	ch_merge := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}

	wg.Add(len(ins))

	for _, in := range ins {
		go ch_merge(in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func test_pipeline() {
	fmt.Println("----------- test pipeline --------------\n")
	//初始化变量
	inPerchaseCh := make(chan int)
	outPerchaseCh1 := make(chan int)
	outPerchaseCh2 := make(chan int)
	outPerchaseCh3 := make(chan int)
	outMachineCh := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(5)

	//启动管线
	go func() {
		defer wg.Done()
		perchase(ctx, inPerchaseCh, outPerchaseCh1)
	}()

	go func() {
		defer wg.Done()
		perchase(ctx, inPerchaseCh, outPerchaseCh2)
	}()

	go func() {
		defer wg.Done()
		perchase(ctx, inPerchaseCh, outPerchaseCh3)
	}()

	outPerchaseCh := merge(outPerchaseCh1, outPerchaseCh2, outPerchaseCh3)

	go func() {
		defer wg.Done()
		machine(ctx, outPerchaseCh, outMachineCh)
	}()

	go func() {
		defer wg.Done()
		sell(ctx, outMachineCh)
	}()

	//开始输入
	inPerchaseCh <- 10
	time.Sleep(5 * time.Second)
	//inPerchaseCh <- 20

	time.Sleep(20 * time.Second)
	//销毁管线
	close(inPerchaseCh)
	cancel()
	cancel()
	cancel()
	wg.Wait()

}

func boilWater() <-chan string {
	waterCh := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("烧水中...%d秒\n", i)
			time.Sleep(1 * time.Second)
		}
		waterCh <- "水烧好了\n"
	}()
	return waterCh
}

func cleanTable() <-chan string {
	waterCh := make(chan string)
	go func() {
		for i := 0; i < 8; i++ {
			fmt.Printf("擦桌子中...%d秒\n", i)
			time.Sleep(1 * time.Second)
		}
		waterCh <- "桌子擦好了\n"
	}()
	return waterCh
}

func test_future() {
	waterCh := boilWater()
	tableCh := cleanTable()

	water := <-waterCh
	fmt.Printf(water)
	table := <-tableCh
	fmt.Printf(table)
	fmt.Printf("Party Start!!!\n")
}

func main() {
	//test_select_pattern()
	//test_select_range_pattern()
	//test_time_pattern()
	//test_pipeline()
	test_future()
}
