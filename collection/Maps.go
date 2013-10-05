package collection

import (
	"fmt"
)

func Map_Create() {
	var a = map[string]int{"1": 3}
	b := map[int]string{1: "a", 2: "b"}

	fmt.Println(a)
	fmt.Println(b)

	test(b)
	fmt.Println(b)

	c := make(map[string]string)
	c["abc"] = "def"
	fmt.Println(c)

}

func test(a map[int]string) {
	a[3] = "x"
	a[2] = "x"
}

type User struct {
	Name string
}

func Struct_Map() {
	a := [2]int{0, 1}
	b := [2]int{0, 1}
	d := map[[2]int]string{a: "ssss"}
	fmt.Println(d, d[b])
	u := User{"User1"}
	u2 := u
	f := map[User]string{u: "tttt"}
	fmt.Println(f, f[u2])
}

func Map_panduan() {
	a := map[string]string{"1": "A", "2": "B", "3": "C", "4": "D"}

	v, ok := a["1"]
	fmt.Println(v, ok)

	v1, ok := a["5"]
	fmt.Println(v1, ok)

	for k, v := range a {
		fmt.Println(k, v)
	}

}

func Map_valueCopy() {
	users := map[string]User{
		"1": User{"user1"},
	}

	fmt.Println(users)
	fmt.Println(users["1"].Name)
	//users["1"].Name = "jack"

	u := users["1"]

	fmt.Println(u.Name)
	u.Name = "jack"

	users["1"] = u
	fmt.Println(users["1"].Name)

	users2 := map[string]*User{
		"a2": &User{"user2"},
	}
	users2["a2"].Name = "Tom"
	fmt.Println(*users2["a2"])

	d := map[string]int{"b": 2, "c": 3, "e": 5}
	for k, v := range d {
		println(k, v)
		if k == "b" {
			delete(d, k)
		}
		if k == "c" {
			d["a"] = 1
		}
	}
	fmt.Println(d)
}
