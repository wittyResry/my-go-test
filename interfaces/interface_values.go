package interfaces

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

//impl T
func (t *T) M() {
	if t == nil {
		fmt.Println("t is nil")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i interface{}) {
	fmt.Printf("(%v,%T)\n", i, i)
}

func MainInterfaceValues() {
	var i I
	describe(i)
	if i != nil {
		i.M() // 要求判断空
	}

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	var t *T
	i = t
	describe(i)
	i.M()

	var inter interface{}
	describe(inter)

}
