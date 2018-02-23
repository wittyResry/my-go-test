package interfaces

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type myFloat float64

func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type vertex struct {
	X, Y float64
}

func (v *vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func MainInterface() {
	var a Abser
	f := myFloat(-math.Sqrt2)
	a = f
	fmt.Println(a.Abs())

	v := vertex{4, 2}
	a = &v
	fmt.Println(a.Abs())
	fmt.Println(v.Abs()) // 而以指针为接收者的方法被调用时，接收者既能为值又能为指针
	fmt.Println((&v).Abs())
}
