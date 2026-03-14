package main

import "fmt"

//1 指针

//1.2 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
//考察点 ：指针运算、切片操作。
func multiply(pointerSlice *[]int) {
	for k, v := range *pointerSlice {
		(*pointerSlice)[k] = v * 2
	}
	// for i := 0; i < len(*pointerSlice); i++ {
	// 	(*pointerSlice)[i] *= 2
	// }
}
func main() {
	var pointerSlice = []int{1, 2, 3}
	fmt.Println("pointerSlice原始值=", pointerSlice)
	multiply(&pointerSlice)
	fmt.Println("pointerSlice修改后的值=", pointerSlice)
}
