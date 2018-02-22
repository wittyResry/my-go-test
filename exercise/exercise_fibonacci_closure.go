package exercise

import "fmt"

func fibonacci() func() int {
    var s[]int
    return func() int {
        if len(s) == 0 {
            s = append(s, 0)
        } else if len(s) == 1 {
            s = append(s, 1)
        } else {
            s = append(s, s[len(s)-1] + s[len(s)-2])
        }
        return s[len(s)-1]
    }
}

func MainFibonacciClosure() {
    f := fibonacci()
    fmt.Println()
    for i := 0; i < 10; i++ {
        fmt.Printf(" %d", f())
    }
    fmt.Println()
}
