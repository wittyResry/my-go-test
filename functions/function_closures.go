package functions

import "fmt"

//返回一个函数值，它可以引用其函数体之外的变量。每个闭包都被绑定在各自的sum变量上。不会受到其他的影响。
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func MainFunctionClosures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d %d,", pos(i), neg(-i))
	}
}
