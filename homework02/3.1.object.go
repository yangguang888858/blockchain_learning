package main

import (
	"fmt"
	"math"
)

//3 面向对象

// 3.1 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。

// 定义Shape接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义Rectangle结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle结构体实现Area()方法
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Rectangle结构体实现Perimeter()方法
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 定义Circle结构体
type Circle struct {
	Radius float64
}

// Circle结构体实现Area()方法
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circle结构体实现Perimeter()方法
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	r := &Rectangle{Width: 1.0, Height: 2.0}
	fmt.Printf("矩形的面积为%.2f,周长为%.2f\n", r.Area(), r.Perimeter())

	c := &Circle{Radius: 1.0}
	fmt.Printf("圆的面积为%.2f,周长为%.2f\n", c.Area(), c.Perimeter())
}
