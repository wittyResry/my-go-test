# https://tour.golang.org/basics/1

# package包
* 每个Go都是由包构成。程序从main包开始运行。package.go通过导入包路径"fmt"和"math/rand"来使用这 两个包
按照约定，包名与导入路径最后一个元素一致。例如`math/rand`包中的源码都是以package main语句开始的

# import规则
* 导出名在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。 例如， Pi 就是个已导出名，它导出自 math 包。pizza 和 pi 并未以大写字母开头，所以它们是未导出的。在导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问。执行代码，观察错误输出。

# function函数
* 函数可以接收多个参数，类型定义在变量名之后。当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。`x int, y int`被缩写为`x, y int`

# 多值返回
* 函数可以返回任意数量的返回值。支持结构体参数透传

# 命名返回值
* Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。返回值的名称应当具有一定的意义，它可以作为文档使用。没有参数的 return 语句返回已命名的返回值。也就是`直接`返回。

# 变量和变量的初始化
* var 语句可以出现在包或函数级别。变量声明可以包含初始值，如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。

# 短变量的声明
* 在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明。函数外的每个语句都必须以关键字开始（ var 、 func 等等）， 因此 := 结构不能在函数外使用。没有明确初始值的变量声明会被赋予它们的值 。常量不能用 := 语法声明。

# switch语句和没有条件的 switch
* Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，不过 Go 只运行选定的 case，而非之后所有的 case。 实际上，Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。 除非以 fallthrough 语句结束，否则分支会自动终止。 Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。
* 没有条件的 switch 同 switch true 一样。这种形式能将一长串 if-then-else 写得更加清晰。

# defer和defer栈
* defer 语句会将函数推迟到外层函数返回之后执行。推迟的函数调用会被压入一个栈中。 当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。

# GO指针
* Go 拥有指针。 指针保存了值的内存地址。类型 *T 是指向 T 类型值的指针。其零值为 nil 。& 操作符会生成一个指向其操作数的指针。这也就是通常所说的“间接引用”或“重定向”。

# 结构体
* 初始化{1, 2}，结构体字段使用点号来访问。结构体指针：如果我们有一个指向结构体的指针 p ，那么可以通过 (*p).X 来访问其字段 X 。

# 数组
* 数组的长度是其类型的一部分，因此数组不能改变大小。每个数组的大小都是固定的。 而切片则为数组元素提供动态大小的、灵活的视角。 在实践中，切片比数组更常用。切片就像数组的引用，更改切片的元素会修改其底层数组中对应的元素。在进行切片时，你可以利用它的默认行为来忽略上下界。切片下界的默认值为 0 ，上界则是该切片的长度。用 make 创建切片make([]int, 5)



