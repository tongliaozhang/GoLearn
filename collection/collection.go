// collection project collection.go
package collection

import (
	"GoLearn/util"
	"fmt"
)

func Array_Test() {
	//var a1 [10]int
	//fmt.Println(a1)

	//a2 := [10]int{1, 2, 3, 4, 5}
	//a3 := [...]int{1, 2, 3, 4}
	//a4 := []int{1, 2, 3}
	//a5 := [5]int{1: 1, 3: 3}

	//fmt.Println(a2)
	//fmt.Println(a3)
	//fmt.Println(a4)
	//fmt.Println(a5)

	//数组指针类型 *[n]T，指针数组类型 [n]*T。
	//*表示指针类型 &取变量的地址 通过指针类型的*取到变量的值
	//x, y := 1, 2
	//var p1 *[2]int = &[2]int{x, y}

	//var p2 [2]*int = [2]*int{&x, &y}

	//fmt.Println(*p2[0])
	//fmt.Println(*p2[1])

	//fmt.Println(p1)
	//fmt.Println(p2)
	//fmt.Printf("%#v\n", p1)
	//fmt.Printf("%#v\n", p2)

	x1 := [3]int{1, 2, 3}

	fmt.Println(x1)

	x2 := x1

	x2[0] = 100

	fmt.Println(x2)
	fmt.Println(x1)

	util.FG_line()
}

//参数是数组的slices或者指针 改变这个对象同时改变原数组内容
func Pointer_Array(x *[4]int) {
	for i := 0; i < len(x); i++ {
		println(x[i])
	}

	x[3] = 300

	util.FG_line()
}

func Slicen_Array(x []int) {
	for _, v := range x {
		println(v)
	}

	x[3] = 999

	util.FG_line()
}
