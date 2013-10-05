// GoLearn project main.go
package main

import (
	//"GoLearn/collection"
	//"GoLearn/jichu"
	"GoLearn/channelLearn"
	"GoLearn/structTest"
	"fmt"
)

func main() {
	fmt.Println("===Go Begin ===")
	//jichu.FuncTest()
	//jichu.BianLiang()
	//jichu.StringLearn()
	//jichu.Pointer(
	//jichu.IF_Test(0)
	//jichu.IF_Test(100)
	//jichu.IF_Test(-200)
	//jichu.Range_Test()
	//a := 1
	//jichu.Switch_test(a)
	//b := 'b'
	//jichu.Switch_test(b)
	//c := true
	//jichu.Switch_test(c)

	//a, b := jichu.Swap(1, 2)
	//c, _ := jichu.Change(100, 20)
	//println(a, b, c)

	//jichu.MagicAgrs("1+2+3=", 1, 2, 3)

	//x := []int{0, 1, 2, 3, 4, 5}

	//jichu.MagicAgrs("sum = ", x[:3]...)

	//f := jichu.Closures(1)
	//println(f(1))
	//println(f(2))

	//jichu.Bibao_func()

	//jichu.FILO_defer()

	//jichu.Defer_copy()

	//rs := jichu.TryCatch_Go(100, 100)

	//defer func() {
	//	if err := recover(); err != nil {
	//		println("Panic AAA")
	//	} else {
	//		println("no panic")
	//	}
	//}()
	//println(rs)
	//println("chidiaole")

	//collection.Array_Test()

	//x := &[4]int{1, 2, 3, 4}

	//collection.Pointer_Array(x)

	//fmt.Println(x)

	//y := [6]int{100, 200, 300, 400, 500, 600}
	//z := y[:4]

	//collection.Slicen_Array(z)

	//collection.Slice_Create()

	//collection.Slice_append()

	//collection.Map_Create()

	//collection.Struct_Map()

	//collection.Map_panduan()

	//collection.Map_valueCopy()

	//structTest.User_Test()

	//structTest.PointStruct()

	//structTest.AnonsStructInit()

	//structTest.Init_Test()

	structTest.Method_struct()

	channelLearn.Channel_One()
}
