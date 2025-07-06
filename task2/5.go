package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。

// Counter 结构体，使用互斥锁保护计数值
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment 方法，对计数器进行递增操作
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value 方法，获取计数器的值
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// 测试互斥锁保护的计数器
func testMutexCounter() {
	fmt.Println("=== 使用互斥锁的计数器 ===")

	// 创建计数器
	counter := Counter{}

	// 创建等待组
	var wg sync.WaitGroup

	// 启动10个协程
	numGoroutines := 10
	numIncrements := 1000
	wg.Add(numGoroutines)

	startTime := time.Now()

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()

			for j := 0; j < numIncrements; j++ {
				counter.Increment()
			}

			fmt.Printf("协程 %d 完成\n", id)
		}(i)
	}

	// 等待所有协程完成
	wg.Wait()

	elapsedTime := time.Since(startTime)

	// 输出计数器的值
	fmt.Printf("预期值: %d, 实际值: %d\n", numGoroutines*numIncrements, counter.Value())
	fmt.Printf("执行时间: %v\n", elapsedTime)
}

// 测试使用原子操作的计数器
func testAtomicCounter() {
	fmt.Println("\n=== 使用原子操作的计数器 ===")

	// 创建原子计数器
	var atomicCounter int64 = 0

	// 创建等待组
	var wg sync.WaitGroup

	// 启动10个协程
	numGoroutines := 10
	numIncrements := 1000
	wg.Add(numGoroutines)

	startTime := time.Now()

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()

			for j := 0; j < numIncrements; j++ {
				// 使用原子操作递增计数器
				atomic.AddInt64(&atomicCounter, 1)
			}

			fmt.Printf("原子操作协程 %d 完成\n", id)
		}(i)
	}

	// 等待所有协程完成
	wg.Wait()

	elapsedTime := time.Since(startTime)

	// 输出计数器的值
	fmt.Printf("预期值: %d, 实际值: %d\n", numGoroutines*numIncrements, atomicCounter)
	fmt.Printf("执行时间: %v\n", elapsedTime)
}

// TestConcurrencySafety 测试并发安全相关功能
func TestConcurrencySafety() {
	// 测试互斥锁保护的计数器
	testMutexCounter()

	// 测试原子操作的计数器
	testAtomicCounter()

	// 比较两种方法的结果
	fmt.Println("\n两种方法都能确保并发安全，但原子操作通常比互斥锁更高效，因为它不需要获取和释放锁的开销。")
}
