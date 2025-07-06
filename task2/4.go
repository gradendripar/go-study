package main

import (
	"fmt"
	"sync"
	"time"
)

// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。

// 生成1到10的整数并发送到通道
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i // 发送数据到通道
		fmt.Printf("生产者: 发送 %d 到通道\n", i)
		time.Sleep(100 * time.Millisecond) // 让输出更易于观察
	}
	close(ch) // 关闭通道，表示不再发送数据
	fmt.Println("生产者: 通道已关闭")
}

// 从通道接收整数并打印
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// 使用for range循环从通道接收数据，直到通道关闭
	for num := range ch {
		fmt.Printf("消费者: 接收到 %d\n", num)
		time.Sleep(200 * time.Millisecond) // 让输出更易于观察
	}
	fmt.Println("消费者: 接收完成")
}

// 使用带缓冲的通道
func bufferedChannelDemo() {
	fmt.Println("\n=== 带缓冲的通道示例 ===")

	// 创建一个容量为10的缓冲通道
	bufCh := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(2)

	// 生产者协程
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			bufCh <- i
			fmt.Printf("缓冲生产者: 发送 %d\n", i)
			// 生产者速度较快
			if i%10 == 0 {
				fmt.Printf("已发送 %d 个数据\n", i)
			}
		}
		close(bufCh)
		fmt.Println("缓冲生产者: 通道已关闭")
	}()

	// 消费者协程
	go func() {
		defer wg.Done()
		count := 0
		for num := range bufCh {
			count++
			fmt.Printf("缓冲消费者: 接收到 %d\n", num)
			time.Sleep(20 * time.Millisecond) // 消费者处理速度较慢
			if count%10 == 0 {
				fmt.Printf("已接收 %d 个数据\n", count)
			}
		}
		fmt.Println("缓冲消费者: 接收完成")
	}()

	wg.Wait()
}

// TestChannels 测试通道相关功能
func TestChannels() {
	fmt.Println("=== 基本通道通信示例 ===")

	// 创建一个无缓冲的通道
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	// 启动生产者和消费者协程
	go producer(ch, &wg)
	go consumer(ch, &wg)

	// 等待两个协程完成
	wg.Wait()

	// 测试带缓冲的通道
	bufferedChannelDemo()
}
