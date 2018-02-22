package arrays

import "fmt"

func MainRangeContinue() {
    pow := make([]int, 10)
    for i := range pow {
        pow[i] = 1 << uint(i)
    }
    for _, value := range pow {
        fmt.Printf("%d,", value)
    }
}
