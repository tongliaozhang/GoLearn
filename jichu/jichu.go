// jichu project jichu.go
package jichu

import (
	"fmt"
	"strings"
)

func BianLiang() {
	var a = 1234
	var b string = "string"
	var c bool = true

	fmt.Println(a, b, c)

	var x, y int

	x = 3
	y = 4

	m := 5
	n := 6

	fmt.Println(x, y, m, n)

	const cvar int = 100
	fmt.Println(cvar)

	type ByteSize int64
	const (
		_           = iota
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
	)

	fmt.Println(GB, GB/1024/1024/1024)

	fmt.Println(`test\nstr`)
	fmt.Println("test\nstr")

	a1, b1, c1 := "a", "b", "c"
	s := strings.Join([]string{a1, b1, c1}, "")
	println(s)
}

func StringLearn() {
	a := "abc"
	b := []byte(a)
	b[2] = 'B'
	fmt.Println(string(b))

	c := []rune(a)
	c[2] = '王'
	fmt.Println(string(c))
}

type Student struct {
	Id   int
	Name string
}

func Pointer() {
	a := 100
	b := &a
	fmt.Println(*b)

	c := Student{1, "zhangql"}
	//d = c的指针
	d := &c
	fmt.Println("d=", d.Id)
	fmt.Println("d=" + d.Name)

	//e = d的值
	e := *d
	e.Id = 2
	e.Name = "shaolize"
	fmt.Println("e=", e.Id)
	fmt.Println("e=" + e.Name)

	//f = c
	f := c
	f.Id = 3
	f.Name = "xiaoyouyou"
	fmt.Println("f=", f.Id)
	fmt.Println("f=" + f.Name)
}
