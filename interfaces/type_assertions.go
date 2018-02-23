package interfaces

import "fmt"

func MainTypeAssertion() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)
	s, ok := i.(string)
	fmt.Println(s, ok)
	s2, ok2 := i.(float64)
	fmt.Println(s2, ok2)
	//f2 = i.(float64) // panic: interface conversion: interface {} is string, not float64
}
