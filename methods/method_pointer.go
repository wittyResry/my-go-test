package methods

import (
    "fmt"
)

type Vertex2 struct {
    X, Y float64
}

//无返回值函数
func (v * Vertex2) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func (v Vertex2) Scale2(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func MainPorinter() {
    v := Vertex2{3, 4}
    v.Scale(10)
    fmt.Println(v)
    v.Scale2(10)
    fmt.Println(v)
}
