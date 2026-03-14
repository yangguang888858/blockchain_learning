package main

import (
	"fmt"
)

//3 面向对象

// 3.2 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者。

// 创建Person结构体
type Person struct {
	Name string
	Age  int
}

// 创建Employee结构体
type Employee struct {
	Person
	EmployeeID string
}

// 定义PrintInfo()方法
func (e *Employee) PrintInfo() {
	fmt.Printf("员工的name为%v,age为%v,employeeID为%v", e.Name, e.Age, e.EmployeeID)
}

func main() {
	e := &Employee{Person: Person{Name: "tom", Age: 18}, EmployeeID: "001"}
	e.PrintInfo()
}
