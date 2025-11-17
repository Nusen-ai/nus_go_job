package task2

import "fmt"

// 6. 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。

// Person 结构体
type Person struct {
	Name string
	Age  int
}

// Employee 结构体，组合了Person
type Employee struct {
	Person     // 匿名嵌入，实现组合
	EmployeeID string
}

// Employee 的 PrintInfo 方法
func (e Employee) PrintInfo() {
	fmt.Printf("员工信息:\n")
	fmt.Printf("  姓名: %s\n", e.Name) // 直接访问Person的字段
	fmt.Printf("  年龄: %d\n", e.Age)  // 直接访问Person的字段
	fmt.Printf("  员工ID: %s\n", e.EmployeeID)
}

func Job6TestDemo() {

	// 创建Employee实例的几种方式

	// 方式1：分别初始化
	emp1 := Employee{
		Person: Person{
			Name: "张三",
			Age:  28,
		},
		EmployeeID: "E2024001",
	}

	// 方式2：直接初始化
	emp2 := Employee{
		Person:     Person{"李四", 32},
		EmployeeID: "E2024002",
	}

	// 方式3：先创建后赋值
	var emp3 Employee
	emp3.Name = "王五" // 可以直接访问嵌入结构体的字段
	emp3.Age = 25
	emp3.EmployeeID = "E2024003"

	fmt.Println("=== 员工信息输出 ===")

	// 调用PrintInfo方法
	emp1.PrintInfo()
	fmt.Println()
	emp2.PrintInfo()
	fmt.Println()
	emp3.PrintInfo()

	// 演示组合的其他特性
	fmt.Println("\n=== 组合特性演示 ===")

	// 可以直接访问嵌入结构体的字段
	fmt.Printf("员工1的姓名: %s, 年龄: %d\n", emp1.Name, emp1.Age)

	// 也可以通过完整路径访问
	fmt.Printf("员工1的姓名(完整路径): %s\n", emp1.Person.Name)
}
