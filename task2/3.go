package main

import (
	"fmt"
	"math"
)

// 1. 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
// 2. 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者。

// Shape 接口定义
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle的Area方法实现
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Rectangle的Perimeter方法实现
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 结构体
type Circle struct {
	Radius float64
}

// Circle的Area方法实现
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circle的Perimeter方法实现
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Person 结构体
type Person struct {
	Name string
	Age  int
}

// Employee 结构体，组合Person
type Employee struct {
	Person     // 匿名字段，表示组合
	EmployeeID string
}

// PrintInfo 方法，输出员工信息
func (e Employee) PrintInfo() {
	fmt.Printf("员工ID: %s\n", e.EmployeeID)
	fmt.Printf("姓名: %s\n", e.Name) // 可以直接访问Person的字段
	fmt.Printf("年龄: %d\n", e.Age)  // 可以直接访问Person的字段
}

// TestShapes 测试Shape接口实现
func TestShapes() {
	// 创建Rectangle实例
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Println("矩形:")
	fmt.Printf("  面积: %.2f\n", rect.Area())
	fmt.Printf("  周长: %.2f\n", rect.Perimeter())

	// 创建Circle实例
	circle := Circle{Radius: 2}
	fmt.Println("圆形:")
	fmt.Printf("  面积: %.2f\n", circle.Area())
	fmt.Printf("  周长: %.2f\n", circle.Perimeter())

	// 使用接口类型的变量
	fmt.Println("\n通过接口调用:")
	shapes := []Shape{rect, circle}
	for i, shape := range shapes {
		if i == 0 {
			fmt.Println("矩形(通过接口):")
		} else {
			fmt.Println("圆形(通过接口):")
		}
		fmt.Printf("  面积: %.2f\n", shape.Area())
		fmt.Printf("  周长: %.2f\n", shape.Perimeter())
	}
}

// TestEmployee 测试组合功能
func TestEmployee() {
	// 创建Employee实例
	emp := Employee{
		Person: Person{
			Name: "张三",
			Age:  30,
		},
		EmployeeID: "EMP001",
	}

	// 调用PrintInfo方法
	fmt.Println("\n员工信息:")
	emp.PrintInfo()

	// 直接访问组合的字段
	fmt.Println("\n直接访问组合字段:")
	fmt.Printf("姓名: %s\n", emp.Name)
	fmt.Printf("年龄: %d\n", emp.Age)
	fmt.Printf("员工ID: %s\n", emp.EmployeeID)
}
