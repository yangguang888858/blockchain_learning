package main

import "fmt"

//1 指针

// 1.1 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
//引用传递
//传递的实参的地址,会改变原始变量
func referencePass(num *int) {
	*num += 10
}

//值传递
//传递的实参的副本,不会改变原始变量
func valuePass(num int) {
	num += 10
}
func main() {
	num := 100
	fmt.Println("num原始值=", num)
	valuePass(num)
	fmt.Println("值传递后num的值=", num)
	referencePass(&num)
	fmt.Println("引用传递后num的值=", num)
}
