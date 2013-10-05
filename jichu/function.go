package jichu

import (
	"fmt"
)

type callback func(s string)

func Test(a, b int, sum func(int, int) int) {
	println(sum(a, b))
}

func FuncTest() {
	var cb callback = func(s string) {
		println(s)
	}

	cb("Hello Zhangql")

	Test(1, 2, func(a, b int) int { return a + b })

}

func Swap(a, b int) (int, int) {
	return b, a
}

func Change(a, b int) (x, y int) {
	x = a * b
	y = a + b

	return x, y
}

func MagicAgrs(s string, args ...int) {
	defer func() {
		println(s + " in defer func")
	}()

	var x int
	for _, n := range args {
		x += n
	}
	println(s, x)

	s = "defer var s"
}

/**闭包**/
func Closures(x int) func(int) int {
	// 返回匿名函数
	return func(y int) int {
		return x + y
	}
}

func Bibao_func() {
	var fs []func() int //存放匿名函数的数组

	for i := 0; i < 3; i++ {
		fs = append(fs, func() int {
			return i
		})
	}

	for inx, f := range fs {
		fmt.Printf("%d %p = %v\n", inx, f, f())
	}
}

func FILO_defer() {
	defer func() {
		println("defer code 1")
	}()

	defer func() {
		println("defer code 2")
	}()
}

func Defer_copy() {
	x := 10

	defer func(x int) {
		fmt.Printf("defer 1 = %d", x)
	}(x)

	x += 100

	fmt.Printf("func run = %d", x)
}

func TryCatch_Go(m, n int) (t int) {
	//var fs = func(x, y int) (z int) {
	//	z = x / y
	//	return z
	//}
	//t = fs(m, n)
	defer func() {
		if err := recover(); err != nil {
			println(err)
		}
	}()

	t = m / n

	return t
}
