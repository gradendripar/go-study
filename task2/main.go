package main

import (
	"fmt"
)

func main() {
	// 测试题目1：整数指针
	num := 5
	fmt.Println("原始值:", num)
	AddTen(&num)
	fmt.Println("增加10后:", num)

	// 测试题目2：切片指针
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片:", numbers)
	MultiplyByTwo(&numbers)
	fmt.Println("每个元素乘以2后:", numbers)

	// 测试协程相关功能
	fmt.Println("\n===== 协程测试 =====")
	TestGoroutines()

	// 测试接口相关功能
	fmt.Println("\n===== 接口测试 =====")
	TestShapes()

	// 测试组合相关功能
	fmt.Println("\n===== 组合测试 =====")
	TestEmployee()

	// 测试通道相关功能
	fmt.Println("\n===== 通道测试 =====")
	TestChannels()

	// 测试并发安全相关功能
	fmt.Println("\n===== 并发安全测试 =====")
	TestConcurrencySafety()
}
