package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
// 2. 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。

// 打印奇数的函数
func printOdd(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("奇数: %d\n", i)
		time.Sleep(100 * time.Millisecond) // 让输出更易于观察
	}
}

// 打印偶数的函数
func printEven(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("偶数: %d\n", i)
		time.Sleep(150 * time.Millisecond) // 让输出更易于观察
	}
}

// 任务类型定义
type Task func() string

// 任务执行结果
type TaskResult struct {
	Name     string
	Result   string
	Duration time.Duration
}

// 任务调度器
func taskScheduler(tasks map[string]Task) []TaskResult {
	results := make([]TaskResult, 0, len(tasks))
	resultChan := make(chan TaskResult, len(tasks))

	// 并发执行所有任务
	for name, task := range tasks {
		go func(taskName string, taskFunc Task) {
			startTime := time.Now()
			result := taskFunc()
			duration := time.Since(startTime)
			resultChan <- TaskResult{
				Name:     taskName,
				Result:   result,
				Duration: duration,
			}
		}(name, task)
	}

	// 收集所有任务的结果
	for i := 0; i < len(tasks); i++ {
		results = append(results, <-resultChan)
	}

	return results
}

// TestGoroutines 测试协程相关功能
func TestGoroutines() {
	fmt.Println("=== 测试协程并发打印奇偶数 ===")
	var wg sync.WaitGroup
	wg.Add(2)

	// 启动两个协程
	go printOdd(&wg)
	go printEven(&wg)

	// 等待两个协程完成
	wg.Wait()
	fmt.Println("奇偶数打印完成")

	fmt.Println("\n=== 测试任务调度器 ===")
	// 定义一些测试任务
	tasks := map[string]Task{
		"任务1": func() string {
			time.Sleep(500 * time.Millisecond)
			return "任务1完成"
		},
		"任务2": func() string {
			time.Sleep(300 * time.Millisecond)
			return "任务2完成"
		},
		"任务3": func() string {
			time.Sleep(800 * time.Millisecond)
			return "任务3完成"
		},
	}

	// 执行任务调度
	results := taskScheduler(tasks)

	// 打印结果
	fmt.Println("任务执行结果:")
	for _, result := range results {
		fmt.Printf("- %s: %s (耗时: %v)\n", result.Name, result.Result, result.Duration)
	}
}
