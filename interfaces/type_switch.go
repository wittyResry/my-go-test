package interfaces

import "fmt"

func do(i interface{}) {

	switch v := i.(type) {
	case int:
		fmt.Printf("int:%T\n", v)
	default:
		fmt.Printf("unknow:%T\n", v)
	}
}

func MainTypeSwitch() {
	do(21)
	do(true)
}
